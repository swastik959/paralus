// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/rpc/system/partner.proto

package systemv3

import (
	context "context"
	v3 "github.com/paralus/paralus/proto/types/systempb/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PartnerService_CreatePartner_FullMethodName  = "/paralus.dev.rpc.system.v3.PartnerService/CreatePartner"
	PartnerService_GetPartner_FullMethodName     = "/paralus.dev.rpc.system.v3.PartnerService/GetPartner"
	PartnerService_GetInitPartner_FullMethodName = "/paralus.dev.rpc.system.v3.PartnerService/GetInitPartner"
	PartnerService_UpdatePartner_FullMethodName  = "/paralus.dev.rpc.system.v3.PartnerService/UpdatePartner"
	PartnerService_DeletePartner_FullMethodName  = "/paralus.dev.rpc.system.v3.PartnerService/DeletePartner"
)

// PartnerServiceClient is the client API for PartnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartnerServiceClient interface {
	CreatePartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error)
	GetPartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error)
	GetInitPartner(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*v3.Partner, error)
	UpdatePartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error)
	DeletePartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error)
}

type partnerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartnerServiceClient(cc grpc.ClientConnInterface) PartnerServiceClient {
	return &partnerServiceClient{cc}
}

func (c *partnerServiceClient) CreatePartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error) {
	out := new(v3.Partner)
	err := c.cc.Invoke(ctx, PartnerService_CreatePartner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) GetPartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error) {
	out := new(v3.Partner)
	err := c.cc.Invoke(ctx, PartnerService_GetPartner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) GetInitPartner(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*v3.Partner, error) {
	out := new(v3.Partner)
	err := c.cc.Invoke(ctx, PartnerService_GetInitPartner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) UpdatePartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error) {
	out := new(v3.Partner)
	err := c.cc.Invoke(ctx, PartnerService_UpdatePartner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) DeletePartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error) {
	out := new(v3.Partner)
	err := c.cc.Invoke(ctx, PartnerService_DeletePartner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartnerServiceServer is the server API for PartnerService service.
// All implementations should embed UnimplementedPartnerServiceServer
// for forward compatibility
type PartnerServiceServer interface {
	CreatePartner(context.Context, *v3.Partner) (*v3.Partner, error)
	GetPartner(context.Context, *v3.Partner) (*v3.Partner, error)
	GetInitPartner(context.Context, *EmptyRequest) (*v3.Partner, error)
	UpdatePartner(context.Context, *v3.Partner) (*v3.Partner, error)
	DeletePartner(context.Context, *v3.Partner) (*v3.Partner, error)
}

// UnimplementedPartnerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPartnerServiceServer struct {
}

func (UnimplementedPartnerServiceServer) CreatePartner(context.Context, *v3.Partner) (*v3.Partner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePartner not implemented")
}
func (UnimplementedPartnerServiceServer) GetPartner(context.Context, *v3.Partner) (*v3.Partner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartner not implemented")
}
func (UnimplementedPartnerServiceServer) GetInitPartner(context.Context, *EmptyRequest) (*v3.Partner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInitPartner not implemented")
}
func (UnimplementedPartnerServiceServer) UpdatePartner(context.Context, *v3.Partner) (*v3.Partner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePartner not implemented")
}
func (UnimplementedPartnerServiceServer) DeletePartner(context.Context, *v3.Partner) (*v3.Partner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePartner not implemented")
}

// UnsafePartnerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartnerServiceServer will
// result in compilation errors.
type UnsafePartnerServiceServer interface {
	mustEmbedUnimplementedPartnerServiceServer()
}

func RegisterPartnerServiceServer(s grpc.ServiceRegistrar, srv PartnerServiceServer) {
	s.RegisterService(&PartnerService_ServiceDesc, srv)
}

func _PartnerService_CreatePartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Partner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).CreatePartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerService_CreatePartner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).CreatePartner(ctx, req.(*v3.Partner))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_GetPartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Partner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).GetPartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerService_GetPartner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).GetPartner(ctx, req.(*v3.Partner))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_GetInitPartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).GetInitPartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerService_GetInitPartner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).GetInitPartner(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_UpdatePartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Partner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).UpdatePartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerService_UpdatePartner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).UpdatePartner(ctx, req.(*v3.Partner))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_DeletePartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Partner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).DeletePartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerService_DeletePartner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).DeletePartner(ctx, req.(*v3.Partner))
	}
	return interceptor(ctx, in, info, handler)
}

// PartnerService_ServiceDesc is the grpc.ServiceDesc for PartnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "paralus.dev.rpc.system.v3.PartnerService",
	HandlerType: (*PartnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePartner",
			Handler:    _PartnerService_CreatePartner_Handler,
		},
		{
			MethodName: "GetPartner",
			Handler:    _PartnerService_GetPartner_Handler,
		},
		{
			MethodName: "GetInitPartner",
			Handler:    _PartnerService_GetInitPartner_Handler,
		},
		{
			MethodName: "UpdatePartner",
			Handler:    _PartnerService_UpdatePartner_Handler,
		},
		{
			MethodName: "DeletePartner",
			Handler:    _PartnerService_DeletePartner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/system/partner.proto",
}
