// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: proto.proto

package gen

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

// MqttMiddlewareV1Client is the client API for MqttMiddlewareV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MqttMiddlewareV1Client interface {
	PushPowerBank(ctx context.Context, in *CommandPush, opts ...grpc.CallOption) (*ResponsePush, error)
	ForcePushPowerBank(ctx context.Context, in *CommandPush, opts ...grpc.CallOption) (*ResponsePush, error)
	QueryInventory(ctx context.Context, in *CommandInventory, opts ...grpc.CallOption) (*ResponseInventory, error)
	QueryServerInformation(ctx context.Context, in *CommandServerInformation, opts ...grpc.CallOption) (*ResponseServerInformation, error)
	QueryCabinetAPN(ctx context.Context, in *CommandCabinetAPN, opts ...grpc.CallOption) (*ResponseCabinetAPN, error)
	QuerySIMCardICCID(ctx context.Context, in *CommandSIMCardICCID, opts ...grpc.CallOption) (*ResponseSIMCardICCID, error)
	QueryNetworkInformation(ctx context.Context, in *CommandNetworkInformation, opts ...grpc.CallOption) (*ResponseNetworkInformation, error)
	ResetCabinet(ctx context.Context, in *CommandResetCabinet, opts ...grpc.CallOption) (*ResponseResetCabinet, error)
	Subscribe(ctx context.Context, in *Device, opts ...grpc.CallOption) (*ResponseString, error)
}

type mqttMiddlewareV1Client struct {
	cc grpc.ClientConnInterface
}

func NewMqttMiddlewareV1Client(cc grpc.ClientConnInterface) MqttMiddlewareV1Client {
	return &mqttMiddlewareV1Client{cc}
}

func (c *mqttMiddlewareV1Client) PushPowerBank(ctx context.Context, in *CommandPush, opts ...grpc.CallOption) (*ResponsePush, error) {
	out := new(ResponsePush)
	err := c.cc.Invoke(ctx, "/wcharge_mqtt.MqttMiddlewareV1/PushPowerBank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttMiddlewareV1Client) ForcePushPowerBank(ctx context.Context, in *CommandPush, opts ...grpc.CallOption) (*ResponsePush, error) {
	out := new(ResponsePush)
	err := c.cc.Invoke(ctx, "/wcharge_mqtt.MqttMiddlewareV1/ForcePushPowerBank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttMiddlewareV1Client) QueryInventory(ctx context.Context, in *CommandInventory, opts ...grpc.CallOption) (*ResponseInventory, error) {
	out := new(ResponseInventory)
	err := c.cc.Invoke(ctx, "/wcharge_mqtt.MqttMiddlewareV1/QueryInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttMiddlewareV1Client) QueryServerInformation(ctx context.Context, in *CommandServerInformation, opts ...grpc.CallOption) (*ResponseServerInformation, error) {
	out := new(ResponseServerInformation)
	err := c.cc.Invoke(ctx, "/wcharge_mqtt.MqttMiddlewareV1/QueryServerInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttMiddlewareV1Client) QueryCabinetAPN(ctx context.Context, in *CommandCabinetAPN, opts ...grpc.CallOption) (*ResponseCabinetAPN, error) {
	out := new(ResponseCabinetAPN)
	err := c.cc.Invoke(ctx, "/wcharge_mqtt.MqttMiddlewareV1/QueryCabinetAPN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttMiddlewareV1Client) QuerySIMCardICCID(ctx context.Context, in *CommandSIMCardICCID, opts ...grpc.CallOption) (*ResponseSIMCardICCID, error) {
	out := new(ResponseSIMCardICCID)
	err := c.cc.Invoke(ctx, "/wcharge_mqtt.MqttMiddlewareV1/QuerySIMCardICCID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttMiddlewareV1Client) QueryNetworkInformation(ctx context.Context, in *CommandNetworkInformation, opts ...grpc.CallOption) (*ResponseNetworkInformation, error) {
	out := new(ResponseNetworkInformation)
	err := c.cc.Invoke(ctx, "/wcharge_mqtt.MqttMiddlewareV1/QueryNetworkInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttMiddlewareV1Client) ResetCabinet(ctx context.Context, in *CommandResetCabinet, opts ...grpc.CallOption) (*ResponseResetCabinet, error) {
	out := new(ResponseResetCabinet)
	err := c.cc.Invoke(ctx, "/wcharge_mqtt.MqttMiddlewareV1/ResetCabinet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttMiddlewareV1Client) Subscribe(ctx context.Context, in *Device, opts ...grpc.CallOption) (*ResponseString, error) {
	out := new(ResponseString)
	err := c.cc.Invoke(ctx, "/wcharge_mqtt.MqttMiddlewareV1/SubscribeMqtt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MqttMiddlewareV1Server is the server API for MqttMiddlewareV1 service.
// All implementations must embed UnimplementedMqttMiddlewareV1Server
// for forward compatibility
type MqttMiddlewareV1Server interface {
	PushPowerBank(context.Context, *CommandPush) (*ResponsePush, error)
	ForcePushPowerBank(context.Context, *CommandPush) (*ResponsePush, error)
	QueryInventory(context.Context, *CommandInventory) (*ResponseInventory, error)
	QueryServerInformation(context.Context, *CommandServerInformation) (*ResponseServerInformation, error)
	QueryCabinetAPN(context.Context, *CommandCabinetAPN) (*ResponseCabinetAPN, error)
	QuerySIMCardICCID(context.Context, *CommandSIMCardICCID) (*ResponseSIMCardICCID, error)
	QueryNetworkInformation(context.Context, *CommandNetworkInformation) (*ResponseNetworkInformation, error)
	ResetCabinet(context.Context, *CommandResetCabinet) (*ResponseResetCabinet, error)
	Subscribe(context.Context, *Device) (*ResponseString, error)
	mustEmbedUnimplementedMqttMiddlewareV1Server()
}

// UnimplementedMqttMiddlewareV1Server must be embedded to have forward compatible implementations.
type UnimplementedMqttMiddlewareV1Server struct {
}

func (UnimplementedMqttMiddlewareV1Server) PushPowerBank(context.Context, *CommandPush) (*ResponsePush, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushPowerBank not implemented")
}
func (UnimplementedMqttMiddlewareV1Server) ForcePushPowerBank(context.Context, *CommandPush) (*ResponsePush, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForcePushPowerBank not implemented")
}
func (UnimplementedMqttMiddlewareV1Server) QueryInventory(context.Context, *CommandInventory) (*ResponseInventory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryInventory not implemented")
}
func (UnimplementedMqttMiddlewareV1Server) QueryServerInformation(context.Context, *CommandServerInformation) (*ResponseServerInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryServerInformation not implemented")
}
func (UnimplementedMqttMiddlewareV1Server) QueryCabinetAPN(context.Context, *CommandCabinetAPN) (*ResponseCabinetAPN, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCabinetAPN not implemented")
}
func (UnimplementedMqttMiddlewareV1Server) QuerySIMCardICCID(context.Context, *CommandSIMCardICCID) (*ResponseSIMCardICCID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySIMCardICCID not implemented")
}
func (UnimplementedMqttMiddlewareV1Server) QueryNetworkInformation(context.Context, *CommandNetworkInformation) (*ResponseNetworkInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryNetworkInformation not implemented")
}
func (UnimplementedMqttMiddlewareV1Server) ResetCabinet(context.Context, *CommandResetCabinet) (*ResponseResetCabinet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetCabinet not implemented")
}
func (UnimplementedMqttMiddlewareV1Server) Subscribe(context.Context, *Device) (*ResponseString, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeMqtt not implemented")
}
func (UnimplementedMqttMiddlewareV1Server) mustEmbedUnimplementedMqttMiddlewareV1Server() {}

// UnsafeMqttMiddlewareV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MqttMiddlewareV1Server will
// result in compilation errors.
type UnsafeMqttMiddlewareV1Server interface {
	mustEmbedUnimplementedMqttMiddlewareV1Server()
}

func RegisterMqttMiddlewareV1Server(s grpc.ServiceRegistrar, srv MqttMiddlewareV1Server) {
	s.RegisterService(&MqttMiddlewareV1_ServiceDesc, srv)
}

func _MqttMiddlewareV1_PushPowerBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandPush)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttMiddlewareV1Server).PushPowerBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcharge_mqtt.MqttMiddlewareV1/PushPowerBank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttMiddlewareV1Server).PushPowerBank(ctx, req.(*CommandPush))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttMiddlewareV1_ForcePushPowerBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandPush)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttMiddlewareV1Server).ForcePushPowerBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcharge_mqtt.MqttMiddlewareV1/ForcePushPowerBank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttMiddlewareV1Server).ForcePushPowerBank(ctx, req.(*CommandPush))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttMiddlewareV1_QueryInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandInventory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttMiddlewareV1Server).QueryInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcharge_mqtt.MqttMiddlewareV1/QueryInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttMiddlewareV1Server).QueryInventory(ctx, req.(*CommandInventory))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttMiddlewareV1_QueryServerInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandServerInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttMiddlewareV1Server).QueryServerInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcharge_mqtt.MqttMiddlewareV1/QueryServerInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttMiddlewareV1Server).QueryServerInformation(ctx, req.(*CommandServerInformation))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttMiddlewareV1_QueryCabinetAPN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandCabinetAPN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttMiddlewareV1Server).QueryCabinetAPN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcharge_mqtt.MqttMiddlewareV1/QueryCabinetAPN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttMiddlewareV1Server).QueryCabinetAPN(ctx, req.(*CommandCabinetAPN))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttMiddlewareV1_QuerySIMCardICCID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandSIMCardICCID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttMiddlewareV1Server).QuerySIMCardICCID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcharge_mqtt.MqttMiddlewareV1/QuerySIMCardICCID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttMiddlewareV1Server).QuerySIMCardICCID(ctx, req.(*CommandSIMCardICCID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttMiddlewareV1_QueryNetworkInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandNetworkInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttMiddlewareV1Server).QueryNetworkInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcharge_mqtt.MqttMiddlewareV1/QueryNetworkInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttMiddlewareV1Server).QueryNetworkInformation(ctx, req.(*CommandNetworkInformation))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttMiddlewareV1_ResetCabinet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandResetCabinet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttMiddlewareV1Server).ResetCabinet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcharge_mqtt.MqttMiddlewareV1/ResetCabinet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttMiddlewareV1Server).ResetCabinet(ctx, req.(*CommandResetCabinet))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttMiddlewareV1_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttMiddlewareV1Server).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcharge_mqtt.MqttMiddlewareV1/SubscribeMqtt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttMiddlewareV1Server).Subscribe(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

// MqttMiddlewareV1_ServiceDesc is the grpc.ServiceDesc for MqttMiddlewareV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MqttMiddlewareV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wcharge_mqtt.MqttMiddlewareV1",
	HandlerType: (*MqttMiddlewareV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushPowerBank",
			Handler:    _MqttMiddlewareV1_PushPowerBank_Handler,
		},
		{
			MethodName: "ForcePushPowerBank",
			Handler:    _MqttMiddlewareV1_ForcePushPowerBank_Handler,
		},
		{
			MethodName: "QueryInventory",
			Handler:    _MqttMiddlewareV1_QueryInventory_Handler,
		},
		{
			MethodName: "QueryServerInformation",
			Handler:    _MqttMiddlewareV1_QueryServerInformation_Handler,
		},
		{
			MethodName: "QueryCabinetAPN",
			Handler:    _MqttMiddlewareV1_QueryCabinetAPN_Handler,
		},
		{
			MethodName: "QuerySIMCardICCID",
			Handler:    _MqttMiddlewareV1_QuerySIMCardICCID_Handler,
		},
		{
			MethodName: "QueryNetworkInformation",
			Handler:    _MqttMiddlewareV1_QueryNetworkInformation_Handler,
		},
		{
			MethodName: "ResetCabinet",
			Handler:    _MqttMiddlewareV1_ResetCabinet_Handler,
		},
		{
			MethodName: "SubscribeMqtt",
			Handler:    _MqttMiddlewareV1_Subscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}

// MainServerV1Client is the client API for MainServerV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MainServerV1Client interface {
	ReturnPowerBank(ctx context.Context, in *RequestReturnPowerBank, opts ...grpc.CallOption) (*ResponseReturnPowerBank, error)
	ReportCabinetLogin(ctx context.Context, in *RequestReportCabinetLogin, opts ...grpc.CallOption) (*ResponseReportCabinetLogin, error)
}

type mainServerV1Client struct {
	cc grpc.ClientConnInterface
}

func NewMainServerV1Client(cc grpc.ClientConnInterface) MainServerV1Client {
	return &mainServerV1Client{cc}
}

func (c *mainServerV1Client) ReturnPowerBank(ctx context.Context, in *RequestReturnPowerBank, opts ...grpc.CallOption) (*ResponseReturnPowerBank, error) {
	out := new(ResponseReturnPowerBank)
	err := c.cc.Invoke(ctx, "/wcharge_mqtt.MainServerV1/ReturnPowerBank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainServerV1Client) ReportCabinetLogin(ctx context.Context, in *RequestReportCabinetLogin, opts ...grpc.CallOption) (*ResponseReportCabinetLogin, error) {
	out := new(ResponseReportCabinetLogin)
	err := c.cc.Invoke(ctx, "/wcharge_mqtt.MainServerV1/ReportCabinetLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainServerV1Server is the server API for MainServerV1 service.
// All implementations must embed UnimplementedMainServerV1Server
// for forward compatibility
type MainServerV1Server interface {
	ReturnPowerBank(context.Context, *RequestReturnPowerBank) (*ResponseReturnPowerBank, error)
	ReportCabinetLogin(context.Context, *RequestReportCabinetLogin) (*ResponseReportCabinetLogin, error)
	mustEmbedUnimplementedMainServerV1Server()
}

// UnimplementedMainServerV1Server must be embedded to have forward compatible implementations.
type UnimplementedMainServerV1Server struct {
}

func (UnimplementedMainServerV1Server) ReturnPowerBank(context.Context, *RequestReturnPowerBank) (*ResponseReturnPowerBank, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnPowerBank not implemented")
}
func (UnimplementedMainServerV1Server) ReportCabinetLogin(context.Context, *RequestReportCabinetLogin) (*ResponseReportCabinetLogin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportCabinetLogin not implemented")
}
func (UnimplementedMainServerV1Server) mustEmbedUnimplementedMainServerV1Server() {}

// UnsafeMainServerV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MainServerV1Server will
// result in compilation errors.
type UnsafeMainServerV1Server interface {
	mustEmbedUnimplementedMainServerV1Server()
}

func RegisterMainServerV1Server(s grpc.ServiceRegistrar, srv MainServerV1Server) {
	s.RegisterService(&MainServerV1_ServiceDesc, srv)
}

func _MainServerV1_ReturnPowerBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestReturnPowerBank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServerV1Server).ReturnPowerBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcharge_mqtt.MainServerV1/ReturnPowerBank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServerV1Server).ReturnPowerBank(ctx, req.(*RequestReturnPowerBank))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainServerV1_ReportCabinetLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestReportCabinetLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServerV1Server).ReportCabinetLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcharge_mqtt.MainServerV1/ReportCabinetLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServerV1Server).ReportCabinetLogin(ctx, req.(*RequestReportCabinetLogin))
	}
	return interceptor(ctx, in, info, handler)
}

// MainServerV1_ServiceDesc is the grpc.ServiceDesc for MainServerV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MainServerV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wcharge_mqtt.MainServerV1",
	HandlerType: (*MainServerV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReturnPowerBank",
			Handler:    _MainServerV1_ReturnPowerBank_Handler,
		},
		{
			MethodName: "ReportCabinetLogin",
			Handler:    _MainServerV1_ReportCabinetLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}
