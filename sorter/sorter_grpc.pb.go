// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: sorter/sorter.proto

package sorter

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

// SorterClient is the client API for Sorter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SorterClient interface {
	QuickSort(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Numbers, error)
}

type sorterClient struct {
	cc grpc.ClientConnInterface
}

func NewSorterClient(cc grpc.ClientConnInterface) SorterClient {
	return &sorterClient{cc}
}

func (c *sorterClient) QuickSort(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Numbers, error) {
	out := new(Numbers)
	err := c.cc.Invoke(ctx, "/sorter.Sorter/QuickSort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SorterServer is the server API for Sorter service.
// All implementations must embed UnimplementedSorterServer
// for forward compatibility
type SorterServer interface {
	QuickSort(context.Context, *Numbers) (*Numbers, error)
	mustEmbedUnimplementedSorterServer()
}

// UnimplementedSorterServer must be embedded to have forward compatible implementations.
type UnimplementedSorterServer struct {
}

func (UnimplementedSorterServer) QuickSort(context.Context, *Numbers) (*Numbers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuickSort not implemented")
}
func (UnimplementedSorterServer) mustEmbedUnimplementedSorterServer() {}

// UnsafeSorterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SorterServer will
// result in compilation errors.
type UnsafeSorterServer interface {
	mustEmbedUnimplementedSorterServer()
}

func RegisterSorterServer(s grpc.ServiceRegistrar, srv SorterServer) {
	s.RegisterService(&Sorter_ServiceDesc, srv)
}

func _Sorter_QuickSort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Numbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SorterServer).QuickSort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sorter.Sorter/QuickSort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SorterServer).QuickSort(ctx, req.(*Numbers))
	}
	return interceptor(ctx, in, info, handler)
}

// Sorter_ServiceDesc is the grpc.ServiceDesc for Sorter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sorter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sorter.Sorter",
	HandlerType: (*SorterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QuickSort",
			Handler:    _Sorter_QuickSort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sorter/sorter.proto",
}
