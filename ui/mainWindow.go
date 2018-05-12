package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os/exec"
)

type MainWindow struct {
	*widgets.QMainWindow
}

func NewMainWindow() *MainWindow {
	w := widgets.NewQMainWindow(nil, 0)

	m := &MainWindow{w}
	m.init()
	m.initMenu()

	return m
}

func (m *MainWindow) init() {
	m.SetWindowTitle("firerain-installer")
	m.SetFixedSize(core.NewQSize2(1000, 700))

	frame := NewMainFrame(m, 0)
	m.SetCentralWidget(frame)
}


func (m *MainWindow) initMenu() {
	toolMenu := m.MenuBar().AddMenu2("&工具")

	pmAction := toolMenu.AddAction("分区管理器")
	terminalAction :=toolMenu.AddAction("终端")

	pmAction.ConnectTriggered(func(checked bool) {
		exec.Command("/usr/bin/partitionmanager").Start()
	})

	terminalAction.ConnectTriggered(func(checked bool) {
		exec.Command("/usr/bin/konsole").Start()
	})
}