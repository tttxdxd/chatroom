package process

import (
	"chatroom/common/message"
	"fmt"
	"net"
)

func UserLoginProcess(conn net.Conn, msg *message.DataLogin) (err error) {
	var response message.Response
	response.Code = message.CodeLoginSuccess
	response.Error = ""
	err = message.WriteResponse(conn, response)
	if err != nil {
		fmt.Println(" message.WriteMsg(conn, response) error:", err)
		return
	}
	return
}
