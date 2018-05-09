package page

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type EndPage struct {
	*widgets.QFrame

	tipsLabel *widgets.QLabel
}

func NewEndPage(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *EndPage {
	frame := widgets.NewQFrame(parent, fo)

	page := &EndPage{QFrame: frame}
	page.init()

	return page
}

func (page *EndPage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(page)

	page.tipsLabel = widgets.NewQLabel2("安装成功", page, 0)

	vboxLayout.AddStretch(1)
	vboxLayout.AddWidget(page.tipsLabel, 0, core.Qt__AlignHCenter)
	vboxLayout.AddStretch(1)

	page.SetLayout(vboxLayout)
}

func (page *EndPage) SetTips(tips string) {
	page.tipsLabel.SetText(tips)
}
