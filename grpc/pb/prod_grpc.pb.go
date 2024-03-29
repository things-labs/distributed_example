// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// ArithClient is the client API for Arith service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArithClient interface {
	Mul(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error)
}

type arithClient struct {
	cc grpc.ClientConnInterface
}

func NewArithClient(cc grpc.ClientConnInterface) ArithClient {
	return &arithClient{cc}
}

func (c *arithClient) Mul(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error) {
	out := new(ArithResponse)
	err := c.cc.Invoke(ctx, "/Arith/Mul", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithServer is the server API for Arith service.
// All implementations must embed UnimplementedArithServer
// for forward compatibility
type ArithServer interface {
	Mul(context.Context, *ArithRequest) (*ArithResponse, error)
	mustEmbedUnimplementedArithServer()
}

// UnimplementedArithServer must be embedded to have forward compatible implementations.
type UnimplementedArithServer struct {
}

func (UnimplementedArithServer) Mul(context.Context, *ArithRequest) (*ArithResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mul not implemented")
}
func (UnimplementedArithServer) mustEmbedUnimplementedArithServer() {}

// UnsafeArithServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArithServer will
// result in compilation errors.
type UnsafeArithServer interface {
	mustEmbedUnimplementedArithServer()
}

func RegisterArithServer(s grpc.ServiceRegistrar, srv ArithServer) {
	s.RegisterService(&Arith_ServiceDesc, srv)
}

func _Arith_Mul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArithRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithServer).Mul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Arith/Mul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithServer).Mul(ctx, req.(*ArithRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Arith_ServiceDesc is the grpc.ServiceDesc for Arith service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Arith_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Arith",
	HandlerType: (*ArithServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mul",
			Handler:    _Arith_Mul_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/prod.proto",
}
