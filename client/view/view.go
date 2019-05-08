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
	loginBtn  *walk.PushButton
	loginBox  *walk.Composite

	ruserInput  *walk.LineEdit
	rpwdInput   *walk.LineEdit
	registerBtn *walk.PushButton
	registerBox *walk.Composite

	selectLoginBtn    *walk.PushButton
	selectRegisterBtn *walk.PushButton

	info  *walk.TextLabel
	rinfo *walk.TextLabel
}

func WindowShow() {

	// 注册 登陆信息
	message.Center.RegisterMsg(message.TypeLogin, loginTrigger)

	// 注册 注册信息
	message.Center.RegisterMsg(message.TypeRegister, registerTrigger)

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "登录",
		MinSize:  Size{300, 200},
		Layout:   VBox{},
		Children: []Widget{
			GroupBox{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						AssignTo: &mw.selectRegisterBtn,
						Text:     "注册界面",
						OnClicked: func() {
							mw.SetTitle("注册")
							mw.loginBox.SetVisible(false)
							mw.registerBox.SetVisible(true)
							mw.selectLoginBtn.SetEnabled(true)
							mw.selectRegisterBtn.SetEnabled(false)
						},
					},
					PushButton{
						AssignTo: &mw.selectLoginBtn,
						Text:     "登陆界面",
						OnClicked: func() {
							mw.SetTitle("登陆")
							mw.loginBox.SetVisible(true)
							mw.registerBox.SetVisible(false)
							mw.selectLoginBtn.SetEnabled(false)
							mw.selectRegisterBtn.SetEnabled(true)
						},
					},
				},
			},
			Composite{
				AssignTo: &mw.loginBox,
				Layout:   Grid{Columns: 2, Spacing: 10},
				Children: []Widget{
					VSplitter{
						Children: []Widget{
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{Text: "用户ID"},
									LineEdit{AssignTo: &mw.userInput},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{Text: "密  码"},
									LineEdit{AssignTo: &mw.pwdInput},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									TextLabel{
										Text:     "当前信息",
										AssignTo: &mw.info,
										Visible:  true,
									},
									PushButton{
										Text:     "登陆",
										AssignTo: &mw.loginBtn,
										OnClicked: func() {
											if id, ok := checkInputUserID(mw.userInput.Text()); ok {
												if ok := checkInputUserPwd(mw.pwdInput.Text()); ok {
													process.Instance.Login(id, mw.pwdInput.Text())

													mw.info.SetText("登陆中...")

												}
											}
										},
									},
								},
							},
						},
					},
				},
			},
			Composite{
				Visible:  false,
				AssignTo: &mw.registerBox,
				Layout:   Grid{Columns: 2, Spacing: 10},
				Children: []Widget{
					VSplitter{
						Children: []Widget{
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{Text: "用户名"},
									LineEdit{AssignTo: &mw.ruserInput},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{Text: "密  码"},
									LineEdit{AssignTo: &mw.rpwdInput},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									TextLabel{
										Text:     "当前信息",
										AssignTo: &mw.rinfo,
										Visible:  true,
									},
									PushButton{
										Text:     "注册",
										AssignTo: &mw.registerBtn,
										OnClicked: func() {
											if ok := checkInputUsername(mw.ruserInput.Text()); ok {
												if ok := checkInputUserPwd(mw.rpwdInput.Text()); ok {
													process.Instance.Register(mw.ruserInput.Text(), mw.rpwdInput.Text())

												}
											}
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}).Create(); err != nil {

	}
	mw.selectRegisterBtn.SetEnabled(false)
	mw.Run()

}

func checkInputUserID(input string) (id uint32, ok bool) {
	ok = false
	if input == "" {
		mw.info.SetText("用户ID为空")
		return
	}
	res, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		mw.info.SetText("用户ID格式错误")
		return
	}
	id = uint32(res)
	ok = true
	return
}

func checkInputUsername(input string) (ok bool) {
	ok = true
	return
}

func checkInputUserPwd(input string) (ok bool) {
	ok = true
	return
}

func loginTrigger(msgType message.MsgType, res *message.Response) (err error) {
	switch msgType {
	case message.CodeLoginSuccess:
		fmt.Println("登陆成功")
		fmt.Println(res)
		mw.Hide()
		OpenChatRoom(res.Infos[0])
	case message.CodeLoginFailed:
		mw.info.SetText(res.Error)
	default:
	}
	return
}

func registerTrigger(msgType message.MsgType, res *message.Response) (err error) {
	switch msgType {
	case message.CodeRegisterSuccess:
		id := strconv.FormatUint(uint64(res.Infos[0].UserId), 10)
		mw.rinfo.SetText("当前ID：" + id + "注册成功，请前往登陆")
		mw.userInput.SetText(id)
	case message.CodeRegisterFailed:
		mw.rinfo.SetText(res.Error)
	default:
	}
	return
}
