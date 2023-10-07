// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service.proto

package logistics_staff

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

// StaffServiceClient is the client API for StaffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaffServiceClient interface {
	RegisterDriver(ctx context.Context, in *RegisterDriverRequest, opts ...grpc.CallOption) (*RegisterDriverResponse, error)
	LoginDriver(ctx context.Context, in *LoginDriverRequest, opts ...grpc.CallOption) (*LoginDriverResponse, error)
	ConfirmSmsCode(ctx context.Context, in *ConfirmSmsCodeRequest, opts ...grpc.CallOption) (*ConfirmSmsCodeResponse, error)
	GetDriverProfile(ctx context.Context, in *GetDriverProfileRequest, opts ...grpc.CallOption) (*GetDriverProfileResponse, error)
	UpdatpDriverProfile(ctx context.Context, in *UpdateDriverProfileRequest, opts ...grpc.CallOption) (*UpdateDriverProfileResponse, error)
}

type staffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffServiceClient(cc grpc.ClientConnInterface) StaffServiceClient {
	return &staffServiceClient{cc}
}

func (c *staffServiceClient) RegisterDriver(ctx context.Context, in *RegisterDriverRequest, opts ...grpc.CallOption) (*RegisterDriverResponse, error) {
	out := new(RegisterDriverResponse)
	err := c.cc.Invoke(ctx, "/logistics_staff.StaffService/RegisterDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) LoginDriver(ctx context.Context, in *LoginDriverRequest, opts ...grpc.CallOption) (*LoginDriverResponse, error) {
	out := new(LoginDriverResponse)
	err := c.cc.Invoke(ctx, "/logistics_staff.StaffService/LoginDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) ConfirmSmsCode(ctx context.Context, in *ConfirmSmsCodeRequest, opts ...grpc.CallOption) (*ConfirmSmsCodeResponse, error) {
	out := new(ConfirmSmsCodeResponse)
	err := c.cc.Invoke(ctx, "/logistics_staff.StaffService/ConfirmSmsCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) GetDriverProfile(ctx context.Context, in *GetDriverProfileRequest, opts ...grpc.CallOption) (*GetDriverProfileResponse, error) {
	out := new(GetDriverProfileResponse)
	err := c.cc.Invoke(ctx, "/logistics_staff.StaffService/GetDriverProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) UpdatpDriverProfile(ctx context.Context, in *UpdateDriverProfileRequest, opts ...grpc.CallOption) (*UpdateDriverProfileResponse, error) {
	out := new(UpdateDriverProfileResponse)
	err := c.cc.Invoke(ctx, "/logistics_staff.StaffService/UpdatpDriverProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServiceServer is the server API for StaffService service.
// All implementations must embed UnimplementedStaffServiceServer
// for forward compatibility
type StaffServiceServer interface {
	RegisterDriver(context.Context, *RegisterDriverRequest) (*RegisterDriverResponse, error)
	LoginDriver(context.Context, *LoginDriverRequest) (*LoginDriverResponse, error)
	ConfirmSmsCode(context.Context, *ConfirmSmsCodeRequest) (*ConfirmSmsCodeResponse, error)
	GetDriverProfile(context.Context, *GetDriverProfileRequest) (*GetDriverProfileResponse, error)
	UpdatpDriverProfile(context.Context, *UpdateDriverProfileRequest) (*UpdateDriverProfileResponse, error)
	mustEmbedUnimplementedStaffServiceServer()
}

// UnimplementedStaffServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStaffServiceServer struct {
}

func (UnimplementedStaffServiceServer) RegisterDriver(context.Context, *RegisterDriverRequest) (*RegisterDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDriver not implemented")
}
func (UnimplementedStaffServiceServer) LoginDriver(context.Context, *LoginDriverRequest) (*LoginDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginDriver not implemented")
}
func (UnimplementedStaffServiceServer) ConfirmSmsCode(context.Context, *ConfirmSmsCodeRequest) (*ConfirmSmsCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmSmsCode not implemented")
}
func (UnimplementedStaffServiceServer) GetDriverProfile(context.Context, *GetDriverProfileRequest) (*GetDriverProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDriverProfile not implemented")
}
func (UnimplementedStaffServiceServer) UpdatpDriverProfile(context.Context, *UpdateDriverProfileRequest) (*UpdateDriverProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatpDriverProfile not implemented")
}
func (UnimplementedStaffServiceServer) mustEmbedUnimplementedStaffServiceServer() {}

// UnsafeStaffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaffServiceServer will
// result in compilation errors.
type UnsafeStaffServiceServer interface {
	mustEmbedUnimplementedStaffServiceServer()
}

func RegisterStaffServiceServer(s grpc.ServiceRegistrar, srv StaffServiceServer) {
	s.RegisterService(&StaffService_ServiceDesc, srv)
}

func _StaffService_RegisterDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).RegisterDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_staff.StaffService/RegisterDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).RegisterDriver(ctx, req.(*RegisterDriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_LoginDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginDriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).LoginDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_staff.StaffService/LoginDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).LoginDriver(ctx, req.(*LoginDriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_ConfirmSmsCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmSmsCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).ConfirmSmsCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_staff.StaffService/ConfirmSmsCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).ConfirmSmsCode(ctx, req.(*ConfirmSmsCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_GetDriverProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDriverProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).GetDriverProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_staff.StaffService/GetDriverProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).GetDriverProfile(ctx, req.(*GetDriverProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_UpdatpDriverProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDriverProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).UpdatpDriverProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics_staff.StaffService/UpdatpDriverProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).UpdatpDriverProfile(ctx, req.(*UpdateDriverProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StaffService_ServiceDesc is the grpc.ServiceDesc for StaffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logistics_staff.StaffService",
	HandlerType: (*StaffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDriver",
			Handler:    _StaffService_RegisterDriver_Handler,
		},
		{
			MethodName: "LoginDriver",
			Handler:    _StaffService_LoginDriver_Handler,
		},
		{
			MethodName: "ConfirmSmsCode",
			Handler:    _StaffService_ConfirmSmsCode_Handler,
		},
		{
			MethodName: "GetDriverProfile",
			Handler:    _StaffService_GetDriverProfile_Handler,
		},
		{
			MethodName: "UpdatpDriverProfile",
			Handler:    _StaffService_UpdatpDriverProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
