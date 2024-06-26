// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: pkg/grpc/server.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NodeISPServiceClient is the client API for NodeISPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeISPServiceClient interface {
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
}

type nodeISPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeISPServiceClient(cc grpc.ClientConnInterface) NodeISPServiceClient {
	return &nodeISPServiceClient{cc}
}

func (c *nodeISPServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/grpc.NodeISPService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeISPServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/grpc.NodeISPService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeISPServiceServer is the server API for NodeISPService service.
// All implementations must embed UnimplementedNodeISPServiceServer
// for forward compatibility
type NodeISPServiceServer interface {
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
	mustEmbedUnimplementedNodeISPServiceServer()
}

// UnimplementedNodeISPServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeISPServiceServer struct {
}

func (UnimplementedNodeISPServiceServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedNodeISPServiceServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedNodeISPServiceServer) mustEmbedUnimplementedNodeISPServiceServer() {}

// UnsafeNodeISPServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeISPServiceServer will
// result in compilation errors.
type UnsafeNodeISPServiceServer interface {
	mustEmbedUnimplementedNodeISPServiceServer()
}

func RegisterNodeISPServiceServer(s grpc.ServiceRegistrar, srv NodeISPServiceServer) {
	s.RegisterService(&NodeISPService_ServiceDesc, srv)
}

func _NodeISPService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeISPServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.NodeISPService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeISPServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeISPService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeISPServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.NodeISPService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeISPServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeISPService_ServiceDesc is the grpc.ServiceDesc for NodeISPService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeISPService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.NodeISPService",
	HandlerType: (*NodeISPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _NodeISPService_GetStatus_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _NodeISPService_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/server.proto",
}
