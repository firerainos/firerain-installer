package main

import (
	"os"
	"github.com/therecipe/qt/widgets"
	"github.com/firerainos/firerain-installer/ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	ui.NewMainWindow().Show()

	widgets.QApplication_Exec()
}
