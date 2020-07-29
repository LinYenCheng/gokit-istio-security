// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foo.proto

package pb

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

type FooRequest struct {
	S                    string   `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FooRequest) Reset()         { *m = FooRequest{} }
func (m *FooRequest) String() string { return proto.CompactTextString(m) }
func (*FooRequest) ProtoMessage()    {}
func (*FooRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{0}
}

func (m *FooRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FooRequest.Unmarshal(m, b)
}
func (m *FooRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FooRequest.Marshal(b, m, deterministic)
}
func (m *FooRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FooRequest.Merge(m, src)
}
func (m *FooRequest) XXX_Size() int {
	return xxx_messageInfo_FooRequest.Size(m)
}
func (m *FooRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FooRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FooRequest proto.InternalMessageInfo

func (m *FooRequest) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

type FooResponse struct {
	Res                  string   `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FooResponse) Reset()         { *m = FooResponse{} }
func (m *FooResponse) String() string { return proto.CompactTextString(m) }
func (*FooResponse) ProtoMessage()    {}
func (*FooResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{1}
}

func (m *FooResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FooResponse.Unmarshal(m, b)
}
func (m *FooResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FooResponse.Marshal(b, m, deterministic)
}
func (m *FooResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FooResponse.Merge(m, src)
}
func (m *FooResponse) XXX_Size() int {
	return xxx_messageInfo_FooResponse.Size(m)
}
func (m *FooResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FooResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FooResponse proto.InternalMessageInfo

func (m *FooResponse) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

func (m *FooResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*FooRequest)(nil), "pb.FooRequest")
	proto.RegisterType((*FooResponse)(nil), "pb.FooResponse")
}

func init() { proto.RegisterFile("foo.proto", fileDescriptor_7ce1e2eec643ca48) }

var fileDescriptor_7ce1e2eec643ca48 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xcb, 0xcf, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe2, 0xe2, 0x72, 0xcb, 0xcf,
	0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe1, 0x62, 0x2c, 0x96, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x62, 0x2c, 0x56, 0x32, 0xe4, 0xe2, 0x06, 0xcb, 0x15, 0x17, 0xe4, 0xe7, 0x15,
	0xa7, 0x0a, 0x09, 0x70, 0x31, 0x17, 0xa5, 0xc2, 0xa4, 0x41, 0x4c, 0x90, 0x48, 0x6a, 0x51, 0x91,
	0x04, 0x13, 0x44, 0x24, 0xb5, 0xa8, 0xc8, 0x48, 0x97, 0x8b, 0xd9, 0x2d, 0x3f, 0x5f, 0x48, 0x0d,
	0x42, 0xf1, 0xe9, 0x15, 0x24, 0xe9, 0x21, 0x8c, 0x97, 0xe2, 0x87, 0xf3, 0x21, 0x46, 0x26, 0xb1,
	0x81, 0x1d, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x67, 0x54, 0xbe, 0x7b, 0x95, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FooClient is the client API for Foo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FooClient interface {
	Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error)
}

type fooClient struct {
	cc *grpc.ClientConn
}

func NewFooClient(cc *grpc.ClientConn) FooClient {
	return &fooClient{cc}
}

func (c *fooClient) Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error) {
	out := new(FooResponse)
	err := c.cc.Invoke(ctx, "/pb.Foo/Foo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FooServer is the server API for Foo service.
type FooServer interface {
	Foo(context.Context, *FooRequest) (*FooResponse, error)
}

// UnimplementedFooServer can be embedded to have forward compatible implementations.
type UnimplementedFooServer struct {
}

func (*UnimplementedFooServer) Foo(ctx context.Context, req *FooRequest) (*FooResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Foo not implemented")
}

func RegisterFooServer(s *grpc.Server, srv FooServer) {
	s.RegisterService(&_Foo_serviceDesc, srv)
}

func _Foo_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServer).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Foo/Foo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServer).Foo(ctx, req.(*FooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Foo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Foo",
	HandlerType: (*FooServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Foo",
			Handler:    _Foo_Foo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foo.proto",
}
