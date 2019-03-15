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

// 获取所有在线用户信息
func (this *userManager) GetAllUsersInfo() (infos []message.UserInfo) {
	infos = make([]message.UserInfo, 0, len(this.onlineUserList))
	for _, v := range this.onlineUserList {
		info := message.UserInfo{
			UserId:   v.userId,
			Username: v.username,
		}
		infos = append(infos, info)
	}
	return
}

// 通知所有在线用户，除了自己
func (this *userManager) NotifyAllUsers(p *UserProcess, msg *message.Msg) (err error) {

	info := &message.UserInfo{
		p.userId,
		p.username,
	}

	for id, v := range this.onlineUserList {
		if p.userId == id {
			continue
		}

		err = v.UserNotifyOneUserProcess(info, msg)
		if err != nil {
			return
		}
	}
	return

}
