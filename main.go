package main

import (
	"github.com/firerainos/firerain-installer/ui"
	"github.com/therecipe/qt/widgets"
	"os"
	"github.com/firerainos/firerain-installer/config"
)

func main() {
	config.InitConfig()

	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetApplicationVersion("0.1.0")
	app.SetApplicationName("FireRain安装器")
	app.SetWindowIcon(gui.NewQIcon5("/usr/share/icons/hicolor/192x192/apps/firerain-installer.svg"))

	ui.NewMainWindow().Show()

	os.Exit(app.Exec())
}
