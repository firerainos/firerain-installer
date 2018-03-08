package page

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
)

type InstallPage struct {
	Frame *widgets.QFrame
}

func NewInstallPage(parent widgets.QWidget_ITF,fo core.Qt__WindowType) *InstallPage {
	frame := widgets.NewQFrame(parent,fo)

	vboxLayout := widgets.NewQVBoxLayout2(frame)

	welcomeLabel := widgets.NewQLabel2("Install",frame,0)

	vboxLayout.AddWidget(welcomeLabel,0,core.Qt__AlignCenter)

	frame.SetLayout(vboxLayout)

	return &InstallPage{frame}
}