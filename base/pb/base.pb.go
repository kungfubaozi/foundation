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
	Longitude            string   `protobuf:"bytes,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             string   `protobuf:"bytes,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
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

type TsMessage struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TsMessage) Reset()         { *m = TsMessage{} }
func (m *TsMessage) String() string { return proto.CompactTextString(m) }
func (*TsMessage) ProtoMessage()    {}
func (*TsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d99073f0a37c125, []int{5}
}

func (m *TsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TsMessage.Unmarshal(m, b)
}
func (m *TsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TsMessage.Marshal(b, m, deterministic)
}
func (m *TsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TsMessage.Merge(m, src)
}
func (m *TsMessage) XXX_Size() int {
	return xxx_messageInfo_TsMessage.Size(m)
}
func (m *TsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TsMessage proto.InternalMessageInfo

func (m *TsMessage) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TsMessage) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TsMessage) GetContent() string {
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
	proto.RegisterType((*TsMessage)(nil), "fs.base.TsMessage")
}

func init() { proto.RegisterFile("base/pb/base.proto", fileDescriptor_0d99073f0a37c125) }

var fileDescriptor_0d99073f0a37c125 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcf, 0x8a, 0xd4, 0x40,
	0x10, 0xc6, 0x99, 0x64, 0x32, 0x93, 0x94, 0xb0, 0x48, 0xb3, 0x48, 0x23, 0x1e, 0x24, 0x78, 0xf0,
	0x94, 0x15, 0x7d, 0x02, 0x0f, 0x1e, 0x82, 0xec, 0x25, 0x8a, 0x07, 0x6f, 0x3d, 0x49, 0xcd, 0x10,
	0xf2, 0xa7, 0x9b, 0x74, 0x65, 0x61, 0xde, 0xd6, 0x47, 0x91, 0xaa, 0xee, 0x89, 0x83, 0xa7, 0xd4,
	0x57, 0x05, 0x55, 0xdf, 0xf7, 0x4b, 0x83, 0x3a, 0x19, 0x8f, 0x4f, 0xee, 0xf4, 0xc4, 0xdf, 0xca,
	0x2d, 0x96, 0xac, 0x3a, 0x9e, 0x7d, 0xc5, 0xb2, 0x1c, 0x61, 0xff, 0x8c, 0x64, 0x94, 0x82, 0xfd,
	0xd9, 0xb4, 0xa8, 0x77, 0xef, 0x77, 0x1f, 0x8b, 0x46, 0x6a, 0xf5, 0x06, 0x0e, 0x1d, 0xbe, 0xf4,
	0x2d, 0xea, 0x44, 0xba, 0x51, 0xa9, 0x77, 0x50, 0x8c, 0x76, 0xbe, 0xf4, 0xb4, 0x76, 0xa8, 0xf7,
	0x32, 0xfa, 0xd7, 0x50, 0x6f, 0x21, 0x1f, 0x0d, 0x85, 0x61, 0x26, 0xc3, 0x4d, 0x97, 0xdf, 0x20,
	0xfb, 0x41, 0x86, 0x90, 0xcf, 0xb5, 0xb6, 0x0b, 0xe7, 0xd2, 0x46, 0x6a, 0xf5, 0x00, 0x89, 0x1d,
	0xe4, 0x54, 0xde, 0x24, 0x76, 0x50, 0x1a, 0x8e, 0x13, 0x7a, 0x6f, 0x2e, 0xa8, 0x53, 0xd9, 0x73,
	0x93, 0xe5, 0x27, 0xc8, 0x1b, 0xf4, 0xce, 0xce, 0x1e, 0xd5, 0x07, 0xc8, 0x3c, 0xaf, 0x94, 0x55,
	0xaf, 0x3e, 0x3f, 0x54, 0x31, 0x59, 0x25, 0x87, 0x9a, 0x30, 0x2c, 0xff, 0xec, 0x20, 0xe7, 0x9c,
	0x9d, 0x21, 0xc3, 0x0e, 0xdb, 0xb1, 0xc7, 0x99, 0xea, 0x2e, 0xe6, 0xdd, 0xb4, 0x7a, 0x84, 0xcc,
	0x38, 0x57, 0x77, 0x31, 0x72, 0x10, 0x4c, 0x62, 0xf5, 0xb8, 0xd4, 0x5d, 0x74, 0x12, 0x15, 0x5b,
	0xee, 0x5d, 0x44, 0x90, 0xf4, 0x6e, 0xa3, 0x98, 0xdd, 0x51, 0x7c, 0x84, 0x8c, 0xec, 0x80, 0xb3,
	0x3e, 0x84, 0x8d, 0x22, 0xee, 0xd8, 0x1e, 0xff, 0x67, 0xcb, 0xbb, 0xbf, 0x5e, 0x70, 0x26, 0x9d,
	0x07, 0xb6, 0x5b, 0x83, 0x9d, 0xbb, 0xd1, 0xd0, 0xd9, 0x2e, 0x93, 0x2e, 0x04, 0xdd, 0xa6, 0xcb,
	0xdf, 0x90, 0x7f, 0xc7, 0xeb, 0x2f, 0x33, 0xae, 0xa8, 0x5e, 0x43, 0x3a, 0xe0, 0x35, 0x86, 0xe3,
	0x92, 0x5d, 0xbc, 0xf0, 0xe8, 0x96, 0x4b, 0x04, 0xfb, 0xa5, 0xab, 0x0b, 0x7c, 0xd3, 0x46, 0x6a,
	0xee, 0x4d, 0xab, 0x27, 0x49, 0x95, 0x37, 0x52, 0x97, 0x35, 0x14, 0x3f, 0xfd, 0x73, 0xa0, 0x2f,
	0x21, 0x17, 0x3b, 0x6d, 0x4f, 0x65, 0xb1, 0x13, 0x83, 0x20, 0x1b, 0x77, 0x27, 0x64, 0xf9, 0xdf,
	0xb5, 0x76, 0x26, 0x0e, 0x11, 0xff, 0x5d, 0x94, 0xa7, 0x83, 0x3c, 0xc0, 0x2f, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x18, 0x3a, 0xc4, 0xf2, 0x96, 0x02, 0x00, 0x00,
}
