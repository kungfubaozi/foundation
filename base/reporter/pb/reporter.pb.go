// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base/reporter/pb/reporter.proto

package fs_base_reporter

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

type WriteRequest struct {
	Svc                  string   `protobuf:"bytes,1,opt,name=svc,proto3" json:"svc,omitempty"`
	Func                 string   `protobuf:"bytes,2,opt,name=func,proto3" json:"func,omitempty"`
	Who                  string   `protobuf:"bytes,3,opt,name=who,proto3" json:"who,omitempty"`
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Where                string   `protobuf:"bytes,5,opt,name=where,proto3" json:"where,omitempty"`
	Date                 string   `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaba7a9b61ea1b4e, []int{0}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetSvc() string {
	if m != nil {
		return m.Svc
	}
	return ""
}

func (m *WriteRequest) GetFunc() string {
	if m != nil {
		return m.Func
	}
	return ""
}

func (m *WriteRequest) GetWho() string {
	if m != nil {
		return m.Who
	}
	return ""
}

func (m *WriteRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *WriteRequest) GetWhere() string {
	if m != nil {
		return m.Where
	}
	return ""
}

func (m *WriteRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func init() {
	proto.RegisterType((*WriteRequest)(nil), "fs.base.reporter.WriteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReporterClient is the client API for Reporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReporterClient interface {
	// 写入
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*pb.Response, error)
}

type reporterClient struct {
	cc *grpc.ClientConn
}

func NewReporterClient(cc *grpc.ClientConn) ReporterClient {
	return &reporterClient{cc}
}

func (c *reporterClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*pb.Response, error) {
	out := new(pb.Response)
	err := c.cc.Invoke(ctx, "/fs.base.reporter.Reporter/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReporterServer is the server API for Reporter service.
type ReporterServer interface {
	// 写入
	Write(context.Context, *WriteRequest) (*pb.Response, error)
}

func RegisterReporterServer(s *grpc.Server, srv ReporterServer) {
	s.RegisterService(&_Reporter_serviceDesc, srv)
}

func _Reporter_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReporterServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.reporter.Reporter/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReporterServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Reporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fs.base.reporter.Reporter",
	HandlerType: (*ReporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _Reporter_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/reporter/pb/reporter.proto",
}

func init() { proto.RegisterFile("base/reporter/pb/reporter.proto", fileDescriptor_aaba7a9b61ea1b4e) }

var fileDescriptor_aaba7a9b61ea1b4e = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0xdd, 0x2e, 0x6e, 0xf0, 0xb0, 0x06, 0x0f, 0x61, 0x11, 0x5d, 0x3c, 0x2d, 0x08,
	0x29, 0xe8, 0xcd, 0x07, 0xf0, 0x01, 0x72, 0xf1, 0x9c, 0x76, 0xa7, 0xb4, 0x2c, 0xed, 0xc4, 0x99,
	0xa9, 0x05, 0x1f, 0xc2, 0x67, 0x96, 0xa4, 0x5a, 0x65, 0x4f, 0xf9, 0xf2, 0xcf, 0x3f, 0x3f, 0xf3,
	0xab, 0xfb, 0xca, 0x33, 0x94, 0x04, 0x01, 0x49, 0x80, 0xca, 0x50, 0x2d, 0x6c, 0x03, 0xa1, 0xa0,
	0xde, 0x36, 0x6c, 0xa3, 0xc7, 0xfe, 0xea, 0xbb, 0xc7, 0x4f, 0x3e, 0x05, 0x4f, 0x27, 0x20, 0x5b,
	0x63, 0x5f, 0x36, 0x38, 0x0e, 0x47, 0x2f, 0x1d, 0x0e, 0x65, 0xca, 0x0a, 0x55, 0x7a, 0xe7, 0xf5,
	0x87, 0xaf, 0x4c, 0x5d, 0xbd, 0x51, 0x27, 0xe0, 0xe0, 0x7d, 0x04, 0x16, 0xbd, 0x55, 0x39, 0x7f,
	0xd4, 0x26, 0xdb, 0x67, 0x87, 0x8d, 0x8b, 0xa8, 0xb5, 0x5a, 0x35, 0xe3, 0x50, 0x9b, 0x8b, 0x24,
	0x25, 0x8e, 0xae, 0xa9, 0x45, 0x93, 0xcf, 0xae, 0xa9, 0x45, 0x7d, 0xab, 0x36, 0xd2, 0xf5, 0xc0,
	0xe2, 0xfb, 0x60, 0x56, 0xfb, 0xec, 0x90, 0xbb, 0x3f, 0x41, 0xdf, 0xa8, 0x62, 0x6a, 0x81, 0xc0,
	0x14, 0x69, 0x63, 0xfe, 0xc4, 0xe4, 0xa3, 0x17, 0x30, 0xeb, 0x39, 0x39, 0xf2, 0xd3, 0xab, 0xba,
	0x74, 0x3f, 0x4d, 0xf4, 0x8b, 0x2a, 0xd2, 0x6d, 0xfa, 0xce, 0x9e, 0xb7, 0xb4, 0xff, 0x8f, 0xde,
	0x5d, 0x2f, 0x73, 0x07, 0x1c, 0x70, 0x60, 0xa8, 0xd6, 0xa9, 0xdf, 0xf3, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc6, 0x4c, 0xc0, 0xbf, 0x41, 0x01, 0x00, 0x00,
}
