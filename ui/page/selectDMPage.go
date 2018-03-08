package page

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
)

type SelectDMPage struct {
	Frame *widgets.QFrame
}

func NewSelectDMPage(parent widgets.QWidget_ITF,fo core.Qt__WindowType) *SelectDMPage {
	frame := widgets.NewQFrame(parent,fo)

	vboxLayout := widgets.NewQVBoxLayout2(frame)

	welcomeLabel := widgets.NewQLabel2("SelectDM",frame,0)

	vboxLayout.AddWidget(welcomeLabel,0,core.Qt__AlignCenter)

	frame.SetLayout(vboxLayout)

	return &SelectDMPage{frame}
}
