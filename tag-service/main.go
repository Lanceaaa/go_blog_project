package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"path"
	"strings"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/go-programming-tour-book/tag-service/internal/middleware"
	"github.com/go-programming-tour-book/tag-service/pkg/swagger"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"github.com/go-programming-tour-book/tag-service/server"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	// clientv3 "go.etcd.io/etcd/client/v3"
	// "go.etcd.io/etcd/proxy/grpcproxy"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/proxy/grpcproxy"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8001", "启动端口号")
	flag.Parse()
}

const SERVICE_NAME = "tag-service"

func main() {
	err := RunServer(port)
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}
}

// 拆分 TCP 的逻辑
func RunTCPServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func RunServer(port string) error {
	httpMux := RunHttpServer()
	grpcS := RunGrpcServer()

	gatewayMux := runGrpcGatewayServer()

	httpMux.Handle("/", gatewayMux)

	// 创建 etcd 实例
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: time.Second * 60,
	})
	if err != nil {
		return err
	}
	defer etcdClient.Close()

	target := fmt.Sprintf("/etcdv3://go-programming-tour-book/grpc/%s", SERVICE_NAME)
	// 服务信息注册
	grpcproxy.Register(etcdClient, target, ":"+port, 60)

	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

// 针对 HTTP 的 RunHttpServer 方法，其作用是初始化一个新的 HTTP 多路复用器，并新增了一个 /ping 路由及其 Handler，可用于做基本的心跳检测
func RunHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	prefix := "/swagger-ui/"
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	serveMux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	// 允许访问本地 proto 目录下的 .swagger.json 文件服务
	serveMux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			http.NotFound(w, r)
			return
		}

		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join("proto", p)

		http.ServeFile(w, r, p)
	})
	return serveMux
}

// 保持实现了 gRPC Server 的相关逻辑，仅是重新封装为 RunGrpcServer 方法
func RunGrpcServer() *grpc.Server {
	// 一元拦截器
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middleware.AccessLog, // 访问日志拦截器
			middleware.ErrorLog,  // 错误日志拦截器
			middleware.Recovery,  // 异常补抓拦截器
			// 链路追踪这个出先出：runtime error: invalid memory address or nil pointer dereference, stack: goroutine 8 [running]这个问题
			// middleware.ServerTracing, // 链路追踪拦截器
			// HelloInterceptor,
			// WorldInterceptor,
		)),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func HelloInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("你好开始")
	resp, err := handler(ctx, req)
	log.Println("你好结束")
	return resp, err
}

func WorldInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("世界开始")
	resp, err := handler(ctx, req)
	log.Println("世界结束")
	return resp, err
}

func runGrpcGatewayServer() *runtime.ServeMux {
	endpoint := "0.0.0.0:" + port
	// 注册错误方法
	runtime.HTTPError = grpcGatewayError
	gwmux := runtime.NewServeMux()
	// 指定了 Server 为非加密模式
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	// TagServiceHandler 事件，其内部会自动转换并拨号到 gRPC Endpoint，并在上下文结束后关闭连接
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)

	return gwmux
}

// 不同协议的分流服务启动端口号
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	// 根据不同的请求流量类型将其劫持并重定向到相应的 Hander 中去处理，最终以此达到同个端口上既提供 HTTP/1.1 又提供 HTTP/2 的功能
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// gRPC 和 HTTP/1.1 的流量区分，对 protoMajor 进行判断，该字段代表客户端请求的版本号，客户端始终使用 HTTP/1.1 或 HTTP/2
		// Header 头 Content-Type的确定：gRPC 的标志位 application/grpc 的确定
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

type httpError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func grpcGatewayError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	httpError := httpError{Code: int32(s.Code()), Message: s.Message()}
	details := s.Details()
	for _, detail := range details {
		if v, ok := detail.(*pb.Error); ok {
			httpError.Code = v.Code
			httpError.Message = v.Message
		}
	}

	resp, _ := json.Marshal(httpError)
	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))
	_, _ = w.Write(resp)
}
