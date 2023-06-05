// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: auth.proto

package main

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

type ReqPqInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce     string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	MessageId int32  `protobuf:"varint,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *ReqPqInput) Reset() {
	*x = ReqPqInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqPqInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPqInput) ProtoMessage() {}

func (x *ReqPqInput) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPqInput.ProtoReflect.Descriptor instead.
func (*ReqPqInput) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ReqPqInput) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *ReqPqInput) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type ReqPqResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce string `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	MessageId   int32  `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	P           int32  `protobuf:"varint,4,opt,name=p,proto3" json:"p,omitempty"`
	G           int32  `protobuf:"varint,5,opt,name=g,proto3" json:"g,omitempty"`
}

func (x *ReqPqResponse) Reset() {
	*x = ReqPqResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqPqResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPqResponse) ProtoMessage() {}

func (x *ReqPqResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPqResponse.ProtoReflect.Descriptor instead.
func (*ReqPqResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ReqPqResponse) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *ReqPqResponse) GetServerNonce() string {
	if x != nil {
		return x.ServerNonce
	}
	return ""
}

func (x *ReqPqResponse) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *ReqPqResponse) GetP() int32 {
	if x != nil {
		return x.P
	}
	return 0
}

func (x *ReqPqResponse) GetG() int32 {
	if x != nil {
		return x.G
	}
	return 0
}

type Req_DHParamsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce string `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	MessageId   int32  `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	A           int32  `protobuf:"varint,4,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *Req_DHParamsInput) Reset() {
	*x = Req_DHParamsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req_DHParamsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req_DHParamsInput) ProtoMessage() {}

func (x *Req_DHParamsInput) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req_DHParamsInput.ProtoReflect.Descriptor instead.
func (*Req_DHParamsInput) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *Req_DHParamsInput) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *Req_DHParamsInput) GetServerNonce() string {
	if x != nil {
		return x.ServerNonce
	}
	return ""
}

func (x *Req_DHParamsInput) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *Req_DHParamsInput) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

type Req_DHParamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce string `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	MessageId   int32  `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	B           int32  `protobuf:"varint,4,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *Req_DHParamsResponse) Reset() {
	*x = Req_DHParamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req_DHParamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req_DHParamsResponse) ProtoMessage() {}

func (x *Req_DHParamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req_DHParamsResponse.ProtoReflect.Descriptor instead.
func (*Req_DHParamsResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *Req_DHParamsResponse) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *Req_DHParamsResponse) GetServerNonce() string {
	if x != nil {
		return x.ServerNonce
	}
	return ""
}

func (x *Req_DHParamsResponse) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *Req_DHParamsResponse) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x43, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x5f, 0x70, 0x71, 0x5f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x5f,
	0x70, 0x71, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e,
	0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x70, 0x12, 0x0c, 0x0a, 0x01, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x67, 0x22,
	0x7b, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x5f, 0x44, 0x48, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x0c,
	0x0a, 0x01, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x22, 0x7e, 0x0a, 0x16,
	0x52, 0x65, 0x71, 0x5f, 0x44, 0x48, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x0c,
	0x0a, 0x01, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x32, 0x89, 0x01, 0x0a,
	0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x70, 0x71, 0x12,
	0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x5f, 0x70, 0x71, 0x5f, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x5f, 0x70,
	0x71, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0d,
	0x72, 0x65, 0x71, 0x5f, 0x44, 0x48, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x19, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x5f, 0x44, 0x48, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x52, 0x65, 0x71, 0x5f, 0x44, 0x48, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x50, 0x01, 0x5a, 0x0c, 0x77, 0x65,
	0x62, 0x2f, 0x68, 0x77, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auth_proto_goTypes = []interface{}{
	(*ReqPqInput)(nil),           // 0: auth.Req_pq_input
	(*ReqPqResponse)(nil),        // 1: auth.Req_pq_response
	(*Req_DHParamsInput)(nil),    // 2: auth.Req_DH_params_input
	(*Req_DHParamsResponse)(nil), // 3: auth.Req_DH_params_response
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: auth.Auth.req_pq:input_type -> auth.Req_pq_input
	2, // 1: auth.Auth.req_DH_params:input_type -> auth.Req_DH_params_input
	1, // 2: auth.Auth.req_pq:output_type -> auth.Req_pq_response
	3, // 3: auth.Auth.req_DH_params:output_type -> auth.Req_DH_params_response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqPqInput); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqPqResponse); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req_DHParamsInput); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req_DHParamsResponse); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
