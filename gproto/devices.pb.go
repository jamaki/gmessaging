// Code generated by protoc-gen-go.
// source: devices.proto
// DO NOT EDIT!

/*
Package gproto is a generated protocol buffer package.

It is generated from these files:
	devices.proto

It has these top-level messages:
	Router
	Routers
	GetAllRequest
	GetByHostnameRequest
	RouterRequest
	RouterResponse
*/
package gproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Structs
type Router struct {
	Hostname string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	IP       []byte `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
}

func (m *Router) Reset()                    { *m = Router{} }
func (m *Router) String() string            { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()               {}
func (*Router) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Router) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Router) GetIP() []byte {
	if m != nil {
		return m.IP
	}
	return nil
}

type Routers struct {
	Router []*Router `protobuf:"bytes,1,rep,name=router" json:"router,omitempty"`
}

func (m *Routers) Reset()                    { *m = Routers{} }
func (m *Routers) String() string            { return proto.CompactTextString(m) }
func (*Routers) ProtoMessage()               {}
func (*Routers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Routers) GetRouter() []*Router {
	if m != nil {
		return m.Router
	}
	return nil
}

// Wrappers
type GetAllRequest struct {
}

func (m *GetAllRequest) Reset()                    { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()               {}
func (*GetAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GetByHostnameRequest struct {
	Hostname string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
}

func (m *GetByHostnameRequest) Reset()                    { *m = GetByHostnameRequest{} }
func (m *GetByHostnameRequest) String() string            { return proto.CompactTextString(m) }
func (*GetByHostnameRequest) ProtoMessage()               {}
func (*GetByHostnameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetByHostnameRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type RouterRequest struct {
	Router *Router `protobuf:"bytes,1,opt,name=router" json:"router,omitempty"`
}

func (m *RouterRequest) Reset()                    { *m = RouterRequest{} }
func (m *RouterRequest) String() string            { return proto.CompactTextString(m) }
func (*RouterRequest) ProtoMessage()               {}
func (*RouterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RouterRequest) GetRouter() *Router {
	if m != nil {
		return m.Router
	}
	return nil
}

type RouterResponse struct {
	Router *Router `protobuf:"bytes,1,opt,name=router" json:"router,omitempty"`
}

func (m *RouterResponse) Reset()                    { *m = RouterResponse{} }
func (m *RouterResponse) String() string            { return proto.CompactTextString(m) }
func (*RouterResponse) ProtoMessage()               {}
func (*RouterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RouterResponse) GetRouter() *Router {
	if m != nil {
		return m.Router
	}
	return nil
}

func init() {
	proto.RegisterType((*Router)(nil), "gproto.Router")
	proto.RegisterType((*Routers)(nil), "gproto.Routers")
	proto.RegisterType((*GetAllRequest)(nil), "gproto.GetAllRequest")
	proto.RegisterType((*GetByHostnameRequest)(nil), "gproto.GetByHostnameRequest")
	proto.RegisterType((*RouterRequest)(nil), "gproto.RouterRequest")
	proto.RegisterType((*RouterResponse)(nil), "gproto.RouterResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DeviceService service

type DeviceServiceClient interface {
	GetByHostname(ctx context.Context, in *GetByHostnameRequest, opts ...grpc.CallOption) (*RouterResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (DeviceService_GetAllClient, error)
	Save(ctx context.Context, in *RouterRequest, opts ...grpc.CallOption) (*RouterResponse, error)
	SaveAll(ctx context.Context, opts ...grpc.CallOption) (DeviceService_SaveAllClient, error)
}

type deviceServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceServiceClient(cc *grpc.ClientConn) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) GetByHostname(ctx context.Context, in *GetByHostnameRequest, opts ...grpc.CallOption) (*RouterResponse, error) {
	out := new(RouterResponse)
	err := grpc.Invoke(ctx, "/gproto.DeviceService/GetByHostname", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (DeviceService_GetAllClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DeviceService_serviceDesc.Streams[0], c.cc, "/gproto.DeviceService/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &deviceServiceGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DeviceService_GetAllClient interface {
	Recv() (*RouterResponse, error)
	grpc.ClientStream
}

type deviceServiceGetAllClient struct {
	grpc.ClientStream
}

func (x *deviceServiceGetAllClient) Recv() (*RouterResponse, error) {
	m := new(RouterResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *deviceServiceClient) Save(ctx context.Context, in *RouterRequest, opts ...grpc.CallOption) (*RouterResponse, error) {
	out := new(RouterResponse)
	err := grpc.Invoke(ctx, "/gproto.DeviceService/Save", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) SaveAll(ctx context.Context, opts ...grpc.CallOption) (DeviceService_SaveAllClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DeviceService_serviceDesc.Streams[1], c.cc, "/gproto.DeviceService/SaveAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &deviceServiceSaveAllClient{stream}
	return x, nil
}

type DeviceService_SaveAllClient interface {
	Send(*RouterRequest) error
	Recv() (*RouterResponse, error)
	grpc.ClientStream
}

type deviceServiceSaveAllClient struct {
	grpc.ClientStream
}

func (x *deviceServiceSaveAllClient) Send(m *RouterRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *deviceServiceSaveAllClient) Recv() (*RouterResponse, error) {
	m := new(RouterResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DeviceService service

type DeviceServiceServer interface {
	GetByHostname(context.Context, *GetByHostnameRequest) (*RouterResponse, error)
	GetAll(*GetAllRequest, DeviceService_GetAllServer) error
	Save(context.Context, *RouterRequest) (*RouterResponse, error)
	SaveAll(DeviceService_SaveAllServer) error
}

func RegisterDeviceServiceServer(s *grpc.Server, srv DeviceServiceServer) {
	s.RegisterService(&_DeviceService_serviceDesc, srv)
}

func _DeviceService_GetByHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByHostnameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetByHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gproto.DeviceService/GetByHostname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetByHostname(ctx, req.(*GetByHostnameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DeviceServiceServer).GetAll(m, &deviceServiceGetAllServer{stream})
}

type DeviceService_GetAllServer interface {
	Send(*RouterResponse) error
	grpc.ServerStream
}

type deviceServiceGetAllServer struct {
	grpc.ServerStream
}

func (x *deviceServiceGetAllServer) Send(m *RouterResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DeviceService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gproto.DeviceService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Save(ctx, req.(*RouterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_SaveAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DeviceServiceServer).SaveAll(&deviceServiceSaveAllServer{stream})
}

type DeviceService_SaveAllServer interface {
	Send(*RouterResponse) error
	Recv() (*RouterRequest, error)
	grpc.ServerStream
}

type deviceServiceSaveAllServer struct {
	grpc.ServerStream
}

func (x *deviceServiceSaveAllServer) Send(m *RouterResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *deviceServiceSaveAllServer) Recv() (*RouterRequest, error) {
	m := new(RouterRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DeviceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gproto.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByHostname",
			Handler:    _DeviceService_GetByHostname_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _DeviceService_Save_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _DeviceService_GetAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SaveAll",
			Handler:       _DeviceService_SaveAll_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "devices.proto",
}

func init() { proto.RegisterFile("devices.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x49, 0x2d, 0xcb,
	0x4c, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x07, 0xd3, 0x4a, 0x26,
	0x5c, 0x6c, 0x41, 0xf9, 0xa5, 0x25, 0xa9, 0x45, 0x42, 0x52, 0x5c, 0x1c, 0x19, 0xf9, 0xc5, 0x25,
	0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x10, 0x1f, 0x17,
	0x93, 0x67, 0x80, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x93, 0x67, 0x80, 0x92, 0x21, 0x17,
	0x3b, 0x44, 0x57, 0xb1, 0x90, 0x1a, 0x17, 0x5b, 0x11, 0x98, 0x29, 0xc1, 0xa8, 0xc0, 0xac, 0xc1,
	0x6d, 0xc4, 0xa7, 0x07, 0x31, 0x59, 0x0f, 0xa2, 0x20, 0x08, 0x2a, 0xab, 0xc4, 0xcf, 0xc5, 0xeb,
	0x9e, 0x5a, 0xe2, 0x98, 0x93, 0x13, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0xa2, 0x64, 0xc4, 0x25,
	0xe2, 0x9e, 0x5a, 0xe2, 0x54, 0xe9, 0x01, 0xb5, 0x04, 0x2a, 0x8e, 0xcf, 0x1d, 0x4a, 0xe6, 0x5c,
	0xbc, 0x50, 0x63, 0xa1, 0x8a, 0x91, 0x6d, 0x67, 0xc4, 0x63, 0xbb, 0x05, 0x17, 0x1f, 0x4c, 0x63,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0xb1, 0x3a, 0x8d, 0xda, 0x98, 0xb8, 0x78, 0x5d, 0xc0, 0x41,
	0x17, 0x9c, 0x5a, 0x04, 0xa2, 0x84, 0x5c, 0xc1, 0x3e, 0x41, 0x38, 0x5c, 0x48, 0x06, 0xa6, 0x15,
	0x9b, 0x7f, 0xa4, 0xc4, 0xd0, 0x0c, 0x86, 0x39, 0xc0, 0x92, 0x8b, 0x0d, 0x12, 0x20, 0x42, 0xa2,
	0x48, 0xfa, 0x11, 0x01, 0x84, 0x4b, 0xa3, 0x01, 0xa3, 0x90, 0x29, 0x17, 0x4b, 0x70, 0x62, 0x59,
	0x2a, 0x42, 0x23, 0x4a, 0xa0, 0xe0, 0xb4, 0xd1, 0x86, 0x8b, 0x1d, 0xa4, 0x0d, 0xc5, 0x4a, 0xa2,
	0x74, 0x6a, 0x30, 0x1a, 0x30, 0x26, 0xb1, 0x81, 0x65, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xaf, 0xaa, 0x18, 0xe9, 0x49, 0x02, 0x00, 0x00,
}
