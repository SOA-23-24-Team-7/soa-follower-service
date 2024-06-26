// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: followerMicroservice.proto

package server

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

const (
	FollowerMicroservice_FollowUser_FullMethodName             = "/server.FollowerMicroservice/FollowUser"
	FollowerMicroservice_UnfollowUser_FullMethodName           = "/server.FollowerMicroservice/UnfollowUser"
	FollowerMicroservice_GetFollowers_FullMethodName           = "/server.FollowerMicroservice/GetFollowers"
	FollowerMicroservice_GetFollowings_FullMethodName          = "/server.FollowerMicroservice/GetFollowings"
	FollowerMicroservice_GetFollowerSuggestions_FullMethodName = "/server.FollowerMicroservice/GetFollowerSuggestions"
)

// FollowerMicroserviceClient is the client API for FollowerMicroservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowerMicroserviceClient interface {
	FollowUser(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowerStringMessage, error)
	UnfollowUser(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowerStringMessage, error)
	GetFollowers(ctx context.Context, in *FollowerIdRequest, opts ...grpc.CallOption) (*FollowerListResponse, error)
	GetFollowings(ctx context.Context, in *FollowerIdRequest, opts ...grpc.CallOption) (*FollowerListResponse, error)
	GetFollowerSuggestions(ctx context.Context, in *FollowerIdRequest, opts ...grpc.CallOption) (*FollowerListResponse, error)
}

type followerMicroserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowerMicroserviceClient(cc grpc.ClientConnInterface) FollowerMicroserviceClient {
	return &followerMicroserviceClient{cc}
}

func (c *followerMicroserviceClient) FollowUser(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowerStringMessage, error) {
	out := new(FollowerStringMessage)
	err := c.cc.Invoke(ctx, FollowerMicroservice_FollowUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followerMicroserviceClient) UnfollowUser(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowerStringMessage, error) {
	out := new(FollowerStringMessage)
	err := c.cc.Invoke(ctx, FollowerMicroservice_UnfollowUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followerMicroserviceClient) GetFollowers(ctx context.Context, in *FollowerIdRequest, opts ...grpc.CallOption) (*FollowerListResponse, error) {
	out := new(FollowerListResponse)
	err := c.cc.Invoke(ctx, FollowerMicroservice_GetFollowers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followerMicroserviceClient) GetFollowings(ctx context.Context, in *FollowerIdRequest, opts ...grpc.CallOption) (*FollowerListResponse, error) {
	out := new(FollowerListResponse)
	err := c.cc.Invoke(ctx, FollowerMicroservice_GetFollowings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followerMicroserviceClient) GetFollowerSuggestions(ctx context.Context, in *FollowerIdRequest, opts ...grpc.CallOption) (*FollowerListResponse, error) {
	out := new(FollowerListResponse)
	err := c.cc.Invoke(ctx, FollowerMicroservice_GetFollowerSuggestions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowerMicroserviceServer is the server API for FollowerMicroservice service.
// All implementations must embed UnimplementedFollowerMicroserviceServer
// for forward compatibility
type FollowerMicroserviceServer interface {
	FollowUser(context.Context, *FollowRequest) (*FollowerStringMessage, error)
	UnfollowUser(context.Context, *FollowRequest) (*FollowerStringMessage, error)
	GetFollowers(context.Context, *FollowerIdRequest) (*FollowerListResponse, error)
	GetFollowings(context.Context, *FollowerIdRequest) (*FollowerListResponse, error)
	GetFollowerSuggestions(context.Context, *FollowerIdRequest) (*FollowerListResponse, error)
	mustEmbedUnimplementedFollowerMicroserviceServer()
}

// UnimplementedFollowerMicroserviceServer must be embedded to have forward compatible implementations.
type UnimplementedFollowerMicroserviceServer struct {
}

func (UnimplementedFollowerMicroserviceServer) FollowUser(context.Context, *FollowRequest) (*FollowerStringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedFollowerMicroserviceServer) UnfollowUser(context.Context, *FollowRequest) (*FollowerStringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfollowUser not implemented")
}
func (UnimplementedFollowerMicroserviceServer) GetFollowers(context.Context, *FollowerIdRequest) (*FollowerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowers not implemented")
}
func (UnimplementedFollowerMicroserviceServer) GetFollowings(context.Context, *FollowerIdRequest) (*FollowerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowings not implemented")
}
func (UnimplementedFollowerMicroserviceServer) GetFollowerSuggestions(context.Context, *FollowerIdRequest) (*FollowerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowerSuggestions not implemented")
}
func (UnimplementedFollowerMicroserviceServer) mustEmbedUnimplementedFollowerMicroserviceServer() {}

// UnsafeFollowerMicroserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowerMicroserviceServer will
// result in compilation errors.
type UnsafeFollowerMicroserviceServer interface {
	mustEmbedUnimplementedFollowerMicroserviceServer()
}

func RegisterFollowerMicroserviceServer(s grpc.ServiceRegistrar, srv FollowerMicroserviceServer) {
	s.RegisterService(&FollowerMicroservice_ServiceDesc, srv)
}

func _FollowerMicroservice_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowerMicroserviceServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FollowerMicroservice_FollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowerMicroserviceServer).FollowUser(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowerMicroservice_UnfollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowerMicroserviceServer).UnfollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FollowerMicroservice_UnfollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowerMicroserviceServer).UnfollowUser(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowerMicroservice_GetFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowerMicroserviceServer).GetFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FollowerMicroservice_GetFollowers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowerMicroserviceServer).GetFollowers(ctx, req.(*FollowerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowerMicroservice_GetFollowings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowerMicroserviceServer).GetFollowings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FollowerMicroservice_GetFollowings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowerMicroserviceServer).GetFollowings(ctx, req.(*FollowerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowerMicroservice_GetFollowerSuggestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowerMicroserviceServer).GetFollowerSuggestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FollowerMicroservice_GetFollowerSuggestions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowerMicroserviceServer).GetFollowerSuggestions(ctx, req.(*FollowerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FollowerMicroservice_ServiceDesc is the grpc.ServiceDesc for FollowerMicroservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FollowerMicroservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FollowerMicroservice",
	HandlerType: (*FollowerMicroserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FollowUser",
			Handler:    _FollowerMicroservice_FollowUser_Handler,
		},
		{
			MethodName: "UnfollowUser",
			Handler:    _FollowerMicroservice_UnfollowUser_Handler,
		},
		{
			MethodName: "GetFollowers",
			Handler:    _FollowerMicroservice_GetFollowers_Handler,
		},
		{
			MethodName: "GetFollowings",
			Handler:    _FollowerMicroservice_GetFollowings_Handler,
		},
		{
			MethodName: "GetFollowerSuggestions",
			Handler:    _FollowerMicroservice_GetFollowerSuggestions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "followerMicroservice.proto",
}
