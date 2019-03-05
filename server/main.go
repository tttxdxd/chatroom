package main

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

// 服务器根据消息类型 决定不同的处理方式
func serverProcessMsg(conn net.Conn, msg *message.Msg) (err error) {
	switch msg.Type {
	case message.TypeLogin: //处理登陆逻辑
		var login message.DataLogin
		err = json.Unmarshal([]byte(msg.Data), &login)
		fmt.Println(login)
		var response message.Response
		response.Code = message.CodeLoginSuccess
		response.Error = ""
		err = message.WriteResponse(conn, response)
		if err != nil {
			fmt.Println(" message.WriteMsg(conn, response) error:", err)
			return
		}
	case message.TypeRegister: //处理注册逻辑

	default:
		fmt.Println("不存在的消息类型")
	}
	return
}

func process(conn net.Conn) {
	defer conn.Close()
	fmt.Println("服务端 开始处理 客户端 数据")
	for {
		var msg message.Msg
		msg, err := message.ReadMsg(conn)
		if err != nil {
			fmt.Println("readMessage(conn) err:", err)
			return
		}
		fmt.Println("msg:", msg)
		serverProcessMsg(conn, &msg)
	}
}

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

		go process(conn)
	}
}
