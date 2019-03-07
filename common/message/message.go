package message

const (
	TypeLogin = iota
	TypeRegister
	TypeClientExit
)

const (
	CodeLoginSuccess    = iota // 登陆成功
	CodeLoginFailed            // 登陆失败 1.密码错误 2.用户名不存在 3.服务器错误
	CodeRegisterSuccess        // 注册成功
	CodeRegisterFailed         // 注册失败 1.用户名已存在
)

type Msg struct {
	Type uint   `json:"type"`
	Data string `json:"data"`
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
	Code  uint32 `json:"code"`
	Error string `json:"error"`
}
