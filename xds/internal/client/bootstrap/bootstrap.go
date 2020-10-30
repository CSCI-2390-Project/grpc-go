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

// Package bootstrap provides the functionality to initialize certain aspects
// of an xDS client by reading a bootstrap file.
package bootstrap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/CSCI-2390-Project/grpc-go"
	"github.com/CSCI-2390-Project/grpc-go/credentials/google"
	"github.com/CSCI-2390-Project/grpc-go/credentials/tls/certprovider"
	"github.com/CSCI-2390-Project/grpc-go/internal"
	"github.com/CSCI-2390-Project/grpc-go/xds/internal/version"
	"github.com/CSCI-2390-Project/protobuf/jsonpb"
	"github.com/CSCI-2390-Project/protobuf/proto"
	v2corepb "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v3corepb "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
)

const (
	// Environment variable which holds the name of the xDS bootstrap file.
	bootstrapFileEnv = "GRPC_XDS_BOOTSTRAP"
	// Environment variable which controls the use of xDS v3 API.
	v3SupportEnv = "GRPC_XDS_EXPERIMENTAL_V3_SUPPORT"
	// The "server_features" field in the bootstrap file contains a list of
	// features supported by the server. A value of "xds_v3" indicates that the
	// server supports the v3 version of the xDS transport protocol.
	serverFeaturesV3 = "xds_v3"

	// Type name for Google default credentials.
	credsGoogleDefault              = "google_default"
	credsInsecure                   = "insecure"
	gRPCUserAgentName               = "gRPC Go"
	clientFeatureNoOverprovisioning = "envoy.lb.does_not_support_overprovisioning"
)

var gRPCVersion = fmt.Sprintf("%s %s", gRPCUserAgentName, grpc.Version)

// For overriding in unit tests.
var bootstrapFileReadFunc = ioutil.ReadFile

// Config provides the xDS client with several key bits of information that it
// requires in its interaction with an xDS server. The Config is initialized
// from the bootstrap file.
type Config struct {
	// BalancerName is the name of the xDS server to connect to.
	//
	// The bootstrap file contains a list of servers (with name+creds), but we
	// pick the first one.
	BalancerName string
	// Creds contains the credentials to be used while talking to the xDS
	// server, as a grpc.DialOption.
	Creds grpc.DialOption
	// TransportAPI indicates the API version of xDS transport protocol to use.
	// This describes the xDS gRPC endpoint and version of
	// DiscoveryRequest/Response used on the wire.
	TransportAPI version.TransportAPI
	// NodeProto contains the Node proto to be used in xDS requests. The actual
	// type depends on the transport protocol version used.
	NodeProto proto.Message
	// CertProviderConfigs contain parsed configs for supported certificate
	// provider plugins found in the bootstrap file.
	CertProviderConfigs map[string]CertProviderConfig
}

// CertProviderConfig wraps the certificate provider plugin name and config
// (corresponding to one plugin instance) found in the bootstrap file.
type CertProviderConfig struct {
	// Name is the registered name of the certificate provider.
	Name string
	// Config is the parsed config to be passed to the certificate provider.
	Config certprovider.StableConfig
}

type channelCreds struct {
	Type   string          `json:"type"`
	Config json.RawMessage `json:"config"`
}

type xdsServer struct {
	ServerURI    string         `json:"server_uri"`
	ChannelCreds []channelCreds `json:"channel_creds"`
}

// NewConfig returns a new instance of Config initialized by reading the
// bootstrap file found at ${GRPC_XDS_BOOTSTRAP}.
//
// The format of the bootstrap file will be as follows:
// {
//    "xds_server": {
//      "server_uri": <string containing URI of xds server>,
//      "channel_creds": [
//        {
//          "type": <string containing channel cred type>,
//          "config": <JSON object containing config for the type>
//        }
//      ],
//      "server_features": [ ... ]
//		"certificate_providers" : {
//			"default": {
//				"plugin_name": "default-plugin-name",
//				"config": { default plugin config in JSON }
//			},
//			"foo": {
//				"plugin_name": "foo",
//				"config": { foo plugin config in JSON }
//			}
//		}
//    },
//    "node": <JSON form of Node proto>
// }
//
// Currently, we support exactly one type of credential, which is
// "google_default", where we use the host's default certs for transport
// credentials and a Google oauth token for call credentials.
//
// This function tries to process as much of the bootstrap file as possible (in
// the presence of the errors) and may return a Config object with certain
// fields left unspecified, in which case the caller should use some sane
// defaults.
func NewConfig() (*Config, error) {
	config := &Config{}

	fName, ok := os.LookupEnv(bootstrapFileEnv)
	if !ok {
		return nil, fmt.Errorf("xds: Environment variable %v not defined", bootstrapFileEnv)
	}
	logger.Infof("Got bootstrap file location from %v environment variable: %v", bootstrapFileEnv, fName)

	data, err := bootstrapFileReadFunc(fName)
	if err != nil {
		return nil, fmt.Errorf("xds: Failed to read bootstrap file %s with error %v", fName, err)
	}
	logger.Debugf("Bootstrap content: %s", data)

	var jsonData map[string]json.RawMessage
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, fmt.Errorf("xds: Failed to parse file %s (content %v) with error: %v", fName, string(data), err)
	}

	serverSupportsV3 := false
	m := jsonpb.Unmarshaler{AllowUnknownFields: true}
	for k, v := range jsonData {
		switch k {
		case "node":
			// We unconditionally convert the JSON into a v3.Node proto. The v3
			// proto does not contain the deprecated field "build_version" from
			// the v2 proto. We do not expect the bootstrap file to contain the
			// "build_version" field. In any case, the unmarshal will succeed
			// because we have set the `AllowUnknownFields` option on the
			// unmarshaler.
			n := &v3corepb.Node{}
			if err := m.Unmarshal(bytes.NewReader(v), n); err != nil {
				return nil, fmt.Errorf("xds: jsonpb.Unmarshal(%v) for field %q failed during bootstrap: %v", string(v), k, err)
			}
			config.NodeProto = n
		case "xds_servers":
			var servers []*xdsServer
			if err := json.Unmarshal(v, &servers); err != nil {
				return nil, fmt.Errorf("xds: json.Unmarshal(%v) for field %q failed during bootstrap: %v", string(v), k, err)
			}
			if len(servers) < 1 {
				return nil, fmt.Errorf("xds: bootstrap file parsing failed during bootstrap: file doesn't contain any xds server to connect to")
			}
			xs := servers[0]
			config.BalancerName = xs.ServerURI
			for _, cc := range xs.ChannelCreds {
				// We stop at the first credential type that we support.
				if cc.Type == credsGoogleDefault {
					config.Creds = grpc.WithCredentialsBundle(google.NewDefaultCredentials())
					break
				} else if cc.Type == credsInsecure {
					config.Creds = grpc.WithInsecure()
					break
				}
			}
		case "server_features":
			var features []string
			if err := json.Unmarshal(v, &features); err != nil {
				return nil, fmt.Errorf("xds: json.Unmarshal(%v) for field %q failed during bootstrap: %v", string(v), k, err)
			}
			for _, f := range features {
				switch f {
				case serverFeaturesV3:
					serverSupportsV3 = true
				}
			}
		case "certificate_providers":
			var providerInstances map[string]json.RawMessage
			if err := json.Unmarshal(v, &providerInstances); err != nil {
				return nil, fmt.Errorf("xds: json.Unmarshal(%v) for field %q failed during bootstrap: %v", string(v), k, err)
			}
			configs := make(map[string]CertProviderConfig)
			getBuilder := internal.GetCertificateProviderBuilder.(func(string) certprovider.Builder)
			for instance, data := range providerInstances {
				var nameAndConfig struct {
					PluginName string          `json:"plugin_name"`
					Config     json.RawMessage `json:"config"`
				}
				if err := json.Unmarshal(data, &nameAndConfig); err != nil {
					return nil, fmt.Errorf("xds: json.Unmarshal(%v) for field %q failed during bootstrap: %v", string(v), instance, err)
				}

				name := nameAndConfig.PluginName
				parser := getBuilder(nameAndConfig.PluginName)
				if parser == nil {
					// We ignore plugins that we do not know about.
					continue
				}
				cfg := nameAndConfig.Config
				c, err := parser.ParseConfig(cfg)
				if err != nil {
					return nil, fmt.Errorf("xds: Config parsing for plugin %q failed: %v", name, err)
				}
				configs[instance] = CertProviderConfig{
					Name:   name,
					Config: c,
				}
			}
			config.CertProviderConfigs = configs
		}
		// Do not fail the xDS bootstrap when an unknown field is seen. This can
		// happen when an older version client reads a newer version bootstrap
		// file with new fields.
	}

	if config.BalancerName == "" {
		return nil, fmt.Errorf("xds: Required field %q not found in bootstrap %s", "xds_servers.server_uri", jsonData["xds_servers"])
	}
	if config.Creds == nil {
		return nil, fmt.Errorf("xds: Required field %q doesn't contain valid value in bootstrap %s", "xds_servers.channel_creds", jsonData["xds_servers"])
	}

	// We end up using v3 transport protocol version only if the following
	// conditions are met:
	// 1. Server supports v3, indicated by the presence of "xds_v3" in
	//    server_features.
	// 2. Environment variable "GRPC_XDS_EXPERIMENTAL_V3_SUPPORT" is set to
	//    true.
	// The default value of the enum type "version.TransportAPI" is v2.
	//
	// TODO: there are multiple env variables, GRPC_XDS_BOOTSTRAP and
	// GRPC_XDS_EXPERIMENTAL_V3_SUPPORT. Move all env variables into a separate
	// package.
	if v3Env := os.Getenv(v3SupportEnv); v3Env == "true" {
		if serverSupportsV3 {
			config.TransportAPI = version.TransportV3
		}
	}

	if err := config.updateNodeProto(); err != nil {
		return nil, err
	}
	logger.Infof("Bootstrap config for creating xds-client: %+v", config)
	return config, nil
}

// updateNodeProto updates the node proto read from the bootstrap file.
//
// Node proto in Config contains a v3.Node protobuf message corresponding to the
// JSON contents found in the bootstrap file. This method performs some post
// processing on it:
// 1. If we don't find a nodeProto in the bootstrap file, we create an empty one
// here. That way, callers of this function can always expect that the NodeProto
// field is non-nil.
// 2. If the transport protocol version to be used is not v3, we convert the
// current v3.Node proto in a v2.Node proto.
// 3. Some additional fields which are not expected to be set in the bootstrap
// file are populated here.
func (c *Config) updateNodeProto() error {
	if c.TransportAPI == version.TransportV3 {
		v3, _ := c.NodeProto.(*v3corepb.Node)
		if v3 == nil {
			v3 = &v3corepb.Node{}
		}
		v3.UserAgentName = gRPCUserAgentName
		v3.UserAgentVersionType = &v3corepb.Node_UserAgentVersion{UserAgentVersion: grpc.Version}
		v3.ClientFeatures = append(v3.ClientFeatures, clientFeatureNoOverprovisioning)
		c.NodeProto = v3
		return nil
	}

	v2 := &v2corepb.Node{}
	if c.NodeProto != nil {
		v3, err := proto.Marshal(c.NodeProto)
		if err != nil {
			return fmt.Errorf("xds: proto.Marshal(%v): %v", c.NodeProto, err)
		}
		if err := proto.Unmarshal(v3, v2); err != nil {
			return fmt.Errorf("xds: proto.Unmarshal(%v): %v", v3, err)
		}
	}
	c.NodeProto = v2

	// BuildVersion is deprecated, and is replaced by user_agent_name and
	// user_agent_version. But the management servers are still using the old
	// field, so we will keep both set.
	v2.BuildVersion = gRPCVersion
	v2.UserAgentName = gRPCUserAgentName
	v2.UserAgentVersionType = &v2corepb.Node_UserAgentVersion{UserAgentVersion: grpc.Version}
	v2.ClientFeatures = append(v2.ClientFeatures, clientFeatureNoOverprovisioning)
	return nil
}
