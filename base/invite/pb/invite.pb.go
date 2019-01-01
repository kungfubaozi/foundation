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

type AddRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Enterprise           string   `protobuf:"bytes,2,opt,name=enterprise,proto3" json:"enterprise,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bf72a1accb58b55, []int{0}
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

func (m *AddRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddRequest) GetEnterprise() string {
	if m != nil {
		return m.Enterprise
	}
	return ""
}

func init() {
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

// InviteServer is the server API for Invite service.
type InviteServer interface {
	// 添加
	Add(context.Context, *AddRequest) (*pb.Response, error)
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

var _Invite_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fs.base.invite.Invite",
	HandlerType: (*InviteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Invite_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/invite/pb/invite.proto",
}

func init() { proto.RegisterFile("base/invite/pb/invite.proto", fileDescriptor_0bf72a1accb58b55) }

var fileDescriptor_0bf72a1accb58b55 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xcd, 0xaa, 0xc2, 0x30,
	0x10, 0x85, 0x69, 0x2f, 0x14, 0xee, 0x2c, 0x0a, 0x66, 0x55, 0x2a, 0x88, 0xb8, 0x12, 0x84, 0x04,
	0xd4, 0xa5, 0x08, 0x5d, 0xba, 0xed, 0x1b, 0x34, 0xce, 0x14, 0x42, 0x31, 0x89, 0x99, 0xd4, 0x85,
	0x4f, 0x2f, 0xa6, 0xe2, 0xcf, 0x6a, 0x86, 0xef, 0x30, 0x67, 0x3e, 0x98, 0xeb, 0x8e, 0x49, 0x19,
	0x7b, 0x33, 0x91, 0x94, 0xd7, 0xaf, 0x4d, 0xfa, 0xe0, 0xa2, 0x13, 0x65, 0xcf, 0xf2, 0x99, 0xcb,
	0x89, 0xd6, 0x9b, 0x3b, 0x0f, 0xbe, 0x0b, 0x03, 0x05, 0x79, 0x76, 0x17, 0xd5, 0xbb, 0xd1, 0x62,
	0x17, 0x8d, 0xb3, 0x2a, 0xb5, 0x78, 0x9d, 0xe6, 0x74, 0xbc, 0x3a, 0x00, 0x34, 0x88, 0x2d, 0x5d,
	0x47, 0xe2, 0x28, 0x4a, 0xc8, 0x0d, 0x56, 0xd9, 0x32, 0x5b, 0xff, 0xb7, 0xb9, 0x41, 0xb1, 0x00,
	0x20, 0x1b, 0x29, 0xf8, 0x60, 0x98, 0xaa, 0x3c, 0xf1, 0x2f, 0xb2, 0x3d, 0x42, 0x71, 0x4a, 0x4f,
	0xc5, 0x1e, 0xfe, 0x1a, 0x44, 0x51, 0xcb, 0x5f, 0x19, 0xf9, 0x29, 0xaf, 0x67, 0xef, 0xac, 0x25,
	0xf6, 0xce, 0x32, 0xe9, 0x22, 0x49, 0xec, 0x1e, 0x01, 0x00, 0x00, 0xff, 0xff, 0x72, 0x41, 0x38,
	0x4c, 0xe0, 0x00, 0x00, 0x00,
}