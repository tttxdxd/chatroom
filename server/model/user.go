package model

import (
	"encoding/json"
)

type User struct {
	UserId   uint32
	UserPwd  string
	UserName string
}

func (this *User) Serializer() (res string, err error) {
	data, err := json.Marshal(*this)
	if err != nil {
		return
	}
	res = string(data)
	return
}

func (this *User) Unserializer(data string) (err error) {
	err = json.Unmarshal([]byte(data), this)
	return
}
