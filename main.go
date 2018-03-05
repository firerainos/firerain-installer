package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

func main() {
	app := gui.NewQGuiApplication(len(os.Args), os.Args)

	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	quickcontrols2.QQuickStyle_SetStyle("material")

	engine := qml.NewQQmlApplicationEngine(nil)

	engine.Load(core.NewQUrl3("qml/main.qml", 0))

	gui.QGuiApplication_Exec()
}
