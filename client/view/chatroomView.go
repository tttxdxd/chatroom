package view

import (
	"chatroom/client/process"
	"chatroom/common/message"
	"fmt"
	"strconv"
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var room = new(ChatRoomWindow)
var currentUser message.UserInfo

type ChatRoomWindow struct {
	*walk.MainWindow

	chatBox *walk.ScrollView //聊天框

	msgInput        *walk.TextEdit
	msgInputBtn     *walk.PushButton
	onlineUserTable *walk.TableView
	onlineUserModel *infoTableModel
}

func OpenChatRoom(user message.UserInfo) {

	defer process.Instance.Exit()

	message.Center.RegisterPassiveMsg(message.CodeNotifyUserOnline, notifyUserOnlineTrigger)
	message.Center.RegisterPassiveMsg(message.CodeNotifyUserLogout, notifyUserLogoutTrigger)

	message.Center.RegisterPassiveMsg(message.CodeRecvMessage, recvMessageTrigger)

	// 注册 获取当前在线用户信息
	message.Center.RegisterMsg(message.TypeGetOnlineUsers, getOnlineUsersTrigger)

	currentUser = user

	room.onlineUserModel = newInfoTableModel()

	if err := (MainWindow{
		AssignTo: &room.MainWindow,
		Title:    "聊天室",
		MinSize:  Size{400, 400},
		Layout:   HBox{MarginsZero: true},
		Children: []Widget{
			Composite{
				MinSize: Size{500, 500},
				Layout:  VBox{},
				Children: []Widget{
					//=========================
					ScrollView{
						MinSize:         Size{500, 300},
						Background:      SolidColorBrush{Color: walk.RGB(255, 255, 255)},
						AssignTo:        &room.chatBox,
						Layout:          Flow{MarginsZero: true, SpacingZero: true},
						VerticalFixed:   false,
						HorizontalFixed: true,
					},
					GroupBox{
						Layout:    VBox{},
						Alignment: AlignHCenterVFar,
						MaxSize:   Size{1000, 150},
						MinSize:   Size{500, 150},
						Children: []Widget{
							TextEdit{
								MaxSize:  Size{1000, 130},
								MinSize:  Size{500, 90},
								AssignTo: &room.msgInput,
								VScroll:  true,
							},
							PushButton{
								Text:     "发送",
								AssignTo: &room.msgInputBtn,
								OnClicked: func() {
									input := room.msgInput.Text()
									if ok := checkInput(input); ok {
										room.msgInput.SetText("")
										process.Instance.SendMessage(input)
										appendMessage(room.chatBox, currentUser, input)
									}
								},
							},
						},
					},
					//================

				},
			},
			Composite{
				MaxSize: Size{220, 500},
				MinSize: Size{220, 500},
				Layout:  VBox{},
				Children: []Widget{
					TableView{
						AssignTo: &room.onlineUserTable,
						Model:    room.onlineUserModel,
						Columns: []TableViewColumn{
							{Title: "用户ID"},
							{Title: "用户名"},
						},
					},
				},
			},
		},
	}).Create(); err != nil {

	}

	process.Instance.GetOnlineUsers()
	process.Instance.NotifyOnlineUsers()

	AddInfo(room.chatBox, time.Now().Format("2006-01-02 15:04:05"), "您已经进入了聊天室~")

	room.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		CloseView()
	})

	room.Run()
}

func checkInput(input string) (ok bool) {
	ok = true
	return
}

func appendMessage(box *walk.ScrollView, user message.UserInfo, msg string) {

	info := strconv.FormatUint(uint64(user.UserId), 10) + " : " + user.Username

	timeS := time.Now().Format("2006-01-02 15:04:05")

	right := true
	if user != currentUser {
		right = false
	}

	AddChatBubble(box, timeS, msg, info, right)
}

func recvMessageTrigger(response *message.Response) {
	user := response.Infos[0]
	appendMessage(room.chatBox, user, response.Text)
}

func notifyUserLogoutTrigger(response *message.Response) {
	user := response.Infos[0]
	info := strconv.FormatUint(uint64(user.UserId), 10) + " : " + user.Username + "下线了"
	timeS := time.Now().Format("2006-01-02 15:04:05")

	process.UserManager.RemoveUser(user.UserId)
	room.onlineUserModel.RemoveItem(user.UserId)
	room.onlineUserModel.PublishRowsReset()
	AddInfo(room.chatBox, timeS, info)

	fmt.Println(response.Infos[0])
}

func notifyUserOnlineTrigger(response *message.Response) {
	fmt.Println(response.Infos[0])

	user := response.Infos[0]
	info := strconv.FormatUint(uint64(user.UserId), 10) + " : " + user.Username + "上线了"
	timeS := time.Now().Format("2006-01-02 15:04:05")

	process.UserManager.AddUser(&user)
	room.onlineUserModel.AddItem(&user)
	room.onlineUserModel.PublishRowsReset()
	AddInfo(room.chatBox, timeS, info)

	fmt.Println(response.Infos[0])
}

func AddInfo(box *walk.ScrollView, time, msg string) {

	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	box.Synchronize(func() {
		//composite, _ := walk.NewComposite(box)
		// composite.SetAlignment(walk.AlignHCenterVFar)
		// composite.SetLayout(walk.NewVBoxLayout())

		// data, _ := walk.NewTextLabel(composite)
		// data.SetTextAlignment(walk.AlignHCenterVFar)
		// data.SetText(time)

		// info, _ := walk.NewTextLabel(composite)
		// info.SetTextAlignment(walk.AlignHCenterVFar)
		// info.SetText(msg)

		// composite.Children().Add(data)
		// composite.Children().Add(info)

		// box.Children().Add(composite)
		// box.Layout().Update(false)

		composite, _ := walk.NewComposite(box)
		Composite{
			MinSize:   Size{500, 32},
			AssignTo:  &composite,
			Layout:    VBox{},
			Alignment: AlignHCenterVFar,
			Children: []Widget{
				TextLabel{

					TextAlignment: AlignHCenterVFar,
					Text:          time,
				},
				TextLabel{
					TextAlignment: AlignHCenterVFar,
					Text:          msg,
				},
			},
		}.Create(NewBuilder(box))
		box.Layout().Update(false)
		box.WndProc(box.Handle(), uint32(277), 3, 0)
	})
}

// 可建池优化
func AddChatBubble(box *walk.ScrollView, time, msg, userinfo string, right bool) {
	// box.Synchronize(func() {
	// 	composite, _ := walk.NewComposite(box)
	// 	composite.SetAlignment(walk.AlignHCenterVFar)
	// 	composite.SetLayout(walk.NewVBoxLayout())
	// 	composite.SetMinMaxSize(walk.Size{400, 130}, walk.Size{400, 330})
	// 	composite.SetSize(walk.Size{400, 130})

	// 	data, _ := walk.NewTextLabel(box)
	// 	data.SetTextAlignment(walk.AlignHCenterVFar)
	// 	data.SetText(time)

	// 	info, _ := walk.NewTextLabel(box)

	// 	if right {
	// 		info.SetTextAlignment(walk.AlignHFarVNear)
	// 		info.SetText(userinfo + "   <<<<")
	// 	} else {
	// 		info.SetTextAlignment(walk.AlignHNearVNear)
	// 		info.SetText(">>>>   " + userinfo)
	// 	}

	// 	bubble, _ := walk.NewComposite(box)
	// 	bubble.SetLayout(walk.NewHBoxLayout())

	// 	infoL, _ := walk.NewTextLabel(box)
	// 	infoL.SetTextAlignment(walk.AlignHNearVNear)
	// 	if right {
	// 		infoL.SetText("　　")
	// 	} else {
	// 		infoL.SetText("　　")
	// 	}
	// 	infoR, _ := walk.NewTextLabel(box)
	// 	infoR.SetTextAlignment(walk.AlignHFarVNear)
	// 	if right {
	// 		infoR.SetText("> ^_^")
	// 	} else {
	// 		infoR.SetText("　　")
	// 	}

	// 	msgBubble, _ := walk.NewTextEdit(box)
	// 	msgBubble.SetEnabled(false)
	// 	msgBubble.SetReadOnly(true)
	// 	msgBubble.SetAlignment(walk.AlignHNearVNear)
	// 	msgBubble.SetText(msg)

	// 	width := getStrWidth(msg) * 8
	// 	msgBubble.SetHeight((width/396 + 1) * 16)

	// 	bubble.Children().Add(infoL)
	// 	bubble.Children().Add(msgBubble)
	// 	bubble.Children().Add(infoR)
	// 	//bubble.Layout().Update(true)
	// 	bubble.Layout().SetSpacing(0)
	// 	bubble.Layout().SetMargins(walk.Margins{0, 0, 0, 0})

	// 	composite.Children().Add(data)
	// 	composite.Children().Add(info)
	// 	composite.Children().Add(bubble)
	// 	composite.Layout().SetSpacing(0)
	// 	composite.Layout().SetMargins(walk.Margins{10, 0, 10, 0})
	// 	//composite.Layout().Update(true)

	// 	box.Children().Add(composite)
	// 	box.Layout().Update(false)

	// 	box.WndProc(box.Handle(), uint32(277), 3, 0)
	// })

	// 以下可以动态改变高度，实现聊天气泡大小的自适应（根据文字量）

	box.Synchronize(func() {

		width := getStrWidth(msg) * 8
		bubbleWidth := 320

		composite, _ := walk.NewComposite(box)

		infoText := ">>>>   " + userinfo
		infoTextAlignment := AlignHNearVNear
		if right {
			infoTextAlignment = AlignHFarVNear
			infoText = userinfo + "   <<<<"
		}

		Composite{
			AssignTo: &composite,
			Layout:   VBox{Margins: Margins{10, 16, 10, 0}},
			MinSize:  Size{400, ((width / bubbleWidth) + 1) * 16},
			Children: []Widget{
				TextLabel{
					TextAlignment: AlignHCenterVFar,
					Text:          time,
				},
				TextLabel{
					TextAlignment: infoTextAlignment,
					Text:          infoText,
				},
				Composite{
					MinSize: Size{400, ((width / bubbleWidth) + 1) * 16},
					Layout:  HBox{MarginsZero: true},
					Children: []Widget{
						TextLabel{
							TextAlignment: AlignHNearVNear,
							Text:          ">>>>",
						},
						TextEdit{

							Enabled:   false,
							Alignment: AlignHNearVNear,
							Text:      msg,
						},
						TextLabel{
							TextAlignment: AlignHFarVNear,
							Text:          "<<<<",
						},
					},
				},
			},
		}.Create(NewBuilder(box))
		box.Layout().Update(false)
		box.WndProc(box.Handle(), uint32(277), 3, 0)
	})
}

func getStrWidth(s string) (width int) {
	width = 0
	for _, c := range []rune(s) {
		if c < 128 {
			width++
		} else {
			width += 2
		}
	}
	return
}

func getOnlineUsersTrigger(msgType message.MsgType, res *message.Response) (err error) {

	switch msgType {
	case message.CodeOnlineUsers:
		for i, _ := range res.Infos {
			process.UserManager.AddUser(&res.Infos[i])
			room.onlineUserModel.AddItem(&res.Infos[i])

		}
		room.onlineUserModel.PublishRowsReset()
	default:
	}
	return
}
