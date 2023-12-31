// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: microservice.proto

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

type PaymentServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PaymentServiceRequest) Reset() {
	*x = PaymentServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentServiceRequest) ProtoMessage() {}

func (x *PaymentServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentServiceRequest.ProtoReflect.Descriptor instead.
func (*PaymentServiceRequest) Descriptor() ([]byte, []int) {
	return file_microservice_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentServiceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PaymentServiceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PaymentServiceReply) Reset() {
	*x = PaymentServiceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentServiceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentServiceReply) ProtoMessage() {}

func (x *PaymentServiceReply) ProtoReflect() protoreflect.Message {
	mi := &file_microservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentServiceReply.ProtoReflect.Descriptor instead.
func (*PaymentServiceReply) Descriptor() ([]byte, []int) {
	return file_microservice_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentServiceReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_microservice_proto protoreflect.FileDescriptor

var file_microservice_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x15, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a,
	0x13, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x49,
	0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x37, 0x0a, 0x05, 0x50, 0x61, 0x79, 0x54, 0x6f, 0x12, 0x16, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x6d, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_microservice_proto_rawDescOnce sync.Once
	file_microservice_proto_rawDescData = file_microservice_proto_rawDesc
)

func file_microservice_proto_rawDescGZIP() []byte {
	file_microservice_proto_rawDescOnce.Do(func() {
		file_microservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_microservice_proto_rawDescData)
	})
	return file_microservice_proto_rawDescData
}

var file_microservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_microservice_proto_goTypes = []interface{}{
	(*PaymentServiceRequest)(nil), // 0: PaymentServiceRequest
	(*PaymentServiceReply)(nil),   // 1: PaymentServiceReply
}
var file_microservice_proto_depIdxs = []int32{
	0, // 0: PaymentService.PayTo:input_type -> PaymentServiceRequest
	1, // 1: PaymentService.PayTo:output_type -> PaymentServiceReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_microservice_proto_init() }
func file_microservice_proto_init() {
	if File_microservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_microservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentServiceRequest); i {
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
		file_microservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentServiceReply); i {
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
			RawDescriptor: file_microservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_microservice_proto_goTypes,
		DependencyIndexes: file_microservice_proto_depIdxs,
		MessageInfos:      file_microservice_proto_msgTypes,
	}.Build()
	File_microservice_proto = out.File
	file_microservice_proto_rawDesc = nil
	file_microservice_proto_goTypes = nil
	file_microservice_proto_depIdxs = nil
}
