// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bancc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BanccClient is the client API for Bancc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BanccClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type banccClient struct {
	cc grpc.ClientConnInterface
}

func NewBanccClient(cc grpc.ClientConnInterface) BanccClient {
	return &banccClient{cc}
}

func (c *banccClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/bancc.Bancc/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banccClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/bancc.Bancc/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BanccServer is the server API for Bancc service.
// All implementations must embed UnimplementedBanccServer
// for forward compatibility
type BanccServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Auth(context.Context, *AuthRequest) (*AuthResponse, error)
	mustEmbedUnimplementedBanccServer()
}

// UnimplementedBanccServer must be embedded to have forward compatible implementations.
type UnimplementedBanccServer struct {
}

func (*UnimplementedBanccServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedBanccServer) Auth(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedBanccServer) mustEmbedUnimplementedBanccServer() {}

func RegisterBanccServer(s *grpc.Server, srv BanccServer) {
	s.RegisterService(&_Bancc_serviceDesc, srv)
}

func _Bancc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanccServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bancc.Bancc/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanccServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bancc_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanccServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bancc.Bancc/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanccServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bancc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bancc.Bancc",
	HandlerType: (*BanccServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Bancc_Ping_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _Bancc_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bancc.proto",
}
