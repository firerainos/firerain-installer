package page

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
)

type WelcomePage struct {
	*widgets.QFrame
}

func NewWelcomePage(parent widgets.QWidget_ITF,fo core.Qt__WindowType) *WelcomePage {
	page := &WelcomePage{widgets.NewQFrame(parent,fo)}

	page.init()

	return page
}

func (w *WelcomePage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(w)

	welcomeLabel := widgets.NewQLabel2("welcome",w,0)


	vboxLayout.AddWidget(welcomeLabel,0,core.Qt__AlignCenter)

	w.SetLayout(vboxLayout)

}