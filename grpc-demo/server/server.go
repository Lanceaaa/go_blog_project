package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net"

	pb "github.com/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
)

var port string

func main() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()

	// 创建 gRPC Server 对象，Server端的抽象对象
	server := grpc.NewServer()
	// 将 GreeterServer（其包含需要被调用的服务端接口）注册到 gRPC Server。 的内部注册中心。
	// 这样可以在接受到请求时，通过内部的 “服务发现”，发现该服务端接口并转接进行逻辑处理。创建 Listen，监听 TCP 端口。
	pb.RegisterGreeterServer(server, &GreeterServer{})
	// gRPC Server 开始 lis.Accept，直到 Stop 或 GracefulStop
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}

type GreeterServer struct{}

// Unary RPC：一元RPC
func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return *pb.HelloReply{Message: "hello.world"}, nil
}

// Server-side streaming RPC：服务端流式 RPC
func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for n := 0; n <= 6; n++ {
		_ = stream.Send(&pb.HelloReply{Message: "hello.list"})
	}

	return nil
}

// Client-side streaming RPC：客户端流式 RPC
func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{Message: "say.record"})
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
	}
	return nil
}

// Bidirectional streaming RPC：双向流式 RPC
func (s *GreeterServer) SayRouter(stream pb.Greeter_SayRouterServer) error {
	n := 0
	for {
		_ = stream.Send(&pb.HelloReply{Message: "say.route"})

		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++
		log.Printf("resp: %v", resp)
	}
}
