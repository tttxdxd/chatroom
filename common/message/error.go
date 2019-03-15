package message

import (
	"errors"
)

var (
	ERROR_LEN_OF_READ = errors.New("读取的长度不匹配")
	ERROR_DISCONNECT  = errors.New("连接断开")
)
