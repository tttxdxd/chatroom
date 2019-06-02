package process

import (
	"chatroom/common/message"
)

// 管理在线用户列表
var UserManager userManager

type userManager struct {
	onlineUserList map[uint32]*message.UserInfo // 需要初始化 分配空间
}

func init() {
	UserManager = userManager{
		onlineUserList: make(map[uint32]*message.UserInfo),
	}
}

func (this *userManager) Clear() {
	this.onlineUserList = make(map[uint32]*message.UserInfo)
}

// 添加用户到在线用户列表
func (this *userManager) AddUser(user *message.UserInfo) {
	this.onlineUserList[user.UserId] = user
}

func (this *userManager) RemoveUser(userId uint32) {
	delete(this.onlineUserList, userId)
}

func (this *userManager) GetUserProcessById(userId uint32) (user *message.UserInfo) {
	if user, ok := this.onlineUserList[userId]; ok {
		return user
	}
	return nil
}
