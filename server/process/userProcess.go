package process

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	conn     net.Conn //每个连接进来的用户都有自己的Conn
	userId   uint32
	username string
}

func (this *UserProcess) UserLoginProcess(msg *message.Msg) (err error) {
	var data message.DataLogin
	err = json.Unmarshal([]byte(msg.Data), &data)
	if err != nil {
		fmt.Printf("格式错误 msg.Data:%s\nerror:%s", msg.Data, err)
		return
	}

	var response message.Response

	// 判断输入的用户id和密码是否匹配
	user, err := model.UserDao.Login(data.UserId, data.UserPwd)
	if err != nil { // 登陆失败
		fmt.Println("model.UserDao.Login(data.UserId, data.UserPwd) error:", err)
		response.Code = message.CodeLoginFailed
		response.Error = err.Error()
	} else { // 登陆成功
		response.Code = message.CodeLoginSuccess
		response.Error = ""

		//添加该用户到在线用户列表（由UserManager维护）
		this.userId = user.UserId
		this.username = user.Username
		UserManager.AddUser(this)
	}
	//fmt.Println(user)

	// 返回 response 到客户端
	err = message.WriteResponse(this.conn, response)
	if err != nil {
		fmt.Println(" message.WriteMsg(conn, response) error:", err)
		return
	}
	return
}

func (this *UserProcess) UserLogoutProcess() {
	UserManager.RemoveUser(this.userId)
}

// 注册处理逻辑
func (this *UserProcess) UserRegisterProcess(msg *message.Msg) (err error) {
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
	err = message.WriteResponse(this.conn, response)
	if err != nil {
		fmt.Println(" message.WriteMsg(conn, response) error:", err)
		return
	}
	return
}

func (this *UserProcess) UserGetAllOnlineUsers(msg *message.Msg) (err error) {
	var response message.Response
	response.Code = 0
	response.Infos = UserManager.GetAllUsersInfo(this.userId)
	fmt.Println(UserManager.onlineUserList)
	fmt.Println(response)
	err = message.WriteResponse(this.conn, response)
	if err != nil {
		fmt.Println(" message.WriteMsg(conn, response) error:", err)
		return
	}
	return
}
