// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: punch.proto

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

// PunchServiceClient is the client API for PunchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PunchServiceClient interface {
	PunchIn(ctx context.Context, in *PunchInRequest, opts ...grpc.CallOption) (*Response, error)
	VerifyOtpPunchin(ctx context.Context, in *PunchInRequestOtp, opts ...grpc.CallOption) (*Response, error)
	PunchOut(ctx context.Context, in *PunchOutRequest, opts ...grpc.CallOption) (*Response, error)
	GetDuty(ctx context.Context, in *GetDutyRequest, opts ...grpc.CallOption) (*Response, error)
}

type punchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPunchServiceClient(cc grpc.ClientConnInterface) PunchServiceClient {
	return &punchServiceClient{cc}
}

func (c *punchServiceClient) PunchIn(ctx context.Context, in *PunchInRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/punch.PunchService/PunchIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *punchServiceClient) VerifyOtpPunchin(ctx context.Context, in *PunchInRequestOtp, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/punch.PunchService/VerifyOtpPunchin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *punchServiceClient) PunchOut(ctx context.Context, in *PunchOutRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/punch.PunchService/PunchOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *punchServiceClient) GetDuty(ctx context.Context, in *GetDutyRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/punch.PunchService/GetDuty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PunchServiceServer is the server API for PunchService service.
// All implementations must embed UnimplementedPunchServiceServer
// for forward compatibility
type PunchServiceServer interface {
	PunchIn(context.Context, *PunchInRequest) (*Response, error)
	VerifyOtpPunchin(context.Context, *PunchInRequestOtp) (*Response, error)
	PunchOut(context.Context, *PunchOutRequest) (*Response, error)
	GetDuty(context.Context, *GetDutyRequest) (*Response, error)
	mustEmbedUnimplementedPunchServiceServer()
}

// UnimplementedPunchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPunchServiceServer struct {
}

func (UnimplementedPunchServiceServer) PunchIn(context.Context, *PunchInRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PunchIn not implemented")
}
func (UnimplementedPunchServiceServer) VerifyOtpPunchin(context.Context, *PunchInRequestOtp) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOtpPunchin not implemented")
}
func (UnimplementedPunchServiceServer) PunchOut(context.Context, *PunchOutRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PunchOut not implemented")
}
func (UnimplementedPunchServiceServer) GetDuty(context.Context, *GetDutyRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDuty not implemented")
}
func (UnimplementedPunchServiceServer) mustEmbedUnimplementedPunchServiceServer() {}

// UnsafePunchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PunchServiceServer will
// result in compilation errors.
type UnsafePunchServiceServer interface {
	mustEmbedUnimplementedPunchServiceServer()
}

func RegisterPunchServiceServer(s grpc.ServiceRegistrar, srv PunchServiceServer) {
	s.RegisterService(&PunchService_ServiceDesc, srv)
}

func _PunchService_PunchIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PunchInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PunchServiceServer).PunchIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/punch.PunchService/PunchIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PunchServiceServer).PunchIn(ctx, req.(*PunchInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PunchService_VerifyOtpPunchin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PunchInRequestOtp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PunchServiceServer).VerifyOtpPunchin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/punch.PunchService/VerifyOtpPunchin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PunchServiceServer).VerifyOtpPunchin(ctx, req.(*PunchInRequestOtp))
	}
	return interceptor(ctx, in, info, handler)
}

func _PunchService_PunchOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PunchOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PunchServiceServer).PunchOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/punch.PunchService/PunchOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PunchServiceServer).PunchOut(ctx, req.(*PunchOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PunchService_GetDuty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDutyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PunchServiceServer).GetDuty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/punch.PunchService/GetDuty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PunchServiceServer).GetDuty(ctx, req.(*GetDutyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PunchService_ServiceDesc is the grpc.ServiceDesc for PunchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PunchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "punch.PunchService",
	HandlerType: (*PunchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PunchIn",
			Handler:    _PunchService_PunchIn_Handler,
		},
		{
			MethodName: "VerifyOtpPunchin",
			Handler:    _PunchService_VerifyOtpPunchin_Handler,
		},
		{
			MethodName: "PunchOut",
			Handler:    _PunchService_PunchOut_Handler,
		},
		{
			MethodName: "GetDuty",
			Handler:    _PunchService_GetDuty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "punch.proto",
}
