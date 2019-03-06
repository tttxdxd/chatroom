package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

func Login(username uint32, password string) (err error) {
	//构建结构体
	msg := message.Msg{
		Type: message.TypeLogin,
		Data: "",
	}
	data := message.DataLogin{
		UserId:  username,
		UserPwd: password,
	}

	byteData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json.Marshal(data) err:", err)
		return
	}
	msg.Data = string(byteData)

	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err", err)
		return
	}
	defer conn.Close()

	//发送用户登陆数据
	err = message.WriteMsg(conn, msg)
	if err != nil {
		fmt.Println("message.WriteMsg(conn, msg) error", err)
		return
	}

	//返回服务端数据
	res, err := message.ReadResponse(conn)
	if err != nil {
		fmt.Println(" message.ReadMsg(conn) error:", err)
		return
	}

	switch res.Code {
	case message.CodeLoginSuccess:
		fmt.Println("登陆成功")

		sendMessage(conn)
		go receiveMessage(conn)
	case message.CodeLoginFailed:
		fmt.Println(res.Error)
	default:
	}

	return nil
}

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
