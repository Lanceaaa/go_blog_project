package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	pb "github.com/go-programming-tour-book/tag-service/proto"
	"github.com/go-programming-tour-book/tag-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// var port string
var grpcPort string
var httpPort string

func init() {
	flag.StringVar(&grpcPort, "grpc_port", "8001", "gRPC 启动端口号")
	flag.StringVar(&httpPort, "http_port", "9001", "HTTP 启动端口号")
	flag.Parse()
}

// 针对 HTTP 的 RunHttpServer 方法，其作用是初始化一个新的 HTTP 多路复用器，并新增了一个 /ping 路由及其 Handler，可用于做基本的心跳检测
func RunHttpServer(port string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return http.ListenAndServe(":"+port, serveMux)
}

// 保持实现了 gRPC Server 的相关逻辑，仅是重新封装为 RunGrpcServer 方法
func RunGrpcServer(port string) error {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	return s.Serve(lis)
}

func main() {
	// 监听 HTTP EndPoint 和 gRPC EndPoint 是一个阻塞的行为
	errs := make(chan error)
	// RunHttpServer 或 RunGrpcServer 方法启动或运行出现了问题，会将 err 写入 chan 中
	go func() {
		err := RunHttpServer(httpPort)
		if err != nil {
			errs <- err
		}
	}()

	go func() {
		err := RunGrpcServer(grpcPort)
		if err != nil {
			errs <- err
		}
	}()

	// 因此我们只需要利用 select 对其进行检测即可
	select {
	case err := <-errs:
		log.Fatalf("Run Server err: %v", err)
	}
}
