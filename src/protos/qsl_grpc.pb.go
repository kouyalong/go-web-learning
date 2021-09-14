// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// QSLEngineClient is the client API for QSLEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QSLEngineClient interface {
	// 异步接口
	Query(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetResult(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error)
	Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error)
	GetFields(ctx context.Context, in *FieldRequest, opts ...grpc.CallOption) (*FieldResponse, error)
	GroupCount(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*GroupCountResponse, error)
	SimpleQuery(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SyncResponse, error)
	GetSourceConfig(ctx context.Context, in *SourceConfigRequest, opts ...grpc.CallOption) (*SourceConfigResponse, error)
	GetShowInfo(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error)
}

type qSLEngineClient struct {
	cc grpc.ClientConnInterface
}

func NewQSLEngineClient(cc grpc.ClientConnInterface) QSLEngineClient {
	return &qSLEngineClient{cc}
}

func (c *qSLEngineClient) Query(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/QSLEngine/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qSLEngineClient) GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/QSLEngine/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qSLEngineClient) GetResult(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/QSLEngine/GetResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qSLEngineClient) Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error) {
	out := new(TerminateResponse)
	err := c.cc.Invoke(ctx, "/QSLEngine/Terminate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qSLEngineClient) GetFields(ctx context.Context, in *FieldRequest, opts ...grpc.CallOption) (*FieldResponse, error) {
	out := new(FieldResponse)
	err := c.cc.Invoke(ctx, "/QSLEngine/GetFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qSLEngineClient) GroupCount(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*GroupCountResponse, error) {
	out := new(GroupCountResponse)
	err := c.cc.Invoke(ctx, "/QSLEngine/GroupCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qSLEngineClient) SimpleQuery(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, "/QSLEngine/SimpleQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qSLEngineClient) GetSourceConfig(ctx context.Context, in *SourceConfigRequest, opts ...grpc.CallOption) (*SourceConfigResponse, error) {
	out := new(SourceConfigResponse)
	err := c.cc.Invoke(ctx, "/QSLEngine/GetSourceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qSLEngineClient) GetShowInfo(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error) {
	out := new(ShowResponse)
	err := c.cc.Invoke(ctx, "/QSLEngine/GetShowInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QSLEngineServer is the server API for QSLEngine service.
// All implementations must embed UnimplementedQSLEngineServer
// for forward compatibility
type QSLEngineServer interface {
	// 异步接口
	Query(context.Context, *SearchRequest) (*SearchResponse, error)
	GetStatus(context.Context, *StatusRequest) (*StatusResponse, error)
	GetResult(context.Context, *ResultRequest) (*ResultResponse, error)
	Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error)
	GetFields(context.Context, *FieldRequest) (*FieldResponse, error)
	GroupCount(context.Context, *SearchRequest) (*GroupCountResponse, error)
	SimpleQuery(context.Context, *SearchRequest) (*SyncResponse, error)
	GetSourceConfig(context.Context, *SourceConfigRequest) (*SourceConfigResponse, error)
	GetShowInfo(context.Context, *ShowRequest) (*ShowResponse, error)
	mustEmbedUnimplementedQSLEngineServer()
}

// UnimplementedQSLEngineServer must be embedded to have forward compatible implementations.
type UnimplementedQSLEngineServer struct {
}

func (UnimplementedQSLEngineServer) Query(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedQSLEngineServer) GetStatus(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedQSLEngineServer) GetResult(context.Context, *ResultRequest) (*ResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (UnimplementedQSLEngineServer) Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Terminate not implemented")
}
func (UnimplementedQSLEngineServer) GetFields(context.Context, *FieldRequest) (*FieldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFields not implemented")
}
func (UnimplementedQSLEngineServer) GroupCount(context.Context, *SearchRequest) (*GroupCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupCount not implemented")
}
func (UnimplementedQSLEngineServer) SimpleQuery(context.Context, *SearchRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimpleQuery not implemented")
}
func (UnimplementedQSLEngineServer) GetSourceConfig(context.Context, *SourceConfigRequest) (*SourceConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSourceConfig not implemented")
}
func (UnimplementedQSLEngineServer) GetShowInfo(context.Context, *ShowRequest) (*ShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShowInfo not implemented")
}
func (UnimplementedQSLEngineServer) mustEmbedUnimplementedQSLEngineServer() {}

// UnsafeQSLEngineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QSLEngineServer will
// result in compilation errors.
type UnsafeQSLEngineServer interface {
	mustEmbedUnimplementedQSLEngineServer()
}

func RegisterQSLEngineServer(s grpc.ServiceRegistrar, srv QSLEngineServer) {
	s.RegisterService(&QSLEngine_ServiceDesc, srv)
}

func _QSLEngine_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QSLEngineServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QSLEngine/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QSLEngineServer).Query(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QSLEngine_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QSLEngineServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QSLEngine/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QSLEngineServer).GetStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QSLEngine_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QSLEngineServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QSLEngine/GetResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QSLEngineServer).GetResult(ctx, req.(*ResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QSLEngine_Terminate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QSLEngineServer).Terminate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QSLEngine/Terminate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QSLEngineServer).Terminate(ctx, req.(*TerminateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QSLEngine_GetFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QSLEngineServer).GetFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QSLEngine/GetFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QSLEngineServer).GetFields(ctx, req.(*FieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QSLEngine_GroupCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QSLEngineServer).GroupCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QSLEngine/GroupCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QSLEngineServer).GroupCount(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QSLEngine_SimpleQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QSLEngineServer).SimpleQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QSLEngine/SimpleQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QSLEngineServer).SimpleQuery(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QSLEngine_GetSourceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SourceConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QSLEngineServer).GetSourceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QSLEngine/GetSourceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QSLEngineServer).GetSourceConfig(ctx, req.(*SourceConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QSLEngine_GetShowInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QSLEngineServer).GetShowInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QSLEngine/GetShowInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QSLEngineServer).GetShowInfo(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QSLEngine_ServiceDesc is the grpc.ServiceDesc for QSLEngine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QSLEngine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "QSLEngine",
	HandlerType: (*QSLEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _QSLEngine_Query_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _QSLEngine_GetStatus_Handler,
		},
		{
			MethodName: "GetResult",
			Handler:    _QSLEngine_GetResult_Handler,
		},
		{
			MethodName: "Terminate",
			Handler:    _QSLEngine_Terminate_Handler,
		},
		{
			MethodName: "GetFields",
			Handler:    _QSLEngine_GetFields_Handler,
		},
		{
			MethodName: "GroupCount",
			Handler:    _QSLEngine_GroupCount_Handler,
		},
		{
			MethodName: "SimpleQuery",
			Handler:    _QSLEngine_SimpleQuery_Handler,
		},
		{
			MethodName: "GetSourceConfig",
			Handler:    _QSLEngine_GetSourceConfig_Handler,
		},
		{
			MethodName: "GetShowInfo",
			Handler:    _QSLEngine_GetShowInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qsl.proto",
}
