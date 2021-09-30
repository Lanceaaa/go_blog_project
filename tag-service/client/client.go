package main

import (
	"context"
	"log"

	"github.com/go-programming-tour-book/tag-service/internal/middleware"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {
	ctx := context.Background()
	// 创建给定目标的客户端连接，另外我们所要请求的服务端是非加密模式的，因此我们调用了 grpc.WithInsecure 方法禁用了此 ClientConn 的传输安全性验证
	// 调用时需要将超时控制的拦截器注册进去
	clientConn, _ := GetClientConn(ctx, "localhost:8001", []grpc.DialOption{grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(middleware.UnaryContextTimeout()),
	)})
	defer clientConn.Close()

	// 初始化指定 RPC Proto Service 的客户端实例对象
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	// 发起指定 RPC 方法的调用
	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "hahahahh"})

	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	// 一元调用和流式调用添加对应的客户端拦截器
	opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			middleware.UnaryContextTimeout(),
		),
	))
	opts = append(opts, grpc.WithStreamInterceptor(
		grpc_middleware.ChainStreamClient(
			middleware.StreamContextTimeout(),
		),
	))
	// 添加重试操作
	opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithMax(2),
				grpc_retry.WithCodes(
					codes.Unknown,
					codes.Internal,
					codes.DeadlineExceeded,
				),
			),
		),
	))
	return grpc.DialContext(ctx, target, opts...)
}
