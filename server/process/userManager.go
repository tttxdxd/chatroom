package process

import (
	"chatroom/common/message"
)

var UserManager userManager

type userManager struct {
	onlineUserList map[uint32]*UserProcess // 需要初始化 分配空间
}

func init() {
	UserManager = userManager{
		onlineUserList: make(map[uint32]*UserProcess),
	}
}

// 添加用户到在线用户列表
func (this *userManager) AddUser(userProcess *UserProcess) {
	this.onlineUserList[userProcess.userId] = userProcess
}

func (this *userManager) RemoveUser(userId uint32) {
	delete(this.onlineUserList, userId)
}

func (this *userManager) GetUserProcessById(userId uint32) (userProcess *UserProcess) {
	if userProcess, ok := this.onlineUserList[userId]; ok {
		return userProcess
	}
	return nil
}

func (this *userManager) GetAllUsersInfo(userId uint32) (infos []message.UserInfo) {
	infos = make([]message.UserInfo, 0, len(this.onlineUserList))
	for _, v := range this.onlineUserList {
		if userId == v.userId {
			continue
		}
		info := message.UserInfo{
			UserId:   v.userId,
			Username: v.username,
		}
		infos = append(infos, info)
	}
	return
}
