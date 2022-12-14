// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

package stream

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

type Command struct {
	In                   string   `protobuf:"bytes,1,opt,name=In,proto3" json:"In,omitempty"`
	Out                  string   `protobuf:"bytes,2,opt,name=Out,proto3" json:"Out,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{0}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetIn() string {
	if m != nil {
		return m.In
	}
	return ""
}

func (m *Command) GetOut() string {
	if m != nil {
		return m.Out
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{1}
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
	proto.RegisterType((*Command)(nil), "stream.Command")
	proto.RegisterType((*Empty)(nil), "stream.Empty")
}

func init() {
	proto.RegisterFile("stream.proto", fileDescriptor_bb17ef3f514bfe54)
}

var fileDescriptor_bb17ef3f514bfe54 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xb4, 0xb9, 0xd8,
	0x9d, 0xf3, 0x73, 0x73, 0x13, 0xf3, 0x52, 0x84, 0xf8, 0xb8, 0x98, 0x3c, 0xf3, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x98, 0x3c, 0xf3, 0x84, 0x04, 0xb8, 0x98, 0xfd, 0x4b, 0x4b, 0x24, 0x98,
	0xc0, 0x02, 0x20, 0xa6, 0x12, 0x3b, 0x17, 0xab, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x51, 0x3a, 0x17,
	0xbb, 0x67, 0x6e, 0x41, 0x4e, 0x62, 0x5e, 0x89, 0x90, 0x1e, 0x17, 0x8f, 0x5b, 0x6a, 0x49, 0x72,
	0x06, 0xcc, 0x14, 0x5e, 0x3d, 0xa8, 0x3d, 0x60, 0x95, 0x52, 0xfc, 0x30, 0x2e, 0x4c, 0x5e, 0x87,
	0x8b, 0x2b, 0x38, 0x35, 0x2f, 0xc5, 0xbf, 0xb4, 0xa4, 0xa0, 0xb4, 0x44, 0x08, 0x5d, 0x5a, 0x0a,
	0x55, 0xbb, 0x91, 0x39, 0x17, 0xab, 0x63, 0x4a, 0x6e, 0x66, 0x9e, 0x90, 0x1e, 0x17, 0x57, 0x50,
	0x69, 0x1e, 0xcc, 0x10, 0x0c, 0x6d, 0xe8, 0x02, 0x49, 0x6c, 0x60, 0x6f, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xb3, 0xe2, 0x85, 0x8f, 0xf6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ImplantClient is the client API for Implant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImplantClient interface {
	FetchCommand(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Command, error)
	SendOutput(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Empty, error)
}

type implantClient struct {
	cc grpc.ClientConnInterface
}

func NewImplantClient(cc grpc.ClientConnInterface) ImplantClient {
	return &implantClient{cc}
}

func (c *implantClient) FetchCommand(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Command, error) {
	out := new(Command)
	err := c.cc.Invoke(ctx, "/stream.Implant/FetchCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *implantClient) SendOutput(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/stream.Implant/SendOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImplantServer is the server API for Implant service.
type ImplantServer interface {
	FetchCommand(context.Context, *Empty) (*Command, error)
	SendOutput(context.Context, *Command) (*Empty, error)
}

// UnimplementedImplantServer can be embedded to have forward compatible implementations.
type UnimplementedImplantServer struct {
}

func (*UnimplementedImplantServer) FetchCommand(ctx context.Context, req *Empty) (*Command, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCommand not implemented")
}
func (*UnimplementedImplantServer) SendOutput(ctx context.Context, req *Command) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOutput not implemented")
}

func RegisterImplantServer(s *grpc.Server, srv ImplantServer) {
	s.RegisterService(&_Implant_serviceDesc, srv)
}

func _Implant_FetchCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImplantServer).FetchCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.Implant/FetchCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImplantServer).FetchCommand(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Implant_SendOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImplantServer).SendOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.Implant/SendOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImplantServer).SendOutput(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

var _Implant_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stream.Implant",
	HandlerType: (*ImplantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchCommand",
			Handler:    _Implant_FetchCommand_Handler,
		},
		{
			MethodName: "SendOutput",
			Handler:    _Implant_SendOutput_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stream.proto",
}

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminClient interface {
	RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Command, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Command, error) {
	out := new(Command)
	err := c.cc.Invoke(ctx, "/stream.Admin/RunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
type AdminServer interface {
	RunCommand(context.Context, *Command) (*Command, error)
}

// UnimplementedAdminServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (*UnimplementedAdminServer) RunCommand(ctx context.Context, req *Command) (*Command, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCommand not implemented")
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_RunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).RunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.Admin/RunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).RunCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stream.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunCommand",
			Handler:    _Admin_RunCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stream.proto",
}
