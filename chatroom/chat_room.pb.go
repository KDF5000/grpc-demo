// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat_room.proto

/*
Package chatroom is a generated protocol buffer package.

It is generated from these files:
	chat_room.proto

It has these top-level messages:
	ChatMessage
*/
package chatroom

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ChatMessage struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Msg      string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *ChatMessage) Reset()                    { *m = ChatMessage{} }
func (m *ChatMessage) String() string            { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()               {}
func (*ChatMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ChatMessage) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ChatMessage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ChatMessage)(nil), "chatroom.ChatMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChatRoom service

type ChatRoomClient interface {
	SaySomething(ctx context.Context, opts ...grpc.CallOption) (ChatRoom_SaySomethingClient, error)
}

type chatRoomClient struct {
	cc *grpc.ClientConn
}

func NewChatRoomClient(cc *grpc.ClientConn) ChatRoomClient {
	return &chatRoomClient{cc}
}

func (c *chatRoomClient) SaySomething(ctx context.Context, opts ...grpc.CallOption) (ChatRoom_SaySomethingClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChatRoom_serviceDesc.Streams[0], c.cc, "/chatroom.ChatRoom/SaySomething", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatRoomSaySomethingClient{stream}
	return x, nil
}

type ChatRoom_SaySomethingClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatRoomSaySomethingClient struct {
	grpc.ClientStream
}

func (x *chatRoomSaySomethingClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatRoomSaySomethingClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ChatRoom service

type ChatRoomServer interface {
	SaySomething(ChatRoom_SaySomethingServer) error
}

func RegisterChatRoomServer(s *grpc.Server, srv ChatRoomServer) {
	s.RegisterService(&_ChatRoom_serviceDesc, srv)
}

func _ChatRoom_SaySomething_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatRoomServer).SaySomething(&chatRoomSaySomethingServer{stream})
}

type ChatRoom_SaySomethingServer interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chatRoomSaySomethingServer struct {
	grpc.ServerStream
}

func (x *chatRoomSaySomethingServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatRoomSaySomethingServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatRoom_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chatroom.ChatRoom",
	HandlerType: (*ChatRoomServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SaySomething",
			Handler:       _ChatRoom_SaySomething_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat_room.proto",
}

func init() { proto.RegisterFile("chat_room.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xce, 0x48, 0x2c,
	0x89, 0x2f, 0xca, 0xcf, 0xcf, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x09, 0x80,
	0xf8, 0x4a, 0xd6, 0x5c, 0xdc, 0xce, 0x19, 0x89, 0x25, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9,
	0x42, 0x52, 0x5c, 0x1c, 0xa5, 0xc5, 0xa9, 0x45, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x00, 0x17, 0x73, 0x6e, 0x71, 0xba, 0x04, 0x13, 0x58, 0x18,
	0xc4, 0x34, 0xf2, 0xe3, 0xe2, 0x00, 0x69, 0x0e, 0xca, 0xcf, 0xcf, 0x15, 0x72, 0xe2, 0xe2, 0x09,
	0x4e, 0xac, 0x0c, 0xce, 0xcf, 0x4d, 0x2d, 0xc9, 0xc8, 0xcc, 0x4b, 0x17, 0x12, 0xd5, 0x83, 0xd9,
	0xa1, 0x87, 0x64, 0x81, 0x14, 0x76, 0x61, 0x25, 0x06, 0x0d, 0x46, 0x03, 0xc6, 0x24, 0x36, 0xb0,
	0xeb, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x30, 0x3a, 0xc1, 0x65, 0xb0, 0x00, 0x00, 0x00,
}
