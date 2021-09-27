package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	pb "github.com/go-programming-tour-book/tag-service/proto"
	"github.com/go-programming-tour-book/tag-service/server"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8001", "启动端口号")
	flag.Parse()
}

// 拆分 TCP 的逻辑
func RunTCPServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

// 针对 HTTP 的 RunHttpServer 方法，其作用是初始化一个新的 HTTP 多路复用器，并新增了一个 /ping 路由及其 Handler，可用于做基本的心跳检测
func RunHttpServer(port string) *http.Server {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
}

// 保持实现了 gRPC Server 的相关逻辑，仅是重新封装为 RunGrpcServer 方法
func RunGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func main() {
	// 初始化 TCP Listener
	l, err := RunTCPServer(port)
	if err != nil {
		log.Fatalf("Run TCP Server err: %v", err)
	}
	m := cmux.New(l)
	// cmux 也是基于 application/grpc 标识去进行分流
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := RunGrpcServer()
	httpS := RunHttpServer(port)
	// 使用 cmux 监听 gRPC 和 HTTP 服务
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	// 开启服务
	err = m.Serve()
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}
}
