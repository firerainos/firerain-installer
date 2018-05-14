package page

import (
	"github.com/firerainos/firerain-installer/config"
	_ "github.com/firerainos/firerain-installer/resources"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/firerainos/firerain-installer/styles"
)

type SelectDEPage struct {
	*widgets.QFrame

	deListWidget *widgets.QListWidget

	deName []string
}

func NewSelectDEPage(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *SelectDEPage {
	frame := widgets.NewQFrame(parent, fo)

	page := &SelectDEPage{QFrame: frame}
	page.init()
	page.initConnect()

	return page
}

func (s *SelectDEPage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(s)

	welcomeLabel := widgets.NewQLabel2("选择要安装的桌面环境", s, 0)
	s.deListWidget = widgets.NewQListWidget(s)

	s.deListWidget.SetMinimumSize2(530, 530)
	s.deListWidget.SetViewMode(widgets.QListView__IconMode)
	s.deListWidget.SetFlow(widgets.QListView__LeftToRight)
	s.deListWidget.SetMovement(widgets.QListView__Static)
	s.deListWidget.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	s.deListWidget.SetIconSize(core.NewQSize2(150, 150))
	s.deListWidget.SetSpacing(40)
	s.deListWidget.SetStyleSheet(styles.DEList)

	widgets.NewQListWidgetItem3(gui.NewQIcon5(":/resources/de-logo/kde.png"), "KDE", s.deListWidget, 0).SetSizeHint(core.NewQSize2(200, 200))
	widgets.NewQListWidgetItem3(gui.NewQIcon5(":/resources/de-logo/dde.png"), "DDE", s.deListWidget, 0).SetSizeHint(core.NewQSize2(200, 200))
	widgets.NewQListWidgetItem3(gui.NewQIcon5(":/resources/de-logo/cinnamon.png"), "Cinnamon", s.deListWidget, 0).SetSizeHint(core.NewQSize2(200, 200))
	widgets.NewQListWidgetItem3(gui.NewQIcon5(":/resources/de-logo/gnome.png"), "GNOME", s.deListWidget, 0).SetSizeHint(core.NewQSize2(200, 200))

	s.deListWidget.SetCurrentRow(0)

	vboxLayout.AddWidget(welcomeLabel, 0, core.Qt__AlignCenter)
	vboxLayout.AddWidget(s.deListWidget, 0, core.Qt__AlignCenter)

	s.SetLayout(vboxLayout)
}

func (s *SelectDEPage) initConnect() {
	s.deListWidget.ConnectCurrentTextChanged(func(currentText string) {
		for _, pkg := range s.deName {
			config.Conf.RemovePackage(pkg)
		}
		switch currentText {
		case "KDE":
			config.Conf.DEApplication = "kde-applications"
			s.deName = []string{"plasma-meta", "sddm", "dolphin", "konsole", "kate"}
		case "DDE":
			config.Conf.DEApplication = "deepin-extra"
			s.deName = []string{"deepin", "lightdm"}
		case "Cinnamon":
			config.Conf.DEApplication = "gnome-extra"
			s.deName = []string{"cinnamon", "cinnamon-translations", "lightdm", "lightdm-gtk-greeter"}
		case "GNOME":
			config.Conf.DEApplication = "gnome-extra"
			s.deName = []string{"gnome", "lightdm", "lightdm-gtk-greeter"}
		}
		for _, pkg := range s.deName {
			config.Conf.AddPackage(pkg)
		}
	})

	s.deListWidget.CurrentTextChanged("KDE")
}
