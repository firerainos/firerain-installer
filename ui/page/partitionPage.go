package page

import (
	widgets2 "github.com/firerainos/firerain-installer/ui/widgets"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type PartitionPage struct {
	*widgets.QFrame

	tipLabel *widgets.QLabel

	partitionList *widgets2.PartitionList
}

func NewPartitionPage(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *PartitionPage {
	frame := widgets.NewQFrame(parent, fo)

	page := &PartitionPage{QFrame: frame}
	page.init()
	page.initConnect()

	return page
}

func (p *PartitionPage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(p)

	logoLabel := widgets.NewQLabel(p, 0)
	welcomeLabel := widgets.NewQLabel2("FireRainOS", p, 0)
	p.tipLabel = widgets.NewQLabel2("请选择安装磁盘", p, 0)
	p.partitionList = widgets2.NewPartitionList(p)

	logoLabel.SetPixmap(gui.NewQPixmap5("/home/linux/go/src/github.com/firerainos/firerain-installer/resources/logo.png", "", 0).Scaled2(200, 200, core.Qt__KeepAspectRatioByExpanding, 0))
	logoLabel.SetFixedSize2(200, 200)

	vboxLayout.AddWidget(logoLabel, 0, core.Qt__AlignCenter)
	vboxLayout.AddWidget(welcomeLabel, 0, core.Qt__AlignCenter)
	vboxLayout.AddWidget(p.tipLabel, 0, core.Qt__AlignCenter)
	vboxLayout.AddWidget(p.partitionList, 0, core.Qt__AlignCenter)

	p.SetLayout(vboxLayout)
}

func (p *PartitionPage) initConnect() {
	p.partitionList.ConnectPartitionItemChange(func() {
		p.tipLabel.SetText("FireRainOS 将安装在磁盘\"" + p.partitionList.CurrendItem().DevPath() + "\"上")
	})
}
