// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: heartbeat/rpc/pb/heartbeat.proto

package pb

import (
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

type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip      string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heartbeat_rpc_pb_heartbeat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_heartbeat_rpc_pb_heartbeat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_heartbeat_rpc_pb_heartbeat_proto_rawDescGZIP(), []int{0}
}

func (x *PingReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *PingReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PongRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip      string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PongRes) Reset() {
	*x = PongRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heartbeat_rpc_pb_heartbeat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongRes) ProtoMessage() {}

func (x *PongRes) ProtoReflect() protoreflect.Message {
	mi := &file_heartbeat_rpc_pb_heartbeat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongRes.ProtoReflect.Descriptor instead.
func (*PongRes) Descriptor() ([]byte, []int) {
	return file_heartbeat_rpc_pb_heartbeat_proto_rawDescGZIP(), []int{1}
}

func (x *PongRes) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *PongRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type OpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op     string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	Number int32  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *OpReq) Reset() {
	*x = OpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heartbeat_rpc_pb_heartbeat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpReq) ProtoMessage() {}

func (x *OpReq) ProtoReflect() protoreflect.Message {
	mi := &file_heartbeat_rpc_pb_heartbeat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpReq.ProtoReflect.Descriptor instead.
func (*OpReq) Descriptor() ([]byte, []int) {
	return file_heartbeat_rpc_pb_heartbeat_proto_rawDescGZIP(), []int{2}
}

func (x *OpReq) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *OpReq) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type OpRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag bool `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *OpRes) Reset() {
	*x = OpRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heartbeat_rpc_pb_heartbeat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpRes) ProtoMessage() {}

func (x *OpRes) ProtoReflect() protoreflect.Message {
	mi := &file_heartbeat_rpc_pb_heartbeat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpRes.ProtoReflect.Descriptor instead.
func (*OpRes) Descriptor() ([]byte, []int) {
	return file_heartbeat_rpc_pb_heartbeat_proto_rawDescGZIP(), []int{3}
}

func (x *OpRes) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

var File_heartbeat_rpc_pb_heartbeat_proto protoreflect.FileDescriptor

var file_heartbeat_rpc_pb_heartbeat_proto_rawDesc = []byte{
	0x0a, 0x20, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x62, 0x2f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x33, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x33, 0x0a, 0x07, 0x50,
	0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x2f, 0x0a, 0x05, 0x4f, 0x70, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x1b, 0x0a, 0x05, 0x4f, 0x70, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c,
	0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x32, 0x54,
	0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x70, 0x12, 0x09,
	0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4f,
	0x70, 0x52, 0x65, 0x73, 0x42, 0x13, 0x5a, 0x11, 0x2f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_heartbeat_rpc_pb_heartbeat_proto_rawDescOnce sync.Once
	file_heartbeat_rpc_pb_heartbeat_proto_rawDescData = file_heartbeat_rpc_pb_heartbeat_proto_rawDesc
)

func file_heartbeat_rpc_pb_heartbeat_proto_rawDescGZIP() []byte {
	file_heartbeat_rpc_pb_heartbeat_proto_rawDescOnce.Do(func() {
		file_heartbeat_rpc_pb_heartbeat_proto_rawDescData = protoimpl.X.CompressGZIP(file_heartbeat_rpc_pb_heartbeat_proto_rawDescData)
	})
	return file_heartbeat_rpc_pb_heartbeat_proto_rawDescData
}

var file_heartbeat_rpc_pb_heartbeat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_heartbeat_rpc_pb_heartbeat_proto_goTypes = []interface{}{
	(*PingReq)(nil), // 0: pb.PingReq
	(*PongRes)(nil), // 1: pb.PongRes
	(*OpReq)(nil),   // 2: pb.OpReq
	(*OpRes)(nil),   // 3: pb.OpRes
}
var file_heartbeat_rpc_pb_heartbeat_proto_depIdxs = []int32{
	0, // 0: pb.HeartbeatService.Ping:input_type -> pb.PingReq
	2, // 1: pb.HeartbeatService.SendOp:input_type -> pb.OpReq
	1, // 2: pb.HeartbeatService.Ping:output_type -> pb.PongRes
	3, // 3: pb.HeartbeatService.SendOp:output_type -> pb.OpRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_heartbeat_rpc_pb_heartbeat_proto_init() }
func file_heartbeat_rpc_pb_heartbeat_proto_init() {
	if File_heartbeat_rpc_pb_heartbeat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heartbeat_rpc_pb_heartbeat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
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
		file_heartbeat_rpc_pb_heartbeat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongRes); i {
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
		file_heartbeat_rpc_pb_heartbeat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpReq); i {
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
		file_heartbeat_rpc_pb_heartbeat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpRes); i {
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
			RawDescriptor: file_heartbeat_rpc_pb_heartbeat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heartbeat_rpc_pb_heartbeat_proto_goTypes,
		DependencyIndexes: file_heartbeat_rpc_pb_heartbeat_proto_depIdxs,
		MessageInfos:      file_heartbeat_rpc_pb_heartbeat_proto_msgTypes,
	}.Build()
	File_heartbeat_rpc_pb_heartbeat_proto = out.File
	file_heartbeat_rpc_pb_heartbeat_proto_rawDesc = nil
	file_heartbeat_rpc_pb_heartbeat_proto_goTypes = nil
	file_heartbeat_rpc_pb_heartbeat_proto_depIdxs = nil
}
