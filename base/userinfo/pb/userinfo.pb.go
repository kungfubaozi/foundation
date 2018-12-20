// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base/userinfo/pb/userinfo.proto

package fs_base_userinfo

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
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df1661b7c2e51838, []int{0}
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

func (m *GetRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetResponse struct {
	State                *pb.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	UserId               string    `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	NickName             string    `protobuf:"bytes,3,opt,name=nickName,proto3" json:"nickName,omitempty"`
	RealName             string    `protobuf:"bytes,4,opt,name=realName,proto3" json:"realName,omitempty"`
	Age                  int64     `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	Icon                 string    `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df1661b7c2e51838, []int{1}
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

func (m *GetResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetResponse) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *GetResponse) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *GetResponse) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *GetResponse) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

type UserBaseInfo struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NickName             string   `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	RealName             string   `protobuf:"bytes,3,opt,name=realName,proto3" json:"realName,omitempty"`
	Age                  int64    `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Icon                 string   `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	Sex                  int64    `protobuf:"varint,6,opt,name=sex,proto3" json:"sex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserBaseInfo) Reset()         { *m = UserBaseInfo{} }
func (m *UserBaseInfo) String() string { return proto.CompactTextString(m) }
func (*UserBaseInfo) ProtoMessage()    {}
func (*UserBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_df1661b7c2e51838, []int{2}
}

func (m *UserBaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserBaseInfo.Unmarshal(m, b)
}
func (m *UserBaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserBaseInfo.Marshal(b, m, deterministic)
}
func (m *UserBaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBaseInfo.Merge(m, src)
}
func (m *UserBaseInfo) XXX_Size() int {
	return xxx_messageInfo_UserBaseInfo.Size(m)
}
func (m *UserBaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserBaseInfo proto.InternalMessageInfo

func (m *UserBaseInfo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserBaseInfo) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *UserBaseInfo) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *UserBaseInfo) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserBaseInfo) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *UserBaseInfo) GetSex() int64 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "fs.base.userinfo.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "fs.base.userinfo.GetResponse")
	proto.RegisterType((*UserBaseInfo)(nil), "fs.base.userinfo.UserBaseInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserInfoClient is the client API for UserInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserInfoClient interface {
	Upsert(ctx context.Context, in *UserBaseInfo, opts ...grpc.CallOption) (*pb.Response, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type userInfoClient struct {
	cc *grpc.ClientConn
}

func NewUserInfoClient(cc *grpc.ClientConn) UserInfoClient {
	return &userInfoClient{cc}
}

func (c *userInfoClient) Upsert(ctx context.Context, in *UserBaseInfo, opts ...grpc.CallOption) (*pb.Response, error) {
	out := new(pb.Response)
	err := c.cc.Invoke(ctx, "/fs.base.userinfo.UserInfo/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/fs.base.userinfo.UserInfo/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServer is the server API for UserInfo service.
type UserInfoServer interface {
	Upsert(context.Context, *UserBaseInfo) (*pb.Response, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterUserInfoServer(s *grpc.Server, srv UserInfoServer) {
	s.RegisterService(&_UserInfo_serviceDesc, srv)
}

func _UserInfo_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBaseInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.userinfo.UserInfo/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).Upsert(ctx, req.(*UserBaseInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfo_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.base.userinfo.UserInfo/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fs.base.userinfo.UserInfo",
	HandlerType: (*UserInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upsert",
			Handler:    _UserInfo_Upsert_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserInfo_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/userinfo/pb/userinfo.proto",
}

func init() { proto.RegisterFile("base/userinfo/pb/userinfo.proto", fileDescriptor_df1661b7c2e51838) }

var fileDescriptor_df1661b7c2e51838 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0x53, 0x0a, 0x0d, 0x5e, 0x8c, 0xc1, 0x59, 0x98, 0xa6, 0xf1, 0x87, 0x10, 0x16, 0x24,
	0x26, 0xd3, 0x04, 0x97, 0xee, 0xd8, 0x10, 0x36, 0x2e, 0x6a, 0x78, 0x80, 0x29, 0xdc, 0x9a, 0x06,
	0x99, 0x19, 0xe7, 0x0e, 0x89, 0xf1, 0x15, 0xdc, 0xfb, 0x0e, 0xbe, 0xa5, 0x99, 0xa9, 0x94, 0x2a,
	0xe2, 0xaa, 0xe7, 0xce, 0x39, 0x73, 0xf2, 0x75, 0x2e, 0xdc, 0xe4, 0x82, 0x30, 0xdd, 0x12, 0x9a,
	0x52, 0x16, 0x2a, 0xd5, 0x79, 0xad, 0xb9, 0x36, 0xca, 0x2a, 0xd6, 0x2f, 0x88, 0xbb, 0x0c, 0xdf,
	0x9d, 0x27, 0xb7, 0x6f, 0xb4, 0xd6, 0xc2, 0xac, 0xd1, 0xf0, 0xa5, 0xda, 0xa4, 0x85, 0xda, 0xca,
	0x95, 0xb0, 0xa5, 0x92, 0xa9, 0xef, 0xd2, 0xb9, 0xff, 0x56, 0xd7, 0x87, 0x23, 0x80, 0x19, 0xda,
	0x0c, 0x5f, 0xb6, 0x48, 0x96, 0x5d, 0x40, 0xe4, 0x6a, 0xe6, 0xab, 0x38, 0x18, 0x04, 0xe3, 0x93,
	0xec, 0x7b, 0x1a, 0x7e, 0x06, 0xd0, 0xf3, 0x31, 0xd2, 0x4a, 0x12, 0xb2, 0x11, 0x74, 0xc8, 0x0a,
	0x8b, 0x3e, 0xd6, 0x9b, 0x9c, 0xf1, 0x1d, 0xc4, 0xa3, 0x3b, 0xcd, 0x2a, 0xb3, 0xd1, 0xd6, 0x6a,
	0xb6, 0xb1, 0x04, 0xba, 0xb2, 0x5c, 0xae, 0x1f, 0xc4, 0x06, 0xe3, 0xd0, 0x3b, 0xf5, 0xec, 0x3c,
	0x83, 0xe2, 0xd9, 0x7b, 0xed, 0xca, 0xdb, 0xcd, 0xac, 0x0f, 0xa1, 0x78, 0xc2, 0xb8, 0x33, 0x08,
	0xc6, 0x61, 0xe6, 0x24, 0x63, 0xd0, 0x2e, 0x97, 0x4a, 0xc6, 0x91, 0x4f, 0x7a, 0x3d, 0xfc, 0x08,
	0xe0, 0x74, 0x41, 0x68, 0xa6, 0x82, 0x70, 0x2e, 0x0b, 0x75, 0xec, 0xa7, 0x7e, 0x60, 0xb4, 0xfe,
	0xc1, 0x08, 0xff, 0xc6, 0x68, 0x1f, 0x62, 0x74, 0xf6, 0x18, 0x2e, 0x45, 0xf8, 0xea, 0xc9, 0xc2,
	0xcc, 0xc9, 0xc9, 0x7b, 0x00, 0x5d, 0x07, 0xe6, 0xa1, 0xee, 0x21, 0x5a, 0x68, 0x42, 0x63, 0xd9,
	0x35, 0xff, 0xbd, 0x41, 0xde, 0xc4, 0x4f, 0xce, 0x6b, 0xbf, 0x7e, 0xfe, 0x29, 0x84, 0x33, 0xb4,
	0xec, 0xf2, 0xf0, 0xe6, 0x7e, 0x97, 0xc9, 0xd5, 0x11, 0xb7, 0xea, 0xc8, 0x23, 0xbf, 0xff, 0xbb,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x22, 0xaf, 0x41, 0x61, 0x02, 0x00, 0x00,
}
