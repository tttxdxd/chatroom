package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

var Instance *Processor

type Processor struct {
	conn net.Conn
}

func InitProcess(conn net.Conn) {
	Instance = &Processor{
		conn: conn,
	}

}

func (this *Processor) Login(user_id uint32, user_pwd string) (err error) {
	// var user_id uint32
	// var user_pwd string
	// fmt.Print("输入用户ID：")
	// fmt.Scanf("%d\n", &user_id)
	// fmt.Print("输入密码：")
	// fmt.Scanf("%s\n", &user_pwd)

	//---------------------------------------------------------

	//构建结构体
	msg := message.Msg{
		Type: message.TypeLogin,
		Data: "",
	}
	data := message.DataLogin{
		UserId:  user_id,
		UserPwd: user_pwd,
	}

	byteData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json.Marshal(data) err:", err)
		return
	}
	msg.Data = string(byteData)

	// 由外界维持TCP连接
	// conn, err := net.Dial("tcp", "localhost:8889")
	// if err != nil {
	// 	fmt.Println("net.Dial err", err)
	// 	return
	// }
	// defer conn.Close()

	message.Center.AddMsg(&msg)

	//发送用户登陆数据
	err = message.WriteMsg(this.conn, &msg)
	if err != nil {
		fmt.Println("message.WriteMsg(conn, msg) error", err)
		return
	}

	// //返回服务端数据
	// res, err := message.ReadResponse(this.conn)
	// if err != nil {
	// 	fmt.Println(" message.ReadMsg(conn) error:", err)
	// 	return
	// }

	return
}

// 注册处理逻辑
func (this *Processor) Register(username string, password string) {

	fmt.Println("开始创建用户（注册）")
	var user message.User
	// fmt.Print("输入用户ID：")
	// fmt.Scanf("%d\n", &user.UserId)
	// fmt.Print("输入用户名：")
	// fmt.Scanf("%s\n", &user.Username)
	// fmt.Print("输入密码：")
	// fmt.Scanf("%s\n", &user.UserPwd)

	user.Username = username
	user.UserPwd = password

	//-------------------------------------------------------------------------

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal(user) error: ", err)
		return
	}

	msg := message.Msg{
		Type: message.TypeRegister,
		Data: string(data),
	}
	fmt.Println(msg)

	message.Center.AddMsg(&msg)

	// 发送将注册的用户信息到服务端
	err = message.WriteMsg(this.conn, &msg)
	if err != nil {
		fmt.Println("message.WriteMsg(conn, msg) error", err)
		return
	}

	// //返回服务端回应数据
	// res, err := message.ReadResponse(this.conn)
	// if err != nil {
	// 	fmt.Println(" message.ReadMsg(conn) error:", err)
	// 	return
	// }

	return
}

func (this *Processor) GetOnlineUsers() (err error) {
	msg := message.Msg{
		Type: message.TypeGetOnlineUsers,
	}
	message.Center.AddMsg(&msg)
	err = message.WriteMsg(this.conn, &msg)
	if err != nil {
		fmt.Println("message.WriteMsg(conn, msg) error: ", err)
		return
	}
	return
}

func (this *Processor) Exit() (err error) {
	defer fmt.Println("断开与服务端的连接")
	msg := message.Msg{
		Type: message.TypeClientExit,
		Data: "",
	}
	err = message.WriteMsg(this.conn, &msg)
	if err != nil {
		fmt.Println("message.WriteMsg(conn, msg) error: ", err)
		return
	}
	return
}

// 通知所有在线用户 （除了自己）
func (this *Processor) NotifyOnlineUsers() (err error) {
	msg := message.Msg{
		Type: message.TypeNotifyOnlineUsers,
	}
	err = message.WriteMsg(this.conn, &msg)
	if err != nil {
		fmt.Println("message.WriteMsg(conn, msg) error: ", err)
		return
	}
	return
}

// 发送消息（暂定群发）
func (this *Processor) SendMessage(text string) (err error) {
	msg := message.Msg{
		Type: message.TypeSendMessage,
		Data: text,
	}
	err = message.WriteMsg(this.conn, &msg)
	if err != nil {
		fmt.Println("message.WriteMsg(conn, msg) error: ", err)
		return
	}
	return
}

// 判断是否继续操作
func JudgeWhetherToContinue(operating string) bool {
	fmt.Println("是否继续" + operating + "（Y/N):")
	var ok string
	fmt.Scanf("%s\n", &ok)
	switch strings.ToLower(ok) {
	case "y":
		fallthrough
	case "yes":
		return true
	case "n":
		fallthrough
	case "no":
		fallthrough
	default:
		return false
	}
	return false
}
