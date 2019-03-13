package message

const (
	PlaceHolder = iota //

	TypeLogin        //
	CodeLoginSuccess // 登陆成功
	CodeLoginFailed  // 登陆失败 1.密码错误 2.用户名不存在 3.服务器错误

	TypeRegister        //
	CodeRegisterSuccess // 注册成功
	CodeRegisterFailed  // 注册失败 1.用户名已存在

	TypeGetOnlineUsers //获取在线用户列表（聊天室内）
	CodeOnlineUsers

	TypeClientExit       //
	CodeNotifyUserLogout //被通知用户下线

	TypeNotifyOnlineUsers //通知在线用户消息
	CodeNotifyUserOnline  //被通知用户上线

	TypeSendMessage //发送聊天信息(群发)
	CodeRecvMessage //发送成功 服务器收到消息

)

type MsgType uint
type MsgID uint

type Msg struct {
	Type MsgType `json:"type"`
	ID   MsgID   `json:"id"`
	Data string  `json:"data"`
}

// User公共的信息
type UserInfo struct {
	UserId   uint32 `json:"user_id"`
	Username string `json:"username"`
}

type User struct {
	UserId   uint32 `json:"user_id"`
	Username string `json:"username"`
	UserPwd  string `json:"uesr_pwd"`
}

type DataLogin struct {
	UserId  uint32 `json:"user_id"`
	UserPwd string `json:"user_pwd"`
}

type DataRegister struct {
	User User `json:"user"`
}

type Response struct {
	Infos []UserInfo `json:"infos"` // 获取多个用户信息时，自此取
	Error string     `json:"error"` // 发生错误时，自此取
	Text  string     `json:"text"`  //发送聊天内容时，自此取
}
