// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.3
// source: black_list.proto

package black_list

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BlackListService_Add_FullMethodName    = "/black_list.BlackListService/Add"
	BlackListService_GetAll_FullMethodName = "/black_list.BlackListService/GetAll"
	BlackListService_Remove_FullMethodName = "/black_list.BlackListService/Remove"
)

// BlackListServiceClient is the client API for BlackListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlackListServiceClient interface {
	Add(ctx context.Context, in *BlackListCreate, opts ...grpc.CallOption) (*Void, error)
	GetAll(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*GetAllBlackListRes, error)
	Remove(ctx context.Context, in *RemoveReq, opts ...grpc.CallOption) (*Void, error)
}

type blackListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlackListServiceClient(cc grpc.ClientConnInterface) BlackListServiceClient {
	return &blackListServiceClient{cc}
}

func (c *blackListServiceClient) Add(ctx context.Context, in *BlackListCreate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, BlackListService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blackListServiceClient) GetAll(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*GetAllBlackListRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllBlackListRes)
	err := c.cc.Invoke(ctx, BlackListService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blackListServiceClient) Remove(ctx context.Context, in *RemoveReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, BlackListService_Remove_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlackListServiceServer is the server API for BlackListService service.
// All implementations must embed UnimplementedBlackListServiceServer
// for forward compatibility.
type BlackListServiceServer interface {
	Add(context.Context, *BlackListCreate) (*Void, error)
	GetAll(context.Context, *Filter) (*GetAllBlackListRes, error)
	Remove(context.Context, *RemoveReq) (*Void, error)
	mustEmbedUnimplementedBlackListServiceServer()
}

// UnimplementedBlackListServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBlackListServiceServer struct{}

func (UnimplementedBlackListServiceServer) Add(context.Context, *BlackListCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedBlackListServiceServer) GetAll(context.Context, *Filter) (*GetAllBlackListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedBlackListServiceServer) Remove(context.Context, *RemoveReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedBlackListServiceServer) mustEmbedUnimplementedBlackListServiceServer() {}
func (UnimplementedBlackListServiceServer) testEmbeddedByValue()                          {}

// UnsafeBlackListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlackListServiceServer will
// result in compilation errors.
type UnsafeBlackListServiceServer interface {
	mustEmbedUnimplementedBlackListServiceServer()
}

func RegisterBlackListServiceServer(s grpc.ServiceRegistrar, srv BlackListServiceServer) {
	// If the following call pancis, it indicates UnimplementedBlackListServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BlackListService_ServiceDesc, srv)
}

func _BlackListService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlackListCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackListServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlackListService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackListServiceServer).Add(ctx, req.(*BlackListCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlackListService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackListServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlackListService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackListServiceServer).GetAll(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlackListService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackListServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlackListService_Remove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackListServiceServer).Remove(ctx, req.(*RemoveReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BlackListService_ServiceDesc is the grpc.ServiceDesc for BlackListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlackListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "black_list.BlackListService",
	HandlerType: (*BlackListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _BlackListService_Add_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _BlackListService_GetAll_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _BlackListService_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "black_list.proto",
}
