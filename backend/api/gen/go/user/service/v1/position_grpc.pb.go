// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: user/service/v1/position.proto

package servicev1

import (
	context "context"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PositionService_ListPosition_FullMethodName   = "/user.service.v1.PositionService/ListPosition"
	PositionService_GetPosition_FullMethodName    = "/user.service.v1.PositionService/GetPosition"
	PositionService_CreatePosition_FullMethodName = "/user.service.v1.PositionService/CreatePosition"
	PositionService_UpdatePosition_FullMethodName = "/user.service.v1.PositionService/UpdatePosition"
	PositionService_DeletePosition_FullMethodName = "/user.service.v1.PositionService/DeletePosition"
)

// PositionServiceClient is the client API for PositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 职位服务
type PositionServiceClient interface {
	// 查询职位列表
	ListPosition(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*ListPositionResponse, error)
	// 查询职位详情
	GetPosition(ctx context.Context, in *GetPositionRequest, opts ...grpc.CallOption) (*Position, error)
	// 创建职位
	CreatePosition(ctx context.Context, in *CreatePositionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新职位
	UpdatePosition(ctx context.Context, in *UpdatePositionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除职位
	DeletePosition(ctx context.Context, in *DeletePositionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type positionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionServiceClient(cc grpc.ClientConnInterface) PositionServiceClient {
	return &positionServiceClient{cc}
}

func (c *positionServiceClient) ListPosition(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*ListPositionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPositionResponse)
	err := c.cc.Invoke(ctx, PositionService_ListPosition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) GetPosition(ctx context.Context, in *GetPositionRequest, opts ...grpc.CallOption) (*Position, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Position)
	err := c.cc.Invoke(ctx, PositionService_GetPosition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) CreatePosition(ctx context.Context, in *CreatePositionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PositionService_CreatePosition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) UpdatePosition(ctx context.Context, in *UpdatePositionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PositionService_UpdatePosition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) DeletePosition(ctx context.Context, in *DeletePositionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PositionService_DeletePosition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionServiceServer is the server API for PositionService service.
// All implementations must embed UnimplementedPositionServiceServer
// for forward compatibility.
//
// 职位服务
type PositionServiceServer interface {
	// 查询职位列表
	ListPosition(context.Context, *v1.PagingRequest) (*ListPositionResponse, error)
	// 查询职位详情
	GetPosition(context.Context, *GetPositionRequest) (*Position, error)
	// 创建职位
	CreatePosition(context.Context, *CreatePositionRequest) (*emptypb.Empty, error)
	// 更新职位
	UpdatePosition(context.Context, *UpdatePositionRequest) (*emptypb.Empty, error)
	// 删除职位
	DeletePosition(context.Context, *DeletePositionRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPositionServiceServer()
}

// UnimplementedPositionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPositionServiceServer struct{}

func (UnimplementedPositionServiceServer) ListPosition(context.Context, *v1.PagingRequest) (*ListPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPosition not implemented")
}
func (UnimplementedPositionServiceServer) GetPosition(context.Context, *GetPositionRequest) (*Position, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPosition not implemented")
}
func (UnimplementedPositionServiceServer) CreatePosition(context.Context, *CreatePositionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePosition not implemented")
}
func (UnimplementedPositionServiceServer) UpdatePosition(context.Context, *UpdatePositionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePosition not implemented")
}
func (UnimplementedPositionServiceServer) DeletePosition(context.Context, *DeletePositionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePosition not implemented")
}
func (UnimplementedPositionServiceServer) mustEmbedUnimplementedPositionServiceServer() {}
func (UnimplementedPositionServiceServer) testEmbeddedByValue()                         {}

// UnsafePositionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionServiceServer will
// result in compilation errors.
type UnsafePositionServiceServer interface {
	mustEmbedUnimplementedPositionServiceServer()
}

func RegisterPositionServiceServer(s grpc.ServiceRegistrar, srv PositionServiceServer) {
	// If the following call pancis, it indicates UnimplementedPositionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PositionService_ServiceDesc, srv)
}

func _PositionService_ListPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).ListPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PositionService_ListPosition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).ListPosition(ctx, req.(*v1.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_GetPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).GetPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PositionService_GetPosition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).GetPosition(ctx, req.(*GetPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_CreatePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).CreatePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PositionService_CreatePosition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).CreatePosition(ctx, req.(*CreatePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_UpdatePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).UpdatePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PositionService_UpdatePosition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).UpdatePosition(ctx, req.(*UpdatePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_DeletePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).DeletePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PositionService_DeletePosition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).DeletePosition(ctx, req.(*DeletePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PositionService_ServiceDesc is the grpc.ServiceDesc for PositionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PositionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.service.v1.PositionService",
	HandlerType: (*PositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPosition",
			Handler:    _PositionService_ListPosition_Handler,
		},
		{
			MethodName: "GetPosition",
			Handler:    _PositionService_GetPosition_Handler,
		},
		{
			MethodName: "CreatePosition",
			Handler:    _PositionService_CreatePosition_Handler,
		},
		{
			MethodName: "UpdatePosition",
			Handler:    _PositionService_UpdatePosition_Handler,
		},
		{
			MethodName: "DeletePosition",
			Handler:    _PositionService_DeletePosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/service/v1/position.proto",
}
