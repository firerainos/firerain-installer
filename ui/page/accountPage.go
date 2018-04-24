package page

import (
	"github.com/firerainos/firerain-installer/config"
	widgets2 "github.com/firerainos/firerain-installer/ui/widgets"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os"
)

type AccountPage struct {
	*widgets.QFrame

	tipsLabel *widgets.QLabel

	username,password *widgets2.LineEdit
}

func NewAccountPage(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *AccountPage {
	frame := widgets.NewQFrame(parent, fo)

	accountPage := &AccountPage{QFrame: frame}
	accountPage.init()

	return accountPage
}

func (page *AccountPage) init() {
	pwd, _ := os.Getwd()

	vboxLayout := widgets.NewQVBoxLayout2(page)

	loginLabel := widgets.NewQLabel2("登陆", page, 0)

	page.tipsLabel = widgets.NewQLabel(page,0)
	page.tipsLabel.SetAlignment(core.Qt__AlignCenter)

	page.username = widgets2.NewLineEdit(pwd+"/resources/username.svg", page)
	page.password = widgets2.NewLineEdit(pwd+"/resources/password.svg", page)
	page.password.SetEchoMode(widgets.QLineEdit__Password)

	vboxLayout.AddStretch(1)
	vboxLayout.AddWidget(loginLabel, 0, core.Qt__AlignCenter)
	vboxLayout.AddWidget(page.tipsLabel, 0, core.Qt__AlignCenter)
	vboxLayout.AddSpacing(100)
	vboxLayout.AddWidget(page.username, 0, core.Qt__AlignHCenter)
	vboxLayout.AddWidget(page.password, 0, core.Qt__AlignHCenter)
	vboxLayout.AddStretch(1)

	page.SetLayout(vboxLayout)

	page.username.ConnectLeaveEvent(func(event *core.QEvent) {
		config.Conf.SetUsername(page.username.Text())
	})

	page.password.ConnectLeaveEvent(func(event *core.QEvent) {
		config.Conf.SetPassword(page.password.Text())
	})
}

func  (page *AccountPage) SetTips(tips string) {
	page.tipsLabel.SetText(tips)
}

func  (page *AccountPage) SetEnableLogin(enable bool) {
	page.username.SetEnabled(enable)
	page.password.SetEnabled(enable)
}

