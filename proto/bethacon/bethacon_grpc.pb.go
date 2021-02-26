// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bethaconproto

import (
	context "context"

	empty "github.com/golang/protobuf/ptypes/empty"
	health "github.com/indiealistic/bethacon/proto/health"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BethaConServiceClient is the client API for BethaConService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BethaConServiceClient interface {
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*health.HealthResponse, error)
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// Attester
	GetAttestation(ctx context.Context, in *GetAttestationRequest, opts ...grpc.CallOption) (*GetAttestationResponse, error)
	ProposeAttestation(ctx context.Context, in *ProposeAttestationRequest, opts ...grpc.CallOption) (*ProposeAttestationResponse, error)
}

type bethaConServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBethaConServiceClient(cc grpc.ClientConnInterface) BethaConServiceClient {
	return &bethaConServiceClient{cc}
}

func (c *bethaConServiceClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*health.HealthResponse, error) {
	out := new(health.HealthResponse)
	err := c.cc.Invoke(ctx, "/grpcapiproto.BethaConService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bethaConServiceClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpcapiproto.BethaConService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bethaConServiceClient) GetAttestation(ctx context.Context, in *GetAttestationRequest, opts ...grpc.CallOption) (*GetAttestationResponse, error) {
	out := new(GetAttestationResponse)
	err := c.cc.Invoke(ctx, "/grpcapiproto.BethaConService/GetAttestation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bethaConServiceClient) ProposeAttestation(ctx context.Context, in *ProposeAttestationRequest, opts ...grpc.CallOption) (*ProposeAttestationResponse, error) {
	out := new(ProposeAttestationResponse)
	err := c.cc.Invoke(ctx, "/grpcapiproto.BethaConService/ProposeAttestation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BethaConServiceServer is the server API for BethaConService service.
// All implementations must embed UnimplementedBethaConServiceServer
// for forward compatibility
type BethaConServiceServer interface {
	Health(context.Context, *empty.Empty) (*health.HealthResponse, error)
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	// Attester
	GetAttestation(context.Context, *GetAttestationRequest) (*GetAttestationResponse, error)
	ProposeAttestation(context.Context, *ProposeAttestationRequest) (*ProposeAttestationResponse, error)
	mustEmbedUnimplementedBethaConServiceServer()
}

// UnimplementedBethaConServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBethaConServiceServer struct {
}

func (UnimplementedBethaConServiceServer) Health(context.Context, *empty.Empty) (*health.HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedBethaConServiceServer) Ping(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBethaConServiceServer) GetAttestation(context.Context, *GetAttestationRequest) (*GetAttestationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttestation not implemented")
}
func (UnimplementedBethaConServiceServer) ProposeAttestation(context.Context, *ProposeAttestationRequest) (*ProposeAttestationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProposeAttestation not implemented")
}
func (UnimplementedBethaConServiceServer) mustEmbedUnimplementedBethaConServiceServer() {}

// UnsafeBethaConServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BethaConServiceServer will
// result in compilation errors.
type UnsafeBethaConServiceServer interface {
	mustEmbedUnimplementedBethaConServiceServer()
}

func RegisterBethaConServiceServer(s grpc.ServiceRegistrar, srv BethaConServiceServer) {
	s.RegisterService(&BethaConService_ServiceDesc, srv)
}

func _BethaConService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BethaConServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapiproto.BethaConService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BethaConServiceServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BethaConService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BethaConServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapiproto.BethaConService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BethaConServiceServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BethaConService_GetAttestation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttestationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BethaConServiceServer).GetAttestation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapiproto.BethaConService/GetAttestation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BethaConServiceServer).GetAttestation(ctx, req.(*GetAttestationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BethaConService_ProposeAttestation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposeAttestationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BethaConServiceServer).ProposeAttestation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapiproto.BethaConService/ProposeAttestation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BethaConServiceServer).ProposeAttestation(ctx, req.(*ProposeAttestationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BethaConService_ServiceDesc is the grpc.ServiceDesc for BethaConService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BethaConService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapiproto.BethaConService",
	HandlerType: (*BethaConServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _BethaConService_Health_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _BethaConService_Ping_Handler,
		},
		{
			MethodName: "GetAttestation",
			Handler:    _BethaConService_GetAttestation_Handler,
		},
		{
			MethodName: "ProposeAttestation",
			Handler:    _BethaConService_ProposeAttestation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/indiealistic/bethacon/proto/bethacon/bethacon.proto",
}
