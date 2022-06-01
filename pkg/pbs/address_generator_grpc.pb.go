// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pbs

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

// WithdrawClient is the client API for Withdraw service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WithdrawClient interface {
	Withdraw(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*WithdrawRes, error)
}

type withdrawClient struct {
	cc grpc.ClientConnInterface
}

func NewWithdrawClient(cc grpc.ClientConnInterface) WithdrawClient {
	return &withdrawClient{cc}
}

func (c *withdrawClient) Withdraw(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*WithdrawRes, error) {
	out := new(WithdrawRes)
	err := c.cc.Invoke(ctx, "/chainerGrpc.Withdraw/Withdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WithdrawServer is the server API for Withdraw service.
// All implementations should embed UnimplementedWithdrawServer
// for forward compatibility
type WithdrawServer interface {
	Withdraw(context.Context, *WithdrawReq) (*WithdrawRes, error)
}

// UnimplementedWithdrawServer should be embedded to have forward compatible implementations.
type UnimplementedWithdrawServer struct {
}

func (UnimplementedWithdrawServer) Withdraw(context.Context, *WithdrawReq) (*WithdrawRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}

// UnsafeWithdrawServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WithdrawServer will
// result in compilation errors.
type UnsafeWithdrawServer interface {
	mustEmbedUnimplementedWithdrawServer()
}

func RegisterWithdrawServer(s grpc.ServiceRegistrar, srv WithdrawServer) {
	s.RegisterService(&Withdraw_ServiceDesc, srv)
}

func _Withdraw_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainerGrpc.Withdraw/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawServer).Withdraw(ctx, req.(*WithdrawReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Withdraw_ServiceDesc is the grpc.ServiceDesc for Withdraw service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Withdraw_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainerGrpc.Withdraw",
	HandlerType: (*WithdrawServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Withdraw",
			Handler:    _Withdraw_Withdraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address_generator.proto",
}
