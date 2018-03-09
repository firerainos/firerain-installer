package page

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
)

type SelectDMPage struct {
	*widgets.QFrame
}

func NewSelectDMPage(parent widgets.QWidget_ITF,fo core.Qt__WindowType) *SelectDMPage {
	frame := widgets.NewQFrame(parent,fo)

	page := &SelectDMPage{frame}
	page.init()

	return page
}

func (s *SelectDMPage) init(){
	vboxLayout := widgets.NewQVBoxLayout2(s)

	welcomeLabel := widgets.NewQLabel2("SelectDM",s,0)

	vboxLayout.AddWidget(welcomeLabel,0,core.Qt__AlignCenter)

	s.SetLayout(vboxLayout)
}
