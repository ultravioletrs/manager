// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package manager

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

// ManagerServiceClient is the client API for ManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerServiceClient interface {
	// rpc Health(HealthRequest) returns (HealthResponse) {}
	// rpc Run(RunRequest) returns (RunResponse) {}
	CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*CreateDomainResponse, error)
}

type managerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerServiceClient(cc grpc.ClientConnInterface) ManagerServiceClient {
	return &managerServiceClient{cc}
}

func (c *managerServiceClient) CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*CreateDomainResponse, error) {
	out := new(CreateDomainResponse)
	err := c.cc.Invoke(ctx, "/manager_proto.ManagerService/CreateDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServiceServer is the server API for ManagerService service.
// All implementations must embed UnimplementedManagerServiceServer
// for forward compatibility
type ManagerServiceServer interface {
	// rpc Health(HealthRequest) returns (HealthResponse) {}
	// rpc Run(RunRequest) returns (RunResponse) {}
	CreateDomain(context.Context, *CreateDomainRequest) (*CreateDomainResponse, error)
	mustEmbedUnimplementedManagerServiceServer()
}

// UnimplementedManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServiceServer struct {
}

func (UnimplementedManagerServiceServer) CreateDomain(context.Context, *CreateDomainRequest) (*CreateDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDomain not implemented")
}
func (UnimplementedManagerServiceServer) mustEmbedUnimplementedManagerServiceServer() {}

// UnsafeManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServiceServer will
// result in compilation errors.
type UnsafeManagerServiceServer interface {
	mustEmbedUnimplementedManagerServiceServer()
}

func RegisterManagerServiceServer(s grpc.ServiceRegistrar, srv ManagerServiceServer) {
	s.RegisterService(&ManagerService_ServiceDesc, srv)
}

func _ManagerService_CreateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).CreateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_proto.ManagerService/CreateDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).CreateDomain(ctx, req.(*CreateDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagerService_ServiceDesc is the grpc.ServiceDesc for ManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager_proto.ManagerService",
	HandlerType: (*ManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDomain",
			Handler:    _ManagerService_CreateDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/manager.proto",
}
