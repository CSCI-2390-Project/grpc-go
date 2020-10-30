// Copyright 2016 gRPC authors.
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

// An integration test service that covers all the method signature permutations
// of unary/streaming requests/responses.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.3.0
// source: benchmark/grpc_testing/services.proto

package grpc_testing

import (
	protoreflect "github.com/CSCI-2390-Project/protobuf-go/reflect/protoreflect"
	protoimpl "github.com/CSCI-2390-Project/protobuf-go/runtime/protoimpl"
	proto "github.com/CSCI-2390-Project/protobuf/proto"
	reflect "reflect"
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

var File_benchmark_grpc_testing_services_proto protoreflect.FileDescriptor

var file_benchmark_grpc_testing_services_proto_rawDesc = []byte{
	0x0a, 0x25, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x25, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x62, 0x65,
	0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x87, 0x02, 0x0a, 0x10, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x55, 0x6e, 0x61, 0x72, 0x79,
	0x43, 0x61, 0x6c, 0x6c, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x6c, 0x6c,
	0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x5b, 0x0a, 0x1a, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x1b, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x32, 0x97, 0x02, 0x0a,
	0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45,
	0x0a, 0x09, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x28, 0x01, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x1a, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x28, 0x01, 0x30, 0x01, 0x12, 0x42, 0x0a, 0x09,
	0x43, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x0a, 0x51, 0x75, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x12,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_benchmark_grpc_testing_services_proto_goTypes = []interface{}{
	(*SimpleRequest)(nil),  // 0: grpc.testing.SimpleRequest
	(*ServerArgs)(nil),     // 1: grpc.testing.ServerArgs
	(*ClientArgs)(nil),     // 2: grpc.testing.ClientArgs
	(*CoreRequest)(nil),    // 3: grpc.testing.CoreRequest
	(*Void)(nil),           // 4: grpc.testing.Void
	(*SimpleResponse)(nil), // 5: grpc.testing.SimpleResponse
	(*ServerStatus)(nil),   // 6: grpc.testing.ServerStatus
	(*ClientStatus)(nil),   // 7: grpc.testing.ClientStatus
	(*CoreResponse)(nil),   // 8: grpc.testing.CoreResponse
}
var file_benchmark_grpc_testing_services_proto_depIdxs = []int32{
	0, // 0: grpc.testing.BenchmarkService.UnaryCall:input_type -> grpc.testing.SimpleRequest
	0, // 1: grpc.testing.BenchmarkService.StreamingCall:input_type -> grpc.testing.SimpleRequest
	0, // 2: grpc.testing.BenchmarkService.UnconstrainedStreamingCall:input_type -> grpc.testing.SimpleRequest
	1, // 3: grpc.testing.WorkerService.RunServer:input_type -> grpc.testing.ServerArgs
	2, // 4: grpc.testing.WorkerService.RunClient:input_type -> grpc.testing.ClientArgs
	3, // 5: grpc.testing.WorkerService.CoreCount:input_type -> grpc.testing.CoreRequest
	4, // 6: grpc.testing.WorkerService.QuitWorker:input_type -> grpc.testing.Void
	5, // 7: grpc.testing.BenchmarkService.UnaryCall:output_type -> grpc.testing.SimpleResponse
	5, // 8: grpc.testing.BenchmarkService.StreamingCall:output_type -> grpc.testing.SimpleResponse
	5, // 9: grpc.testing.BenchmarkService.UnconstrainedStreamingCall:output_type -> grpc.testing.SimpleResponse
	6, // 10: grpc.testing.WorkerService.RunServer:output_type -> grpc.testing.ServerStatus
	7, // 11: grpc.testing.WorkerService.RunClient:output_type -> grpc.testing.ClientStatus
	8, // 12: grpc.testing.WorkerService.CoreCount:output_type -> grpc.testing.CoreResponse
	4, // 13: grpc.testing.WorkerService.QuitWorker:output_type -> grpc.testing.Void
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_benchmark_grpc_testing_services_proto_init() }
func file_benchmark_grpc_testing_services_proto_init() {
	if File_benchmark_grpc_testing_services_proto != nil {
		return
	}
	file_benchmark_grpc_testing_messages_proto_init()
	file_benchmark_grpc_testing_control_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_benchmark_grpc_testing_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_benchmark_grpc_testing_services_proto_goTypes,
		DependencyIndexes: file_benchmark_grpc_testing_services_proto_depIdxs,
	}.Build()
	File_benchmark_grpc_testing_services_proto = out.File
	file_benchmark_grpc_testing_services_proto_rawDesc = nil
	file_benchmark_grpc_testing_services_proto_goTypes = nil
	file_benchmark_grpc_testing_services_proto_depIdxs = nil
}
