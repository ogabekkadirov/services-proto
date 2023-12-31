// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: restaurant_support/service.proto

package restaurant_support

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

// SupportServiceClient is the client API for SupportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupportServiceClient interface {
	ListRating(ctx context.Context, in *ListRatingRequest, opts ...grpc.CallOption) (*ListRatingResponse, error)
}

type supportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSupportServiceClient(cc grpc.ClientConnInterface) SupportServiceClient {
	return &supportServiceClient{cc}
}

func (c *supportServiceClient) ListRating(ctx context.Context, in *ListRatingRequest, opts ...grpc.CallOption) (*ListRatingResponse, error) {
	out := new(ListRatingResponse)
	err := c.cc.Invoke(ctx, "/restaurant_support.SupportService/ListRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupportServiceServer is the server API for SupportService service.
// All implementations must embed UnimplementedSupportServiceServer
// for forward compatibility
type SupportServiceServer interface {
	ListRating(context.Context, *ListRatingRequest) (*ListRatingResponse, error)
	mustEmbedUnimplementedSupportServiceServer()
}

// UnimplementedSupportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSupportServiceServer struct {
}

func (UnimplementedSupportServiceServer) ListRating(context.Context, *ListRatingRequest) (*ListRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRating not implemented")
}
func (UnimplementedSupportServiceServer) mustEmbedUnimplementedSupportServiceServer() {}

// UnsafeSupportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupportServiceServer will
// result in compilation errors.
type UnsafeSupportServiceServer interface {
	mustEmbedUnimplementedSupportServiceServer()
}

func RegisterSupportServiceServer(s grpc.ServiceRegistrar, srv SupportServiceServer) {
	s.RegisterService(&SupportService_ServiceDesc, srv)
}

func _SupportService_ListRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportServiceServer).ListRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant_support.SupportService/ListRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportServiceServer).ListRating(ctx, req.(*ListRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SupportService_ServiceDesc is the grpc.ServiceDesc for SupportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "restaurant_support.SupportService",
	HandlerType: (*SupportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRating",
			Handler:    _SupportService_ListRating_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "restaurant_support/service.proto",
}
