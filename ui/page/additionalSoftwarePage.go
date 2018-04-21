package page

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/firerainos/firerain-installer/api"
)

type AdditionalSoftwarePage struct {
	*widgets.QFrame

	itemList,installList *widgets.QListWidget
	packageList *widgets.QListView

	account *api.Account
}

func NewAdditionalSoftwarePage(account *api.Account,parent widgets.QWidget_ITF, fo core.Qt__WindowType) *AdditionalSoftwarePage {
	frame := widgets.NewQFrame(parent, fo)

	page := &AdditionalSoftwarePage{QFrame:frame,account:account}
	page.init()
	page.initConnect()

	return page
}

func (a *AdditionalSoftwarePage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(a)

	welcomeLabel := widgets.NewQLabel2("选择软件包", a, 0)

	hboxLayout := widgets.NewQHBoxLayout2(a)
	hboxLayout.SetSpacing(0)

	a.itemList = widgets.NewQListWidget(a)
	a.packageList = widgets.NewQListView(a)
	a.installList = widgets.NewQListWidget(a)

	a.itemList.SetFixedSize2(150,500)
	a.packageList.SetFixedSize2(430,500)
	a.installList.SetFixedSize2(250,500)

	buttonLayout := widgets.NewQVBoxLayout2(a)
	buttonLayout.SetSpacing(16)

	leftButton := widgets.NewQPushButton2("<",a)
	rightButton := widgets.NewQPushButton2(">",a)

	leftButton.SetFixedSize2(32,32)
	rightButton.SetFixedSize2(32,32)

	buttonLayout.AddStretch(1)
	buttonLayout.AddWidget(leftButton,0,core.Qt__AlignCenter)
	buttonLayout.AddWidget(rightButton,0,core.Qt__AlignCenter)
	buttonLayout.AddStretch(1)

	hboxLayout.AddWidget(a.itemList,0,core.Qt__AlignLeft)
	hboxLayout.AddWidget(a.packageList,0,core.Qt__AlignLeft)
	hboxLayout.AddLayout(buttonLayout,0)
	hboxLayout.AddWidget(a.installList,0,core.Qt__AlignRight)

	vboxLayout.AddStretch(1)
	vboxLayout.AddWidget(welcomeLabel, 0, core.Qt__AlignCenter)
	vboxLayout.AddSpacing(50)
	vboxLayout.AddLayout(hboxLayout, 1)
	vboxLayout.AddStretch(1)

	a.SetLayout(vboxLayout)
}

func (a *AdditionalSoftwarePage) initConnect() {
}

func (a *AdditionalSoftwarePage) LoadData() {
}