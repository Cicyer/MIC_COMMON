// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package MessageService

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	//平台消息发送
	SendEmails(ctx context.Context, in *SendEmailsReq, opts ...grpc.CallOption) (*SendEmailsResp, error)
	SendSms(ctx context.Context, in *SendEmailsReq, opts ...grpc.CallOption) (*SendSmsResp, error)
	SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error)
	//平台消息管理
	GetMessageRecords(ctx context.Context, in *GetMessageRecordsReq, opts ...grpc.CallOption) (*MessageListResp, error)
	GetMessageBusiness(ctx context.Context, in *GetMessageBusinessReq, opts ...grpc.CallOption) (*MessageBusinessResp, error)
	ConsumeMessage(ctx context.Context, in *ConsumeMessageReq, opts ...grpc.CallOption) (*ConsumeMessageResp, error)
	UnreadCount(ctx context.Context, in *GetUnreadCountReq, opts ...grpc.CallOption) (*UnreadCountResp, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) SendEmails(ctx context.Context, in *SendEmailsReq, opts ...grpc.CallOption) (*SendEmailsResp, error) {
	out := new(SendEmailsResp)
	err := c.cc.Invoke(ctx, "/MessageService.MessageService/SendEmails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendSms(ctx context.Context, in *SendEmailsReq, opts ...grpc.CallOption) (*SendSmsResp, error) {
	out := new(SendSmsResp)
	err := c.cc.Invoke(ctx, "/MessageService.MessageService/SendSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error) {
	out := new(SendMessageResp)
	err := c.cc.Invoke(ctx, "/MessageService.MessageService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetMessageRecords(ctx context.Context, in *GetMessageRecordsReq, opts ...grpc.CallOption) (*MessageListResp, error) {
	out := new(MessageListResp)
	err := c.cc.Invoke(ctx, "/MessageService.MessageService/GetMessageRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetMessageBusiness(ctx context.Context, in *GetMessageBusinessReq, opts ...grpc.CallOption) (*MessageBusinessResp, error) {
	out := new(MessageBusinessResp)
	err := c.cc.Invoke(ctx, "/MessageService.MessageService/GetMessageBusiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) ConsumeMessage(ctx context.Context, in *ConsumeMessageReq, opts ...grpc.CallOption) (*ConsumeMessageResp, error) {
	out := new(ConsumeMessageResp)
	err := c.cc.Invoke(ctx, "/MessageService.MessageService/ConsumeMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) UnreadCount(ctx context.Context, in *GetUnreadCountReq, opts ...grpc.CallOption) (*UnreadCountResp, error) {
	out := new(UnreadCountResp)
	err := c.cc.Invoke(ctx, "/MessageService.MessageService/UnreadCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	//平台消息发送
	SendEmails(context.Context, *SendEmailsReq) (*SendEmailsResp, error)
	SendSms(context.Context, *SendEmailsReq) (*SendSmsResp, error)
	SendMessage(context.Context, *SendMessageReq) (*SendMessageResp, error)
	//平台消息管理
	GetMessageRecords(context.Context, *GetMessageRecordsReq) (*MessageListResp, error)
	GetMessageBusiness(context.Context, *GetMessageBusinessReq) (*MessageBusinessResp, error)
	ConsumeMessage(context.Context, *ConsumeMessageReq) (*ConsumeMessageResp, error)
	UnreadCount(context.Context, *GetUnreadCountReq) (*UnreadCountResp, error)
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) SendEmails(context.Context, *SendEmailsReq) (*SendEmailsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmails not implemented")
}
func (UnimplementedMessageServiceServer) SendSms(context.Context, *SendEmailsReq) (*SendSmsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}
func (UnimplementedMessageServiceServer) SendMessage(context.Context, *SendMessageReq) (*SendMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessageServiceServer) GetMessageRecords(context.Context, *GetMessageRecordsReq) (*MessageListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageRecords not implemented")
}
func (UnimplementedMessageServiceServer) GetMessageBusiness(context.Context, *GetMessageBusinessReq) (*MessageBusinessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageBusiness not implemented")
}
func (UnimplementedMessageServiceServer) ConsumeMessage(context.Context, *ConsumeMessageReq) (*ConsumeMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsumeMessage not implemented")
}
func (UnimplementedMessageServiceServer) UnreadCount(context.Context, *GetUnreadCountReq) (*UnreadCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnreadCount not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_SendEmails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendEmails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService.MessageService/SendEmails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendEmails(ctx, req.(*SendEmailsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService.MessageService/SendSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendSms(ctx, req.(*SendEmailsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService.MessageService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendMessage(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetMessageRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRecordsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetMessageRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService.MessageService/GetMessageRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetMessageRecords(ctx, req.(*GetMessageRecordsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetMessageBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageBusinessReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetMessageBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService.MessageService/GetMessageBusiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetMessageBusiness(ctx, req.(*GetMessageBusinessReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_ConsumeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).ConsumeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService.MessageService/ConsumeMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).ConsumeMessage(ctx, req.(*ConsumeMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_UnreadCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnreadCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).UnreadCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService.MessageService/UnreadCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).UnreadCount(ctx, req.(*GetUnreadCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MessageService.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmails",
			Handler:    _MessageService_SendEmails_Handler,
		},
		{
			MethodName: "SendSms",
			Handler:    _MessageService_SendSms_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _MessageService_SendMessage_Handler,
		},
		{
			MethodName: "GetMessageRecords",
			Handler:    _MessageService_GetMessageRecords_Handler,
		},
		{
			MethodName: "GetMessageBusiness",
			Handler:    _MessageService_GetMessageBusiness_Handler,
		},
		{
			MethodName: "ConsumeMessage",
			Handler:    _MessageService_ConsumeMessage_Handler,
		},
		{
			MethodName: "UnreadCount",
			Handler:    _MessageService_UnreadCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "MessageService.proto",
}
