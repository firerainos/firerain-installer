package page

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
)

type WelcomePage struct {
	Frame *widgets.QFrame
}

func NewWelcomePage(parent widgets.QWidget_ITF,fo core.Qt__WindowType) *WelcomePage {
	frame := widgets.NewQFrame(parent,fo)

	vboxLayout := widgets.NewQVBoxLayout2(frame)

	welcomeLabel := widgets.NewQLabel2("welcome",frame,0)

	vboxLayout.AddWidget(welcomeLabel,0,core.Qt__AlignCenter)

	frame.SetLayout(vboxLayout)

	return &WelcomePage{frame}
}