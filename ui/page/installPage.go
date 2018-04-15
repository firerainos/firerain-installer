package page

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type InstallPage struct {
	*widgets.QFrame
}

func NewInstallPage(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *InstallPage {
	frame := widgets.NewQFrame(parent, fo)

	page := &InstallPage{frame}
	page.init()

	return page
}

func (i *InstallPage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(i)

	welcomeLabel := widgets.NewQLabel2("Install", i, 0)

	vboxLayout.AddWidget(welcomeLabel, 0, core.Qt__AlignCenter)

	i.SetLayout(vboxLayout)
}
