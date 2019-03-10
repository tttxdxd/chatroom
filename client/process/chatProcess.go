package process

import (
	"chatroom/common/message"
	"fmt"
)

// TODO 需对应同步退出 使用chan
func (this *Processor) sendMessage() {
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
			fmt.Println("---------获取在线用户列表----------")
			var msg message.Msg
			msg.Type = message.TypeGetOnlineUsers
			message.Center.AddMsg(&msg)
			err := message.WriteMsg(this.conn, &msg)
			if err != nil {
				fmt.Println("message.WriteMsg(conn, msg) error: ", err)
				break
			}

			block <- true

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

func (this *Processor) ReceiveMessage() {
	for {
		fmt.Println("start readMSg")
		msg, err := message.ReadMsg(this.conn)
		if err != nil {
			fmt.Println(" message.ReadMsg(conn) error:", err)
			return
		}
		fmt.Println("msg=", msg)

		err = message.Center.Distribute(msg)
		if err != nil {
			fmt.Println("Distribute(&msg) error:", err)
			return
		}
		//<-block
		fmt.Println("end readMSg")
	}
}
