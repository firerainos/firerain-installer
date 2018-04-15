package page

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type AdditionalSoftwarePage struct {
	*widgets.QFrame
}

func NewAdditionalSoftwarePage(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *AdditionalSoftwarePage {
	frame := widgets.NewQFrame(parent, fo)

	page := &AdditionalSoftwarePage{frame}
	page.init()

	return page
}

func (a *AdditionalSoftwarePage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(a)

	welcomeLabel := widgets.NewQLabel2("AdditionalSoftware", a, 0)

	vboxLayout.AddWidget(welcomeLabel, 0, core.Qt__AlignCenter)

	a.SetLayout(vboxLayout)
}
