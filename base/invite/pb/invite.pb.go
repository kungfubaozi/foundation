// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base/invite/pb/invite.proto

package fs_base_invite

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	pb "zskparker.com/foundation/base/pb"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type MoveToUserRequest struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	InviteId             string   `protobuf:"bytes,2,opt,name=inviteId,proto3" json:"inviteId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoveToUserRequest) Reset()         { *m = MoveToUserRequest{} }
func (m *MoveToUserRequest) String() string { return proto.CompactTextString(m) }
func (*MoveToUserRequest) ProtoMessage()    {}
func (*MoveToUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bf72a1accb58b55, []int{0}
}

func (m *MoveToUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveToUserRequest.Unmarshal(m, b)
}
func (m *MoveToUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveToUserRequest.Marshal(b, m, deterministic)
}
func (m *MoveToUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveToUserRequest.Merge(m, src)
}
func (m *MoveToUserRequest) XXX_Size() int {
	return xxx_messageInfo_MoveToUserRequest.Size(m)
}
func (m *MoveToUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveToUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MoveToUserRequest proto.InternalMessageInfo

func (m *MoveToUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *MoveToUserRequest) GetInviteId() string {
	if m != nil {
		return m.InviteId
	}
	return ""
}

type GetResponse struct {
	State                *pb.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	InviteId             string    `protobuf:"bytes,2,opt,name=inviteId,proto3" json:"inviteId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bf72a1accb58b55, []int{1}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetState() *pb.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *GetResponse) GetInviteId() string {
	if m != nil {
		return m.InviteId
	}
	return ""
}

type GetRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bf72a1accb58b55, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *GetRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type AddRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Enterprise           string   `protobuf:"bytes,2,opt,name=enterprise,proto3" json:"enterprise,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	RealName             string   `protobuf:"bytes,4,opt,name=realName,proto3" json:"realName,omitempty"`
	Meta                 *pb.Meta `protobuf:"bytes,5,opt,name=meta,proto3" json:"meta,omitempty"`
	Level                int64    `protobuf:"varint,6,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bf72a1accb58b55, []int{3}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *AddRequest) GetEnterprise() string {
	if m != nil {
		return m.Enterprise
	}
	return ""
}

func (m *AddRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddRequest) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *AddRequest) GetMeta() *pb.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AddRequest) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*MoveToUserRequest)(nil), "fs.base.invite.MoveToUserRequest")
	proto.RegisterType((*GetResponse)(nil), "fs.base.invite.GetResponse")
	proto.RegisterType((*GetRequest)(nil), "fs.base.invite.GetRequest")
	proto.RegisterType((*AddRequest)(nil), "fs.base.invite.AddRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InviteClient is the client API for Invite service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InviteClient interface {
	// 添加
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*pb.Response, error)
	// 获取
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// 移动到用户组里
	MoveToUser(ctx context.Context, in *MoveToUserRequest, opts ...grpc.CallOption) (*pb.Response, error)
}

type inviteClient struct {
	cc *grpc.ClientConn
}

func NewInviteClient(cc *grpc.ClientConn) InviteClient {
	return &inviteClient{cc}
}

func (c *inviteClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*pb.Response, error) {
	out := new(pb.Response)
	err := c.cc.Invoke(ctx, "/fs.base.invite.Invite/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/fs.base.invite.Invite/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteClient) MoveToUser(ctx context.Context, in *MoveToUserRequest, opts ...grpc.CallOption) (*pb.Response, error) {
	out := new(pb.Response)
	err := c.cc.Invoke(ctx, "/fs.base.invite.Invite/MoveToUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InviteServer is the server API for Invite service.
type InviteServer interface {
	// 添加
	Add(context.Context, *AddRequest) (*pb.Response, error)
	// 获取
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// 移动到用户组里
	MoveToUser(context.Context, *MoveToUserRequest) (*pb.Response, error)
}

func RegisterInviteServer(s *grpc.Server, srv InviteServer) {
	s.RegisterService(&_Invite_serviceDesc, srv)
}

func _Invite_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.invite.Invite/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invite_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.invite.Invite/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invite_MoveToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveToUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteServer).MoveToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.invite.Invite/MoveToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteServer).MoveToUser(ctx, req.(*MoveToUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Invite_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fs.base.invite.Invite",
	HandlerType: (*InviteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Invite_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Invite_Get_Handler,
		},
		{
			MethodName: "MoveToUser",
			Handler:    _Invite_MoveToUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/invite/pb/invite.proto",
}

func init() { proto.RegisterFile("base/invite/pb/invite.proto", fileDescriptor_0bf72a1accb58b55) }

var fileDescriptor_0bf72a1accb58b55 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x69, 0xab, 0x4e, 0xb1, 0xd0, 0xc5, 0x43, 0x48, 0x41, 0xda, 0xe2, 0xa1, 0x20,
	0x24, 0x50, 0x3d, 0x79, 0x10, 0xea, 0xa5, 0x14, 0xa9, 0x42, 0xd4, 0x07, 0xd8, 0x64, 0xa7, 0x10,
	0xda, 0x66, 0xd7, 0xdd, 0x4d, 0x05, 0x5f, 0xcc, 0xbb, 0x4f, 0x26, 0xd9, 0x4d, 0xd2, 0xd6, 0x16,
	0x4f, 0xdd, 0x7f, 0xbe, 0x99, 0xe9, 0x3f, 0x93, 0x81, 0x5e, 0x4c, 0x15, 0x86, 0x69, 0xb6, 0x49,
	0x35, 0x86, 0x22, 0x2e, 0x5f, 0x81, 0x90, 0x5c, 0x73, 0xd2, 0x59, 0xa8, 0xa0, 0xe0, 0x81, 0x8d,
	0xfa, 0x37, 0x5f, 0x6a, 0x29, 0xa8, 0x5c, 0xa2, 0x0c, 0x12, 0xbe, 0x0e, 0x17, 0x3c, 0xcf, 0x18,
	0xd5, 0x29, 0xcf, 0x42, 0xd3, 0x45, 0xc4, 0xe6, 0xd7, 0x16, 0x0f, 0x9f, 0xa0, 0x3b, 0xe7, 0x1b,
	0x7c, 0xe3, 0xef, 0x0a, 0x65, 0x84, 0x1f, 0x39, 0x2a, 0x4d, 0x7c, 0x38, 0x13, 0x54, 0xa9, 0x4f,
	0x2e, 0x99, 0xe7, 0xf4, 0x9d, 0xd1, 0x79, 0x54, 0xeb, 0x82, 0xd9, 0xff, 0x99, 0x31, 0xef, 0xc4,
	0xb2, 0x4a, 0x0f, 0x5f, 0xa0, 0x3d, 0x45, 0x1d, 0xa1, 0x12, 0x3c, 0x53, 0x48, 0xae, 0xa1, 0xa9,
	0x34, 0xd5, 0x68, 0x7a, 0xb4, 0xc7, 0x9d, 0xa0, 0x32, 0xfa, 0x5a, 0x44, 0x23, 0x0b, 0xff, 0x6d,
	0x78, 0x0f, 0x60, 0x1a, 0x5a, 0x5b, 0x04, 0x1a, 0x09, 0x67, 0x58, 0x5a, 0x32, 0x6f, 0xe2, 0xc1,
	0x29, 0x4d, 0x12, 0x9e, 0x67, 0xba, 0x2c, 0xae, 0xe4, 0xf0, 0xdb, 0x01, 0x98, 0x30, 0x56, 0x15,
	0xef, 0x24, 0x3a, 0x7b, 0x89, 0xe4, 0x0a, 0x00, 0x33, 0x8d, 0x52, 0xc8, 0x54, 0x61, 0xd9, 0x65,
	0x27, 0x52, 0x18, 0xcc, 0x15, 0xca, 0x8c, 0xae, 0xd1, 0x73, 0xad, 0xc1, 0x4a, 0x17, 0x4c, 0x22,
	0x5d, 0x3d, 0x17, 0xac, 0x61, 0x59, 0xa5, 0xc9, 0x00, 0x1a, 0x6b, 0xd4, 0xd4, 0x6b, 0x9a, 0xe9,
	0x2f, 0xea, 0xe9, 0xe7, 0xa8, 0x69, 0x64, 0x10, 0xb9, 0x84, 0xe6, 0x0a, 0x37, 0xb8, 0xf2, 0x5a,
	0x7d, 0x67, 0xe4, 0x46, 0x56, 0x8c, 0x7f, 0x1c, 0x68, 0xcd, 0xcc, 0x0a, 0xc8, 0x1d, 0xb8, 0x13,
	0xc6, 0x88, 0x1f, 0xec, 0x7f, 0xe3, 0x60, 0x3b, 0x98, 0xdf, 0xad, 0x59, 0xbd, 0xf8, 0x07, 0x70,
	0xa7, 0xa8, 0x0f, 0xab, 0xb6, 0xbb, 0xf4, 0x7b, 0x47, 0x59, 0x59, 0xff, 0x08, 0xb0, 0x3d, 0x0a,
	0x32, 0xf8, 0x9b, 0x7a, 0x70, 0x30, 0x47, 0x3c, 0xc4, 0x2d, 0x73, 0x5f, 0xb7, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x0b, 0x68, 0x31, 0x20, 0xbb, 0x02, 0x00, 0x00,
}
