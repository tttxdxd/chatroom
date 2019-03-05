package main

import (
	"chatroom/server/process"
	"fmt"
	"net"
)

func main() {
	//开启监听8889
	fmt.Println("服务端 开始监听 8889 端口")
	listen, err := net.Listen("tcp", "localhost:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("listen err:", err)
	}
	for {
		fmt.Println("服务端 开始等待 客户端 连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
		}
		go process.Process(conn)
	}
}
