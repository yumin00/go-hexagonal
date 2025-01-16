// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: core/user/user.proto

package user

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
	UserData_GetUserInfoById_FullMethodName = "/user.UserData/GetUserInfoById"
)

// UserDataClient is the client API for UserData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserDataClient interface {
	GetUserInfoById(ctx context.Context, in *GetUserInfoByIdRequest, opts ...grpc.CallOption) (*GetUserInfoByIdResponse, error)
}

type userDataClient struct {
	cc grpc.ClientConnInterface
}

func NewUserDataClient(cc grpc.ClientConnInterface) UserDataClient {
	return &userDataClient{cc}
}

func (c *userDataClient) GetUserInfoById(ctx context.Context, in *GetUserInfoByIdRequest, opts ...grpc.CallOption) (*GetUserInfoByIdResponse, error) {
	out := new(GetUserInfoByIdResponse)
	err := c.cc.Invoke(ctx, UserData_GetUserInfoById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDataServer is the server API for UserData service.
// All implementations must embed UnimplementedUserDataServer
// for forward compatibility
type UserDataServer interface {
	GetUserInfoById(context.Context, *GetUserInfoByIdRequest) (*GetUserInfoByIdResponse, error)
	mustEmbedUnimplementedUserDataServer()
}

// UnimplementedUserDataServer must be embedded to have forward compatible implementations.
type UnimplementedUserDataServer struct {
}

func (UnimplementedUserDataServer) GetUserInfoById(context.Context, *GetUserInfoByIdRequest) (*GetUserInfoByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoById not implemented")
}
func (UnimplementedUserDataServer) mustEmbedUnimplementedUserDataServer() {}

// UnsafeUserDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserDataServer will
// result in compilation errors.
type UnsafeUserDataServer interface {
	mustEmbedUnimplementedUserDataServer()
}

func RegisterUserDataServer(s grpc.ServiceRegistrar, srv UserDataServer) {
	s.RegisterService(&UserData_ServiceDesc, srv)
}

func _UserData_GetUserInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServer).GetUserInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserData_GetUserInfoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServer).GetUserInfoById(ctx, req.(*GetUserInfoByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserData_ServiceDesc is the grpc.ServiceDesc for UserData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserData",
	HandlerType: (*UserDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfoById",
			Handler:    _UserData_GetUserInfoById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/user/user.proto",
}