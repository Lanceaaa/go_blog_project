package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-programming-tour-book/chatroom/global"
	"github.com/go-programming-tour-book/chatroom/server"
)

var (
	addr   = ":2022"
	banner = `
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    | 
   |    |    | /----\   |
   |____|    |/      \  |

Go 语言编程之旅 —— 一起用 Go 做项目：ChatRoom，start on：%s
`
)

func main() {
	fmt.Printf(banner+"\n", addr)

	server.RegisterHandle()

	global.Init()

	log.Fatal(http.ListenAndServe(addr, nil))
}
