package view

import (
	"chatroom/common/message"

	"github.com/lxn/walk"
)

type infoTableModel struct {
	walk.TableModelBase
	items []*message.UserInfo
}

func newInfoTableModel() *infoTableModel {
	onlineUserModel := new(infoTableModel)
	onlineUserModel.items = []*message.UserInfo{}
	return onlineUserModel
}

func (m *infoTableModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.UserId

	case 1:
		return item.Username

	}

	panic("unexpected col")
}

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *infoTableModel) RowCount() int {
	return len(m.items)
}

func (m *infoTableModel) AddItem(item *message.UserInfo) {
	m.items = append(m.items, item)
}

func (m *infoTableModel)RemoveItem(userId uint32){
	for i,item:=range m.items{
		if item.UserId==userId{
			temp:=m.items[i+1:]
			m.items=m.items[:i]
			m.items=append(m.items,temp...)
			break
		}
	}
}