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

type GetInvitesRequest struct {
	Meta                 *pb.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int64    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInvitesRequest) Reset()         { *m = GetInvitesRequest{} }
func (m *GetInvitesRequest) String() string { return proto.CompactTextString(m) }
func (*GetInvitesRequest) ProtoMessage()    {}
func (*GetInvitesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bf72a1accb58b55, []int{0}
}

func (m *GetInvitesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInvitesRequest.Unmarshal(m, b)
}
func (m *GetInvitesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInvitesRequest.Marshal(b, m, deterministic)
}
func (m *GetInvitesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInvitesRequest.Merge(m, src)
}
func (m *GetInvitesRequest) XXX_Size() int {
	return xxx_messageInfo_GetInvitesRequest.Size(m)
}
func (m *GetInvitesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInvitesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInvitesRequest proto.InternalMessageInfo

func (m *GetInvitesRequest) GetMeta() *pb.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetInvitesRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetInvitesRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type GetInvitesResponse struct {
	State                *pb.State     `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Info                 []*InviteInfo `protobuf:"bytes,2,rep,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetInvitesResponse) Reset()         { *m = GetInvitesResponse{} }
func (m *GetInvitesResponse) String() string { return proto.CompactTextString(m) }
func (*GetInvitesResponse) ProtoMessage()    {}
func (*GetInvitesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bf72a1accb58b55, []int{1}
}

func (m *GetInvitesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInvitesResponse.Unmarshal(m, b)
}
func (m *GetInvitesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInvitesResponse.Marshal(b, m, deterministic)
}
func (m *GetInvitesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInvitesResponse.Merge(m, src)
}
func (m *GetInvitesResponse) XXX_Size() int {
	return xxx_messageInfo_GetInvitesResponse.Size(m)
}
func (m *GetInvitesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInvitesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInvitesResponse proto.InternalMessageInfo

func (m *GetInvitesResponse) GetState() *pb.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *GetInvitesResponse) GetInfo() []*InviteInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type InviteInfo struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	CreateAt             string   `protobuf:"bytes,2,opt,name=createAt,proto3" json:"createAt,omitempty"`
	OkAt                 int64    `protobuf:"varint,3,opt,name=okAt,proto3" json:"okAt,omitempty"`
	OperateUserId        string   `protobuf:"bytes,4,opt,name=operateUserId,proto3" json:"operateUserId,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Enterprise           string   `protobuf:"bytes,6,opt,name=enterprise,proto3" json:"enterprise,omitempty"`
	Username             string   `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
	RealName             string   `protobuf:"bytes,8,opt,name=realName,proto3" json:"realName,omitempty"`
	Level                int64    `protobuf:"varint,9,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InviteInfo) Reset()         { *m = InviteInfo{} }
func (m *InviteInfo) String() string { return proto.CompactTextString(m) }
func (*InviteInfo) ProtoMessage()    {}
func (*InviteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bf72a1accb58b55, []int{2}
}

func (m *InviteInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InviteInfo.Unmarshal(m, b)
}
func (m *InviteInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InviteInfo.Marshal(b, m, deterministic)
}
func (m *InviteInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InviteInfo.Merge(m, src)
}
func (m *InviteInfo) XXX_Size() int {
	return xxx_messageInfo_InviteInfo.Size(m)
}
func (m *InviteInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_InviteInfo.DiscardUnknown(m)
}

var xxx_messageInfo_InviteInfo proto.InternalMessageInfo

func (m *InviteInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *InviteInfo) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *InviteInfo) GetOkAt() int64 {
	if m != nil {
		return m.OkAt
	}
	return 0
}

func (m *InviteInfo) GetOperateUserId() string {
	if m != nil {
		return m.OperateUserId
	}
	return ""
}

func (m *InviteInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *InviteInfo) GetEnterprise() string {
	if m != nil {
		return m.Enterprise
	}
	return ""
}

func (m *InviteInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *InviteInfo) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *InviteInfo) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

type UpdateRequest struct {
	InviteId             string   `protobuf:"bytes,1,opt,name=inviteId,proto3" json:"inviteId,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bf72a1accb58b55, []int{3}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetInviteId() string {
	if m != nil {
		return m.InviteId
	}
	return ""
}

func (m *UpdateRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type GetResponse struct {
	State                *pb.State   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	InviteId             string      `protobuf:"bytes,2,opt,name=inviteId,proto3" json:"inviteId,omitempty"`
	Detail               *InviteInfo `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bf72a1accb58b55, []int{4}
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

func (m *GetResponse) GetDetail() *InviteInfo {
	if m != nil {
		return m.Detail
	}
	return nil
}

type GetRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bf72a1accb58b55, []int{5}
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
	return fileDescriptor_0bf72a1accb58b55, []int{6}
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
	proto.RegisterType((*GetInvitesRequest)(nil), "fs.base.invite.GetInvitesRequest")
	proto.RegisterType((*GetInvitesResponse)(nil), "fs.base.invite.GetInvitesResponse")
	proto.RegisterType((*InviteInfo)(nil), "fs.base.invite.InviteInfo")
	proto.RegisterType((*UpdateRequest)(nil), "fs.base.invite.UpdateRequest")
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
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*pb.Response, error)
	GetInvites(ctx context.Context, in *GetInvitesRequest, opts ...grpc.CallOption) (*GetInvitesResponse, error)
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

func (c *inviteClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*pb.Response, error) {
	out := new(pb.Response)
	err := c.cc.Invoke(ctx, "/fs.base.invite.Invite/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteClient) GetInvites(ctx context.Context, in *GetInvitesRequest, opts ...grpc.CallOption) (*GetInvitesResponse, error) {
	out := new(GetInvitesResponse)
	err := c.cc.Invoke(ctx, "/fs.base.invite.Invite/GetInvites", in, out, opts...)
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
	Update(context.Context, *UpdateRequest) (*pb.Response, error)
	GetInvites(context.Context, *GetInvitesRequest) (*GetInvitesResponse, error)
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

func _Invite_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.invite.Invite/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invite_GetInvites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvitesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteServer).GetInvites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.invite.Invite/GetInvites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteServer).GetInvites(ctx, req.(*GetInvitesRequest))
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
			MethodName: "Update",
			Handler:    _Invite_Update_Handler,
		},
		{
			MethodName: "GetInvites",
			Handler:    _Invite_GetInvites_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/invite/pb/invite.proto",
}

func init() { proto.RegisterFile("base/invite/pb/invite.proto", fileDescriptor_0bf72a1accb58b55) }

var fileDescriptor_0bf72a1accb58b55 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x8e, 0xda, 0x3c,
	0x10, 0x56, 0x48, 0xc8, 0xc2, 0x20, 0x56, 0x5a, 0xeb, 0x3f, 0x44, 0x59, 0xfd, 0x15, 0x1b, 0xed,
	0x01, 0xa9, 0x52, 0x90, 0x68, 0x4f, 0x3d, 0x54, 0xe2, 0x50, 0x21, 0x0e, 0xed, 0x21, 0xab, 0xbd,
	0x56, 0x32, 0x78, 0x68, 0x53, 0xc0, 0x76, 0x6d, 0xb3, 0x87, 0xbd, 0xf6, 0x01, 0xfa, 0x38, 0x7d,
	0xb9, 0x1e, 0x2a, 0xdb, 0x49, 0x20, 0x2c, 0x45, 0xea, 0x89, 0x99, 0xf9, 0x66, 0xe6, 0x9b, 0x99,
	0x0f, 0x07, 0x6e, 0x97, 0x54, 0xe3, 0xa4, 0xe4, 0x4f, 0xa5, 0xc1, 0x89, 0x5c, 0x56, 0x56, 0x2e,
	0x95, 0x30, 0x82, 0x5c, 0xaf, 0x75, 0x6e, 0xf1, 0xdc, 0x47, 0xd3, 0xd7, 0xcf, 0x7a, 0x23, 0xa9,
	0xda, 0xa0, 0xca, 0x57, 0x62, 0x37, 0x59, 0x8b, 0x3d, 0x67, 0xd4, 0x94, 0x82, 0x4f, 0x5c, 0x17,
	0xb9, 0x74, 0xbf, 0xbe, 0x38, 0xfb, 0x0c, 0x37, 0x73, 0x34, 0x0b, 0x57, 0xa9, 0x0b, 0xfc, 0xbe,
	0x47, 0x6d, 0xc8, 0x1d, 0x44, 0x3b, 0x34, 0x34, 0x09, 0x46, 0xc1, 0x78, 0x30, 0x1d, 0xe6, 0x35,
	0xc1, 0x47, 0x34, 0xb4, 0x70, 0x10, 0x21, 0x10, 0x49, 0xfa, 0x05, 0x93, 0xce, 0x28, 0x18, 0x87,
	0x85, 0xb3, 0x6d, 0x4c, 0x97, 0xcf, 0x98, 0x84, 0x3e, 0x66, 0xed, 0xec, 0x1b, 0x90, 0xe3, 0xfe,
	0x5a, 0x0a, 0xae, 0x91, 0xdc, 0x43, 0x57, 0x1b, 0x6a, 0xb0, 0x62, 0xb8, 0x6e, 0x18, 0x1e, 0x6c,
	0xb4, 0xf0, 0x20, 0xc9, 0x21, 0x2a, 0xf9, 0x5a, 0x24, 0x9d, 0x51, 0x38, 0x1e, 0x4c, 0xd3, 0xbc,
	0xbd, 0x67, 0xee, 0x9b, 0x2e, 0xf8, 0x5a, 0x14, 0x2e, 0x2f, 0xfb, 0x1d, 0x00, 0x1c, 0x82, 0xe4,
	0x3f, 0xe8, 0xca, 0xaf, 0x82, 0x7b, 0x92, 0x7e, 0xe1, 0x1d, 0x92, 0x42, 0x6f, 0xa5, 0x90, 0x1a,
	0x9c, 0x19, 0x37, 0x7c, 0xbf, 0x68, 0x7c, 0xbb, 0x80, 0xd8, 0xcc, 0x4c, 0xbd, 0x80, 0xb5, 0xc9,
	0x3d, 0x0c, 0x85, 0x44, 0x45, 0x0d, 0x3e, 0x6a, 0x54, 0x0b, 0x96, 0x44, 0xae, 0xa8, 0x1d, 0xb4,
	0x5c, 0xb8, 0xa3, 0xe5, 0x36, 0xe9, 0x7a, 0x2e, 0xe7, 0x90, 0x57, 0x00, 0xc8, 0x0d, 0x2a, 0xa9,
	0x4a, 0x8d, 0x49, 0xec, 0xa0, 0xa3, 0x88, 0x9d, 0x65, 0xaf, 0x51, 0x71, 0xba, 0xc3, 0xe4, 0xca,
	0xcf, 0x52, 0xfb, 0x16, 0x53, 0x48, 0xb7, 0x9f, 0x2c, 0xd6, 0xf3, 0x58, 0xed, 0x5b, 0xb6, 0x2d,
	0x3e, 0xe1, 0x36, 0xe9, 0xbb, 0x41, 0xbd, 0x93, 0x7d, 0x80, 0xe1, 0xa3, 0x64, 0xf6, 0x7e, 0x95,
	0x8c, 0x29, 0xf4, 0xfc, 0xa9, 0x16, 0xac, 0xba, 0x41, 0xe3, 0x93, 0x04, 0xae, 0xe8, 0x6a, 0x25,
	0xf6, 0xbc, 0xbe, 0x42, 0xed, 0x66, 0x3f, 0x02, 0x18, 0xcc, 0xd1, 0xfc, 0xa3, 0x56, 0xc7, 0x5c,
	0x9d, 0x13, 0xae, 0x29, 0xc4, 0x0c, 0x8d, 0xbd, 0x4e, 0xe8, 0x5a, 0x5c, 0x52, 0xb2, 0xca, 0xcc,
	0x46, 0x00, 0x6e, 0x08, 0xbf, 0x09, 0x81, 0x68, 0x25, 0x58, 0xad, 0xa4, 0xb3, 0xb3, 0x5f, 0x01,
	0xc0, 0x8c, 0xb1, 0x3a, 0xe5, 0x68, 0xa1, 0xa0, 0xb5, 0xd0, 0x89, 0x0a, 0x9d, 0x8b, 0x2a, 0x84,
	0x17, 0x54, 0x88, 0x4e, 0x54, 0xa8, 0x5f, 0x49, 0xf7, 0xef, 0xaf, 0xa4, 0x11, 0x2a, 0x3e, 0x12,
	0x6a, 0xfa, 0xb3, 0x03, 0xb1, 0x5f, 0x99, 0xbc, 0x85, 0x70, 0xc6, 0x18, 0x79, 0x71, 0x91, 0xc3,
	0x62, 0xe9, 0x4d, 0x83, 0x35, 0x92, 0xbc, 0x87, 0x70, 0x8e, 0xe6, 0x65, 0xd5, 0xe1, 0x62, 0xe9,
	0xed, 0x59, 0xac, 0xaa, 0x7f, 0x07, 0xb1, 0xff, 0xa7, 0x90, 0xff, 0x4f, 0xd3, 0x5a, 0xff, 0xa0,
	0x73, 0xdc, 0x0f, 0x4e, 0x98, 0xea, 0x41, 0x93, 0xbb, 0x33, 0x34, 0xed, 0x8f, 0x49, 0x9a, 0x5d,
	0x4a, 0xf1, 0x4d, 0x97, 0xb1, 0xfb, 0x18, 0xbd, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x7a,
	0xac, 0xe6, 0xe8, 0x04, 0x00, 0x00,
}
