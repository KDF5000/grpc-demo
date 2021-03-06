package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/chat_room/chatroom"
)

type Server struct {
	host    string
	port    int
	msg     []pb.ChatMessage
	outChan chan *pb.ChatMessage
}

// SaySomething  An interface generated by grpc
func (s *Server) SaySomething(stream pb.ChatRoom_SaySomethingServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Recv io.EOF!")
			return nil
		}
		if err != nil {
			return err
		}
		s.outChan <- in
		//send random k messages
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(s.msg))
		count := rand.Intn(2) + 1
		for i := index; i < len(s.msg) && i < index+count; i++ {
			newMsg := &pb.ChatMessage{Username: s.msg[i].Username, Msg: in.Msg + "-" + s.msg[i].Msg}
			if err := stream.Send(newMsg); err != nil {
				// fmt.Println(err)
				return err
			}
		}
	}
}

func main() {
	s := &Server{host: "localhost", port: 9001, outChan: make(chan *pb.ChatMessage)}
	for i := 1; i <= 10; i++ {
		m := pb.ChatMessage{Username: "Nginx", Msg: fmt.Sprintf("Response Message %d", i)}
		s.msg = append(s.msg, m)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		for {
			select {
			case cmsg := <-s.outChan:
				fmt.Printf("%s: %s\n", cmsg.Username, cmsg.Msg)
			default:
				break
			}
		}
	}()
	grpcServer := grpc.NewServer()
	pb.RegisterChatRoomServer(grpcServer, s)
	grpcServer.Serve(lis)
}
