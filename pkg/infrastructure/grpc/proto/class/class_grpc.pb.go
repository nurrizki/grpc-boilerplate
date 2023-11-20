// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: class.proto

package miniProject

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

// ClassServiceClient is the client API for ClassService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClassServiceClient interface {
	GetClass(ctx context.Context, in *GetClassRequest, opts ...grpc.CallOption) (*GetClassResponse, error)
	GetClassDistinct(ctx context.Context, in *GetClassDistinctRequest, opts ...grpc.CallOption) (*GetClassDistinctResponse, error)
}

type classServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClassServiceClient(cc grpc.ClientConnInterface) ClassServiceClient {
	return &classServiceClient{cc}
}

func (c *classServiceClient) GetClass(ctx context.Context, in *GetClassRequest, opts ...grpc.CallOption) (*GetClassResponse, error) {
	out := new(GetClassResponse)
	err := c.cc.Invoke(ctx, "/miniProject.ClassService/GetClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) GetClassDistinct(ctx context.Context, in *GetClassDistinctRequest, opts ...grpc.CallOption) (*GetClassDistinctResponse, error) {
	out := new(GetClassDistinctResponse)
	err := c.cc.Invoke(ctx, "/miniProject.ClassService/GetClassDistinct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClassServiceServer is the server API for ClassService service.
// All implementations should embed UnimplementedClassServiceServer
// for forward compatibility
type ClassServiceServer interface {
	GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error)
	GetClassDistinct(context.Context, *GetClassDistinctRequest) (*GetClassDistinctResponse, error)
}

// UnimplementedClassServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClassServiceServer struct {
}

func (UnimplementedClassServiceServer) GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClass not implemented")
}
func (UnimplementedClassServiceServer) GetClassDistinct(context.Context, *GetClassDistinctRequest) (*GetClassDistinctResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassDistinct not implemented")
}

// UnsafeClassServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClassServiceServer will
// result in compilation errors.
type UnsafeClassServiceServer interface {
	mustEmbedUnimplementedClassServiceServer()
}

func RegisterClassServiceServer(s grpc.ServiceRegistrar, srv ClassServiceServer) {
	s.RegisterService(&ClassService_ServiceDesc, srv)
}

func _ClassService_GetClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).GetClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miniProject.ClassService/GetClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).GetClass(ctx, req.(*GetClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_GetClassDistinct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassDistinctRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).GetClassDistinct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miniProject.ClassService/GetClassDistinct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).GetClassDistinct(ctx, req.(*GetClassDistinctRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClassService_ServiceDesc is the grpc.ServiceDesc for ClassService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClassService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miniProject.ClassService",
	HandlerType: (*ClassServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClass",
			Handler:    _ClassService_GetClass_Handler,
		},
		{
			MethodName: "GetClassDistinct",
			Handler:    _ClassService_GetClassDistinct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "class.proto",
}