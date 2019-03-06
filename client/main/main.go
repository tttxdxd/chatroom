package main

import (
	"chatroom/client/process"
	"fmt"
)

var loop bool

func main() {

	loop = true

	for loop {
		fmt.Println("\t\t\t聊天室 客户端：")
		fmt.Println("\t\t\t1. 进入聊天室")
		fmt.Println("\t\t\t2. 创建用户")
		fmt.Println("\t\t\t3. 销毁用户")
		fmt.Println("\t\t\t4. 退出聊天室")
		fmt.Println("————————————————————————————————————")
		index := 0
		fmt.Print("输入选择（1，2，3，4）：")
		fmt.Scanf("%d\n", &index)
		switch index {
		case 1:
			fmt.Println("进入聊天室")
			var username uint32
			var password string
			fmt.Print("输入用户ID：")
			fmt.Scanf("%d\n", &username)
			fmt.Print("输入密码：")
			fmt.Scanf("%s\n", &password)
			process.Login(username, password)
			loop = false
		case 2:
		case 3:
		case 4:
			fmt.Println("退出成功")
			loop = false
		default:
			fmt.Println("输入错误")
		}
	}
}

func display() {

}

func adduser() {

}

func removeuser() {

}
