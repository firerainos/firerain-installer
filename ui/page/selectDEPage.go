package page

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"os"
	"fmt"
)

type SelectDEPage struct {
	*widgets.QFrame

	deListWidget *widgets.QListWidget
}

func NewSelectDEPage(parent widgets.QWidget_ITF,fo core.Qt__WindowType) *SelectDEPage {
	frame := widgets.NewQFrame(parent,fo)

	page := &SelectDEPage{QFrame:frame}
	page.init()
	page.initConnect()

	return page
}

func (s *SelectDEPage) init(){
	vboxLayout := widgets.NewQVBoxLayout2(s)

	welcomeLabel := widgets.NewQLabel2("选择要安装的桌面环境",s,0)
	s.deListWidget = widgets.NewQListWidget(s)

	s.deListWidget.SetMinimumSize2(530,530)
	s.deListWidget.SetViewMode(widgets.QListView__IconMode)
	s.deListWidget.SetFlow(widgets.QListView__LeftToRight)
	s.deListWidget.SetMovement(widgets.QListView__Static)
	s.deListWidget.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	s.deListWidget.SetIconSize(core.NewQSize2(180,180))
	s.deListWidget.SetSpacing(40)

	path,_ := os.Getwd()

	widgets.NewQListWidgetItem3(gui.NewQIcon5(path+ "/resources/de-logo/kde.png"),"KDE",s.deListWidget,0).SetSizeHint(core.NewQSize2(200,200))
	widgets.NewQListWidgetItem3(gui.NewQIcon5(path+ "/resources/de-logo/dde.png"),"DDE",s.deListWidget,0).SetSizeHint(core.NewQSize2(200,200))
	widgets.NewQListWidgetItem3(gui.NewQIcon5(path+ "/resources/de-logo/cinnamon.png"),"Cinnamon",s.deListWidget,0).SetSizeHint(core.NewQSize2(200,200))
	widgets.NewQListWidgetItem3(gui.NewQIcon5(path+ "/resources/de-logo/gnome.png"),"GNOME",s.deListWidget,0).SetSizeHint(core.NewQSize2(200,200))

	vboxLayout.AddWidget(welcomeLabel,0,core.Qt__AlignCenter)
	vboxLayout.AddWidget(s.deListWidget,0,core.Qt__AlignCenter)

	s.SetLayout(vboxLayout)
}

func (s *SelectDEPage) initConnect (){
	s.deListWidget.ConnectCurrentTextChanged(func(currentText string) {
		fmt.Println(currentText)
	})
}