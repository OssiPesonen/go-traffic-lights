// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/providers_app_service.proto

package providers_app

import (
	context "context"
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
	ProvidersAppService_ListProviders_FullMethodName          = "/proto.ProvidersAppService/ListProviders"
	ProvidersAppService_ReadProvider_FullMethodName           = "/proto.ProvidersAppService/ReadProvider"
	ProvidersAppService_CreateProvider_FullMethodName         = "/proto.ProvidersAppService/CreateProvider"
	ProvidersAppService_GetToken_FullMethodName               = "/proto.ProvidersAppService/GetToken"
	ProvidersAppService_RegisterUser_FullMethodName           = "/proto.ProvidersAppService/RegisterUser"
	ProvidersAppService_RefreshToken_FullMethodName           = "/proto.ProvidersAppService/RefreshToken"
	ProvidersAppService_RevokeRefreshToken_FullMethodName     = "/proto.ProvidersAppService/RevokeRefreshToken"
	ProvidersAppService_RevokeAllRefreshTokens_FullMethodName = "/proto.ProvidersAppService/RevokeAllRefreshTokens"
)

// ProvidersAppServiceClient is the client API for ProvidersAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProvidersAppServiceClient interface {
	// Providers
	ListProviders(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListProviderResponse, error)
	ReadProvider(ctx context.Context, in *ReadProviderRequest, opts ...grpc.CallOption) (*ReadProviderResponse, error)
	CreateProvider(ctx context.Context, in *CreateProviderRequest, opts ...grpc.CallOption) (*CreateProviderResponse, error)
	// Users
	GetToken(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	RegisterUser(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	RevokeRefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RevokeAllRefreshTokens(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type providersAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProvidersAppServiceClient(cc grpc.ClientConnInterface) ProvidersAppServiceClient {
	return &providersAppServiceClient{cc}
}

func (c *providersAppServiceClient) ListProviders(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListProviderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProviderResponse)
	err := c.cc.Invoke(ctx, ProvidersAppService_ListProviders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providersAppServiceClient) ReadProvider(ctx context.Context, in *ReadProviderRequest, opts ...grpc.CallOption) (*ReadProviderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadProviderResponse)
	err := c.cc.Invoke(ctx, ProvidersAppService_ReadProvider_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providersAppServiceClient) CreateProvider(ctx context.Context, in *CreateProviderRequest, opts ...grpc.CallOption) (*CreateProviderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProviderResponse)
	err := c.cc.Invoke(ctx, ProvidersAppService_CreateProvider_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providersAppServiceClient) GetToken(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, ProvidersAppService_GetToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providersAppServiceClient) RegisterUser(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProvidersAppService_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providersAppServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, ProvidersAppService_RefreshToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providersAppServiceClient) RevokeRefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProvidersAppService_RevokeRefreshToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providersAppServiceClient) RevokeAllRefreshTokens(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProvidersAppService_RevokeAllRefreshTokens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvidersAppServiceServer is the server API for ProvidersAppService service.
// All implementations must embed UnimplementedProvidersAppServiceServer
// for forward compatibility.
type ProvidersAppServiceServer interface {
	// Providers
	ListProviders(context.Context, *emptypb.Empty) (*ListProviderResponse, error)
	ReadProvider(context.Context, *ReadProviderRequest) (*ReadProviderResponse, error)
	CreateProvider(context.Context, *CreateProviderRequest) (*CreateProviderResponse, error)
	// Users
	GetToken(context.Context, *LoginRequest) (*TokenResponse, error)
	RegisterUser(context.Context, *RegistrationRequest) (*emptypb.Empty, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*TokenResponse, error)
	RevokeRefreshToken(context.Context, *RefreshTokenRequest) (*emptypb.Empty, error)
	RevokeAllRefreshTokens(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedProvidersAppServiceServer()
}

// UnimplementedProvidersAppServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProvidersAppServiceServer struct{}

func (UnimplementedProvidersAppServiceServer) ListProviders(context.Context, *emptypb.Empty) (*ListProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProviders not implemented")
}
func (UnimplementedProvidersAppServiceServer) ReadProvider(context.Context, *ReadProviderRequest) (*ReadProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadProvider not implemented")
}
func (UnimplementedProvidersAppServiceServer) CreateProvider(context.Context, *CreateProviderRequest) (*CreateProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProvider not implemented")
}
func (UnimplementedProvidersAppServiceServer) GetToken(context.Context, *LoginRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedProvidersAppServiceServer) RegisterUser(context.Context, *RegistrationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedProvidersAppServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedProvidersAppServiceServer) RevokeRefreshToken(context.Context, *RefreshTokenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeRefreshToken not implemented")
}
func (UnimplementedProvidersAppServiceServer) RevokeAllRefreshTokens(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeAllRefreshTokens not implemented")
}
func (UnimplementedProvidersAppServiceServer) mustEmbedUnimplementedProvidersAppServiceServer() {}
func (UnimplementedProvidersAppServiceServer) testEmbeddedByValue()                             {}

// UnsafeProvidersAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProvidersAppServiceServer will
// result in compilation errors.
type UnsafeProvidersAppServiceServer interface {
	mustEmbedUnimplementedProvidersAppServiceServer()
}

func RegisterProvidersAppServiceServer(s grpc.ServiceRegistrar, srv ProvidersAppServiceServer) {
	// If the following call pancis, it indicates UnimplementedProvidersAppServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProvidersAppService_ServiceDesc, srv)
}

func _ProvidersAppService_ListProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvidersAppServiceServer).ListProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProvidersAppService_ListProviders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvidersAppServiceServer).ListProviders(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvidersAppService_ReadProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvidersAppServiceServer).ReadProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProvidersAppService_ReadProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvidersAppServiceServer).ReadProvider(ctx, req.(*ReadProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvidersAppService_CreateProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvidersAppServiceServer).CreateProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProvidersAppService_CreateProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvidersAppServiceServer).CreateProvider(ctx, req.(*CreateProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvidersAppService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvidersAppServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProvidersAppService_GetToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvidersAppServiceServer).GetToken(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvidersAppService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvidersAppServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProvidersAppService_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvidersAppServiceServer).RegisterUser(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvidersAppService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvidersAppServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProvidersAppService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvidersAppServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvidersAppService_RevokeRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvidersAppServiceServer).RevokeRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProvidersAppService_RevokeRefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvidersAppServiceServer).RevokeRefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvidersAppService_RevokeAllRefreshTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvidersAppServiceServer).RevokeAllRefreshTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProvidersAppService_RevokeAllRefreshTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvidersAppServiceServer).RevokeAllRefreshTokens(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ProvidersAppService_ServiceDesc is the grpc.ServiceDesc for ProvidersAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProvidersAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProvidersAppService",
	HandlerType: (*ProvidersAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProviders",
			Handler:    _ProvidersAppService_ListProviders_Handler,
		},
		{
			MethodName: "ReadProvider",
			Handler:    _ProvidersAppService_ReadProvider_Handler,
		},
		{
			MethodName: "CreateProvider",
			Handler:    _ProvidersAppService_CreateProvider_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _ProvidersAppService_GetToken_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _ProvidersAppService_RegisterUser_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _ProvidersAppService_RefreshToken_Handler,
		},
		{
			MethodName: "RevokeRefreshToken",
			Handler:    _ProvidersAppService_RevokeRefreshToken_Handler,
		},
		{
			MethodName: "RevokeAllRefreshTokens",
			Handler:    _ProvidersAppService_RevokeAllRefreshTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/providers_app_service.proto",
}