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
	app.SetApplicationVersion("0.0.1")
	app.SetApplicationName("FireRain安装器")

	ui.NewMainWindow().Show()

	os.Exit(app.Exec())
}
