// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cellayv1

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

// GamesServiceClient is the client API for GamesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GamesServiceClient interface {
	GetAll(ctx context.Context, in *GamesServiceGetAllRequest, opts ...grpc.CallOption) (*GamesServiceGetAllResponse, error)
}

type gamesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGamesServiceClient(cc grpc.ClientConnInterface) GamesServiceClient {
	return &gamesServiceClient{cc}
}

func (c *gamesServiceClient) GetAll(ctx context.Context, in *GamesServiceGetAllRequest, opts ...grpc.CallOption) (*GamesServiceGetAllResponse, error) {
	out := new(GamesServiceGetAllResponse)
	err := c.cc.Invoke(ctx, "/cellay.v1.GamesService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GamesServiceServer is the server API for GamesService service.
// All implementations must embed UnimplementedGamesServiceServer
// for forward compatibility
type GamesServiceServer interface {
	GetAll(context.Context, *GamesServiceGetAllRequest) (*GamesServiceGetAllResponse, error)
	mustEmbedUnimplementedGamesServiceServer()
}

// UnimplementedGamesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGamesServiceServer struct {
}

func (UnimplementedGamesServiceServer) GetAll(context.Context, *GamesServiceGetAllRequest) (*GamesServiceGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedGamesServiceServer) mustEmbedUnimplementedGamesServiceServer() {}

// UnsafeGamesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GamesServiceServer will
// result in compilation errors.
type UnsafeGamesServiceServer interface {
	mustEmbedUnimplementedGamesServiceServer()
}

func RegisterGamesServiceServer(s grpc.ServiceRegistrar, srv GamesServiceServer) {
	s.RegisterService(&GamesService_ServiceDesc, srv)
}

func _GamesService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GamesServiceGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GamesServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cellay.v1.GamesService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GamesServiceServer).GetAll(ctx, req.(*GamesServiceGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GamesService_ServiceDesc is the grpc.ServiceDesc for GamesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GamesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cellay.v1.GamesService",
	HandlerType: (*GamesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _GamesService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cellay/v1/cellay.proto",
}

// MatchesServiceClient is the client API for MatchesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchesServiceClient interface {
	Start(ctx context.Context, in *MatchesServiceStartRequest, opts ...grpc.CallOption) (*MatchesServiceStartResponse, error)
}

type matchesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchesServiceClient(cc grpc.ClientConnInterface) MatchesServiceClient {
	return &matchesServiceClient{cc}
}

func (c *matchesServiceClient) Start(ctx context.Context, in *MatchesServiceStartRequest, opts ...grpc.CallOption) (*MatchesServiceStartResponse, error) {
	out := new(MatchesServiceStartResponse)
	err := c.cc.Invoke(ctx, "/cellay.v1.MatchesService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchesServiceServer is the server API for MatchesService service.
// All implementations must embed UnimplementedMatchesServiceServer
// for forward compatibility
type MatchesServiceServer interface {
	Start(context.Context, *MatchesServiceStartRequest) (*MatchesServiceStartResponse, error)
	mustEmbedUnimplementedMatchesServiceServer()
}

// UnimplementedMatchesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMatchesServiceServer struct {
}

func (UnimplementedMatchesServiceServer) Start(context.Context, *MatchesServiceStartRequest) (*MatchesServiceStartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedMatchesServiceServer) mustEmbedUnimplementedMatchesServiceServer() {}

// UnsafeMatchesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchesServiceServer will
// result in compilation errors.
type UnsafeMatchesServiceServer interface {
	mustEmbedUnimplementedMatchesServiceServer()
}

func RegisterMatchesServiceServer(s grpc.ServiceRegistrar, srv MatchesServiceServer) {
	s.RegisterService(&MatchesService_ServiceDesc, srv)
}

func _MatchesService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchesServiceStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchesServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cellay.v1.MatchesService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchesServiceServer).Start(ctx, req.(*MatchesServiceStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MatchesService_ServiceDesc is the grpc.ServiceDesc for MatchesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cellay.v1.MatchesService",
	HandlerType: (*MatchesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _MatchesService_Start_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cellay/v1/cellay.proto",
}

// AssetsServiceClient is the client API for AssetsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetsServiceClient interface {
}

type assetsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetsServiceClient(cc grpc.ClientConnInterface) AssetsServiceClient {
	return &assetsServiceClient{cc}
}

// AssetsServiceServer is the server API for AssetsService service.
// All implementations must embed UnimplementedAssetsServiceServer
// for forward compatibility
type AssetsServiceServer interface {
	mustEmbedUnimplementedAssetsServiceServer()
}

// UnimplementedAssetsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssetsServiceServer struct {
}

func (UnimplementedAssetsServiceServer) mustEmbedUnimplementedAssetsServiceServer() {}

// UnsafeAssetsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetsServiceServer will
// result in compilation errors.
type UnsafeAssetsServiceServer interface {
	mustEmbedUnimplementedAssetsServiceServer()
}

func RegisterAssetsServiceServer(s grpc.ServiceRegistrar, srv AssetsServiceServer) {
	s.RegisterService(&AssetsService_ServiceDesc, srv)
}

// AssetsService_ServiceDesc is the grpc.ServiceDesc for AssetsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cellay.v1.AssetsService",
	HandlerType: (*AssetsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cellay/v1/cellay.proto",
}