package ui

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
)

type MainWindow struct {
	mainWindow *widgets.QMainWindow
}

func NewMainWindow() *MainWindow {
	mainWindow := widgets.NewQMainWindow(nil,0)

	m := MainWindow{mainWindow}
	m.initUI()

	return &m
}

func (m *MainWindow) initUI(){
	m.mainWindow.SetWindowTitle("firerain-installer")
	m.mainWindow.SetFixedSize(core.NewQSize2(900,700))
}

func (m *MainWindow) Show() {
	m.mainWindow.Show()
}