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

	mainWindow.SetWindowTitle("firerain-installer")
	mainWindow.SetFixedSize(core.NewQSize2(900,700))

	frame := NewMainFrame(mainWindow,0)
	mainWindow.SetCentralWidget(frame.frame)

	m := MainWindow{mainWindow}

	return &m
}

func (m *MainWindow) Show() {
	m.mainWindow.Show()
}