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

// AddressGeneratorClient is the client API for AddressGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressGeneratorClient interface {
	GetDepositAddress(ctx context.Context, in *DepositAddressReq, opts ...grpc.CallOption) (*DepositAddressRes, error)
}

type addressGeneratorClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressGeneratorClient(cc grpc.ClientConnInterface) AddressGeneratorClient {
	return &addressGeneratorClient{cc}
}

func (c *addressGeneratorClient) GetDepositAddress(ctx context.Context, in *DepositAddressReq, opts ...grpc.CallOption) (*DepositAddressRes, error) {
	out := new(DepositAddressRes)
	err := c.cc.Invoke(ctx, "/chainerGrpc.AddressGenerator/GetDepositAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressGeneratorServer is the server API for AddressGenerator service.
// All implementations should embed UnimplementedAddressGeneratorServer
// for forward compatibility
type AddressGeneratorServer interface {
	GetDepositAddress(context.Context, *DepositAddressReq) (*DepositAddressRes, error)
}

// UnimplementedAddressGeneratorServer should be embedded to have forward compatible implementations.
type UnimplementedAddressGeneratorServer struct {
}

func (UnimplementedAddressGeneratorServer) GetDepositAddress(context.Context, *DepositAddressReq) (*DepositAddressRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepositAddress not implemented")
}

// UnsafeAddressGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressGeneratorServer will
// result in compilation errors.
type UnsafeAddressGeneratorServer interface {
	mustEmbedUnimplementedAddressGeneratorServer()
}

func RegisterAddressGeneratorServer(s grpc.ServiceRegistrar, srv AddressGeneratorServer) {
	s.RegisterService(&AddressGenerator_ServiceDesc, srv)
}

func _AddressGenerator_GetDepositAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressGeneratorServer).GetDepositAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainerGrpc.AddressGenerator/GetDepositAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressGeneratorServer).GetDepositAddress(ctx, req.(*DepositAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressGenerator_ServiceDesc is the grpc.ServiceDesc for AddressGenerator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressGenerator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainerGrpc.AddressGenerator",
	HandlerType: (*AddressGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDepositAddress",
			Handler:    _AddressGenerator_GetDepositAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "withdraw.proto",
}
