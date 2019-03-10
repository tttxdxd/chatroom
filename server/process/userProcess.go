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

	var resType message.MsgType
	var response message.Response

	// 判断输入的用户id和密码是否匹配
	user, err := model.UserDao.Login(data.UserId, data.UserPwd)
	if err != nil { // 登陆失败
		fmt.Println("model.UserDao.Login(data.UserId, data.UserPwd) error:", err)
		resType = message.CodeLoginFailed
		response.Error = err.Error()
	} else { // 登陆成功
		resType = message.CodeLoginSuccess
		response.Error = ""

		//添加该用户到在线用户列表（由UserManager维护）
		this.userId = user.UserId
		this.username = user.Username
		UserManager.AddUser(this)
	}
	//fmt.Println(user)

	res, ok := message.NewMsg(msg.ID, resType, response)
	if !ok {
		fmt.Println("message.NewMsg(msg.ID, resType, response) error")
		return
	}

	// 返回 response 到客户端
	err = message.WriteMsg(this.conn, res)
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

	var resType message.MsgType
	var response message.Response

	err = model.UserDao.Register(&user)
	if err != nil { // 注册失败
		fmt.Println("model.UserDao.Login(data.UserId, data.UserPwd) error:", err)
		resType = message.CodeRegisterFailed
		response.Error = err.Error()
	} else { // 注册成功
		resType = message.CodeRegisterSuccess
		response.Error = ""
	}

	res, ok := message.NewMsg(msg.ID, resType, response)
	if !ok {
		fmt.Println("message.NewMsg(msg.ID, resType, response) error")
		return
	}

	// 返回 response 到客户端
	err = message.WriteMsg(this.conn, res)
	if err != nil {
		fmt.Println(" message.WriteMsg(conn, response) error:", err)
		return
	}
	return
}

// 获取所有用户信息的处理逻辑
func (this *UserProcess) UserGetAllOnlineUsersProcess(msg *message.Msg) (err error) {
	var resType message.MsgType
	var response message.Response

	resType = message.PlaceHolder
	response.Infos = UserManager.GetAllUsersInfo()
	fmt.Println(UserManager.onlineUserList)
	fmt.Println(response)

	res, ok := message.NewMsg(msg.ID, resType, response)
	if !ok {
		fmt.Println("message.NewMsg(msg.ID, resType, response) error")
		return
	}

	fmt.Println("res", res)
	err = message.WriteMsg(this.conn, res)
	if err != nil {
		fmt.Println(" message.WriteMsg(conn, response) error:", err)
		return
	}
	return
}

func (this *UserProcess) UserNotifyAllUsersProcess(msg *message.Msg) {

}
