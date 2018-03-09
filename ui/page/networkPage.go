package page

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	"github.com/firerainos/firerain-installer/core/networkmanager"
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

	wifiListWidget := widgets.NewQListWidget(n)

	nm := networkmanager.NewNetworkManager()
	nm.WifiScan()

	list := nm.WifiList()

	for _,wifi := range *list  {
		wifiListWidget.AddItem(wifi.Ssid)
	}

	vboxLayout.AddWidget(networkLabel,0,core.Qt__AlignCenter)

	n.SetLayout(vboxLayout)
}
