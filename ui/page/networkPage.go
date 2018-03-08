package page

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
)

type NetworkPage struct {
	Frame *widgets.QFrame
}

func NewNetworkPage(parent widgets.QWidget_ITF,fo core.Qt__WindowType) *NetworkPage {
	frame := widgets.NewQFrame(parent,fo)

	vboxLayout := widgets.NewQVBoxLayout2(frame)

	welcomeLabel := widgets.NewQLabel2("Network",frame,0)

	vboxLayout.AddWidget(welcomeLabel,0,core.Qt__AlignCenter)

	frame.SetLayout(vboxLayout)

	return &NetworkPage{frame}
}