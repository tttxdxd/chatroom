package main

import (
	"chatroom/server/model"
	"chatroom/server/process"
	"fmt"
	"net"
)

func main() {

	//开启redis连接池
	err := model.InitPool(":6379", "")
	if err != nil {
		return
	}
	// 初始化DAO
	model.InitDAO()

	// model.UserDao.Register(&message.User{
	// 	UserId:   100,
	// 	UserPwd:  "123456",
	// 	Username: "Alex",
	// })

	//开启监听8889
	fmt.Println("服务端 开始监听 8889 端口")
	listen, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("服务端 开始等待 客户端 连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
		}
		go process.Process(conn)
	}
}
