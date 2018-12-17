// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base/pb/base.proto

package fs_base

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Meta struct {
	Face                 string   `protobuf:"bytes,1,opt,name=face,proto3" json:"face,omitempty"`
	Device               string   `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Meta) Reset()         { *m = Meta{} }
func (m *Meta) String() string { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()    {}
func (*Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d99073f0a37c125, []int{0}
}

func (m *Meta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meta.Unmarshal(m, b)
}
func (m *Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meta.Marshal(b, m, deterministic)
}
func (m *Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meta.Merge(m, src)
}
func (m *Meta) XXX_Size() int {
	return xxx_messageInfo_Meta.Size(m)
}
func (m *Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_Meta proto.InternalMessageInfo

func (m *Meta) GetFace() string {
	if m != nil {
		return m.Face
	}
	return ""
}

func (m *Meta) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

type State struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Ok                   bool     `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d99073f0a37c125, []int{1}
}

func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *State) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *State) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Response struct {
	State                *State   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d99073f0a37c125, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetState() *State {
	if m != nil {
		return m.State
	}
	return nil
}

type Metadata struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	AppId                string   `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId,omitempty"`
	UserId               string   `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Ip                   string   `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	Face                 string   `protobuf:"bytes,5,opt,name=face,proto3" json:"face,omitempty"`
	Token                string   `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	Device               string   `protobuf:"bytes,7,opt,name=device,proto3" json:"device,omitempty"`
	UserAgent            string   `protobuf:"bytes,8,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
	Platform             int64    `protobuf:"varint,9,opt,name=platform,proto3" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d99073f0a37c125, []int{3}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Metadata) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Metadata) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Metadata) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Metadata) GetFace() string {
	if m != nil {
		return m.Face
	}
	return ""
}

func (m *Metadata) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Metadata) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *Metadata) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *Metadata) GetPlatform() int64 {
	if m != nil {
		return m.Platform
	}
	return 0
}

type KeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Type                 int64    `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d99073f0a37c125, []int{4}
}

func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (m *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(m, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *KeyValue) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*Meta)(nil), "fs.base.Meta")
	proto.RegisterType((*State)(nil), "fs.base.State")
	proto.RegisterType((*Response)(nil), "fs.base.Response")
	proto.RegisterType((*Metadata)(nil), "fs.base.Metadata")
	proto.RegisterType((*KeyValue)(nil), "fs.base.KeyValue")
}

func init() { proto.RegisterFile("base/pb/base.proto", fileDescriptor_0d99073f0a37c125) }

var fileDescriptor_0d99073f0a37c125 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xc1, 0x6a, 0xc2, 0x40,
	0x10, 0x86, 0x31, 0x31, 0x1a, 0xa7, 0x20, 0x65, 0x90, 0xb2, 0x94, 0x1e, 0x4a, 0xe8, 0xa1, 0xa7,
	0x58, 0xec, 0x13, 0xf4, 0xd0, 0x82, 0x94, 0x5e, 0x52, 0xe8, 0x7d, 0xcd, 0x8e, 0x12, 0xa2, 0xd9,
	0xc5, 0x5d, 0x05, 0xdf, 0xb6, 0x8f, 0x52, 0x66, 0xb2, 0x5a, 0x4f, 0x99, 0x7f, 0x86, 0xf9, 0xf3,
	0xcd, 0xbf, 0x80, 0x2b, 0xed, 0x69, 0xee, 0x56, 0x73, 0xfe, 0x96, 0x6e, 0x6f, 0x83, 0xc5, 0xf1,
	0xda, 0x97, 0x2c, 0x8b, 0x05, 0x0c, 0xbf, 0x28, 0x68, 0x44, 0x18, 0xae, 0x75, 0x4d, 0x6a, 0xf0,
	0x38, 0x78, 0x9e, 0x54, 0x52, 0xe3, 0x1d, 0x8c, 0x0c, 0x1d, 0x9b, 0x9a, 0x54, 0x22, 0xdd, 0xa8,
	0x8a, 0x77, 0xc8, 0xbe, 0x83, 0x0e, 0xc4, 0x4b, 0xb5, 0x35, 0xfd, 0x52, 0x5a, 0x49, 0x8d, 0x53,
	0x48, 0x6c, 0x2b, 0x0b, 0x79, 0x95, 0xd8, 0x16, 0x15, 0x8c, 0x77, 0xe4, 0xbd, 0xde, 0x90, 0x4a,
	0xc5, 0xe5, 0x2c, 0x8b, 0x17, 0xc8, 0x2b, 0xf2, 0xce, 0x76, 0x9e, 0xf0, 0x09, 0x32, 0xcf, 0x96,
	0x62, 0x75, 0xb3, 0x98, 0x96, 0x91, 0xaf, 0x94, 0x1f, 0x55, 0xfd, 0xb0, 0xf8, 0x1d, 0x40, 0xce,
	0xb4, 0x46, 0x07, 0x8d, 0xf7, 0x90, 0xd7, 0xdb, 0x86, 0xba, 0xb0, 0x34, 0x91, 0xfa, 0xa2, 0x71,
	0x06, 0x99, 0x76, 0x6e, 0x69, 0x22, 0x78, 0x2f, 0xf8, 0x9e, 0x83, 0xa7, 0xfd, 0xd2, 0x44, 0x92,
	0xa8, 0x18, 0xb9, 0x71, 0x6a, 0x28, 0xbd, 0xa4, 0x71, 0x97, 0x2c, 0xb2, 0xab, 0x2c, 0x66, 0x90,
	0x05, 0xdb, 0x52, 0xa7, 0x46, 0xbd, 0xa3, 0x88, 0xab, 0x84, 0xc6, 0xd7, 0x09, 0xe1, 0x03, 0x4c,
	0xd8, 0xfb, 0x6d, 0x43, 0x5d, 0x50, 0xb9, 0x8c, 0xfe, 0x1b, 0x4c, 0xee, 0xb6, 0x3a, 0xac, 0xed,
	0x7e, 0xa7, 0x26, 0x12, 0xdd, 0x45, 0x17, 0x1f, 0x90, 0x7f, 0xd2, 0xe9, 0x47, 0x6f, 0x0f, 0x84,
	0xb7, 0x90, 0xb6, 0x74, 0x8a, 0xc7, 0x71, 0xc9, 0x14, 0x47, 0x1e, 0x9d, 0xef, 0x12, 0xc1, 0xbc,
	0xe1, 0xe4, 0xfa, 0x7c, 0xd3, 0x4a, 0xea, 0xd5, 0x48, 0xde, 0xf9, 0xf5, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x78, 0xc1, 0x39, 0x91, 0xfd, 0x01, 0x00, 0x00,
}
