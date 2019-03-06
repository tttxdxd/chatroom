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
