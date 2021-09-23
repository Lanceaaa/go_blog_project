package main

import (
	"context"
	"io"
	"log"

	pb "github.com/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
)

func main() {
	// 创建与给定目标（服务端）的连接句柄
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	// 创建 Greeter 的客户端对象
	client := pb.NewGreeterClient(conn)
	// 发送 RPC 请求，等待同步响应，得到回调后返回响应结果
	_ = SayHello(client)
}

// Unary RPC：一元 RPC
func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "lance"}))
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil 
}

// Server-side streaming RPC：服务端流式 RPC
func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
	}
	return nil
}

// Client-side streaming RPC：客户端流式 RPC
func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	for n := 0; n < 6; n++ {
	    _ = stream.Send(r)
	}
	resp, _ := stream.CloseAndRecv()
	log.Printf("resp err: %v", resp)

	return nil 
}

// Bidirectional streaming RPC：双向流式 RPC
func SayRouter(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRouter(context.Background())
	if n := 0; n <= 6; n++ {
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp err: %v", resp)
	}

	_ = stream.CloseSend()
}