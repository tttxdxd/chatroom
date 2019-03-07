package process

import (
	"chatroom/common/message"
	"fmt"
	"net"
)

// 服务器信息分发 处理中心
func Process(conn net.Conn) {
	defer fmt.Println("断开与客户端的连接")
	defer conn.Close()

	fmt.Println("服务端 开始处理 客户端 数据")
	for {
		var msg message.Msg
		msg, err := message.ReadMsg(conn)
		if err != nil {
			fmt.Println("readMessage(conn) err:", err)
			return
		}
		if msg.Type == message.TypeClientExit { //退出
			return
		}
		err = serverProcessMsg(conn, &msg)
		if err != nil {
			fmt.Println("readMessage(conn) err:", err)
			return
		}
	}
}

// 服务器根据消息类型 分发给不同的处理器 决定不同的处理方式
func serverProcessMsg(conn net.Conn, msg *message.Msg) (err error) {
	if msg == nil {
		return
	}

	switch msg.Type {
	case message.TypeLogin: //处理登陆逻辑
		err = UserLoginProcess(conn, msg)
	case message.TypeRegister: //处理注册逻辑
		err = UserRegisterProcess(conn, msg)
	default:
		fmt.Println("不存在的消息类型")
	}
	return
}
