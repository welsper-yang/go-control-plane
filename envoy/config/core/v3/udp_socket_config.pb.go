// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/config/core/v3/udp_socket_config.proto

package envoy_config_core_v3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Generic UDP socket configuration.
type UdpSocketConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum size of received UDP datagrams. Using a larger size will cause Envoy to allocate
	// more memory per socket. Received datagrams above this size will be dropped. If not set
	// defaults to 1500 bytes.
	MaxRxDatagramSize *wrappers.UInt64Value `protobuf:"bytes,1,opt,name=max_rx_datagram_size,json=maxRxDatagramSize,proto3" json:"max_rx_datagram_size,omitempty"`
}

func (x *UdpSocketConfig) Reset() {
	*x = UdpSocketConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_udp_socket_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UdpSocketConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UdpSocketConfig) ProtoMessage() {}

func (x *UdpSocketConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_udp_socket_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UdpSocketConfig.ProtoReflect.Descriptor instead.
func (*UdpSocketConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_udp_socket_config_proto_rawDescGZIP(), []int{0}
}

func (x *UdpSocketConfig) GetMaxRxDatagramSize() *wrappers.UInt64Value {
	if x != nil {
		return x.MaxRxDatagramSize
	}
	return nil
}

var File_envoy_config_core_v3_udp_socket_config_proto protoreflect.FileDescriptor

var file_envoy_config_core_v3_udp_socket_config_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x64, 0x70, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x0f,
	0x55, 0x64, 0x70, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x5a, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x78, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x67, 0x72,
	0x61, 0x6d, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0xfa, 0x42, 0x08,
	0x32, 0x06, 0x10, 0x80, 0x80, 0x04, 0x20, 0x00, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x52, 0x78, 0x44,
	0x61, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x44, 0x0a, 0x22, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x33, 0x42, 0x14, 0x55, 0x64, 0x70, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_core_v3_udp_socket_config_proto_rawDescOnce sync.Once
	file_envoy_config_core_v3_udp_socket_config_proto_rawDescData = file_envoy_config_core_v3_udp_socket_config_proto_rawDesc
)

func file_envoy_config_core_v3_udp_socket_config_proto_rawDescGZIP() []byte {
	file_envoy_config_core_v3_udp_socket_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_core_v3_udp_socket_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_core_v3_udp_socket_config_proto_rawDescData)
	})
	return file_envoy_config_core_v3_udp_socket_config_proto_rawDescData
}

var file_envoy_config_core_v3_udp_socket_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_core_v3_udp_socket_config_proto_goTypes = []interface{}{
	(*UdpSocketConfig)(nil),      // 0: envoy.config.core.v3.UdpSocketConfig
	(*wrappers.UInt64Value)(nil), // 1: google.protobuf.UInt64Value
}
var file_envoy_config_core_v3_udp_socket_config_proto_depIdxs = []int32{
	1, // 0: envoy.config.core.v3.UdpSocketConfig.max_rx_datagram_size:type_name -> google.protobuf.UInt64Value
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_core_v3_udp_socket_config_proto_init() }
func file_envoy_config_core_v3_udp_socket_config_proto_init() {
	if File_envoy_config_core_v3_udp_socket_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_core_v3_udp_socket_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UdpSocketConfig); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_core_v3_udp_socket_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_core_v3_udp_socket_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_core_v3_udp_socket_config_proto_depIdxs,
		MessageInfos:      file_envoy_config_core_v3_udp_socket_config_proto_msgTypes,
	}.Build()
	File_envoy_config_core_v3_udp_socket_config_proto = out.File
	file_envoy_config_core_v3_udp_socket_config_proto_rawDesc = nil
	file_envoy_config_core_v3_udp_socket_config_proto_goTypes = nil
	file_envoy_config_core_v3_udp_socket_config_proto_depIdxs = nil
}
