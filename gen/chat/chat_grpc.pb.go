// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chat

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

// ChatApiClient is the client API for ChatApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatApiClient interface {
	// The Connect RPC method takes a User as input and returns a stream of Messages
	Connect(ctx context.Context, in *User, opts ...grpc.CallOption) (ChatApi_ConnectClient, error)
	// The Broadcast RPC method takes a Message as input and returns a Message
	Broadcast(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type chatApiClient struct {
	cc grpc.ClientConnInterface
}

func NewChatApiClient(cc grpc.ClientConnInterface) ChatApiClient {
	return &chatApiClient{cc}
}

func (c *chatApiClient) Connect(ctx context.Context, in *User, opts ...grpc.CallOption) (ChatApi_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatApi_ServiceDesc.Streams[0], "/main.ChatApi/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatApiConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatApi_ConnectClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatApiConnectClient struct {
	grpc.ClientStream
}

func (x *chatApiConnectClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatApiClient) Broadcast(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/main.ChatApi/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatApiServer is the server API for ChatApi service.
// All implementations must embed UnimplementedChatApiServer
// for forward compatibility
type ChatApiServer interface {
	// The Connect RPC method takes a User as input and returns a stream of Messages
	Connect(*User, ChatApi_ConnectServer) error
	// The Broadcast RPC method takes a Message as input and returns a Message
	Broadcast(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedChatApiServer()
}

// UnimplementedChatApiServer must be embedded to have forward compatible implementations.
type UnimplementedChatApiServer struct {
}

func (UnimplementedChatApiServer) Connect(*User, ChatApi_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedChatApiServer) Broadcast(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedChatApiServer) mustEmbedUnimplementedChatApiServer() {}

// UnsafeChatApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatApiServer will
// result in compilation errors.
type UnsafeChatApiServer interface {
	mustEmbedUnimplementedChatApiServer()
}

func RegisterChatApiServer(s grpc.ServiceRegistrar, srv ChatApiServer) {
	s.RegisterService(&ChatApi_ServiceDesc, srv)
}

func _ChatApi_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatApiServer).Connect(m, &chatApiConnectServer{stream})
}

type ChatApi_ConnectServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatApiConnectServer struct {
	grpc.ServerStream
}

func (x *chatApiConnectServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatApi_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatApiServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ChatApi/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatApiServer).Broadcast(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatApi_ServiceDesc is the grpc.ServiceDesc for ChatApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.ChatApi",
	HandlerType: (*ChatApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _ChatApi_Broadcast_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _ChatApi_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}
