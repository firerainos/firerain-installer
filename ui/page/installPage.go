package page

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type InstallPage struct {
	*widgets.QFrame

	tipsLabel *widgets.QLabel

	messageText *widgets.QTextBrowser
}

func NewInstallPage(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *InstallPage {
	frame := widgets.NewQFrame(parent, fo)

	page := &InstallPage{QFrame: frame}
	page.init()

	return page
}

func (i *InstallPage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(i)

	welcomeLabel := widgets.NewQLabel2("FireRainOS安装", i, 0)

	i.tipsLabel = widgets.NewQLabel(i, 0)
	i.tipsLabel.SetFixedWidth(200)
	i.tipsLabel.SetAlignment(core.Qt__AlignCenter)

	i.messageText = widgets.NewQTextBrowser(i)
	i.messageText.SetFixedSize2(800, 500)

	vboxLayout.AddSpacing(80)
	vboxLayout.AddWidget(welcomeLabel, 0, core.Qt__AlignCenter)
	vboxLayout.AddStretch(1)
	vboxLayout.AddWidget(i.tipsLabel, 0, core.Qt__AlignHCenter)
	vboxLayout.AddStretch(1)
	vboxLayout.AddWidget(i.messageText, 0, core.Qt__AlignCenter)
	vboxLayout.AddSpacing(50)

	i.SetLayout(vboxLayout)
}

func (i *InstallPage) SetTips(tips string) {
	i.tipsLabel.SetText(tips)
}

func (i *InstallPage) SetMessageVisible() {
	i.messageText.SetVisible(false)
}

func (i *InstallPage) AddMessage(message string) {
	i.messageText.Append(message)
}
