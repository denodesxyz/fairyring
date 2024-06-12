// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: fairyring/pep/tx.proto

package pep

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
	Msg_UpdateParams_FullMethodName             = "/fairyring.pep.Msg/UpdateParams"
	Msg_SubmitEncryptedTx_FullMethodName        = "/fairyring.pep.Msg/SubmitEncryptedTx"
	Msg_SubmitGeneralEncryptedTx_FullMethodName = "/fairyring.pep.Msg/SubmitGeneralEncryptedTx"
	Msg_CreateAggregatedKeyShare_FullMethodName = "/fairyring.pep.Msg/CreateAggregatedKeyShare"
	Msg_RequestGeneralKeyshare_FullMethodName   = "/fairyring.pep.Msg/RequestGeneralKeyshare"
	Msg_GetGeneralKeyshare_FullMethodName       = "/fairyring.pep.Msg/GetGeneralKeyshare"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	SubmitEncryptedTx(ctx context.Context, in *MsgSubmitEncryptedTx, opts ...grpc.CallOption) (*MsgSubmitEncryptedTxResponse, error)
	SubmitGeneralEncryptedTx(ctx context.Context, in *MsgSubmitGeneralEncryptedTx, opts ...grpc.CallOption) (*MsgSubmitEncryptedTxResponse, error)
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateAggregatedKeyShare(ctx context.Context, in *MsgCreateAggregatedKeyShare, opts ...grpc.CallOption) (*MsgCreateAggregatedKeyShareResponse, error)
	RequestGeneralKeyshare(ctx context.Context, in *MsgRequestGeneralKeyshare, opts ...grpc.CallOption) (*MsgRequestGeneralKeyshareResponse, error)
	GetGeneralKeyshare(ctx context.Context, in *MsgGetGeneralKeyshare, opts ...grpc.CallOption) (*MsgGetGeneralKeyshareResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitEncryptedTx(ctx context.Context, in *MsgSubmitEncryptedTx, opts ...grpc.CallOption) (*MsgSubmitEncryptedTxResponse, error) {
	out := new(MsgSubmitEncryptedTxResponse)
	err := c.cc.Invoke(ctx, Msg_SubmitEncryptedTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitGeneralEncryptedTx(ctx context.Context, in *MsgSubmitGeneralEncryptedTx, opts ...grpc.CallOption) (*MsgSubmitEncryptedTxResponse, error) {
	out := new(MsgSubmitEncryptedTxResponse)
	err := c.cc.Invoke(ctx, Msg_SubmitGeneralEncryptedTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateAggregatedKeyShare(ctx context.Context, in *MsgCreateAggregatedKeyShare, opts ...grpc.CallOption) (*MsgCreateAggregatedKeyShareResponse, error) {
	out := new(MsgCreateAggregatedKeyShareResponse)
	err := c.cc.Invoke(ctx, Msg_CreateAggregatedKeyShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RequestGeneralKeyshare(ctx context.Context, in *MsgRequestGeneralKeyshare, opts ...grpc.CallOption) (*MsgRequestGeneralKeyshareResponse, error) {
	out := new(MsgRequestGeneralKeyshareResponse)
	err := c.cc.Invoke(ctx, Msg_RequestGeneralKeyshare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetGeneralKeyshare(ctx context.Context, in *MsgGetGeneralKeyshare, opts ...grpc.CallOption) (*MsgGetGeneralKeyshareResponse, error) {
	out := new(MsgGetGeneralKeyshareResponse)
	err := c.cc.Invoke(ctx, Msg_GetGeneralKeyshare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	SubmitEncryptedTx(context.Context, *MsgSubmitEncryptedTx) (*MsgSubmitEncryptedTxResponse, error)
	SubmitGeneralEncryptedTx(context.Context, *MsgSubmitGeneralEncryptedTx) (*MsgSubmitEncryptedTxResponse, error)
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateAggregatedKeyShare(context.Context, *MsgCreateAggregatedKeyShare) (*MsgCreateAggregatedKeyShareResponse, error)
	RequestGeneralKeyshare(context.Context, *MsgRequestGeneralKeyshare) (*MsgRequestGeneralKeyshareResponse, error)
	GetGeneralKeyshare(context.Context, *MsgGetGeneralKeyshare) (*MsgGetGeneralKeyshareResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) SubmitEncryptedTx(context.Context, *MsgSubmitEncryptedTx) (*MsgSubmitEncryptedTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitEncryptedTx not implemented")
}
func (UnimplementedMsgServer) SubmitGeneralEncryptedTx(context.Context, *MsgSubmitGeneralEncryptedTx) (*MsgSubmitEncryptedTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitGeneralEncryptedTx not implemented")
}
func (UnimplementedMsgServer) CreateAggregatedKeyShare(context.Context, *MsgCreateAggregatedKeyShare) (*MsgCreateAggregatedKeyShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAggregatedKeyShare not implemented")
}
func (UnimplementedMsgServer) RequestGeneralKeyshare(context.Context, *MsgRequestGeneralKeyshare) (*MsgRequestGeneralKeyshareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestGeneralKeyshare not implemented")
}
func (UnimplementedMsgServer) GetGeneralKeyshare(context.Context, *MsgGetGeneralKeyshare) (*MsgGetGeneralKeyshareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeneralKeyshare not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitEncryptedTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitEncryptedTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitEncryptedTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SubmitEncryptedTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitEncryptedTx(ctx, req.(*MsgSubmitEncryptedTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitGeneralEncryptedTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitGeneralEncryptedTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitGeneralEncryptedTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SubmitGeneralEncryptedTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitGeneralEncryptedTx(ctx, req.(*MsgSubmitGeneralEncryptedTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateAggregatedKeyShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateAggregatedKeyShare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateAggregatedKeyShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateAggregatedKeyShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateAggregatedKeyShare(ctx, req.(*MsgCreateAggregatedKeyShare))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RequestGeneralKeyshare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestGeneralKeyshare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestGeneralKeyshare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RequestGeneralKeyshare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestGeneralKeyshare(ctx, req.(*MsgRequestGeneralKeyshare))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetGeneralKeyshare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGetGeneralKeyshare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetGeneralKeyshare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_GetGeneralKeyshare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetGeneralKeyshare(ctx, req.(*MsgGetGeneralKeyshare))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fairyring.pep.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "SubmitEncryptedTx",
			Handler:    _Msg_SubmitEncryptedTx_Handler,
		},
		{
			MethodName: "SubmitGeneralEncryptedTx",
			Handler:    _Msg_SubmitGeneralEncryptedTx_Handler,
		},
		{
			MethodName: "CreateAggregatedKeyShare",
			Handler:    _Msg_CreateAggregatedKeyShare_Handler,
		},
		{
			MethodName: "RequestGeneralKeyshare",
			Handler:    _Msg_RequestGeneralKeyshare_Handler,
		},
		{
			MethodName: "GetGeneralKeyshare",
			Handler:    _Msg_GetGeneralKeyshare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fairyring/pep/tx.proto",
}
