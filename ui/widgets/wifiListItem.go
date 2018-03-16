package widgets

import (
	"github.com/firerainos/firerain-installer/core/networkmanager"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"strconv"
)

type WifiListItem struct {
	widgets.QWidget

	wifInfo networkmanager.WifiInfo

	btn *widgets.QPushButton

	signalLabel, ssidLabel, securityLabel *widgets.QLabel

	_ func() `constructor:"init()"`

	_ func(ssid, security string, inUse bool) `signal:"wifiListItemClicked"`
}

//func NewWifiListItem2(wifiInfo networkmanager.WifiInfo,parent widgets.QWidget_ITF, fo core.Qt__WindowType) *WifiListItem {
//	widget := widgets.NewQWidget(parent, fo)
//
//	wifiListItem := &WifiListItem{QWidget: *widget,wifInfo:wifiInfo}
//	wifiListItem.init()
//
//	return wifiListItem
//}

func (w *WifiListItem) init() {
	w.SetMinimumSize2(460, 50)

	hboxLayout := widgets.NewQHBoxLayout()
	hboxLayout.SetContentsMargins(10, 10, 10, 10)
	hboxLayout.SetSpacing(50)

	w.signalLabel = widgets.NewQLabel(w, 0)
	w.ssidLabel = widgets.NewQLabel(w, 0)
	w.securityLabel = widgets.NewQLabel(w, 0)

	w.btn = widgets.NewQPushButton(w)

	hboxLayout.AddWidget(w.signalLabel, 0, core.Qt__AlignLeft)
	hboxLayout.AddWidget(w.ssidLabel, 0, core.Qt__AlignLeft)
	hboxLayout.AddStretch(10)
	hboxLayout.AddWidget(w.securityLabel, 0, core.Qt__AlignRight)
	hboxLayout.AddWidget(w.btn, 0, core.Qt__AlignRight)

	w.SetLayout(hboxLayout)

	w.btn.ConnectClicked(func(checked bool) {
		w.btn.SetText("连接中")
		w.btn.SetEnabled(false)
		w.WifiListItemClicked(w.wifInfo.Ssid, w.wifInfo.Security, w.wifInfo.InUse)
	})
}

func (w *WifiListItem) SetWifiInfo(wifiInfo networkmanager.WifiInfo) {
	w.wifInfo = wifiInfo

	w.signalLabel.SetText(strconv.Itoa(w.wifInfo.Signal))
	w.ssidLabel.SetText(w.wifInfo.Ssid)
	w.securityLabel.SetText(w.wifInfo.Security)

	if w.wifInfo.InUse {
		w.btn.SetText("已连接")
		w.btn.SetEnabled(false)
	} else {
		w.btn.SetText("连接")
		w.btn.SetEnabled(true)
	}
}
