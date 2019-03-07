package process

import (
	"chatroom/common/message"
	"fmt"
	"net"
)

// TODO 需对应同步退出 使用chan
func sendMessage(conn net.Conn) {
	for {
		fmt.Println("---------当前客户端上线----------")
		fmt.Println("---------1.在线用户列表----------")
		fmt.Println("---------2.发送消息----------")
		fmt.Println("---------3.----------")
		fmt.Println("---------4.退出聊天室----------")
		fmt.Println("------------------------------")
		fmt.Print("输入1，2，3，4选择：")
		var index int
		fmt.Scanf("%d\n", &index)
		switch index {
		case 1:
			fmt.Println("---------1.在线用户列表----------")
		case 2:
			fmt.Println("---------2.发送消息----------")
		case 3:
			fmt.Println("---------3.----------")
		case 4:
			fmt.Println("---------4.退出聊天室----------")
			return
		default:
		}
	}
}

func receiveMessage(conn net.Conn) {
	for {
		msg, err := message.ReadMsg(conn)
		if err != nil {
			fmt.Println(" message.ReadMsg(conn) error:", err)
			return
		}
		fmt.Println("msg=", msg)
		// 处理服务端发送的消息
		// switch msg.Type{
		// 	case
		// }
	}
}
