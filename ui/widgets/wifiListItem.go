package widgets

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	"github.com/firerainos/firerain-installer/core/networkmanager"
	"strconv"
)

type WifiListItem struct {
	*widgets.QWidget

	wifInfo networkmanager.WifiInfo

	btn *widgets.QPushButton
}

func NewWifiListItem(wifiInfo networkmanager.WifiInfo,parent widgets.QWidget_ITF, fo core.Qt__WindowType) *WifiListItem {
	widget := widgets.NewQWidget(parent, fo)

	wifiListItem := &WifiListItem{QWidget: widget,wifInfo:wifiInfo}
	wifiListItem.init()

	return wifiListItem
}

func (w *WifiListItem) init() {
	w.SetMinimumSize2(460,50)

	hboxLayout := widgets.NewQHBoxLayout()
	hboxLayout.SetContentsMargins(10,10,10,10)
	hboxLayout.SetSpacing(50)

	signalLabel := widgets.NewQLabel2(strconv.Itoa(w.wifInfo.Signal),w,0)
	ssidLabel := widgets.NewQLabel2(w.wifInfo.Ssid,w,0)
	securityLabel := widgets.NewQLabel2(w.wifInfo.Security,w,0)

	if w.wifInfo.InUse {
		w.btn = widgets.NewQPushButton2("断开",w)
	} else {
		w.btn = widgets.NewQPushButton2("连接",w)
	}

	hboxLayout.AddWidget(signalLabel,0,core.Qt__AlignLeft)
	hboxLayout.AddWidget(ssidLabel,0,core.Qt__AlignLeft)
	hboxLayout.AddStretch(10)
	hboxLayout.AddWidget(securityLabel,0,core.Qt__AlignRight)
	hboxLayout.AddWidget(w.btn,0,core.Qt__AlignRight)

	w.SetLayout(hboxLayout)
}