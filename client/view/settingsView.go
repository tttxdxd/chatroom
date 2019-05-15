package view

import (
	"chatroom/client/process"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io"
	"os"
	"path/filepath"
)

type SettingsView struct {
	*walk.MainWindow

	hostInput *walk.TextEdit
}

var setting = new(SettingsView)

func OpenSettingsView() {
	if err := (MainWindow{
		AssignTo: &setting.MainWindow,
		Title:    "设置",
		MinSize:  Size{400, 100},
		Layout:   HBox{MarginsZero: true},
		Children: []Widget{
			GroupBox{
				Layout: VBox{},
				Children: []Widget{
					TextLabel{
						Text: "请输入 ip:port，默认为 127.0.0.1:8889",
					},
					GroupBox{
						Layout: HBox{},
						Children: []Widget{
							TextEdit{
								AssignTo: &setting.hostInput,
							},
						},
					},
					PushButton{
						Text: "保存并连接",
						OnClicked: func() {
							Address = saveToFile()
							if Connected == "已连接" {
								process.ChangeCh()
							}
							StateChange()
							_ = setting.Close()
						},
					},
				},
			},
		},
	}).Create(); err != nil {

	}

	readHostFromFile()
	setting.Run()
}

func camp(value, min, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}

func readHost() (res string) {
	res = "127.0.0.1:8889"
	path, err := os.Executable()
	if err != nil {
		println(err)
		return
	}
	configPath := filepath.Join(path, "../settings.txt")
	println(configPath)
	file, err := os.OpenFile(configPath, os.O_CREATE, 0666)
	if err != nil {
		println(err)
		return
	}
	defer file.Close()

	buf := make([]byte, 64)
	n, err := file.Read(buf)
	if err != nil {
		println(err)
	}
	if n > 0 {
		res = string(buf[:n])
	}
	return
}

func readHostFromFile() {
	path, err := os.Executable()
	if err != nil {
		println(err)
		return
	}
	configPath := filepath.Join(path, "../settings.txt")
	println(configPath)
	file, err := os.OpenFile(configPath, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		println(err)
		return
	}
	defer file.Close()

	buf := make([]byte, 64)
	n, err := file.Read(buf)
	if err != nil {
		println(err)
	}
	if n > 0 {
		_ = setting.hostInput.SetText(string(buf[:n]))
		if err != nil {
			println(err)
			return
		}
	}else{
		_ = setting.hostInput.SetText("127.0.0.1:8889")
	}

}

func saveToFile() string {
	s := setting.hostInput.Text()
	path, err := os.Executable()
	if err != nil {
		println(err)
		return s
	}
	configPath := filepath.Join(path, "../settings.txt")
	println(configPath)
	file, err := os.OpenFile(configPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		println(err)
		return s
	}
	defer file.Close()

	println("write string ", s)
	_, err = io.WriteString(file, s)
	//_, err = file.WriteString(s)
	if err != nil {
		println(err)
		return s
	}
	return s
}
