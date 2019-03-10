package view

import (
	"chatroom/client/process"
	"chatroom/common/message"
	"fmt"
	"strconv"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var mw = new(MyMainWindow)

type MyMainWindow struct {
	*walk.MainWindow

	userInput *walk.LineEdit
	pwdInput  *walk.LineEdit

	loginBtn *walk.PushButton

	info *walk.TextLabel
}

func WindowShow() {

	// 注册 登陆信息
	message.Center.RegisterMsg(message.TypeLogin, loginTrigger)

	// 注册 注册信息
	message.Center.RegisterMsg(message.TypeRegister, registerTrigger)

	// 注册 获取当前在线用户信息
	message.Center.RegisterMsg(message.TypeGetOnlineUsers, getOnlineUsersTrigger)

	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "登录",
		MinSize:  Size{300, 200},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2, Spacing: 10},
				Children: []Widget{
					VSplitter{
						Children: []Widget{
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{Text: "用户名"},
									LineEdit{AssignTo: &mw.userInput},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{Text: "密码"},
									LineEdit{AssignTo: &mw.pwdInput},
								},
							},
							PushButton{
								Text:     "登陆",
								AssignTo: &mw.loginBtn,
								OnClicked: func() {
									if id, ok := checkInputUserID(mw.userInput.Text()); ok {
										if pwd, ok := checkInputUserPwd(mw.pwdInput.Text()); ok {
											process.Instance.Login(id, pwd)
											mw.info.SetVisible(true)
											// mw.Hide()
											// OpenChatRoom()
										}
									}
								},
							},
						},
					},
				},
			},
			TextLabel{
				Text:     "当前信息",
				AssignTo: &mw.info,
				Visible:  false,
			},
		},
	}.Run()

}

func checkInputUserID(input string) (id uint32, ok bool) {
	ok = false
	if input == "" {
		var tmp walk.Form
		walk.MsgBox(tmp, "用户名为空", "", walk.MsgBoxIconInformation)
		return
	}
	res, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return
	}
	id = uint32(res)
	ok = true
	return
}

func checkInputUserPwd(input string) (pwd string, ok bool) {
	ok = true
	pwd = input
	return
}

func loginTrigger(msgType message.MsgType, res *message.Response) (err error) {
	switch msgType {
	case message.CodeLoginSuccess:
		fmt.Println("登陆成功")

		mw.info.SetVisible(false)
		mw.Hide()
		OpenChatRoom()
	case message.CodeLoginFailed:
		mw.info.SetText(res.Error)
	default:
	}
	return
}

func registerTrigger(msgType message.MsgType, res *message.Response) (err error) {
	switch msgType {
	case message.CodeRegisterSuccess:
		fmt.Println("注册成功")

	case message.CodeRegisterFailed:

	default:
	}
	return
}

func getOnlineUsersTrigger(msgType message.MsgType, res *message.Response) (err error) {

	for _, info := range res.Infos {
		fmt.Printf("\tid:%d\tname:%s\n", info.UserId, info.Username)
	}
	return
}
