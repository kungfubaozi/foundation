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

// 作为API请求用，放在body里
type Meta struct {
	Face                 string   `protobuf:"bytes,1,opt,name=face,proto3" json:"face,omitempty"`
	Longitude            string   `protobuf:"bytes,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             string   `protobuf:"bytes,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Validate             string   `protobuf:"bytes,4,opt,name=validate,proto3" json:"validate,omitempty"`
	Id                   string   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
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

func (m *Meta) GetLongitude() string {
	if m != nil {
		return m.Longitude
	}
	return ""
}

func (m *Meta) GetLatitude() string {
	if m != nil {
		return m.Latitude
	}
	return ""
}

func (m *Meta) GetValidate() string {
	if m != nil {
		return m.Validate
	}
	return ""
}

func (m *Meta) GetId() string {
	if m != nil {
		return m.Id
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

// 作为服务间传输用
type Metadata struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	ProjectId            string   `protobuf:"bytes,2,opt,name=projectId,proto3" json:"projectId,omitempty"`
	UserId               string   `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Ip                   string   `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	Token                string   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	Device               string   `protobuf:"bytes,6,opt,name=device,proto3" json:"device,omitempty"`
	UserAgent            string   `protobuf:"bytes,7,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
	Platform             int64    `protobuf:"varint,8,opt,name=platform,proto3" json:"platform,omitempty"`
	Api                  string   `protobuf:"bytes,9,opt,name=api,proto3" json:"api,omitempty"`
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

func (m *Metadata) GetProjectId() string {
	if m != nil {
		return m.ProjectId
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

func (m *Metadata) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

type KeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Type                 int64    `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Must                 bool     `protobuf:"varint,4,opt,name=must,proto3" json:"must,omitempty"`
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

func (m *KeyValue) GetMust() bool {
	if m != nil {
		return m.Must
	}
	return false
}

type DirectMessage struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DirectMessage) Reset()         { *m = DirectMessage{} }
func (m *DirectMessage) String() string { return proto.CompactTextString(m) }
func (*DirectMessage) ProtoMessage()    {}
func (*DirectMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d99073f0a37c125, []int{5}
}

func (m *DirectMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectMessage.Unmarshal(m, b)
}
func (m *DirectMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectMessage.Marshal(b, m, deterministic)
}
func (m *DirectMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectMessage.Merge(m, src)
}
func (m *DirectMessage) XXX_Size() int {
	return xxx_messageInfo_DirectMessage.Size(m)
}
func (m *DirectMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DirectMessage proto.InternalMessageInfo

func (m *DirectMessage) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *DirectMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*Meta)(nil), "fs.base.Meta")
	proto.RegisterType((*State)(nil), "fs.base.State")
	proto.RegisterType((*Response)(nil), "fs.base.Response")
	proto.RegisterType((*Metadata)(nil), "fs.base.Metadata")
	proto.RegisterType((*KeyValue)(nil), "fs.base.KeyValue")
	proto.RegisterType((*DirectMessage)(nil), "fs.base.DirectMessage")
}

func init() { proto.RegisterFile("base/pb/base.proto", fileDescriptor_0d99073f0a37c125) }

var fileDescriptor_0d99073f0a37c125 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x52, 0xcf, 0xcb, 0xd3, 0x40,
	0x10, 0x25, 0xbf, 0xda, 0x74, 0xc4, 0x0f, 0x59, 0x44, 0x16, 0xf1, 0x20, 0xc1, 0x83, 0xa7, 0x7c,
	0xa2, 0x27, 0x8f, 0x82, 0x1e, 0x3e, 0xa4, 0x97, 0x08, 0x1e, 0xbc, 0x6d, 0xb3, 0xd3, 0xb2, 0x26,
	0xcd, 0x2e, 0xd9, 0x49, 0xa1, 0x37, 0xff, 0x5a, 0xff, 0x0e, 0x99, 0xdd, 0x4d, 0x3c, 0xf5, 0xbd,
	0x79, 0xdd, 0x79, 0xf3, 0x66, 0x02, 0xe2, 0xa4, 0x3c, 0x3e, 0xba, 0xd3, 0x23, 0xff, 0xb6, 0x6e,
	0xb6, 0x64, 0xc5, 0xfe, 0xec, 0x5b, 0xa6, 0xcd, 0x9f, 0x0c, 0xca, 0x23, 0x92, 0x12, 0x02, 0xca,
	0xb3, 0xea, 0x51, 0x66, 0x6f, 0xb3, 0xf7, 0x87, 0x2e, 0x60, 0xf1, 0x06, 0x0e, 0xa3, 0x9d, 0x2e,
	0x86, 0x16, 0x8d, 0x32, 0x0f, 0xc2, 0xff, 0x82, 0x78, 0x0d, 0xf5, 0xa8, 0x28, 0x8a, 0x45, 0x10,
	0x37, 0xce, 0xda, 0x4d, 0x8d, 0x46, 0x2b, 0x42, 0x59, 0x46, 0x6d, 0xe5, 0xe2, 0x01, 0x72, 0xa3,
	0x65, 0x15, 0xaa, 0xb9, 0xd1, 0xcd, 0x37, 0xa8, 0x7e, 0x10, 0x0b, 0x02, 0xca, 0xde, 0xea, 0x38,
	0x42, 0xd1, 0x05, 0xcc, 0x7f, 0xb6, 0x43, 0xf0, 0xae, 0xbb, 0xdc, 0x0e, 0x42, 0xc2, 0xfe, 0x8a,
	0xde, 0xab, 0xcb, 0xea, 0xb9, 0xd2, 0xe6, 0x03, 0xd4, 0x1d, 0x7a, 0x67, 0x27, 0x8f, 0xe2, 0x1d,
	0x54, 0x9e, 0x5b, 0x86, 0x56, 0xcf, 0x3e, 0x3e, 0xb4, 0x29, 0x6e, 0x1b, 0x8c, 0xba, 0x28, 0x36,
	0x7f, 0x33, 0xa8, 0x39, 0xbb, 0x56, 0xa4, 0x78, 0xe2, 0x7e, 0x34, 0x38, 0xd1, 0x93, 0x4e, 0x3b,
	0xd8, 0x38, 0xef, 0xc1, 0xcd, 0xf6, 0x37, 0xf6, 0x2c, 0xa6, 0x3d, 0x6c, 0x05, 0xf1, 0x0a, 0x76,
	0x8b, 0xc7, 0xf9, 0x49, 0xa7, 0x89, 0x12, 0x0b, 0x39, 0x5d, 0x4a, 0x9f, 0x1b, 0x27, 0x5e, 0x42,
	0x45, 0x76, 0xc0, 0x29, 0x45, 0x8f, 0x84, 0x5f, 0x6b, 0xbc, 0x99, 0x1e, 0xe5, 0x2e, 0xbe, 0x8e,
	0x8c, 0x3d, 0xb9, 0xcf, 0x97, 0x0b, 0x4e, 0x24, 0xf7, 0xd1, 0x73, 0x2b, 0xf0, 0xb4, 0x6e, 0x54,
	0x74, 0xb6, 0xf3, 0x55, 0xd6, 0x61, 0x5d, 0x1b, 0x17, 0x2f, 0xa0, 0x50, 0xce, 0xc8, 0x43, 0x78,
	0xc3, 0xb0, 0xf9, 0x05, 0xf5, 0x77, 0xbc, 0xff, 0x54, 0xe3, 0x82, 0xac, 0x0e, 0x78, 0x4f, 0x11,
	0x19, 0xf2, 0x5c, 0x37, 0x96, 0x52, 0xb2, 0x48, 0xf8, 0x18, 0x74, 0x77, 0x71, 0xcb, 0x45, 0x17,
	0x30, 0xd7, 0xae, 0x8b, 0xa7, 0x90, 0xa9, 0xee, 0x02, 0x6e, 0x3e, 0xc3, 0xf3, 0xaf, 0x66, 0xc6,
	0x9e, 0x8e, 0xf1, 0x0e, 0x1c, 0x9b, 0x6c, 0xea, 0x9f, 0x93, 0xe5, 0x8b, 0xf5, 0x76, 0x22, 0x8e,
	0x11, 0x0d, 0x56, 0x7a, 0xda, 0x85, 0x6f, 0xf1, 0xd3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2,
	0xd7, 0xf8, 0xa5, 0xa1, 0x02, 0x00, 0x00,
}
