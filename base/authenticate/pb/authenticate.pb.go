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

type RefreshRequest struct {
	Meta                 *pb.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshRequest) Reset()         { *m = RefreshRequest{} }
func (m *RefreshRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshRequest) ProtoMessage()    {}
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{0}
}

func (m *RefreshRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshRequest.Unmarshal(m, b)
}
func (m *RefreshRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshRequest.Marshal(b, m, deterministic)
}
func (m *RefreshRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshRequest.Merge(m, src)
}
func (m *RefreshRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshRequest.Size(m)
}
func (m *RefreshRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshRequest proto.InternalMessageInfo

func (m *RefreshRequest) GetMeta() *pb.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RefreshRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type RefreshResponse struct {
	State                *pb.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	AccessToken          string    `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RefreshResponse) Reset()         { *m = RefreshResponse{} }
func (m *RefreshResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshResponse) ProtoMessage()    {}
func (*RefreshResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{1}
}

func (m *RefreshResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshResponse.Unmarshal(m, b)
}
func (m *RefreshResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshResponse.Marshal(b, m, deterministic)
}
func (m *RefreshResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshResponse.Merge(m, src)
}
func (m *RefreshResponse) XXX_Size() int {
	return xxx_messageInfo_RefreshResponse.Size(m)
}
func (m *RefreshResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshResponse proto.InternalMessageInfo

func (m *RefreshResponse) GetState() *pb.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *RefreshResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type RouteRequest struct {
	Meta                 *pb.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteRequest) Reset()         { *m = RouteRequest{} }
func (m *RouteRequest) String() string { return proto.CompactTextString(m) }
func (*RouteRequest) ProtoMessage()    {}
func (*RouteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{2}
}

func (m *RouteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteRequest.Unmarshal(m, b)
}
func (m *RouteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteRequest.Marshal(b, m, deterministic)
}
func (m *RouteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteRequest.Merge(m, src)
}
func (m *RouteRequest) XXX_Size() int {
	return xxx_messageInfo_RouteRequest.Size(m)
}
func (m *RouteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RouteRequest proto.InternalMessageInfo

func (m *RouteRequest) GetMeta() *pb.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RouteRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

// 跳转不用返回refreshToken
type RouteResponse struct {
	State                *pb.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	AccessToken          string    `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	Session              string    `protobuf:"bytes,3,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RouteResponse) Reset()         { *m = RouteResponse{} }
func (m *RouteResponse) String() string { return proto.CompactTextString(m) }
func (*RouteResponse) ProtoMessage()    {}
func (*RouteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{3}
}

func (m *RouteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteResponse.Unmarshal(m, b)
}
func (m *RouteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteResponse.Marshal(b, m, deterministic)
}
func (m *RouteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteResponse.Merge(m, src)
}
func (m *RouteResponse) XXX_Size() int {
	return xxx_messageInfo_RouteResponse.Size(m)
}
func (m *RouteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RouteResponse proto.InternalMessageInfo

func (m *RouteResponse) GetState() *pb.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *RouteResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *RouteResponse) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type OfflineUserRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfflineUserRequest) Reset()         { *m = OfflineUserRequest{} }
func (m *OfflineUserRequest) String() string { return proto.CompactTextString(m) }
func (*OfflineUserRequest) ProtoMessage()    {}
func (*OfflineUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{4}
}

func (m *OfflineUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfflineUserRequest.Unmarshal(m, b)
}
func (m *OfflineUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfflineUserRequest.Marshal(b, m, deterministic)
}
func (m *OfflineUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfflineUserRequest.Merge(m, src)
}
func (m *OfflineUserRequest) XXX_Size() int {
	return xxx_messageInfo_OfflineUserRequest.Size(m)
}
func (m *OfflineUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OfflineUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OfflineUserRequest proto.InternalMessageInfo

func (m *OfflineUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type CheckRequest struct {
	Metadata                     *pb.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	MaxOnlineCount               int64        `protobuf:"varint,2,opt,name=maxOnlineCount,proto3" json:"maxOnlineCount,omitempty"`
	AllowOtherProjectUserToLogin bool         `protobuf:"varint,3,opt,name=allowOtherProjectUserToLogin,proto3" json:"allowOtherProjectUserToLogin,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}     `json:"-"`
	XXX_unrecognized             []byte       `json:"-"`
	XXX_sizecache                int32        `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{5}
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

func (m *CheckRequest) GetMetadata() *pb.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CheckRequest) GetMaxOnlineCount() int64 {
	if m != nil {
		return m.MaxOnlineCount
	}
	return 0
}

func (m *CheckRequest) GetAllowOtherProjectUserToLogin() bool {
	if m != nil {
		return m.AllowOtherProjectUserToLogin
	}
	return false
}

type CheckResponse struct {
	State                *pb.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	UserId               string    `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ProjectId            string    `protobuf:"bytes,3,opt,name=projectId,proto3" json:"projectId,omitempty"`
	ClientId             string    `protobuf:"bytes,4,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Relation             string    `protobuf:"bytes,5,opt,name=relation,proto3" json:"relation,omitempty"`
	Level                int64     `protobuf:"varint,6,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{6}
}

func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetState() *pb.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *CheckResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CheckResponse) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *CheckResponse) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *CheckResponse) GetRelation() string {
	if m != nil {
		return m.Relation
	}
	return ""
}

func (m *CheckResponse) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

type NewResponse struct {
	State                *pb.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	RefreshToken         string    `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	AccessToken          string    `protobuf:"bytes,3,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	Session              string    `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *NewResponse) Reset()         { *m = NewResponse{} }
func (m *NewResponse) String() string { return proto.CompactTextString(m) }
func (*NewResponse) ProtoMessage()    {}
func (*NewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{7}
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

func (m *NewResponse) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type NewRequest struct {
	MaxOnlineCount       int64      `protobuf:"varint,1,opt,name=maxOnlineCount,proto3" json:"maxOnlineCount,omitempty"`
	Authorize            *Authorize `protobuf:"bytes,2,opt,name=authorize,proto3" json:"authorize,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NewRequest) Reset()         { *m = NewRequest{} }
func (m *NewRequest) String() string { return proto.CompactTextString(m) }
func (*NewRequest) ProtoMessage()    {}
func (*NewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{8}
}

func (m *NewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewRequest.Unmarshal(m, b)
}
func (m *NewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewRequest.Marshal(b, m, deterministic)
}
func (m *NewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewRequest.Merge(m, src)
}
func (m *NewRequest) XXX_Size() int {
	return xxx_messageInfo_NewRequest.Size(m)
}
func (m *NewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewRequest proto.InternalMessageInfo

func (m *NewRequest) GetMaxOnlineCount() int64 {
	if m != nil {
		return m.MaxOnlineCount
	}
	return 0
}

func (m *NewRequest) GetAuthorize() *Authorize {
	if m != nil {
		return m.Authorize
	}
	return nil
}

type Authorize struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ProjectId            string   `protobuf:"bytes,3,opt,name=projectId,proto3" json:"projectId,omitempty"`
	ClientId             string   `protobuf:"bytes,4,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Device               string   `protobuf:"bytes,5,opt,name=device,proto3" json:"device,omitempty"`
	Platform             int64    `protobuf:"varint,6,opt,name=platform,proto3" json:"platform,omitempty"`
	UserAgent            string   `protobuf:"bytes,7,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
	Ip                   string   `protobuf:"bytes,8,opt,name=ip,proto3" json:"ip,omitempty"`
	AccessTokenUUID      string   `protobuf:"bytes,9,opt,name=accessTokenUUID,proto3" json:"accessTokenUUID,omitempty"`
	RefreshTokenUUID     string   `protobuf:"bytes,10,opt,name=refreshTokenUUID,proto3" json:"refreshTokenUUID,omitempty"`
	Relation             string   `protobuf:"bytes,11,opt,name=relation,proto3" json:"relation,omitempty"`
	Level                int64    `protobuf:"varint,12,opt,name=level,proto3" json:"level,omitempty"`
	LoginMode            string   `protobuf:"bytes,13,opt,name=loginMode,proto3" json:"loginMode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Authorize) Reset()         { *m = Authorize{} }
func (m *Authorize) String() string { return proto.CompactTextString(m) }
func (*Authorize) ProtoMessage()    {}
func (*Authorize) Descriptor() ([]byte, []int) {
	return fileDescriptor_47da8fcd76e8b636, []int{9}
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

func (m *Authorize) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Authorize) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Authorize) GetClientId() string {
	if m != nil {
		return m.ClientId
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

func (m *Authorize) GetAccessTokenUUID() string {
	if m != nil {
		return m.AccessTokenUUID
	}
	return ""
}

func (m *Authorize) GetRefreshTokenUUID() string {
	if m != nil {
		return m.RefreshTokenUUID
	}
	return ""
}

func (m *Authorize) GetRelation() string {
	if m != nil {
		return m.Relation
	}
	return ""
}

func (m *Authorize) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Authorize) GetLoginMode() string {
	if m != nil {
		return m.LoginMode
	}
	return ""
}

func init() {
	proto.RegisterType((*RefreshRequest)(nil), "fs.base.authenticate.RefreshRequest")
	proto.RegisterType((*RefreshResponse)(nil), "fs.base.authenticate.RefreshResponse")
	proto.RegisterType((*RouteRequest)(nil), "fs.base.authenticate.RouteRequest")
	proto.RegisterType((*RouteResponse)(nil), "fs.base.authenticate.RouteResponse")
	proto.RegisterType((*OfflineUserRequest)(nil), "fs.base.authenticate.OfflineUserRequest")
	proto.RegisterType((*CheckRequest)(nil), "fs.base.authenticate.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "fs.base.authenticate.CheckResponse")
	proto.RegisterType((*NewResponse)(nil), "fs.base.authenticate.NewResponse")
	proto.RegisterType((*NewRequest)(nil), "fs.base.authenticate.NewRequest")
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
	New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*NewResponse, error)
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	OfflineUser(ctx context.Context, in *OfflineUserRequest, opts ...grpc.CallOption) (*pb.Response, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
}

type authenticateClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticateClient(cc *grpc.ClientConn) AuthenticateClient {
	return &authenticateClient{cc}
}

func (c *authenticateClient) New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*NewResponse, error) {
	out := new(NewResponse)
	err := c.cc.Invoke(ctx, "/fs.base.authenticate.Authenticate/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/fs.base.authenticate.Authenticate/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateClient) OfflineUser(ctx context.Context, in *OfflineUserRequest, opts ...grpc.CallOption) (*pb.Response, error) {
	out := new(pb.Response)
	err := c.cc.Invoke(ctx, "/fs.base.authenticate.Authenticate/OfflineUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, "/fs.base.authenticate.Authenticate/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticateServer is the server API for Authenticate service.
type AuthenticateServer interface {
	New(context.Context, *NewRequest) (*NewResponse, error)
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	OfflineUser(context.Context, *OfflineUserRequest) (*pb.Response, error)
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
}

func RegisterAuthenticateServer(s *grpc.Server, srv AuthenticateServer) {
	s.RegisterService(&_Authenticate_serviceDesc, srv)
}

func _Authenticate_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRequest)
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
		return srv.(AuthenticateServer).New(ctx, req.(*NewRequest))
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

func _Authenticate_OfflineUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfflineUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServer).OfflineUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.authenticate.Authenticate/OfflineUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServer).OfflineUser(ctx, req.(*OfflineUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticate_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.authenticate.Authenticate/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServer).Refresh(ctx, req.(*RefreshRequest))
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
			MethodName: "OfflineUser",
			Handler:    _Authenticate_OfflineUser_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Authenticate_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/authenticate/pb/authenticate.proto",
}

func init() {
	proto.RegisterFile("base/authenticate/pb/authenticate.proto", fileDescriptor_47da8fcd76e8b636)
}

var fileDescriptor_47da8fcd76e8b636 = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x96, 0x93, 0x26, 0x4d, 0x26, 0x87, 0xfe, 0xff, 0xaa, 0xaa, 0xac, 0xa8, 0x12, 0xa9, 0x29,
	0x10, 0x71, 0x70, 0xa4, 0x72, 0xcd, 0x45, 0x29, 0x37, 0xad, 0xe8, 0x41, 0xa6, 0x01, 0x71, 0xb9,
	0x71, 0x26, 0x8d, 0x89, 0xe3, 0x75, 0x77, 0x37, 0x2d, 0xea, 0x2d, 0x8f, 0xc0, 0x53, 0xf0, 0x08,
	0xbc, 0x05, 0x8f, 0x84, 0x76, 0xbd, 0x71, 0x9c, 0xc6, 0x0d, 0x54, 0xf4, 0xca, 0x9a, 0xc3, 0x7e,
	0xf3, 0xcd, 0xce, 0xe7, 0x59, 0x78, 0xd6, 0xa7, 0x02, 0xbb, 0x74, 0x2a, 0x47, 0x18, 0xc9, 0xc0,
	0xa7, 0x12, 0xbb, 0x71, 0x7f, 0xc1, 0x76, 0x63, 0xce, 0x24, 0x23, 0x9b, 0x43, 0xe1, 0xaa, 0x5c,
	0x37, 0x1b, 0x6b, 0xbd, 0xb8, 0x11, 0xe3, 0x98, 0xf2, 0x31, 0x72, 0xd7, 0x67, 0x93, 0xee, 0x90,
	0x4d, 0xa3, 0x01, 0x95, 0x01, 0x8b, 0xba, 0x1a, 0x37, 0xee, 0xeb, 0x6f, 0x02, 0xe1, 0x7c, 0x82,
	0xa6, 0x87, 0x43, 0x8e, 0x62, 0xe4, 0xe1, 0xe5, 0x14, 0x85, 0x24, 0x3b, 0xb0, 0x36, 0x41, 0x49,
	0x6d, 0xab, 0x6d, 0x75, 0x6a, 0x7b, 0x0d, 0x77, 0x56, 0xe3, 0x18, 0x25, 0xf5, 0x74, 0x88, 0x38,
	0x50, 0xe7, 0xc9, 0xa1, 0x73, 0x36, 0xc6, 0xc8, 0x2e, 0xb4, 0xad, 0x4e, 0xd5, 0x5b, 0xf0, 0x39,
	0x9f, 0x61, 0x23, 0x05, 0x16, 0x31, 0x8b, 0x04, 0x92, 0x5d, 0x28, 0x09, 0x49, 0x25, 0x1a, 0xe8,
	0x66, 0x0a, 0xfd, 0x41, 0x79, 0xbd, 0x24, 0x48, 0xda, 0x50, 0xa3, 0xbe, 0x8f, 0x42, 0x64, 0xb1,
	0xb3, 0x2e, 0xa7, 0x07, 0x75, 0x8f, 0x4d, 0x25, 0x3e, 0x30, 0xe3, 0x4b, 0x68, 0x18, 0xd8, 0x87,
	0xe5, 0x4b, 0x6c, 0x58, 0x17, 0x28, 0x44, 0xc0, 0x22, 0xbb, 0xa8, 0xa3, 0x33, 0xd3, 0x79, 0x09,
	0xe4, 0x74, 0x38, 0x0c, 0x83, 0x08, 0x7b, 0x02, 0xf9, 0xac, 0x9f, 0x2d, 0x28, 0x4f, 0x05, 0xf2,
	0xc3, 0x81, 0x2e, 0x5c, 0xf5, 0x8c, 0xe5, 0xfc, 0xb0, 0xa0, 0x7e, 0x30, 0x42, 0x7f, 0x3c, 0x4b,
	0x7c, 0x05, 0x15, 0xd5, 0xdd, 0x80, 0xa6, 0xcd, 0xff, 0xbf, 0xd0, 0xbc, 0x0a, 0x78, 0x69, 0x0a,
	0x79, 0x0a, 0xcd, 0x09, 0xfd, 0x7a, 0x1a, 0xa9, 0x7a, 0x07, 0x6c, 0x1a, 0x49, 0x4d, 0xb6, 0xe8,
	0xdd, 0xf2, 0x92, 0xb7, 0xb0, 0x4d, 0xc3, 0x90, 0x5d, 0x9f, 0xca, 0x11, 0xf2, 0x33, 0xce, 0xbe,
	0xa0, 0x2f, 0x15, 0xbf, 0x73, 0xf6, 0x9e, 0x5d, 0x04, 0x49, 0x13, 0x15, 0x6f, 0x65, 0x8e, 0xf3,
	0xd3, 0x82, 0x86, 0xe1, 0x7a, 0xaf, 0xdb, 0x9c, 0xf7, 0x5e, 0xc8, 0xf6, 0x4e, 0xb6, 0xa1, 0x1a,
	0x27, 0x55, 0x0e, 0x07, 0xe6, 0x16, 0xe7, 0x0e, 0xd2, 0x82, 0x8a, 0x1f, 0x06, 0x18, 0xa9, 0xe0,
	0x9a, 0x0e, 0xa6, 0xb6, 0x8a, 0x71, 0x0c, 0xf5, 0x0f, 0x60, 0x97, 0x92, 0xd8, 0xcc, 0x26, 0x9b,
	0x50, 0x0a, 0xf1, 0x0a, 0x43, 0xbb, 0xac, 0x2f, 0x22, 0x31, 0x9c, 0xef, 0x16, 0xd4, 0x4e, 0xf0,
	0xfa, 0x9e, 0xcc, 0xff, 0x42, 0x62, 0xb7, 0xb5, 0x52, 0x5c, 0xa9, 0x95, 0xb5, 0x45, 0xad, 0x08,
	0x00, 0x4d, 0x2a, 0x19, 0xfd, 0xf2, 0x2c, 0xad, 0xdc, 0x59, 0xbe, 0x81, 0xaa, 0x5a, 0x0e, 0x8c,
	0x07, 0x37, 0xa8, 0x29, 0xd5, 0xf6, 0x1e, 0xb9, 0x79, 0x6b, 0xc3, 0xdd, 0x9f, 0xa5, 0x79, 0xf3,
	0x13, 0xce, 0xb7, 0x22, 0x54, 0xd3, 0xc0, 0x5d, 0xc2, 0x54, 0xc3, 0x91, 0xc1, 0x04, 0x85, 0xa4,
	0x93, 0xd8, 0x68, 0x6a, 0xee, 0xf8, 0x87, 0xd1, 0x6d, 0x41, 0x79, 0x80, 0x57, 0x81, 0x8f, 0x66,
	0x70, 0xc6, 0x52, 0x67, 0xe2, 0x90, 0xca, 0x21, 0xe3, 0x13, 0x33, 0xb9, 0xd4, 0x56, 0xd5, 0x14,
	0xab, 0xfd, 0x0b, 0x8c, 0xa4, 0xbd, 0x9e, 0x54, 0x4b, 0x1d, 0xa4, 0x09, 0x85, 0x20, 0xb6, 0x2b,
	0xda, 0x5d, 0x08, 0x62, 0xd2, 0x81, 0x8d, 0xcc, 0xed, 0xf7, 0x7a, 0x87, 0xef, 0xec, 0xaa, 0x0e,
	0xde, 0x76, 0x93, 0xe7, 0xf0, 0x5f, 0x76, 0x94, 0x3a, 0x15, 0x74, 0xea, 0x92, 0x7f, 0x41, 0x72,
	0xb5, 0xbb, 0x24, 0x57, 0xcf, 0x48, 0x4e, 0xb1, 0x0e, 0xd5, 0x7f, 0x73, 0xcc, 0x06, 0x68, 0x37,
	0x12, 0xd6, 0xa9, 0x63, 0xef, 0x57, 0x01, 0xea, 0xfb, 0x99, 0x59, 0x91, 0x23, 0x28, 0x9e, 0xe0,
	0x35, 0x69, 0xe7, 0x4f, 0x72, 0x2e, 0x93, 0xd6, 0xce, 0x8a, 0x0c, 0xa3, 0xee, 0x33, 0x28, 0xe9,
	0x1f, 0x95, 0x38, 0xf9, 0xb9, 0xd9, 0x8d, 0xd3, 0x7a, 0xbc, 0x32, 0xc7, 0x20, 0x1e, 0x41, 0x2d,
	0xb3, 0xd5, 0x48, 0x27, 0xff, 0xcc, 0xf2, 0xe2, 0x6b, 0xcd, 0xb7, 0x57, 0x8a, 0xf5, 0x11, 0xd6,
	0xcd, 0x33, 0x42, 0x76, 0xf3, 0x71, 0x16, 0x9f, 0xaf, 0xd6, 0x93, 0x3f, 0x64, 0x25, 0xb8, 0xfd,
	0xb2, 0x7e, 0xfe, 0x5e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x4e, 0xe5, 0xd5, 0x6c, 0x07,
	0x00, 0x00,
}
