// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.2
// source: proto/traffic_lights_service.proto

package go_traffic_lights

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TrafficLightsServiceClient is the client API for TrafficLightsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrafficLightsServiceClient interface {
	// Providers
	ListProviders(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListProviderResponse, error)
	ReadProvider(ctx context.Context, in *ReadProviderRequest, opts ...grpc.CallOption) (*ReadProviderResponse, error)
	// Users
	GetToken(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	RegisterUser(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	RevokeRefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RevokeAllRefreshToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type trafficLightsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrafficLightsServiceClient(cc grpc.ClientConnInterface) TrafficLightsServiceClient {
	return &trafficLightsServiceClient{cc}
}

func (c *trafficLightsServiceClient) ListProviders(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListProviderResponse, error) {
	out := new(ListProviderResponse)
	err := c.cc.Invoke(ctx, "/proto.TrafficLightsService/ListProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trafficLightsServiceClient) ReadProvider(ctx context.Context, in *ReadProviderRequest, opts ...grpc.CallOption) (*ReadProviderResponse, error) {
	out := new(ReadProviderResponse)
	err := c.cc.Invoke(ctx, "/proto.TrafficLightsService/ReadProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trafficLightsServiceClient) GetToken(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/proto.TrafficLightsService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trafficLightsServiceClient) RegisterUser(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.TrafficLightsService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trafficLightsServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/proto.TrafficLightsService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trafficLightsServiceClient) RevokeRefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.TrafficLightsService/RevokeRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trafficLightsServiceClient) RevokeAllRefreshToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.TrafficLightsService/RevokeAllRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrafficLightsServiceServer is the server API for TrafficLightsService service.
// All implementations must embed UnimplementedTrafficLightsServiceServer
// for forward compatibility
type TrafficLightsServiceServer interface {
	// Providers
	ListProviders(context.Context, *emptypb.Empty) (*ListProviderResponse, error)
	ReadProvider(context.Context, *ReadProviderRequest) (*ReadProviderResponse, error)
	// Users
	GetToken(context.Context, *LoginRequest) (*TokenResponse, error)
	RegisterUser(context.Context, *RegistrationRequest) (*emptypb.Empty, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*TokenResponse, error)
	RevokeRefreshToken(context.Context, *RefreshTokenRequest) (*emptypb.Empty, error)
	RevokeAllRefreshToken(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedTrafficLightsServiceServer()
}

// UnimplementedTrafficLightsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrafficLightsServiceServer struct {
}

func (UnimplementedTrafficLightsServiceServer) ListProviders(context.Context, *emptypb.Empty) (*ListProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProviders not implemented")
}
func (UnimplementedTrafficLightsServiceServer) ReadProvider(context.Context, *ReadProviderRequest) (*ReadProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadProvider not implemented")
}
func (UnimplementedTrafficLightsServiceServer) GetToken(context.Context, *LoginRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedTrafficLightsServiceServer) RegisterUser(context.Context, *RegistrationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedTrafficLightsServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedTrafficLightsServiceServer) RevokeRefreshToken(context.Context, *RefreshTokenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeRefreshToken not implemented")
}
func (UnimplementedTrafficLightsServiceServer) RevokeAllRefreshToken(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeAllRefreshToken not implemented")
}
func (UnimplementedTrafficLightsServiceServer) mustEmbedUnimplementedTrafficLightsServiceServer() {}

// UnsafeTrafficLightsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrafficLightsServiceServer will
// result in compilation errors.
type UnsafeTrafficLightsServiceServer interface {
	mustEmbedUnimplementedTrafficLightsServiceServer()
}

func RegisterTrafficLightsServiceServer(s grpc.ServiceRegistrar, srv TrafficLightsServiceServer) {
	s.RegisterService(&TrafficLightsService_ServiceDesc, srv)
}

func _TrafficLightsService_ListProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrafficLightsServiceServer).ListProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TrafficLightsService/ListProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrafficLightsServiceServer).ListProviders(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrafficLightsService_ReadProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrafficLightsServiceServer).ReadProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TrafficLightsService/ReadProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrafficLightsServiceServer).ReadProvider(ctx, req.(*ReadProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrafficLightsService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrafficLightsServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TrafficLightsService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrafficLightsServiceServer).GetToken(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrafficLightsService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrafficLightsServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TrafficLightsService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrafficLightsServiceServer).RegisterUser(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrafficLightsService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrafficLightsServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TrafficLightsService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrafficLightsServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrafficLightsService_RevokeRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrafficLightsServiceServer).RevokeRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TrafficLightsService/RevokeRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrafficLightsServiceServer).RevokeRefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrafficLightsService_RevokeAllRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrafficLightsServiceServer).RevokeAllRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TrafficLightsService/RevokeAllRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrafficLightsServiceServer).RevokeAllRefreshToken(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TrafficLightsService_ServiceDesc is the grpc.ServiceDesc for TrafficLightsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrafficLightsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TrafficLightsService",
	HandlerType: (*TrafficLightsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProviders",
			Handler:    _TrafficLightsService_ListProviders_Handler,
		},
		{
			MethodName: "ReadProvider",
			Handler:    _TrafficLightsService_ReadProvider_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _TrafficLightsService_GetToken_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _TrafficLightsService_RegisterUser_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _TrafficLightsService_RefreshToken_Handler,
		},
		{
			MethodName: "RevokeRefreshToken",
			Handler:    _TrafficLightsService_RevokeRefreshToken_Handler,
		},
		{
			MethodName: "RevokeAllRefreshToken",
			Handler:    _TrafficLightsService_RevokeAllRefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/traffic_lights_service.proto",
}
