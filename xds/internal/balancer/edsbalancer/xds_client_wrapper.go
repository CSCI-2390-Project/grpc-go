/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package edsbalancer

import (
	"fmt"
	"sync"

	"github.com/CSCI-2390-Project/grpc-go"
	"github.com/CSCI-2390-Project/grpc-go/attributes"
	"github.com/CSCI-2390-Project/grpc-go/balancer"
	"github.com/CSCI-2390-Project/grpc-go/internal/grpclog"
	xdsinternal "github.com/CSCI-2390-Project/grpc-go/xds/internal"
	xdsclient "github.com/CSCI-2390-Project/grpc-go/xds/internal/client"
	"github.com/CSCI-2390-Project/grpc-go/xds/internal/client/bootstrap"
	"github.com/CSCI-2390-Project/grpc-go/xds/internal/client/load"
)

// xdsClientInterface contains only the xds_client methods needed by EDS
// balancer. It's defined so we can override xdsclientNew function in tests.
type xdsClientInterface interface {
	WatchEndpoints(clusterName string, edsCb func(xdsclient.EndpointsUpdate, error)) (cancel func())
	ReportLoad(server string) (loadStore *load.Store, cancel func())
	Close()
}

var (
	xdsclientNew = func(opts xdsclient.Options) (xdsClientInterface, error) {
		return xdsclient.New(opts)
	}
	bootstrapConfigNew = bootstrap.NewConfig
)

type loadStoreWrapper struct {
	mu      sync.RWMutex
	service string
	// Both store and perCluster will be nil if load reporting is disabled (EDS
	// response doesn't have LRS server name). Note that methods on Store and
	// perCluster all handle nil, so there's no need to check nil before calling
	// them.
	store      *load.Store
	perCluster load.PerClusterReporter
}

func (lsw *loadStoreWrapper) updateServiceName(service string) {
	lsw.mu.Lock()
	defer lsw.mu.Unlock()
	if lsw.service == service {
		return
	}
	lsw.service = service
	lsw.perCluster = lsw.store.PerCluster(lsw.service, "")
}

func (lsw *loadStoreWrapper) updateLoadStore(store *load.Store) {
	lsw.mu.Lock()
	defer lsw.mu.Unlock()
	if store == lsw.store {
		return
	}
	lsw.store = store
	lsw.perCluster = nil
	lsw.perCluster = lsw.store.PerCluster(lsw.service, "")
}

func (lsw *loadStoreWrapper) CallStarted(locality string) {
	lsw.mu.RLock()
	defer lsw.mu.RUnlock()
	lsw.perCluster.CallStarted(locality)
}

func (lsw *loadStoreWrapper) CallFinished(locality string, err error) {
	lsw.mu.RLock()
	defer lsw.mu.RUnlock()
	lsw.perCluster.CallFinished(locality, err)
}

func (lsw *loadStoreWrapper) CallServerLoad(locality, name string, val float64) {
	lsw.mu.RLock()
	defer lsw.mu.RUnlock()
	lsw.perCluster.CallServerLoad(locality, name, val)
}

func (lsw *loadStoreWrapper) CallDropped(category string) {
	lsw.mu.RLock()
	defer lsw.mu.RUnlock()
	lsw.perCluster.CallDropped(category)
}

// xdsclientWrapper is responsible for getting the xds client from attributes or
// creating a new xds client, and start watching EDS. The given callbacks will
// be called with EDS updates or errors.
type xdsClientWrapper struct {
	logger *grpclog.PrefixLogger

	newEDSUpdate func(xdsclient.EndpointsUpdate, error)
	bbo          balancer.BuildOptions

	balancerName string
	// xdsClient could come from attributes, or created with balancerName.
	xdsClient xdsClientInterface

	// loadWrapper is a wrapper with loadOriginal, with clusterName and
	// edsServiceName. It's used children to report loads.
	loadWrapper *loadStoreWrapper
	// edsServiceName is the edsServiceName currently being watched, not
	// necessary the edsServiceName from service config.
	//
	// If edsServiceName from service config is an empty, this will be user's
	// dial target (because that's what we use to watch EDS).
	//
	// TODO: remove the empty string related behavior, when we switch to always
	// do CDS.
	edsServiceName       string
	cancelEndpointsWatch func()
	loadReportServer     *string // LRS is disabled if loadReporterServer is nil.
	cancelLoadReport     func()
}

// newXDSClientWrapper creates an empty xds_client wrapper that does nothing. It
// can accept xds_client configs, to new/switch xds_client to use.
//
// The given callbacks won't be called until the underlying xds_client is
// working and sends updates.
func newXDSClientWrapper(newEDSUpdate func(xdsclient.EndpointsUpdate, error), bbo balancer.BuildOptions, logger *grpclog.PrefixLogger) *xdsClientWrapper {
	return &xdsClientWrapper{
		logger:       logger,
		newEDSUpdate: newEDSUpdate,
		bbo:          bbo,
		loadWrapper:  &loadStoreWrapper{},
	}
}

// replaceXDSClient replaces xdsClient fields to the newClient if they are
// different. If xdsClient is replaced, the balancerName field will also be
// updated to newBalancerName.
//
// If the old xdsClient is replaced, and was created locally (not from
// attributes), it will be closed.
//
// It returns whether xdsClient is replaced.
func (c *xdsClientWrapper) replaceXDSClient(newClient xdsClientInterface, newBalancerName string) bool {
	if c.xdsClient == newClient {
		return false
	}
	oldClient := c.xdsClient
	oldBalancerName := c.balancerName
	c.xdsClient = newClient
	c.balancerName = newBalancerName
	if oldBalancerName != "" {
		// OldBalancerName!="" means if the old client was not from attributes.
		oldClient.Close()
	}
	return true
}

// updateXDSClient sets xdsClient in wrapper to the correct one based on the
// attributes and service config.
//
// If client is found in attributes, it will be used, but we also need to decide
// whether to close the old client.
// - if old client was created locally (balancerName is not ""), close it and
// replace it
// - if old client was from previous attributes, only replace it, but don't
// close it
//
// If client is not found in attributes, will need to create a new one only if
// the balancerName (from bootstrap file or from service config) changed.
// - if balancer names are the same, do nothing, and return false
// - if balancer names are different, create new one, and return true
func (c *xdsClientWrapper) updateXDSClient(config *EDSConfig, attr *attributes.Attributes) (bool, error) {
	if attr != nil {
		if clientFromAttr, _ := attr.Value(xdsinternal.XDSClientID).(xdsClientInterface); clientFromAttr != nil {
			// This will also clear balancerName, to indicate that client is
			// from attributes.
			return c.replaceXDSClient(clientFromAttr, ""), nil
		}
	}

	clientConfig, err := bootstrapConfigNew()
	if err != nil {
		// TODO: propagate this error to ClientConn, and fail RPCs if necessary.
		clientConfig = &bootstrap.Config{BalancerName: config.BalancerName}
	}

	if c.balancerName == clientConfig.BalancerName {
		return false, nil
	}

	var dopts []grpc.DialOption
	if dialer := c.bbo.Dialer; dialer != nil {
		dopts = []grpc.DialOption{grpc.WithContextDialer(dialer)}
	}

	// TODO: there's no longer a need to read bootstrap file and create a new
	// xds client. The EDS balancer should always get the xds client from
	// attributes. Otherwise, this function should just fail. Also, xdsclient
	// will be shared by multiple clients, so trying to make an xds client is
	// just the wrong move.
	newClient, err := xdsclientNew(xdsclient.Options{Config: *clientConfig, DialOpts: dopts})
	if err != nil {
		// This should never fail. xdsclientnew does a non-blocking dial, and
		// all the config passed in should be validated.
		//
		// This could leave c.xdsClient as nil if this is the first update.
		return false, fmt.Errorf("eds: failed to create xdsClient, error: %v", err)
	}
	return c.replaceXDSClient(newClient, clientConfig.BalancerName), nil
}

// startEndpointsWatch starts the EDS watch. Caller can call this when the
// xds_client is updated, or the edsServiceName is updated.
//
// Note that if there's already a watch in progress, it's not explicitly
// canceled. Because for each xds_client, there should be only one EDS watch in
// progress. So a new EDS watch implicitly cancels the previous one.
//
// This usually means load report needs to be restarted, but this function does
// NOT do that. Caller needs to call startLoadReport separately.
func (c *xdsClientWrapper) startEndpointsWatch() {
	if c.cancelEndpointsWatch != nil {
		c.cancelEndpointsWatch()
	}
	cancelEDSWatch := c.xdsClient.WatchEndpoints(c.edsServiceName, func(update xdsclient.EndpointsUpdate, err error) {
		c.logger.Infof("Watch update from xds-client %p, content: %+v", c.xdsClient, update)
		c.newEDSUpdate(update, err)
	})
	c.logger.Infof("Watch started on resource name %v with xds-client %p", c.edsServiceName, c.xdsClient)
	c.cancelEndpointsWatch = func() {
		cancelEDSWatch()
		c.logger.Infof("Watch cancelled on resource name %v with xds-client %p", c.edsServiceName, c.xdsClient)
	}
}

// startLoadReport starts load reporting. If there's already a load reporting in
// progress, it cancels that.
//
// Caller can cal this when the loadReportServer name changes, but
// edsServiceName doesn't (so we only need to restart load reporting, not EDS
// watch).
func (c *xdsClientWrapper) startLoadReport(loadReportServer *string) *load.Store {
	if c.cancelLoadReport != nil {
		c.cancelLoadReport()
	}
	c.loadReportServer = loadReportServer
	var loadStore *load.Store
	if c.loadReportServer != nil {
		loadStore, c.cancelLoadReport = c.xdsClient.ReportLoad(*c.loadReportServer)
	}
	return loadStore
}

func (c *xdsClientWrapper) loadStore() load.PerClusterReporter {
	if c == nil {
		return nil
	}
	return c.loadWrapper
}

// handleUpdate applies the service config and attributes updates to the client,
// including updating the xds_client to use, and updating the EDS name to watch.
func (c *xdsClientWrapper) handleUpdate(config *EDSConfig, attr *attributes.Attributes) error {
	clientChanged, err := c.updateXDSClient(config, attr)
	if err != nil {
		return err
	}

	// Need to restart EDS watch when one of the following happens:
	// - the xds_client is updated
	// - the xds_client didn't change, but the edsServiceName changed
	if clientChanged || c.edsServiceName != config.EDSServiceName {
		c.edsServiceName = config.EDSServiceName
		c.startEndpointsWatch()
		// TODO: this update for the LRS service name is too early. It should
		// only apply to the new EDS response. But this is applied to the RPCs
		// before the new EDS response. To fully fix this, the EDS balancer
		// needs to do a graceful switch to another EDS implementation.
		//
		// This is OK for now, because we don't actually expect edsServiceName
		// to change. Fix this (a bigger change) will happen later.
		c.loadWrapper.updateServiceName(c.edsServiceName)
	}

	// Only need to restart load reporting when:
	// - the loadReportServer name changed
	if !equalStringPointers(c.loadReportServer, config.LrsLoadReportingServerName) {
		loadStore := c.startLoadReport(config.LrsLoadReportingServerName)
		c.loadWrapper.updateLoadStore(loadStore)
	}

	return nil
}

func (c *xdsClientWrapper) cancelWatch() {
	c.loadReportServer = nil
	if c.cancelLoadReport != nil {
		c.cancelLoadReport()
	}
	c.edsServiceName = ""
	if c.cancelEndpointsWatch != nil {
		c.cancelEndpointsWatch()
	}
}

func (c *xdsClientWrapper) close() {
	c.cancelWatch()
	if c.xdsClient != nil && c.balancerName != "" {
		// Only close xdsClient if it's not from attributes.
		c.xdsClient.Close()
	}
}

// equalStringPointers returns true if
// - a and b are both nil OR
// - *a == *b (and a and b are both non-nil)
func equalStringPointers(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
