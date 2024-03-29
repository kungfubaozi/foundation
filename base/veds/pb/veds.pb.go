// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base/veds/pb/veds.proto

package fs_base_veds

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

type CryptRequest struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CryptRequest) Reset()         { *m = CryptRequest{} }
func (m *CryptRequest) String() string { return proto.CompactTextString(m) }
func (*CryptRequest) ProtoMessage()    {}
func (*CryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbf58f26039ab754, []int{0}
}

func (m *CryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptRequest.Unmarshal(m, b)
}
func (m *CryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptRequest.Marshal(b, m, deterministic)
}
func (m *CryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptRequest.Merge(m, src)
}
func (m *CryptRequest) XXX_Size() int {
	return xxx_messageInfo_CryptRequest.Size(m)
}
func (m *CryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CryptRequest proto.InternalMessageInfo

func (m *CryptRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CryptResponse struct {
	State                *pb.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Value                string    `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CryptResponse) Reset()         { *m = CryptResponse{} }
func (m *CryptResponse) String() string { return proto.CompactTextString(m) }
func (*CryptResponse) ProtoMessage()    {}
func (*CryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbf58f26039ab754, []int{1}
}

func (m *CryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptResponse.Unmarshal(m, b)
}
func (m *CryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptResponse.Marshal(b, m, deterministic)
}
func (m *CryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptResponse.Merge(m, src)
}
func (m *CryptResponse) XXX_Size() int {
	return xxx_messageInfo_CryptResponse.Size(m)
}
func (m *CryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CryptResponse proto.InternalMessageInfo

func (m *CryptResponse) GetState() *pb.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *CryptResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*CryptRequest)(nil), "fs.base.veds.CryptRequest")
	proto.RegisterType((*CryptResponse)(nil), "fs.base.veds.CryptResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VEDSClient is the client API for VEDS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VEDSClient interface {
	Encrypt(ctx context.Context, in *CryptRequest, opts ...grpc.CallOption) (*CryptResponse, error)
	Decrypt(ctx context.Context, in *CryptRequest, opts ...grpc.CallOption) (*CryptResponse, error)
}

type vEDSClient struct {
	cc *grpc.ClientConn
}

func NewVEDSClient(cc *grpc.ClientConn) VEDSClient {
	return &vEDSClient{cc}
}

func (c *vEDSClient) Encrypt(ctx context.Context, in *CryptRequest, opts ...grpc.CallOption) (*CryptResponse, error) {
	out := new(CryptResponse)
	err := c.cc.Invoke(ctx, "/fs.base.veds.VEDS/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vEDSClient) Decrypt(ctx context.Context, in *CryptRequest, opts ...grpc.CallOption) (*CryptResponse, error) {
	out := new(CryptResponse)
	err := c.cc.Invoke(ctx, "/fs.base.veds.VEDS/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VEDSServer is the server API for VEDS service.
type VEDSServer interface {
	Encrypt(context.Context, *CryptRequest) (*CryptResponse, error)
	Decrypt(context.Context, *CryptRequest) (*CryptResponse, error)
}

func RegisterVEDSServer(s *grpc.Server, srv VEDSServer) {
	s.RegisterService(&_VEDS_serviceDesc, srv)
}

func _VEDS_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VEDSServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.veds.VEDS/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VEDSServer).Encrypt(ctx, req.(*CryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VEDS_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VEDSServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.veds.VEDS/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VEDSServer).Decrypt(ctx, req.(*CryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VEDS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fs.base.veds.VEDS",
	HandlerType: (*VEDSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encrypt",
			Handler:    _VEDS_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _VEDS_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/veds/pb/veds.proto",
}

func init() { proto.RegisterFile("base/veds/pb/veds.proto", fileDescriptor_dbf58f26039ab754) }

var fileDescriptor_dbf58f26039ab754 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2f, 0x4b, 0x4d, 0x29, 0xd6, 0x2f, 0x48, 0x02, 0xd3, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x3c, 0x69, 0xc5, 0x7a, 0x20, 0x39, 0x3d, 0x90, 0x98, 0x94, 0x76, 0x55, 0x71, 0x76, 0x41,
	0x62, 0x51, 0x76, 0x6a, 0x91, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x5a, 0x7e, 0x69, 0x5e, 0x4a, 0x62,
	0x49, 0x66, 0x7e, 0x9e, 0x3e, 0x58, 0x7f, 0x41, 0x12, 0x98, 0x86, 0x68, 0x55, 0x52, 0xe1, 0xe2,
	0x71, 0x2e, 0xaa, 0x2c, 0x28, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62,
	0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0xbc,
	0xb9, 0x78, 0xa1, 0xaa, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x54, 0xb8, 0x58, 0x8b, 0x4b,
	0x12, 0x4b, 0x20, 0xca, 0xb8, 0x8d, 0xf8, 0xf4, 0x60, 0x2e, 0x08, 0x06, 0x89, 0x06, 0x41, 0x24,
	0x11, 0x86, 0x31, 0x21, 0x19, 0x66, 0xd4, 0xc7, 0xc8, 0xc5, 0x12, 0xe6, 0xea, 0x12, 0x2c, 0xe4,
	0xc4, 0xc5, 0xee, 0x9a, 0x97, 0x0c, 0x32, 0x57, 0x48, 0x4a, 0x0f, 0xd9, 0x0b, 0x7a, 0xc8, 0x4e,
	0x92, 0x92, 0xc6, 0x2a, 0x07, 0x75, 0x88, 0x13, 0x17, 0xbb, 0x4b, 0x2a, 0x65, 0x66, 0x24, 0xb1,
	0x81, 0x83, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x90, 0x17, 0xbe, 0xab, 0x60, 0x01, 0x00,
	0x00,
}
