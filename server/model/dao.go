package model

import (
	"chatroom/common/message"
	"fmt"
)

// data assace object

var UserDao UserDAO

type UserDAO interface {
	Login(userId uint32, userPwd string) (user *message.User, err error)
	Register(user *message.User) (err error)
}

func InitDAO() {
	fmt.Println("开始初始化UserDAO...")
	UserDao = &RedisDAO{
		key: "users",
	}
	fmt.Println("初始化UserDAO成功...")
	fmt.Println("-------------------------------------")
}
