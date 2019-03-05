package main

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func login(username uint32, password string) (err error) {
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
	fmt.Println(res)

	time.Sleep(3 * time.Second)

	return nil
}
