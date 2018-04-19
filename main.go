package main

import (
	"github.com/firerainos/firerain-installer/ui"
	"github.com/therecipe/qt/widgets"
	"os"
	"github.com/firerainos/firerain-installer/config"
)

func main() {
	config.InitConfig()

	widgets.NewQApplication(len(os.Args), os.Args)

	ui.NewMainWindow().Show()

	widgets.QApplication_Exec()
}
