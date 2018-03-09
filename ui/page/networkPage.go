package page

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	widgets2 "github.com/firerainos/firerain-installer/ui/widgets"
)

type NetworkPage struct {
	*widgets.QFrame
}

func NewNetworkPage(parent widgets.QWidget_ITF,fo core.Qt__WindowType) *NetworkPage {
	page := &NetworkPage{widgets.NewQFrame(parent,fo)}
	page.init()

	return page
}

func (n *NetworkPage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(n)

	networkLabel := widgets.NewQLabel2("Network",n,0)

	wifiList := widgets2.NewWifiList(n)

	vboxLayout.AddWidget(networkLabel,0,core.Qt__AlignCenter)
	vboxLayout.AddWidget(wifiList,0,core.Qt__AlignCenter)

	n.SetLayout(vboxLayout)
}
