package main

import (
	"context"
	"log"

	pb "github.com/go-programming-tour-book/tag-service/proto"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	// 创建给定目标的客户端连接，另外我们所要请求的服务端是非加密模式的，因此我们调用了 grpc.WithInsecure 方法禁用了此 ClientConn 的传输安全性验证
	clientConn, _ := GetClientConn(ctx, "localhost:8001", nil)
	defer clientConn.Close()

	// 初始化指定 RPC Proto Service 的客户端实例对象
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	// 发起指定 RPC 方法的调用
	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "hahahahh"})

	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
