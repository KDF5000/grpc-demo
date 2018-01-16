package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/chat_room/chatroom"
)

var (
	name = "KDF5000"
)

func recvMessage(stream pb.ChatRoom_SaySomethingClient, wg *sync.WaitGroup) {
	msgs := make([]*pb.ChatMessage, 0)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			// fmt.Println("Client recv io.EOF")
			break
		}
		if err != nil {
			break
		}
		msgs = append(msgs, in)
	}
	if len(msgs) > 0 {
		for i := 0; i < len(msgs); i++ {
			fmt.Printf("%s:%s\n", msgs[i].Username, msgs[i].Msg)
		}
	}
	wg.Done()
}

func sendMsg(msgs []string, conn *grpc.ClientConn) {
	if len(msgs) <= 0 {
		return
	}
	client := pb.NewChatRoomClient(conn)
	stream, err := client.SaySomething(context.Background())
	if err != nil {
		log.Fatalf("get stream error: %v", err)
	}
	wg := sync.WaitGroup{}
	go recvMessage(stream, &wg)
	wg.Add(1)
	for i := 0; i < len(msgs); i++ {
		stream.Send(&pb.ChatMessage{Username: "KDF5000", Msg: msgs[i]})
	}
	stream.CloseSend()
	wg.Wait()
}
func main() {
	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 随机产生1-3次的输入，然后使用stream发送给server，server端也是随机产生几个回复消息
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(userName string) {
		inputReader := bufio.NewReader(os.Stdin)
		for {
			done := false
			msgs := make([]string, 0)
			rand.Seed(time.Now().UnixNano())
			count := rand.Intn(3) + 1
			for i := 1; i <= count; i++ {
				fmt.Printf("%s> ", userName)
				in, err := inputReader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
					break
				}
				if in == "\n" {
					done = true
					break
				}
				msgs = append(msgs, strings.TrimSpace(in))
			}
			if len(msgs) > 0 {
				sendMsg(msgs, conn)
			}
			if done {
				break
			}
		}
		fmt.Println("GodBye!")
		wg.Done()
	}("KDF5000")
	wg.Wait()
}
