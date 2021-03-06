## This is go blog project
used gin frame

### blog-service目录结构如下：
#### ├──configs           配置文件
#### ├──docs              文档集合
#### ├──global            全局变量
#### ├──internal          内部模块
#### │   └──dao             数据访问层，所有与数据相关的操作都会在dao层进行，例如MySQL，Elasticsearch等
#### │   └──middleware      http中间件
#### │   └──model           模型层，存放model对象
#### │   └──roters          路由相关的逻辑
#### │   └──service         项目核心业务逻辑
#### ├──pkg               项目相关的模块包
#### ├──storage           项目生成的临时文件
#### ├──scripts           各类构建、安装、分析等操作的脚本
#### ├──third_party       第三方的资源工具、如Swagger UI

```bash
go get -u github.com/gin-gonic/gin@v1.6.3
go get -u github.com/spf13/viper@v1.4.0
go get -u github.com/jinzhu/gorm@v1.9.12
go get -u gopkg.in/natefinch/lumberjack.v2
# swagger
go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
go get -u github.com/swaggo/gin-swagger@v1.2.0
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template
# validator接口校验
go get -u github.com/go-playground/validator/v10
# jwt
go get -u github.com/dgrijalva/jwt-go@v3.2.0
# 发送电子邮件
go get -u gopkg.in/gomail.v2
# 限流器（令牌桶）
go get -u github.com/juju/ratelimit@v1.0.1
```

# 链路追踪
```bash
docker run -d --name jaeger \
-e COLLECTOR_ZIPLIN_HTTP_PORT=9411 \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832:udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14268:14268 \
-p 9411:9411 \
jaegertracing/all-in-one:1.16 --reporter.grpc.host-port=127.0.0.1:8001
```
```bash
# OpenTracing API
go get -u github.com/opentracing/opentracing-go@v1.1.0
# Jaeger Client
go get -u github.com/uber/jaeger-client-go@v2.22.1
# sql 追踪
go get -u github.com/eddycjy/opentracing-gorm
```

# 打包进二进制文件
```bash
# go-bindata ...要使用最新版本
go get -u github.com/go-bindata/go-bindata/...
# 命令
go-bindata -o configs/config.go -pkg=configs configs/config.yaml
# 读取对应文件内容
b, _ := configs.Asset("configs/config.yaml")
```
把第三方文件打包进二进制文件后，二进制文件必然增大，而且常规方法下无法做文件的热更新和监听，必须重启并且重新打包才能使用最新的内容。

# 配置热更新
```bash
# 安装开源库fsnotify
go get -u golang.org/x/sys
go get -u github.com/fsnotify/fsnotify
```

# Protobuf
```bash
# protoc安装 Protobuf的编译器 编译.proto文件
wget https://github.com/google/protobuf/releases/download/v3.11.2/protobuf-all-3.11.2.zip
unzip protobuf-all-3.11.2.zip && cd protobuf-3.11.2/
./configure
make
make install
# 检查是否安装成功
protoc --version
```
```bash
# protoc插件安装 方式1
go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2
# 方式2
GIT_TAG="v1.3.2"
go get -d -u github.com/golang/protobuf/protoc-gen-go
git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
go install github.com/golang/protobuf/protoc-gen-go
# 将所编译安装的 Protoc Plugin 的可执行文件中移动到相应的 bin 目录
mv $GOPATH/bin/protoc-gen-go /usr/local/go/bin/
```

# 初始化Demo项目
```bash
mkdir grpc-demo
cd grpc-demo
go mod init github.com/go-programming-tour-book/grpc-demo
```

### grpc-demo目录结构如下：
#### ├──client           配置文件
#### ├──proto            文档集合
#### ├──server           全局变量
###  └──go.mod           

# 编译和生成 proto 文件
```bash
protoc --go_out=plugins=grpc:. ./proto/*.proto
```
# 如果出现如下类似报错
```
protoc: error while loading shared libraries: libprotobuf.so.15: cannot open shared object file: No such file or directory
# 运行这个即可
ldconfig
```
- –go_out：设置所生成 Go 代码输出的目录，该指令会加载 protoc-gen-go 插件达到生成 Go 代码的目的，生成的文件以 .pb.go 为文件后缀，在这里 “:”（冒号）号充当分隔符的作用，后跟命令所需要的参数集，在这里代表着要将所生成的 Go 代码输出到所指向 protoc 编译的当前目录
- plugins=plugin1+plugin2：指定要加载的子插件列表，我们定义的 proto 文件是涉及了 RPC 服务的，而默认是不会生成 RPC 代码的，因此需要在 go_out 中给出 plugins 参数传递给 protoc-gen-go，告诉编译器，请支持 RPC（这里指定了内置的 grpc 插件）。

# 安装 gRPC 库
```bash
go get -u google.golang.org/grpc@v1.29.1
```

# 初始化 gRPC 服务项目
```bash
go mod init github.com/go-programming-tour-book/tag-service
```

# 调试 gRPC 接口
```bash
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
# 使用
grpcurl -plaintext localhost:8001 proto.TagService.GetTagList
```

# 使用第三方开源库 cmux 实现多协议支持的功能
```bash
go get -u github.com/soheilhy/cmux@v0.1.4
```

# 使用 grpc-gateway 将 RESTful 转换为 gRPC 请求，实现同一个 RPC 方法提供 gRPC 协议和 HTTP/1.1 的双流量支持
```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.14.5
mv $GOPATH/bin/protoc-gen-grpc-gateway /usr/local/go/bin/
```

# 执行命令在 proto 目录下生成 .pb.go 和 .pb.gw.go 两种文件，对应两类功能支持
```bash
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway\@v1.14.5/third_party/googleapis --grpc-gateway_out=logtostderr=true:. ./proto/*.proto
```

## 使用 protoc 的插件 protoc-gen-swagger 来生成 swagger 定义
```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

## 下载 Swagger UI 文件
> 到 https://github.com/swagger-api/swagger-ui 将其源码压缩包下载下来，然后将 dist 目录下的资源文件拷贝到项目 swagger-ui 目录中去

## 使用 go-bindata 讲 Swagger UI 资源文件转换为 Go 代码
```bash
go get -u github.com/go-bindata/go-bindata/...
```

## 在项目 pkg 项目新建 swagger 目录，在项目根目录执行命令
```bash
go-bindata --nocompress -pkg swagger -o pkg/swagger/data.go third_party/swagger-ui/...
```

## 使用 go-bindata-assetfs 能结合 net/http 标准库和 go-bindata 所生成的 Swagger UI 的 Go 代码两者来供外部访问
```bash
go get -u github.com/elazarl/go-bindata-assetfs/...
```

## 执行命令生成 swagger.json 文件
```bash
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway\@v1.14.5/third_party/googleapis --swagger_out=logtostderr=true:. ./proto/*.proto
```

# 拦截器的使用
## 使用 go-grpc-middleware 实现多拦截器
```bash
go get -u github.com/grpc-ecosystem/go-grpc-middleware@v1.1.0
```

# 使用 Jaeger 进行链路追踪
```bash
go get -u github.com/opentracing/opentracing-go@v1.1.0
go get -u github.com/uber/jaeger-client-go@v2.22.1
```

# 下载、安装和启动 etcd
```bash
wget https://github.com/etcd-io/etcd/releases/download/v3.4.17/etcd-v3.4.17-linux-amd64.tar.gz
tar -zxf etcd-v3.4.17-linux-amd64.tar.gz
cd etcd-v3.4.17-linux-amd64
mv etcd /usr/local/bin
ETCDTL_API=3 && etcd
```

# 安装 etcd client sdk
```bash
# go get go.etcd.io/etcd/client/v3
go get google.golang.org/grpc@1.26.0
go get github.com/coreos/etcd/clientv3@v3.3.18
# 如果拉取过程中出现/go-system模块的相关报错，则在go.mod添加replace来解决这个问题
replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22v22.0.0
```

# 初始化 protoc 的项目
```bash
go mod init github.com/go-programming-tour-book/protoc-gen-go-tour
```

# 编译生成 pb.go 文件
```bash
go build .
mv ./protoc-gen-go-tour /go/bin
# 拷贝 tag.proto 命名为tour.proto
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway\@v1.14.5/third_party/googleapis --go-tour_out=plugins=tour:. ./proto/tour.proto
```