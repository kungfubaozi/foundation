// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base/project/pb/project.proto

package fs_base_project

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

type GetRequest struct {
	PlatformId           string   `protobuf:"bytes,1,opt,name=platformId,proto3" json:"platformId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eee2cafb576e5e1c, []int{0}
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

func (m *GetRequest) GetPlatformId() string {
	if m != nil {
		return m.PlatformId
	}
	return ""
}

type GetResponse struct {
	State                *pb.State    `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Info                 *ProjectInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eee2cafb576e5e1c, []int{1}
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

func (m *GetResponse) GetInfo() *ProjectInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type NewRequest struct {
	Logo                 string   `protobuf:"bytes,1,opt,name=logo,proto3" json:"logo,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	En                   string   `protobuf:"bytes,3,opt,name=en,proto3" json:"en,omitempty"`
	Zh                   string   `protobuf:"bytes,4,opt,name=zh,proto3" json:"zh,omitempty"`
	Platforms            []string `protobuf:"bytes,5,rep,name=platforms,proto3" json:"platforms,omitempty"`
	Meta                 *pb.Meta `protobuf:"bytes,6,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewRequest) Reset()         { *m = NewRequest{} }
func (m *NewRequest) String() string { return proto.CompactTextString(m) }
func (*NewRequest) ProtoMessage()    {}
func (*NewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eee2cafb576e5e1c, []int{2}
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

func (m *NewRequest) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *NewRequest) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *NewRequest) GetEn() string {
	if m != nil {
		return m.En
	}
	return ""
}

func (m *NewRequest) GetZh() string {
	if m != nil {
		return m.Zh
	}
	return ""
}

func (m *NewRequest) GetPlatforms() []string {
	if m != nil {
		return m.Platforms
	}
	return nil
}

func (m *NewRequest) GetMeta() *pb.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type ProjectInfo struct {
	Logo                 string         `protobuf:"bytes,1,opt,name=logo,proto3" json:"logo,omitempty"`
	Desc                 string         `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	En                   string         `protobuf:"bytes,3,opt,name=en,proto3" json:"en,omitempty"`
	Zh                   string         `protobuf:"bytes,4,opt,name=zh,proto3" json:"zh,omitempty"`
	Platforms            []*AppPlatform `protobuf:"bytes,5,rep,name=platforms,proto3" json:"platforms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ProjectInfo) Reset()         { *m = ProjectInfo{} }
func (m *ProjectInfo) String() string { return proto.CompactTextString(m) }
func (*ProjectInfo) ProtoMessage()    {}
func (*ProjectInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_eee2cafb576e5e1c, []int{3}
}

func (m *ProjectInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectInfo.Unmarshal(m, b)
}
func (m *ProjectInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectInfo.Marshal(b, m, deterministic)
}
func (m *ProjectInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectInfo.Merge(m, src)
}
func (m *ProjectInfo) XXX_Size() int {
	return xxx_messageInfo_ProjectInfo.Size(m)
}
func (m *ProjectInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectInfo proto.InternalMessageInfo

func (m *ProjectInfo) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *ProjectInfo) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *ProjectInfo) GetEn() string {
	if m != nil {
		return m.En
	}
	return ""
}

func (m *ProjectInfo) GetZh() string {
	if m != nil {
		return m.Zh
	}
	return ""
}

func (m *ProjectInfo) GetPlatforms() []*AppPlatform {
	if m != nil {
		return m.Platforms
	}
	return nil
}

type AppPlatform struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Enabled              bool     `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppPlatform) Reset()         { *m = AppPlatform{} }
func (m *AppPlatform) String() string { return proto.CompactTextString(m) }
func (*AppPlatform) ProtoMessage()    {}
func (*AppPlatform) Descriptor() ([]byte, []int) {
	return fileDescriptor_eee2cafb576e5e1c, []int{4}
}

func (m *AppPlatform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppPlatform.Unmarshal(m, b)
}
func (m *AppPlatform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppPlatform.Marshal(b, m, deterministic)
}
func (m *AppPlatform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppPlatform.Merge(m, src)
}
func (m *AppPlatform) XXX_Size() int {
	return xxx_messageInfo_AppPlatform.Size(m)
}
func (m *AppPlatform) XXX_DiscardUnknown() {
	xxx_messageInfo_AppPlatform.DiscardUnknown(m)
}

var xxx_messageInfo_AppPlatform proto.InternalMessageInfo

func (m *AppPlatform) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *AppPlatform) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "fs.base.project.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "fs.base.project.GetResponse")
	proto.RegisterType((*NewRequest)(nil), "fs.base.project.NewRequest")
	proto.RegisterType((*ProjectInfo)(nil), "fs.base.project.ProjectInfo")
	proto.RegisterType((*AppPlatform)(nil), "fs.base.project.AppPlatform")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProjectClient is the client API for Project service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectClient interface {
	New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*pb.Response, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type projectClient struct {
	cc *grpc.ClientConn
}

func NewProjectClient(cc *grpc.ClientConn) ProjectClient {
	return &projectClient{cc}
}

func (c *projectClient) New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*pb.Response, error) {
	out := new(pb.Response)
	err := c.cc.Invoke(ctx, "/fs.base.project.Project/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/fs.base.project.Project/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServer is the server API for Project service.
type ProjectServer interface {
	New(context.Context, *NewRequest) (*pb.Response, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterProjectServer(s *grpc.Server, srv ProjectServer) {
	s.RegisterService(&_Project_serviceDesc, srv)
}

func _Project_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.project.Project/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).New(ctx, req.(*NewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.project.Project/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Project_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fs.base.project.Project",
	HandlerType: (*ProjectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _Project_New_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Project_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/project/pb/project.proto",
}

func init() { proto.RegisterFile("base/project/pb/project.proto", fileDescriptor_eee2cafb576e5e1c) }

var fileDescriptor_eee2cafb576e5e1c = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0xae, 0xd3, 0x30,
	0x10, 0x85, 0x95, 0x26, 0xfd, 0xc9, 0x44, 0x14, 0xe1, 0x55, 0x14, 0x0a, 0x2a, 0x11, 0x8b, 0x4a,
	0xa0, 0x04, 0x15, 0xb1, 0x61, 0x05, 0x62, 0x51, 0x75, 0x41, 0x55, 0x99, 0x27, 0x70, 0x92, 0x09,
	0x2d, 0x4d, 0x6d, 0x13, 0xbb, 0xaa, 0xd4, 0x25, 0x6f, 0xc0, 0xe2, 0xbe, 0xef, 0x55, 0x9c, 0x9f,
	0x46, 0xbd, 0x5d, 0xde, 0x55, 0x66, 0xce, 0x39, 0x19, 0x7d, 0xb6, 0x07, 0xde, 0x24, 0x4c, 0x61,
	0x2c, 0x4b, 0xf1, 0x07, 0x53, 0x1d, 0xcb, 0xa4, 0x2d, 0x23, 0x59, 0x0a, 0x2d, 0xc8, 0xcb, 0x5c,
	0x45, 0x55, 0x22, 0x6a, 0xe4, 0xe0, 0xc3, 0x45, 0x1d, 0x24, 0x2b, 0x0f, 0x58, 0x46, 0xa9, 0x38,
	0xc6, 0xb9, 0x38, 0xf1, 0x8c, 0xe9, 0xbd, 0xe0, 0x71, 0x3d, 0x28, 0x89, 0xdb, 0xb8, 0x16, 0xe1,
	0x47, 0x80, 0x15, 0x6a, 0x8a, 0x7f, 0x4f, 0xa8, 0x34, 0x79, 0x0b, 0x20, 0x0b, 0xa6, 0x73, 0x51,
	0x1e, 0xd7, 0x99, 0x6f, 0xcd, 0xad, 0x85, 0x4b, 0x7b, 0x4a, 0x88, 0xe0, 0x99, 0xb4, 0x92, 0x82,
	0x2b, 0x24, 0xef, 0x61, 0xa8, 0x34, 0xd3, 0x68, 0x92, 0xde, 0x72, 0x1a, 0xb5, 0x28, 0xbf, 0x2a,
	0x95, 0xd6, 0x26, 0xf9, 0x04, 0xce, 0x9e, 0xe7, 0xc2, 0x1f, 0x98, 0xd0, 0x2c, 0xba, 0xe1, 0x8d,
	0xb6, 0xf5, 0x77, 0xcd, 0x73, 0x41, 0x4d, 0x32, 0x7c, 0xb0, 0x00, 0x36, 0x78, 0x6e, 0xa9, 0x08,
	0x38, 0x85, 0xf8, 0x2d, 0x1a, 0x1e, 0x53, 0x57, 0x5a, 0x86, 0x2a, 0x35, 0x43, 0x5d, 0x6a, 0x6a,
	0x32, 0x85, 0x01, 0x72, 0xdf, 0x36, 0xca, 0x00, 0x79, 0xd5, 0x5f, 0x76, 0xbe, 0x53, 0xf7, 0x97,
	0x1d, 0x99, 0x81, 0xdb, 0x9e, 0x45, 0xf9, 0xc3, 0xb9, 0xbd, 0x70, 0xe9, 0x55, 0x20, 0xef, 0xc0,
	0x39, 0xa2, 0x66, 0xfe, 0xc8, 0x60, 0xbe, 0xe8, 0x30, 0x7f, 0xa2, 0x66, 0xd4, 0x58, 0xe1, 0x7f,
	0x0b, 0xbc, 0x1e, 0xed, 0xb3, 0x81, 0x7d, 0xbd, 0x05, 0xbb, 0x77, 0x4d, 0xdf, 0xa5, 0xdc, 0x36,
	0xa1, 0x1e, 0x76, 0xf8, 0x03, 0xbc, 0x9e, 0x43, 0x02, 0x98, 0xa4, 0xc5, 0x1e, 0xb9, 0xee, 0xde,
	0xaf, 0xeb, 0x89, 0x0f, 0x63, 0xe4, 0x2c, 0x29, 0x30, 0x33, 0x74, 0x13, 0xda, 0xb6, 0xcb, 0x7f,
	0x16, 0x8c, 0x9b, 0x83, 0x91, 0x2f, 0x60, 0x6f, 0xf0, 0x4c, 0x5e, 0x3f, 0x01, 0xb8, 0xbe, 0x48,
	0xf0, 0xaa, 0x33, 0xbb, 0x5d, 0xf8, 0x06, 0xf6, 0x0a, 0xf5, 0x9d, 0xdf, 0xae, 0xeb, 0x15, 0xcc,
	0xee, 0x9b, 0xf5, 0x84, 0x64, 0x64, 0x36, 0xf2, 0xf3, 0x63, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16,
	0x99, 0xf8, 0xdb, 0xf0, 0x02, 0x00, 0x00,
}
