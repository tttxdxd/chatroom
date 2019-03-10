package view

import (
	"chatroom/client/process"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type ChatRoomWindow struct {
	*walk.MainWindow

	chatBox     *walk.ListBox //聊天框
	msgInput    *walk.LineEdit
	msgInputBtn *walk.PushButton
}

func OpenChatRoom() {
	mw := new(ChatRoomWindow)

	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "聊天室",
		MinSize:  Size{400, 600},
		Layout:   VBox{},
		Children: []Widget{
			ListBox{
				AssignTo: &mw.chatBox,
				Row:      5,
			},
			LineEdit{
				AssignTo: &mw.msgInput,
			},
			PushButton{
				AssignTo: &mw.msgInputBtn,
				OnClicked: func() {
					process.Instance.Exit()
				},
			},
		},
	}.Run()
}
