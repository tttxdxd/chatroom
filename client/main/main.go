package main

import (
	"chatroom/client/process"
	"chatroom/client/view"
	"fmt"
	"net"
	"time"
)

var Conn net.Conn

func tcpReconnect() {
	time.Sleep(2*time.Second)
	count := 0
	for {
		count++
		Conn, err := net.Dial("tcp", view.Address)
		if err != nil {
			fmt.Printf("第%d次连接失败 error:%v\n", count, err)
			view.Connected = "未连接"
			view.StateChange()
			time.Sleep(3 * time.Second)
			continue
		}
		view.Connected = "已连接"
		fmt.Printf("连接 %s 成功\n",view.Address)
		view.StateChange()
		process.InitProcess(Conn)

		process.Instance.ReceiveMessage()
	}
}

func tryConnectServer(address string) (err error) {
	count := 0    //尝试连接的次数
	maxCount := 3 //最大尝试次数
	for count < maxCount {
		count++
		fmt.Println("客户端尝试发起连接...")
		Conn, err = net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("第%d次连接失败 error:%v", count, err)
			continue
		}
		//Conn = conn
		fmt.Println("客户端-服务端连接成功...")
		return nil
	}
	return
}

func main() {

	//err := tryConnectServer("localhost:8889")
	//if err != nil {
	//	fmt.Println("客户端发起连接失败...")
	//	return
	//}

	//process.InitProcess(Conn)
	//
	//go process.Instance.ReceiveMessage()

	go tcpReconnect()

	view.WindowShow()

	// for {
	// 	fmt.Println("\t\t\t聊天室 客户端：")
	// 	fmt.Println("\t\t\t1. 进入聊天室")
	// 	fmt.Println("\t\t\t2. 创建用户")
	// 	fmt.Println("\t\t\t3. 销毁用户")
	// 	fmt.Println("\t\t\t4. 退出聊天室客户端")
	// 	fmt.Println("————————————————————————————————————")
	// 	index := 0
	// 	fmt.Print("输入选择（1，2，3，4）：")
	// 	fmt.Scanf("%d\n", &index)
	// 	switch index {
	// 	case 1:
	// 		fmt.Println("进入聊天室")

	// 		process.Instance.Login()
	// 		fmt.Println("退出聊天室")
	// 	case 2:
	// 		process.Instance.Register()
	// 	case 3:
	// 	case 4:
	// 		process.Instance.Exit()
	// 		fmt.Println("退出成功")
	// 		return
	// 	default:
	// 		fmt.Println("输入错误")
	// 	}
	// }
}
