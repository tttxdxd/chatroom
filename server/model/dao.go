package model

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

// data assace object

var UserDao *UserDAO

type UserDAO struct {
	key string
}

func InitDAO() {
	fmt.Println("开始初始化UserDAO...")
	UserDao = &UserDAO{
		key: "users",
	}
	fmt.Println("初始化UserDAO成功...")
	fmt.Println("-------------------------------------")
}

// 根据用户id获取数据库里的用户信息
func (this *UserDAO) getUserById(id uint32) (user *User, err error) {
	client := redisPool.Get()
	res, err := client.HGet(this.key, string(id)).Result()
	if err != nil { //需判断错误是否是因连接断开引起的还是未找到该用户
		fmt.Println("client.HGet(this.key, string(userId)).Result() error:", err)
		if err == redis.Nil {
			err = ERROR_USER_NOTEXISTS
		} else {
			err = ERROR_SERVER
		}
		return
	}
	user = &User{}
	err = user.Unserializer(res)
	if err != nil {
		fmt.Println("user.Unserializer(res) error:", err)
		err = ERROR_USER_FORMAT
		return
	}
	return
}

func (this *UserDAO) addUser(user *User) (err error) {
	client := redisPool.Get()
	data, err := json.Marshal(*user)
	err = client.HSet(this.key, string(user.UserId), data).Err()
	if err != nil {
		fmt.Println(" client.HSet(this.key, string(user.UserId), user).Result() error:", err)
		return
	}
	return
}

func (this *UserDAO) Login(userId uint32, userPwd string) (user *User, err error) {
	user, err = this.getUserById(userId)
	if err != nil {
		fmt.Println("getUserById(userId) error:", err)
		return
	}
	if userPwd != user.UserPwd {
		user = nil
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDAO) Register(user *User) (err error) {
	_, err = this.getUserById(user.UserId)
	if err != nil { //数据库中无此账号id，可以注册
		err = this.addUser(user)
		return
	}
	return
}
