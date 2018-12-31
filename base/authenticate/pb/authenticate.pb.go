// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base/authenticate/pb/authenticate.proto

package fs_base_authenticate

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

type OfflineRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfflineRequest) Reset()         { *m = OfflineRequest{} }
func (m *OfflineRequest) String() string { return proto.CompactTextString(m) }
func (*OfflineRequest) ProtoMessage()    {}
func (*OfflineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{0}
}

func (m *OfflineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfflineRequest.Unmarshal(m, b)
}
func (m *OfflineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfflineRequest.Marshal(b, m, deterministic)
}
func (m *OfflineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfflineRequest.Merge(m, src)
}
func (m *OfflineRequest) XXX_Size() int {
	return xxx_messageInfo_OfflineRequest.Size(m)
}
func (m *OfflineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OfflineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OfflineRequest proto.InternalMessageInfo

func (m *OfflineRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type CheckRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{1}
}

func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (m *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(m, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type NewResponse struct {
	State                *pb.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	RefreshToken         string    `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	AccessToken          string    `protobuf:"bytes,3,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	Timestamp            string    `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *NewResponse) Reset()         { *m = NewResponse{} }
func (m *NewResponse) String() string { return proto.CompactTextString(m) }
func (*NewResponse) ProtoMessage()    {}
func (*NewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{2}
}

func (m *NewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewResponse.Unmarshal(m, b)
}
func (m *NewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewResponse.Marshal(b, m, deterministic)
}
func (m *NewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewResponse.Merge(m, src)
}
func (m *NewResponse) XXX_Size() int {
	return xxx_messageInfo_NewResponse.Size(m)
}
func (m *NewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewResponse proto.InternalMessageInfo

func (m *NewResponse) GetState() *pb.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *NewResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *NewResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *NewResponse) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type Authorize struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	AppId                string   `protobuf:"bytes,3,opt,name=appId,proto3" json:"appId,omitempty"`
	ProjectId            string   `protobuf:"bytes,4,opt,name=projectId,proto3" json:"projectId,omitempty"`
	Device               string   `protobuf:"bytes,5,opt,name=device,proto3" json:"device,omitempty"`
	Platform             int64    `protobuf:"varint,6,opt,name=platform,proto3" json:"platform,omitempty"`
	UserAgent            string   `protobuf:"bytes,7,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
	Ip                   string   `protobuf:"bytes,8,opt,name=ip,proto3" json:"ip,omitempty"`
	Token                string   `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty"`
	Id                   string   `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	Ab                   string   `protobuf:"bytes,11,opt,name=ab,proto3" json:"ab,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Authorize) Reset()         { *m = Authorize{} }
func (m *Authorize) String() string { return proto.CompactTextString(m) }
func (*Authorize) ProtoMessage()    {}
func (*Authorize) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{3}
}

func (m *Authorize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authorize.Unmarshal(m, b)
}
func (m *Authorize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authorize.Marshal(b, m, deterministic)
}
func (m *Authorize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authorize.Merge(m, src)
}
func (m *Authorize) XXX_Size() int {
	return xxx_messageInfo_Authorize.Size(m)
}
func (m *Authorize) XXX_DiscardUnknown() {
	xxx_messageInfo_Authorize.DiscardUnknown(m)
}

var xxx_messageInfo_Authorize proto.InternalMessageInfo

func (m *Authorize) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Authorize) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Authorize) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Authorize) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Authorize) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *Authorize) GetPlatform() int64 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *Authorize) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *Authorize) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Authorize) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Authorize) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Authorize) GetAb() string {
	if m != nil {
		return m.Ab
	}
	return ""
}

func init() {
	proto.RegisterType((*OfflineRequest)(nil), "fs.base.authenticate.OfflineRequest")
	proto.RegisterType((*CheckRequest)(nil), "fs.base.authenticate.CheckRequest")
	proto.RegisterType((*NewResponse)(nil), "fs.base.authenticate.NewResponse")
	proto.RegisterType((*Authorize)(nil), "fs.base.authenticate.Authorize")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticateClient is the client API for Authenticate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticateClient interface {
	New(ctx context.Context, in *Authorize, opts ...grpc.CallOption) (*NewResponse, error)
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*pb.Response, error)
	Offline(ctx context.Context, in *OfflineRequest, opts ...grpc.CallOption) (*pb.Response, error)
}

type authenticateClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticateClient(cc *grpc.ClientConn) AuthenticateClient {
	return &authenticateClient{cc}
}

func (c *authenticateClient) New(ctx context.Context, in *Authorize, opts ...grpc.CallOption) (*NewResponse, error) {
	out := new(NewResponse)
	err := c.cc.Invoke(ctx, "/fs.base.authenticate.Authenticate/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*pb.Response, error) {
	out := new(pb.Response)
	err := c.cc.Invoke(ctx, "/fs.base.authenticate.Authenticate/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateClient) Offline(ctx context.Context, in *OfflineRequest, opts ...grpc.CallOption) (*pb.Response, error) {
	out := new(pb.Response)
	err := c.cc.Invoke(ctx, "/fs.base.authenticate.Authenticate/Offline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticateServer is the server API for Authenticate service.
type AuthenticateServer interface {
	New(context.Context, *Authorize) (*NewResponse, error)
	Check(context.Context, *CheckRequest) (*pb.Response, error)
	Offline(context.Context, *OfflineRequest) (*pb.Response, error)
}

func RegisterAuthenticateServer(s *grpc.Server, srv AuthenticateServer) {
	s.RegisterService(&_Authenticate_serviceDesc, srv)
}

func _Authenticate_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Authorize)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.authenticate.Authenticate/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServer).New(ctx, req.(*Authorize))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticate_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.authenticate.Authenticate/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticate_Offline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfflineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServer).Offline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.authenticate.Authenticate/Offline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServer).Offline(ctx, req.(*OfflineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authenticate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fs.base.authenticate.Authenticate",
	HandlerType: (*AuthenticateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _Authenticate_New_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Authenticate_Check_Handler,
		},
		{
			MethodName: "Offline",
			Handler:    _Authenticate_Offline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/authenticate/pb/authenticate.proto",
}

func init() {
	proto.RegisterFile("base/authenticate/pb/authenticate.proto", fileDescriptor_47da8fcd76e8b636)
}

var fileDescriptor_47da8fcd76e8b636 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xd1, 0x6a, 0xdb, 0x30,
	0x18, 0x85, 0xb1, 0xb3, 0xa4, 0xcd, 0x9f, 0x10, 0x98, 0x28, 0x43, 0x84, 0xc1, 0x32, 0x13, 0x58,
	0x60, 0xe0, 0x40, 0x77, 0x3f, 0xc8, 0x76, 0x95, 0x9b, 0x0e, 0xbc, 0xbd, 0x80, 0x6c, 0xff, 0x9e,
	0x35, 0xd7, 0x96, 0x26, 0xc9, 0x2b, 0xf4, 0x1d, 0x76, 0xbd, 0xa7, 0xdb, 0xbb, 0x0c, 0x49, 0x8e,
	0x6b, 0x0f, 0xf7, 0xca, 0x9c, 0xa3, 0x4f, 0xe7, 0x97, 0xac, 0x03, 0xef, 0x52, 0xa6, 0xf1, 0xc8,
	0x5a, 0x53, 0x62, 0x63, 0x78, 0xc6, 0x0c, 0x1e, 0x65, 0x3a, 0xd2, 0xb1, 0x54, 0xc2, 0x08, 0x72,
	0x53, 0xe8, 0xd8, 0xb2, 0xf1, 0x70, 0x6d, 0xfb, 0xfe, 0x51, 0x57, 0x92, 0xa9, 0x0a, 0x55, 0x9c,
	0x89, 0xfa, 0x58, 0x88, 0xb6, 0xc9, 0x99, 0xe1, 0xa2, 0x39, 0xba, 0x5c, 0x99, 0xba, 0xaf, 0x8f,
	0x88, 0x0e, 0xb0, 0xf9, 0x52, 0x14, 0xf7, 0xbc, 0xc1, 0x04, 0x7f, 0xb6, 0xa8, 0x0d, 0x79, 0x05,
	0x8b, 0x56, 0xa3, 0x3a, 0xe7, 0x34, 0xd8, 0x05, 0x87, 0x65, 0xd2, 0xa9, 0x68, 0x0f, 0xeb, 0xcf,
	0x25, 0x66, 0xd5, 0x85, 0xbb, 0x81, 0xb9, 0x11, 0x15, 0x36, 0x1d, 0xe6, 0x45, 0xf4, 0x27, 0x80,
	0xd5, 0x1d, 0x3e, 0x24, 0xa8, 0xa5, 0x68, 0x34, 0x92, 0x3d, 0xcc, 0xb5, 0x61, 0x06, 0x1d, 0xb5,
	0xba, 0xdd, 0xc4, 0x97, 0x23, 0x7f, 0xb5, 0x6e, 0xe2, 0x17, 0x49, 0x04, 0x6b, 0x85, 0x85, 0x42,
	0x5d, 0x7e, 0x73, 0x91, 0xa1, 0x8b, 0x1c, 0x79, 0x64, 0x07, 0x2b, 0x96, 0x65, 0xa8, 0xb5, 0x47,
	0x66, 0x0e, 0x19, 0x5a, 0xe4, 0x35, 0x2c, 0x0d, 0xaf, 0x51, 0x1b, 0x56, 0x4b, 0xfa, 0xc2, 0xad,
	0x3f, 0x19, 0xd1, 0xef, 0x10, 0x96, 0xa7, 0xd6, 0x94, 0x42, 0xf1, 0x47, 0x7c, 0xee, 0x96, 0xe3,
	0x8c, 0xf0, 0xbf, 0x0c, 0x7b, 0x67, 0x26, 0xe5, 0x39, 0xef, 0xa6, 0x7b, 0x61, 0xf7, 0x48, 0x25,
	0x7e, 0x60, 0x66, 0xce, 0xf9, 0x65, 0x6e, 0x6f, 0xd8, 0x49, 0x39, 0xfe, 0xe2, 0x19, 0xd2, 0xb9,
	0x9f, 0xe4, 0x15, 0xd9, 0xc2, 0xb5, 0xbc, 0x67, 0xa6, 0x10, 0xaa, 0xa6, 0x8b, 0x5d, 0x70, 0x98,
	0x25, 0xbd, 0xb6, 0x89, 0xf6, 0x3c, 0xa7, 0xef, 0xd8, 0x18, 0x7a, 0xe5, 0x13, 0x7b, 0x83, 0x6c,
	0x20, 0xe4, 0x92, 0x5e, 0x3b, 0x3b, 0xe4, 0xf2, 0xe9, 0x25, 0x96, 0x83, 0x97, 0x70, 0x54, 0x4e,
	0xa1, 0xa3, 0x72, 0xab, 0x59, 0x4a, 0x57, 0x5e, 0xb3, 0xf4, 0xf6, 0x6f, 0x00, 0xeb, 0xd3, 0xa0,
	0x37, 0xe4, 0x0c, 0xb3, 0x3b, 0x7c, 0x20, 0x6f, 0xe2, 0xa9, 0x56, 0xc5, 0xfd, 0xaf, 0xdb, 0xbe,
	0x9d, 0x06, 0x86, 0xaf, 0xfe, 0x11, 0xe6, 0xae, 0x2b, 0x24, 0x9a, 0x66, 0x87, 0x45, 0xda, 0xbe,
	0xec, 0x99, 0x7e, 0xff, 0x27, 0xb8, 0xea, 0x5a, 0x49, 0xf6, 0xd3, 0x09, 0xe3, 0xd2, 0x4e, 0x64,
	0xa4, 0x0b, 0x57, 0xf0, 0x0f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xce, 0xab, 0x70, 0xed, 0x4e,
	0x03, 0x00, 0x00,
}
