package main

import (
	"github.com/firerainos/firerain-installer/config"
	"github.com/firerainos/firerain-installer/ui"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	config.InitConfig()

	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetApplicationVersion("0.2.0")
	app.SetApplicationName("FireRain安装器")
	app.SetWindowIcon(gui.NewQIcon5("/usr/share/icons/hicolor/192x192/apps/firerain-installer.svg"))

	ui.NewMainWindow().Show()

	os.Exit(app.Exec())
}
