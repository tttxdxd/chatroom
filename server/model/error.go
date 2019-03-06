package model

import (
	_ "errors"
)

var (
	ERROR_SERVER         = New("服务器内部错误")
	ERROR_USER_PWD       = New("密码错误")
	ERROR_USER_NOTEXISTS = New("该用户不存在")
	ERROR_USER_EXISTS    = New("该用户已存在")
	ERROR_USER_FORMAT    = New("用户存储格式错误")
)
