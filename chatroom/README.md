### chatroom目录结构如下：
#### ├── README.md        readme 文件
#### ├── cmd              该目录几乎是 Go 圈约定俗成的，Go 官方以及开源界推荐的方式，用于存放 main.main；
#### │   ├── chatroom
#### │       └── main.go
#### ├── go.mod           
#### ├── go.sum           
#### ├── logic            用于存放项目核心业务逻辑代码，和 service 目录是类似的作用；
#### │   ├── broadcast.go 
#### │   ├── message.go   
#### │   └── user.go      
#### ├── server           存放 server 相关代码，虽然这是 WebSocket 项目，但也可以看成是 Web 项目，因此可以理解成存放类似 controller 的代码；
#### │   ├── handle.go    
#### │   ├── home.go      
#### │   └── websocket.go
#### └── template         存放静态模板文件；
####     └── home.html

# 安装聊天室项目
```bash
go get -v github.com/go-programming-tour-book/chatroom/cmd/chatroom
```