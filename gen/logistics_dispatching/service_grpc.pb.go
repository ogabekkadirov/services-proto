// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service.proto

package logistics_dispatching

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

// DispatchingServiceClient is the client API for DispatchingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DispatchingServiceClient interface {
	TrackDriverLocation(ctx context.Context, in *TrackDriverLocationRequest, opts ...grpc.CallOption) (*TrackDriverLocationResponse, error)
	ListDriverLocation(ctx context.Context, in *ListDriverLocationRequest, opts ...grpc.CallOption) (*ListDriverLocationResponse, error)
	UpdateDriverStatus(ctx context.Context, in *UpdateDriverStatusRequest, opts ...grpc.CallOption) (*UpdateDriverStatusResponse, error)
	GetDriverStatus(ctx context.Context, in *GetDriverStatusRequest, opts ...grpc.CallOption) (*GetDriverStatusResponse, error)
	// order
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error)
	PickupOrder(ctx context.Context, in *PickupOrderRequest, opts ...grpc.CallOption) (*PickupOrderResponse, error)
	DeliverOrder(ctx context.Context, in *DeliverOrderRequest, opts ...grpc.CallOption) (*DeliverOrderResponse, error)
	CompleteOrder(ctx context.Context, in *CompleteOrderRequest, opts ...grpc.CallOption) (*CompleteOrderResponse, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	ListOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderResponse, error)
}

type dispatchingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatchingServiceClient(cc grpc.ClientConnInterface) DispatchingServiceClient {
	return &dispatchingServiceClient{cc}
}

func (c *dispatchingServiceClient) TrackDriverLocation(ctx context.Context, in *TrackDriverLocationRequest, opts ...grpc.CallOption) (*TrackDriverLocationResponse, error) {
	out := new(TrackDriverLocationResponse)
	err := c.cc.Invoke(ctx, "/logistics_dispatching.DispatchingService/TrackDriverLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatchingServiceClient) ListDriverLocation(ctx context.Context, in *ListDriverLocationRequest, opts ...grpc.CallOption) (*ListDriverLocationResponse, error) {
	out := new(ListDriverLocationResponse)
	err := c.cc.Invoke(ctx, "/logistics_dispatching.DispatchingService/ListDriverLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatchingServiceClient) UpdateDriverStatus(ctx context.Context, in *UpdateDriverStatusRequest, opts ...grpc.CallOption) (*UpdateDriverStatusResponse, error) {
	out := new(UpdateDriverStatusResponse)
	err := c.cc.Invoke(ctx, "/logistics_dispatching.DispatchingService/UpdateDriverStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatchingServiceClient) GetDriverStatus(ctx context.Context, in *GetDriverStatusRequest, opts ...grpc.CallOption) (*GetDriverStatusResponse, error) {
	out := new(GetDriverStatusResponse)
	err := c.cc.Invoke(ctx, "/logistics_dispatching.DispatchingService/GetDriverStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatchingServiceClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error) {
	out := new(UpdateOrderResponse)
	err := c.cc.Invoke(ctx, "/logistics_dispatching.DispatchingService/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatchingServiceClient) PickupOrder(ctx context.Context, in *PickupOrderRequest, opts ...grpc.CallOption) (*PickupOrderResponse, error) {
	out := new(PickupOrderResponse)
	err := c.cc.Invoke(ctx, "/logistics_dispatching.DispatchingService/PickupOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatchingServiceClient) DeliverOrder(ctx context.Context, in *DeliverOrderRequest, opts ...grpc.CallOption) (*DeliverOrderResponse, error) {
	out := new(DeliverOrderResponse)
	err := c.cc.Invoke(ctx, "/logistics_dispatching.DispatchingService/DeliverOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatchingServiceClient) CompleteOrder(ctx context.Context, in *CompleteOrderRequest, opts ...grpc.CallOption) (*CompleteOrderResponse, error) {
	out := new(CompleteOrderResponse)
	err := c.cc.Invoke(ctx, "/logistics_dispatching.DispatchingService/CompleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatchingServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, "/logistics_dispatching.DispatchingService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatchingServiceClient) ListOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderResponse, error) {
	out := new(ListOrderResponse)
	err := c.cc.Invoke(ctx, "/logistics_dispatching.DispatchingService/ListOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatchingServiceServer is the server API for DispatchingService service.
// All implementations must embed UnimplementedDispatchingServiceServer
// for forward compatibility
type DispatchingServiceServer interface {
	TrackDriverLocation(context.Context, *TrackDriverLocationRequest) (*TrackDriverLocationResponse, error)
	ListDriverLocation(context.Context, *ListDriverLocationRequest) (*ListDriverLocationResponse, error)
	UpdateDriverStatus(context.Context, *UpdateDriverStatusRequest) (*UpdateDriverStatusResponse, error)
	GetDriverStatus(context.Context, *GetDriverStatusRequest) (*GetDriverStatusResponse, error)
	// order
	UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error)
	PickupOrder(context.Context, *PickupOrderRequest) (*PickupOrderResponse, error)
	DeliverOrder(context.Context, *DeliverOrderRequest) (*DeliverOrderResponse, error)
	CompleteOrder(context.Context, *CompleteOrderRequest) (*CompleteOrderResponse, error)
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	ListOrder(context.Context, *ListOrderRequest) (*ListOrderResponse, error)
	mustEmbedUnimplementedDispatchingServiceServer()
}

// UnimplementedDispatchingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDispatchingServiceServer struct {
}

func (UnimplementedDispatchingServiceServer) TrackDriverLocation(context.Context, *TrackDriverLocationRequest) (*TrackDriverLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackDriverLocation not implemented")
}
func (UnimplementedDispatchingServiceServer) ListDriverLocation(context.Context, *ListDriverLocationRequest) (*ListDriverLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDriverLocation not implemented")
}
func (UnimplementedDispatchingServiceServer) UpdateDriverStatus(context.Context, *UpdateDriverStatusRequest) (*UpdateDriverStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDriverStatus not implemented")
}
func (UnimplementedDispatchingServiceServer) GetDriverStatus(context.Context, *GetDriverStatusRequest) (*GetDriverStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDriverStatus not implemented")
}
func (UnimplementedDispatchingServiceServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedDispatchingServiceServer) PickupOrder(context.Context, *PickupOrderRequest) (*PickupOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PickupOrder not implemented")
}
func (UnimplementedDispatchingServiceServer) DeliverOrder(context.Context, *DeliverOrderRequest) (*DeliverOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverOrder not implemented")
}
func (UnimplementedDispatchingServiceServer) CompleteOrder(context.Context, *CompleteOrderRequest) (*CompleteOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteOrder not implemented")
}
func (UnimplementedDispatchingServiceServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedDispatchingServiceServer) ListOrder(context.Context, *ListOrderRequest) (*ListOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (UnimplementedDispatchingServiceServer) mustEmbedUnimplementedDispatchingServiceServer() {}

// UnsafeDispatchingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DispatchingServiceServer will
// result in compilation errors.
type UnsafeDispatchingServiceServer interface {
	mustEmbedUnimplementedDispatchingServiceServer()
}

func RegisterDispatchingServiceServer(s grpc.ServiceRegistrar, srv DispatchingServiceServer) {
	s.RegisterService(&DispatchingService_ServiceDesc, srv)
}

func _DispatchingService_TrackDriverLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackDriverLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchingServiceServer).TrackDriverLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_dispatching.DispatchingService/TrackDriverLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchingServiceServer).TrackDriverLocation(ctx, req.(*TrackDriverLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatchingService_ListDriverLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDriverLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchingServiceServer).ListDriverLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_dispatching.DispatchingService/ListDriverLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchingServiceServer).ListDriverLocation(ctx, req.(*ListDriverLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatchingService_UpdateDriverStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDriverStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchingServiceServer).UpdateDriverStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_dispatching.DispatchingService/UpdateDriverStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchingServiceServer).UpdateDriverStatus(ctx, req.(*UpdateDriverStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatchingService_GetDriverStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDriverStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchingServiceServer).GetDriverStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_dispatching.DispatchingService/GetDriverStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchingServiceServer).GetDriverStatus(ctx, req.(*GetDriverStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatchingService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchingServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_dispatching.DispatchingService/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchingServiceServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatchingService_PickupOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PickupOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchingServiceServer).PickupOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_dispatching.DispatchingService/PickupOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchingServiceServer).PickupOrder(ctx, req.(*PickupOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatchingService_DeliverOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchingServiceServer).DeliverOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_dispatching.DispatchingService/DeliverOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchingServiceServer).DeliverOrder(ctx, req.(*DeliverOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatchingService_CompleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchingServiceServer).CompleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_dispatching.DispatchingService/CompleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchingServiceServer).CompleteOrder(ctx, req.(*CompleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatchingService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchingServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_dispatching.DispatchingService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchingServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatchingService_ListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchingServiceServer).ListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_dispatching.DispatchingService/ListOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchingServiceServer).ListOrder(ctx, req.(*ListOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DispatchingService_ServiceDesc is the grpc.ServiceDesc for DispatchingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DispatchingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logistics_dispatching.DispatchingService",
	HandlerType: (*DispatchingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TrackDriverLocation",
			Handler:    _DispatchingService_TrackDriverLocation_Handler,
		},
		{
			MethodName: "ListDriverLocation",
			Handler:    _DispatchingService_ListDriverLocation_Handler,
		},
		{
			MethodName: "UpdateDriverStatus",
			Handler:    _DispatchingService_UpdateDriverStatus_Handler,
		},
		{
			MethodName: "GetDriverStatus",
			Handler:    _DispatchingService_GetDriverStatus_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _DispatchingService_UpdateOrder_Handler,
		},
		{
			MethodName: "PickupOrder",
			Handler:    _DispatchingService_PickupOrder_Handler,
		},
		{
			MethodName: "DeliverOrder",
			Handler:    _DispatchingService_DeliverOrder_Handler,
		},
		{
			MethodName: "CompleteOrder",
			Handler:    _DispatchingService_CompleteOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _DispatchingService_GetOrder_Handler,
		},
		{
			MethodName: "ListOrder",
			Handler:    _DispatchingService_ListOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}