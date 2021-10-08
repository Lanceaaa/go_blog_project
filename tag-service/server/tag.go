package server

import (
	"context"
	"encoding/json"

	"github.com/go-programming-tour-book/tag-service/pkg/bapi"
	"github.com/go-programming-tour-book/tag-service/pkg/errcode"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"google.golang.org/grpc/metadata"
)

type TagServer struct {
	auth *Auth
}

type Auth struct{}

func (a *Auth) GetAppKey() string {
	return "go-programming-tour-book"
}

func (a *Auth) GetAppSecret() string {
	return "lance"
}

func (a *Auth) Check(ctx context.Context) error {
	md, _ := metadata.FromIncomingContext(ctx)

	var appKey, appSecret string
	if value, ok := md["app_key"]; ok {
		appKey = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}
	if appKey != a.GetAppKey() || appSecret != a.GetAppSecret() {
		return errcode.TogRPCError(errcode.Unauthorized)
	}

	return nil
}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	// panic("测试抛出异常！")
	// 进行 Token 校验
	if err := t.auth.Check(ctx); err != nil {
		return nil, err
	}
	// 指定博客后端的服务地址
	api := bapi.NewAPI("http://127.0.0.1:8080")
	// 调用 GetTagList 方法的 API 获取标签列表数据
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, err
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}

	return &tagList, nil
}
