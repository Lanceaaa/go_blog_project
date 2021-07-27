# This is go blog project
## used gin frame

### blog-service目录结构如下：
#### --configs           配置文件
#### --docs              文档集合
#### --global            全局变量
#### --internal          内部模块
##### ----dao             数据访问层，所有与数据相关的操作都会在dao层进行，例如MySQL，Elasticsearch等
##### ----middleware      http中间件
##### ----model           模型层，存放model对象
##### ----roters          路由相关的逻辑
##### ----service         项目核心业务逻辑
#### --pkg               项目相关的模块包
#### --storage           项目生成的临时文件
#### --scripts           各类构建、安装、分析等操作的脚本
#### --third_party       第三方的资源工具、如Swagger UI

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
