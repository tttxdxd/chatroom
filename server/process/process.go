package process

import (
	"chatroom/common/message"
	"fmt"
	"net"
)

// 服务器信息处理中心
func Process(conn net.Conn) {
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

// 服务器根据消息类型 决定不同的处理方式
func serverProcessMsg(conn net.Conn, msg *message.Msg) (err error) {
	if msg == nil {
		return
	}

	switch msg.Type {
	case message.TypeLogin: //处理登陆逻辑
		UserLoginProcess(conn, msg)
	case message.TypeRegister: //处理注册逻辑

	default:
		fmt.Println("不存在的消息类型")
	}
	return
}
