package process

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"encoding/json"
	"fmt"
	"net"
)

func UserLoginProcess(conn net.Conn, msg *message.Msg) (err error) {
	var data message.DataLogin
	err = json.Unmarshal([]byte(msg.Data), &data)
	if err != nil {
		fmt.Printf("格式错误 msg.Data:%s\nerror:%s", msg.Data, err)
		return
	}

	var response message.Response

	// 判断输入的用户名和密码是否匹配
	user, err := model.UserDao.Login(data.UserId, data.UserPwd)
	if err != nil { // 登陆失败
		fmt.Println("model.UserDao.Login(data.UserId, data.UserPwd) error:", err)
		response.Code = message.CodeLoginFailed
		response.Error = err.Error()
	} else { // 登陆成功
		response.Code = message.CodeLoginSuccess
		response.Error = ""
	}
	fmt.Println(user)

	// 返回 response 到客户端
	err = message.WriteResponse(conn, response)
	if err != nil {
		fmt.Println(" message.WriteMsg(conn, response) error:", err)
		return
	}
	return
}

// 注册处理逻辑
func UserRegisterProcess(conn net.Conn, msg *message.Msg) (err error) {
	// 反序列化msg.Data为User
	var user message.User
	err = json.Unmarshal([]byte(msg.Data), &user)
	if err != nil {
		fmt.Printf("格式错误 msg.Data:%s\nerror:%s", msg.Data, err)
		return
	}

	var response message.Response

	err = model.UserDao.Register(&user)
	if err != nil { // 注册失败
		fmt.Println("model.UserDao.Login(data.UserId, data.UserPwd) error:", err)
		response.Code = message.CodeRegisterFailed
		response.Error = err.Error()
	} else { // 注册成功
		response.Code = message.CodeRegisterSuccess
		response.Error = ""
	}
	// 返回 response 到客户端
	err = message.WriteResponse(conn, response)
	if err != nil {
		fmt.Println(" message.WriteMsg(conn, response) error:", err)
		return
	}
	return
}
