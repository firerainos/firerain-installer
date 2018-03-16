package page

import (
	"github.com/firerainos/firerain-installer/core/networkmanager"
	widgets2 "github.com/firerainos/firerain-installer/ui/widgets"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type NetworkPage struct {
	*widgets.QFrame

	wifiList *widgets.QListWidget

	nm *networkmanager.NetworkManager

	stopScan bool
}

func NewNetworkPage(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *NetworkPage {
	page := &NetworkPage{QFrame: widgets.NewQFrame(parent, fo), nm: networkmanager.NewNetworkManager()}
	page.init()

	page.StartTimer(2000, core.Qt__PreciseTimer)
	page.ConnectTimerEvent(func(event *core.QTimerEvent) {
		if !page.stopScan {
			page.scanWifi()
		}
	})

	return page
}

func (n *NetworkPage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(n)

	networkLabel := widgets.NewQLabel2("Network", n, 0)
	n.wifiList = widgets.NewQListWidget(n)

	n.wifiList.SetMinimumSize2(480, 500)

	vboxLayout.AddWidget(networkLabel, 0, core.Qt__AlignCenter)
	vboxLayout.AddWidget(n.wifiList, 0, core.Qt__AlignCenter)

	n.SetLayout(vboxLayout)
}

func (n *NetworkPage) scanWifi() {
	n.nm.WifiScan()

	wifiList := *n.nm.WifiList()

	n.wifiList.Clear()

	for _, wifiInfo := range wifiList {
		item := widgets.NewQListWidgetItem(n.wifiList, 0)
		item.SetSizeHint(core.NewQSize2(460, 50))
		listItem := widgets2.NewWifiListItem(n, 0)
		listItem.SetWifiInfo(wifiInfo)
		n.wifiList.SetItemWidget(item, listItem)

		listItem.ConnectWifiListItemClicked(n.onWifiListItemClicked)
	}
}

func (n *NetworkPage) onWifiListItemClicked(ssid string, security string, inUse bool) {
	n.stopScan = true
	if security == "" {
		n.nm.ConnectWifi(ssid, "")
	} else {
		dialog := widgets.NewQInputDialog(n, 0)
		password := dialog.GetText(n, "提示", "请输入"+ssid+"的密码", widgets.QLineEdit__Password, "", true, 0, 0)
		if password == "" {
			n.scanWifi()
			n.stopScan = false
			return
		}
		n.nm.ConnectWifi(ssid, password)
	}
	n.stopScan = false
	n.scanWifi()
}
