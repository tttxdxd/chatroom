package message

import (
	"encoding/json"
	"errors"
)

// 实现客户端，服务端收发消息的解耦
// 原模式 writeMsg() 后接readMsg()
// 更改后 如下

var (
	ERROR_TRIGGER_EMPTY     = errors.New("此类型匹配的触发器为空") //消息类型不匹配
	ERROR_CENTER_MISMATCHES = errors.New("无此类型匹配的消息中心,该消息类型未注册")
	ERROR_QUEUE_MISMATCHES  = errors.New("待处理消息队列中无此ID")
	ERROR_MSG_PARSE         = errors.New("消息解析出错")
)

var Center = MsgCenter{
	center:        make(map[MsgType]MsgTrigger),
	msgQueue:      make(map[MsgID]MsgType),
	passiveCenter: make(map[MsgType]func(*Response)),
}

// 收到此消息类型后，触发回调，简称触发器
type MsgTrigger func(MsgType, *Response) error

// message 处理中心
// 消息处理流程
// 1. 初始化并注册 根据信息类型处理信息的所有回调
// 2. 发送消息到服务端（带ID）,同时将该ID及Type加入消息待处理队列
// 3. 服务端回送一条回应消息（带对应ID），根据ID找到此回应消息对应的发送消息类型
// 3. 根据发送消息类型，找到对应消息处理中心
// 4. 根据回应消息类型，在对应消息处理中心中触发对应触发器，将信息从待处理队列中移除
type MsgCenter struct {
	center        map[MsgType]MsgTrigger
	msgQueue      map[MsgID]MsgType
	passiveCenter map[MsgType]func(*Response) //被动消息中心
}

// 注册模板
// MsgCenter.RegisterMsg(message.TypeXXX,func(MsgType, *Response) error{
//
// })
//
// 注册消息
func (this *MsgCenter) RegisterMsg(msgType MsgType, trigger MsgTrigger) {
	this.center[msgType] = trigger
}

// 添加消息到待处理队列 //只有客户端发送消息到服务端且需要回应的消息才需要加入待处理队列
func (this *MsgCenter) AddMsg(msg *Msg) {
	msg.ID = getNextId()
	this.msgQueue[msg.ID] = msg.Type
}

// 移除消息到待处理队列
func (this *MsgCenter) RemoveMsg(msg *Msg) {
	delete(this.msgQueue, msg.ID)
}

// 注册 被动 消息
func (this *MsgCenter) RegisterPassiveMsg(msgType MsgType, callback func(*Response)) {
	this.passiveCenter[msgType] = callback
}

// 分发收到的消息与待处理队列匹配，然后由各自对应的触发器处理
func (this *MsgCenter) Distribute(msg *Msg) (err error) {

	//不在消息队列中,为接受到的消息而非回应
	if callback, ok := this.passiveCenter[msg.Type]; ok {
		var response Response
		err = json.Unmarshal([]byte(msg.Data), &response)
		if err != nil {
			return
		}

		callback(&response)
		return
	}

	if msgType, ok := this.msgQueue[msg.ID]; ok {

		this.RemoveMsg(msg)

		if trigger, ok := this.center[msgType]; ok {
			if trigger != nil {
				var response Response
				err = json.Unmarshal([]byte(msg.Data), &response)
				if err != nil {
					return
				}

				err = trigger(msg.Type, &response)
			} else {
				err = ERROR_TRIGGER_EMPTY
			}
			return
		}
		err = ERROR_CENTER_MISMATCHES
		return
	}
	err = ERROR_QUEUE_MISMATCHES
	return
}

var currentId MsgID

// 获取不同的ID号
func getNextId() (id MsgID) {
	id = currentId
	currentId++
	return
}
