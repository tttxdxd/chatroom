package process

import (
	"chatroom/common/message"
	"fmt"
	"net"
)

var ERROR_EXIT = fmt.Errorf("客户端发送断开请求")

// 服务器信息分发 处理中心
func Process(conn net.Conn) {
	defer fmt.Println("断开与客户端的连接")
	defer conn.Close()

	fmt.Println("服务端 开始处理 客户端 数据")
	userProcess := UserProcess{
		conn:   conn,
		userId: 0,
	}
	for {
		msg, err := message.ReadMsg(conn)
		if err != nil {
			fmt.Println("readMessage(conn) err:", err)
			return
		}

		err = serverProcessMsg(&userProcess, msg)
		if err == ERROR_EXIT {
			fmt.Println(err.Error())
			return
		} else if err != nil {
			fmt.Println("readMessage(conn) err:", err)
			return
		}
	}
}

// 服务器根据消息类型 分发给不同的处理器 决定不同的处理方式
func serverProcessMsg(userProcess *UserProcess, msg *message.Msg) (err error) {
	if msg == nil {
		return
	}

	switch msg.Type {
	case message.TypeLogin: //处理登陆逻辑
		err = userProcess.UserLoginProcess(msg)
	case message.TypeRegister: //处理注册逻辑
		err = userProcess.UserRegisterProcess(msg)
	case message.TypeClientExit: //处理退出逻辑
		msg.Type = message.CodeNotifyUserLogout
		userProcess.UserNotifyAllUsersProcess(msg)
		userProcess.UserLogoutProcess()
		err = ERROR_EXIT
	case message.TypeGetOnlineUsers: //处理获取当前在线用户列表逻辑
		err = userProcess.UserGetAllOnlineUsersProcess(msg)
	case message.TypeNotifyOnlineUsers: //处理 通知在线用户信息 逻辑
		msg.Type = message.CodeNotifyUserOnline
		userProcess.UserNotifyAllUsersProcess(msg)
	case message.TypeSendMessage: //处理 发送群发消息 逻辑
		msg.Type = message.CodeRecvMessage
		userProcess.UserNotifyAllUsersProcess(msg)
	default:
		fmt.Println("不存在的消息类型")
	}
	return
}
