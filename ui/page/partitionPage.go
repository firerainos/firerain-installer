package page

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	widgets2 "github.com/firerainos/firerain-installer/ui/widgets"
)

type PartitionPage struct {
	*widgets.QFrame
}

func NewPartitionPage(parent widgets.QWidget_ITF,fo core.Qt__WindowType) *PartitionPage {
	frame := widgets.NewQFrame(parent,fo)

	page := &PartitionPage{frame}
	page.init()

	return page
}

func (p *PartitionPage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(p)

	welcomeLabel := widgets.NewQLabel2("Partition",p,0)
	partitionList := widgets2.NewPartitionList(p)

	vboxLayout.AddWidget(welcomeLabel,0,core.Qt__AlignCenter)
	vboxLayout.AddWidget(partitionList,0,core.Qt__AlignCenter)

	p.SetLayout(vboxLayout)
}