package widgets

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	"github.com/firerainos/firerain-installer/core/networkmanager"
)

type WifiList struct {
	*widgets.QScrollArea

	vboxLayout *widgets.QVBoxLayout

	nm *networkmanager.NetworkManager
}

func NewWifiList(parent widgets.QWidget_ITF) *WifiList {
	scrollArea := widgets.NewQScrollArea(parent)

	wifiList := &WifiList{QScrollArea: scrollArea,nm:networkmanager.NewNetworkManager()}
	wifiList.init()

	wifiList.StartTimer(5000,core.Qt__PreciseTimer)
	wifiList.ConnectTimerEvent(func(event *core.QTimerEvent) {
		wifiList.scanWifi()
	})

	return wifiList
}

func (w *WifiList) init() {
	w.SetMinimumSize2(480,500)
	w.SetContentsMargins(10,10,10,10)
	w.SetWidgetResizable(true)
	w.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	w.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)

	frame := widgets.NewQFrame(w,0)

	w.vboxLayout = widgets.NewQVBoxLayout2(frame)

	frame.SetLayout(w.vboxLayout)

	w.SetWidget(frame)
}

func (w *WifiList) scanWifi() {
	w.nm.WifiScan()

	wifiList := *w.nm.WifiList()

	for {
		item := w.vboxLayout.TakeAt(0)
		if item.Pointer() == nil {
			break
		}

		if item.Widget().Pointer() != nil {
			item.Widget().DestroyQWidget()
		}

		if item.Layout().Pointer() != nil {
			item.Layout().DestroyQLayoutItem()
		}
	}

	for _,wifiInfo := range wifiList {
		item := NewWifiListItem(wifiInfo,w,0)
		w.vboxLayout.AddWidget(item,0,core.Qt__AlignHCenter)
	}

	w.vboxLayout.AddStretch(1)

}