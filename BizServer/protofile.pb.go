// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protofile.proto

package biz

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetUserInput struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AuthKey              string   `protobuf:"bytes,2,opt,name=auth_key,json=authKey,proto3" json:"auth_key,omitempty"`
	MessageId            int32    `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInput) Reset()         { *m = GetUserInput{} }
func (m *GetUserInput) String() string { return proto.CompactTextString(m) }
func (*GetUserInput) ProtoMessage()    {}
func (*GetUserInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_b28b201cc727c47d, []int{0}
}

func (m *GetUserInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInput.Unmarshal(m, b)
}
func (m *GetUserInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInput.Marshal(b, m, deterministic)
}
func (m *GetUserInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInput.Merge(m, src)
}
func (m *GetUserInput) XXX_Size() int {
	return xxx_messageInfo_GetUserInput.Size(m)
}
func (m *GetUserInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInput proto.InternalMessageInfo

func (m *GetUserInput) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetUserInput) GetAuthKey() string {
	if m != nil {
		return m.AuthKey
	}
	return ""
}

func (m *GetUserInput) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type GetUserOutput struct {
	User                 []*User  `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserOutput) Reset()         { *m = GetUserOutput{} }
func (m *GetUserOutput) String() string { return proto.CompactTextString(m) }
func (*GetUserOutput) ProtoMessage()    {}
func (*GetUserOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_b28b201cc727c47d, []int{1}
}

func (m *GetUserOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserOutput.Unmarshal(m, b)
}
func (m *GetUserOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserOutput.Marshal(b, m, deterministic)
}
func (m *GetUserOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserOutput.Merge(m, src)
}
func (m *GetUserOutput) XXX_Size() int {
	return xxx_messageInfo_GetUserOutput.Size(m)
}
func (m *GetUserOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserOutput proto.InternalMessageInfo

func (m *GetUserOutput) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Family               string   `protobuf:"bytes,2,opt,name=family,proto3" json:"family,omitempty"`
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Age                  int32    `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Sex                  string   `protobuf:"bytes,5,opt,name=sex,proto3" json:"sex,omitempty"`
	Timestamp            int64    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_b28b201cc727c47d, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetFamily() string {
	if m != nil {
		return m.Family
	}
	return ""
}

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

func (m *User) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*GetUserInput)(nil), "biz.get_user_input")
	proto.RegisterType((*GetUserOutput)(nil), "biz.get_user_output")
	proto.RegisterType((*User)(nil), "biz.User")
}

func init() {
	proto.RegisterFile("protofile.proto", fileDescriptor_b28b201cc727c47d)
}

var fileDescriptor_b28b201cc727c47d = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4f, 0xc2, 0x40,
	0x14, 0xc6, 0x3d, 0x5a, 0x8a, 0x7d, 0x24, 0x60, 0x9e, 0x46, 0xab, 0x81, 0xa4, 0xe9, 0xd4, 0x09,
	0x14, 0x17, 0x67, 0x37, 0xc2, 0x62, 0x9a, 0xb8, 0xb8, 0x5c, 0xae, 0xf4, 0x01, 0xa7, 0xb4, 0x45,
	0xee, 0x1a, 0x84, 0xd9, 0xd1, 0x3f, 0xda, 0xdc, 0x01, 0x35, 0x8c, 0x6e, 0xdf, 0xf7, 0x4b, 0xbe,
	0x7b, 0xdf, 0x7b, 0x07, 0xdd, 0xd5, 0xba, 0xd4, 0xe5, 0x4c, 0x2e, 0x69, 0x60, 0x15, 0x3a, 0xa9,
	0xdc, 0x45, 0x53, 0xe8, 0xcc, 0x49, 0xf3, 0x4a, 0xd1, 0x9a, 0xcb, 0x62, 0x55, 0x69, 0xbc, 0x81,
	0xd6, 0xde, 0x65, 0x01, 0x0b, 0x59, 0xdc, 0x4c, 0x3c, 0x63, 0xc7, 0x19, 0xde, 0xc2, 0xb9, 0xa8,
	0xf4, 0x82, 0x7f, 0xd0, 0x36, 0x68, 0x84, 0x2c, 0xf6, 0x93, 0x96, 0xf1, 0x13, 0xda, 0x62, 0x1f,
	0x20, 0x27, 0xa5, 0xc4, 0x9c, 0x4c, 0xcc, 0xb1, 0x31, 0xff, 0x40, 0xc6, 0x59, 0x74, 0x0f, 0xdd,
	0x7a, 0x48, 0x59, 0x69, 0x33, 0xa5, 0x0f, 0xae, 0xb1, 0x01, 0x0b, 0x9d, 0xb8, 0x3d, 0xf2, 0x07,
	0xa9, 0xdc, 0x0d, 0x5e, 0x15, 0xad, 0x13, 0x8b, 0xa3, 0x6f, 0x06, 0xae, 0xb1, 0x88, 0xe0, 0x16,
	0x22, 0x27, 0x5b, 0xc5, 0x4f, 0xac, 0xc6, 0x6b, 0xf0, 0x66, 0x22, 0x97, 0xcb, 0x63, 0x8d, 0x83,
	0xc3, 0x0e, 0x34, 0xea, 0xe9, 0x0d, 0x99, 0xe1, 0x05, 0x38, 0x62, 0x4e, 0x81, 0x6b, 0x81, 0x91,
	0x86, 0x28, 0xfa, 0x0a, 0x9a, 0x36, 0x66, 0x24, 0xf6, 0xc0, 0xd7, 0x32, 0x27, 0xa5, 0x45, 0xbe,
	0x0a, 0xbc, 0x90, 0xc5, 0x4e, 0xf2, 0x07, 0x46, 0x3f, 0x0c, 0xcc, 0x95, 0xf0, 0x09, 0xfc, 0xe3,
	0x02, 0x0a, 0x2f, 0x6d, 0xd9, 0xd3, 0xab, 0xdd, 0x5d, 0x9d, 0xc2, 0xfd, 0x96, 0xd1, 0x19, 0x4e,
	0xa0, 0x57, 0x27, 0xf9, 0x46, 0xea, 0x05, 0x57, 0x9f, 0x4b, 0x2e, 0x8b, 0x77, 0x9a, 0x6a, 0x59,
	0x16, 0xff, 0x7a, 0xec, 0xb9, 0xfb, 0xc2, 0xde, 0xda, 0x1b, 0x4a, 0x87, 0x8b, 0xcd, 0xc3, 0x30,
	0x95, 0xbb, 0xd4, 0xb3, 0x3f, 0xf9, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x78, 0x03, 0x62, 0x1f,
	0xdc, 0x01, 0x00, 0x00,
}
