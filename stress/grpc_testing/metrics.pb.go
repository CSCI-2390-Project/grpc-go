// Copyright 2015-2016 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Contains the definitions for a metrics service and the type of metrics
// exposed by the service.
//
// Currently, 'Gauge' (i.e a metric that represents the measured value of
// something at an instant of time) is the only metric type supported by the
// service.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.3.0
// source: stress/grpc_testing/metrics.proto

package grpc_testing

import (
	protoreflect "github.com/CSCI-2390-Project/protobuf-go/reflect/protoreflect"
	protoimpl "github.com/CSCI-2390-Project/protobuf-go/runtime/protoimpl"
	proto "github.com/CSCI-2390-Project/protobuf/proto"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Response message containing the gauge name and value
type GaugeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Value:
	//	*GaugeResponse_LongValue
	//	*GaugeResponse_DoubleValue
	//	*GaugeResponse_StringValue
	Value isGaugeResponse_Value `protobuf_oneof:"value"`
}

func (x *GaugeResponse) Reset() {
	*x = GaugeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stress_grpc_testing_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GaugeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GaugeResponse) ProtoMessage() {}

func (x *GaugeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stress_grpc_testing_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GaugeResponse.ProtoReflect.Descriptor instead.
func (*GaugeResponse) Descriptor() ([]byte, []int) {
	return file_stress_grpc_testing_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *GaugeResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *GaugeResponse) GetValue() isGaugeResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *GaugeResponse) GetLongValue() int64 {
	if x, ok := x.GetValue().(*GaugeResponse_LongValue); ok {
		return x.LongValue
	}
	return 0
}

func (x *GaugeResponse) GetDoubleValue() float64 {
	if x, ok := x.GetValue().(*GaugeResponse_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *GaugeResponse) GetStringValue() string {
	if x, ok := x.GetValue().(*GaugeResponse_StringValue); ok {
		return x.StringValue
	}
	return ""
}

type isGaugeResponse_Value interface {
	isGaugeResponse_Value()
}

type GaugeResponse_LongValue struct {
	LongValue int64 `protobuf:"varint,2,opt,name=long_value,json=longValue,proto3,oneof"`
}

type GaugeResponse_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type GaugeResponse_StringValue struct {
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*GaugeResponse_LongValue) isGaugeResponse_Value() {}

func (*GaugeResponse_DoubleValue) isGaugeResponse_Value() {}

func (*GaugeResponse_StringValue) isGaugeResponse_Value() {}

// Request message containing the gauge name
type GaugeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GaugeRequest) Reset() {
	*x = GaugeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stress_grpc_testing_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GaugeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GaugeRequest) ProtoMessage() {}

func (x *GaugeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stress_grpc_testing_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GaugeRequest.ProtoReflect.Descriptor instead.
func (*GaugeRequest) Descriptor() ([]byte, []int) {
	return file_stress_grpc_testing_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *GaugeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stress_grpc_testing_metrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_stress_grpc_testing_metrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_stress_grpc_testing_metrics_proto_rawDescGZIP(), []int{2}
}

var File_stress_grpc_testing_metrics_proto protoreflect.FileDescriptor

var file_stress_grpc_testing_metrics_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x74, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0x97, 0x01, 0x0a, 0x0d, 0x47, 0x61, 0x75, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x6c, 0x6f, 0x6e, 0x67, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00,
	0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a,
	0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x47,
	0x61, 0x75, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0xa0, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x61, 0x75, 0x67,
	0x65, 0x73, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1b,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x61,
	0x75, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x43, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x47, 0x61, 0x75, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x61, 0x75, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x61, 0x75, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x73, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stress_grpc_testing_metrics_proto_rawDescOnce sync.Once
	file_stress_grpc_testing_metrics_proto_rawDescData = file_stress_grpc_testing_metrics_proto_rawDesc
)

func file_stress_grpc_testing_metrics_proto_rawDescGZIP() []byte {
	file_stress_grpc_testing_metrics_proto_rawDescOnce.Do(func() {
		file_stress_grpc_testing_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_stress_grpc_testing_metrics_proto_rawDescData)
	})
	return file_stress_grpc_testing_metrics_proto_rawDescData
}

var file_stress_grpc_testing_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_stress_grpc_testing_metrics_proto_goTypes = []interface{}{
	(*GaugeResponse)(nil), // 0: grpc.testing.GaugeResponse
	(*GaugeRequest)(nil),  // 1: grpc.testing.GaugeRequest
	(*EmptyMessage)(nil),  // 2: grpc.testing.EmptyMessage
}
var file_stress_grpc_testing_metrics_proto_depIdxs = []int32{
	2, // 0: grpc.testing.MetricsService.GetAllGauges:input_type -> grpc.testing.EmptyMessage
	1, // 1: grpc.testing.MetricsService.GetGauge:input_type -> grpc.testing.GaugeRequest
	0, // 2: grpc.testing.MetricsService.GetAllGauges:output_type -> grpc.testing.GaugeResponse
	0, // 3: grpc.testing.MetricsService.GetGauge:output_type -> grpc.testing.GaugeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stress_grpc_testing_metrics_proto_init() }
func file_stress_grpc_testing_metrics_proto_init() {
	if File_stress_grpc_testing_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stress_grpc_testing_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GaugeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stress_grpc_testing_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GaugeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stress_grpc_testing_metrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_stress_grpc_testing_metrics_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GaugeResponse_LongValue)(nil),
		(*GaugeResponse_DoubleValue)(nil),
		(*GaugeResponse_StringValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stress_grpc_testing_metrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stress_grpc_testing_metrics_proto_goTypes,
		DependencyIndexes: file_stress_grpc_testing_metrics_proto_depIdxs,
		MessageInfos:      file_stress_grpc_testing_metrics_proto_msgTypes,
	}.Build()
	File_stress_grpc_testing_metrics_proto = out.File
	file_stress_grpc_testing_metrics_proto_rawDesc = nil
	file_stress_grpc_testing_metrics_proto_goTypes = nil
	file_stress_grpc_testing_metrics_proto_depIdxs = nil
}
