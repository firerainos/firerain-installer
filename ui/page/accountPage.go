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

	username := widgets2.NewLineEdit(pwd+"/resources/username.svg", page)
	password := widgets2.NewLineEdit(pwd+"/resources/password.svg", page)
	password.SetEchoMode(widgets.QLineEdit__Password)

	vboxLayout.AddStretch(1)
	vboxLayout.AddWidget(loginLabel, 0, core.Qt__AlignCenter)
	vboxLayout.AddSpacing(100)
	vboxLayout.AddWidget(username, 0, core.Qt__AlignHCenter)
	vboxLayout.AddWidget(password, 0, core.Qt__AlignHCenter)
	vboxLayout.AddStretch(1)

	page.SetLayout(vboxLayout)

	username.ConnectLeaveEvent(func(event *core.QEvent) {
		config.Conf.SetUsername(username.Text())
	})

	password.ConnectLeaveEvent(func(event *core.QEvent) {
		config.Conf.SetPassword(password.Text())
	})
}
