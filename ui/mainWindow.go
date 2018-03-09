package ui

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
)

type MainWindow struct {
	*widgets.QMainWindow
}

func NewMainWindow() *MainWindow {
	w := widgets.NewQMainWindow(nil,0)

	m := &MainWindow{w}
	m.init()

	return m
}

func (m *MainWindow) init() {
	m.SetWindowTitle("firerain-installer")
	m.SetFixedSize(core.NewQSize2(900,700))

	frame := NewMainFrame(m,0)
	m.SetCentralWidget(frame)
}