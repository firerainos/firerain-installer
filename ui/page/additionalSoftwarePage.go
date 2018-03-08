package page

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
)

type AdditionalSoftwarePage struct {
	Frame *widgets.QFrame
}

func NewAdditionalSoftwarePage(parent widgets.QWidget_ITF,fo core.Qt__WindowType) *AdditionalSoftwarePage {
	frame := widgets.NewQFrame(parent,fo)

	vboxLayout := widgets.NewQVBoxLayout2(frame)

	welcomeLabel := widgets.NewQLabel2("AdditionalSoftware",frame,0)

	vboxLayout.AddWidget(welcomeLabel,0,core.Qt__AlignCenter)

	frame.SetLayout(vboxLayout)

	return &AdditionalSoftwarePage{frame}
}