// Code generated by protoc-gen-go. DO NOT EDIT.
// source: posts.proto

package grpc_api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_b14bd1586479c33d, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() { proto.RegisterFile("posts.proto", fileDescriptor_b14bd1586479c33d) }

var fileDescriptor_b14bd1586479c33d = []byte{
	// 88 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xc8, 0x2f, 0x2e,
	0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x62, 0xe7, 0x62, 0x75, 0xcd, 0x2d, 0x28, 0xa9,
	0x34, 0x52, 0xe1, 0x62, 0x0d, 0x00, 0x89, 0x0b, 0x49, 0x73, 0xb1, 0xbb, 0xa7, 0x96, 0x80, 0xd8,
	0x42, 0x6c, 0x7a, 0x60, 0x39, 0x29, 0x28, 0xad, 0xc4, 0xe0, 0xc4, 0x15, 0xc5, 0x91, 0x5e, 0x54,
	0x90, 0x1c, 0x9f, 0x58, 0x90, 0x99, 0xc4, 0x06, 0x36, 0xc1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xce, 0x39, 0x18, 0x86, 0x50, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostsClient is the client API for Posts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostsClient interface {
	GetPost(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type postsClient struct {
	cc *grpc.ClientConn
}

func NewPostsClient(cc *grpc.ClientConn) PostsClient {
	return &postsClient{cc}
}

func (c *postsClient) GetPost(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Posts/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostsServer is the server API for Posts service.
type PostsServer interface {
	GetPost(context.Context, *Empty) (*Empty, error)
}

// UnimplementedPostsServer can be embedded to have forward compatible implementations.
type UnimplementedPostsServer struct {
}

func (*UnimplementedPostsServer) GetPost(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}

func RegisterPostsServer(s *grpc.Server, srv PostsServer) {
	s.RegisterService(&_Posts_serviceDesc, srv)
}

func _Posts_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Posts/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).GetPost(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Posts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Posts",
	HandlerType: (*PostsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPost",
			Handler:    _Posts_GetPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "posts.proto",
}