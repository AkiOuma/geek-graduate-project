// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// RecordServiceClient is the client API for RecordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordServiceClient interface {
	ClockInOnWork(ctx context.Context, in *ClockInOnWorkRequest, opts ...grpc.CallOption) (*ClockInOnWorkReply, error)
	ClockInOffWork(ctx context.Context, in *ClockInOffWorkRequest, opts ...grpc.CallOption) (*ClockInOffWorkReply, error)
}

type recordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordServiceClient(cc grpc.ClientConnInterface) RecordServiceClient {
	return &recordServiceClient{cc}
}

func (c *recordServiceClient) ClockInOnWork(ctx context.Context, in *ClockInOnWorkRequest, opts ...grpc.CallOption) (*ClockInOnWorkReply, error) {
	out := new(ClockInOnWorkReply)
	err := c.cc.Invoke(ctx, "/api.record.v1.RecordService/ClockInOnWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) ClockInOffWork(ctx context.Context, in *ClockInOffWorkRequest, opts ...grpc.CallOption) (*ClockInOffWorkReply, error) {
	out := new(ClockInOffWorkReply)
	err := c.cc.Invoke(ctx, "/api.record.v1.RecordService/ClockInOffWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordServiceServer is the server API for RecordService service.
// All implementations must embed UnimplementedRecordServiceServer
// for forward compatibility
type RecordServiceServer interface {
	ClockInOnWork(context.Context, *ClockInOnWorkRequest) (*ClockInOnWorkReply, error)
	ClockInOffWork(context.Context, *ClockInOffWorkRequest) (*ClockInOffWorkReply, error)
	mustEmbedUnimplementedRecordServiceServer()
}

// UnimplementedRecordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecordServiceServer struct {
}

func (UnimplementedRecordServiceServer) ClockInOnWork(context.Context, *ClockInOnWorkRequest) (*ClockInOnWorkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClockInOnWork not implemented")
}
func (UnimplementedRecordServiceServer) ClockInOffWork(context.Context, *ClockInOffWorkRequest) (*ClockInOffWorkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClockInOffWork not implemented")
}
func (UnimplementedRecordServiceServer) mustEmbedUnimplementedRecordServiceServer() {}

// UnsafeRecordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordServiceServer will
// result in compilation errors.
type UnsafeRecordServiceServer interface {
	mustEmbedUnimplementedRecordServiceServer()
}

func RegisterRecordServiceServer(s grpc.ServiceRegistrar, srv RecordServiceServer) {
	s.RegisterService(&RecordService_ServiceDesc, srv)
}

func _RecordService_ClockInOnWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClockInOnWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).ClockInOnWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.record.v1.RecordService/ClockInOnWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).ClockInOnWork(ctx, req.(*ClockInOnWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_ClockInOffWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClockInOffWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).ClockInOffWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.record.v1.RecordService/ClockInOffWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).ClockInOffWork(ctx, req.(*ClockInOffWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecordService_ServiceDesc is the grpc.ServiceDesc for RecordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.record.v1.RecordService",
	HandlerType: (*RecordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClockInOnWork",
			Handler:    _RecordService_ClockInOnWork_Handler,
		},
		{
			MethodName: "ClockInOffWork",
			Handler:    _RecordService_ClockInOffWork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/record/v1/record.proto",
}
