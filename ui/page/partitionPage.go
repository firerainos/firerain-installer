package page

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
)

type PartitionPage struct {
	Frame *widgets.QFrame
}

func NewPartitionPage(parent widgets.QWidget_ITF,fo core.Qt__WindowType) *PartitionPage {
	frame := widgets.NewQFrame(parent,fo)

	vboxLayout := widgets.NewQVBoxLayout2(frame)

	welcomeLabel := widgets.NewQLabel2("Partition",frame,0)

	vboxLayout.AddWidget(welcomeLabel,0,core.Qt__AlignCenter)

	frame.SetLayout(vboxLayout)

	return &PartitionPage{frame}
}