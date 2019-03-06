package message

const (
	TypeLogin = iota
	TypeRegister
	TypeResponseLogin
)

const (
	CodeLoginSuccess = iota
	CodeLoginFailed
)

type Msg struct {
	Type uint   `json:"type"`
	Data string `json:"data"`
}

type DataLogin struct {
	UserId  uint32 `json:"user_id"`
	UserPwd string `json:"user_pwd"`
}

type DataRegister struct {
	UserId   uint32 `json:"user_id"`
	Username string `json:"username"`
	UserPwd  string `json:"uesr_pwd"`
}

type Response struct {
	Code  uint32 `json:"code"`
	Error string `json:"error"`
}
