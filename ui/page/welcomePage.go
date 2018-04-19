package page

import (
	widgets2 "github.com/firerainos/firerain-installer/ui/widgets"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type WelcomePage struct {
	*widgets.QFrame
}

func NewWelcomePage(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *WelcomePage {
	page := &WelcomePage{widgets.NewQFrame(parent, fo)}

	page.init()

	return page
}

func (w *WelcomePage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(w)
	vboxLayout.SetContentsMargins(0, 0, 0, 0)

	hellWidget := widgets2.NewHelloWidget(w, 0)

	vboxLayout.AddWidget(hellWidget, 0, core.Qt__AlignCenter)
	vboxLayout.AddStretch(1)

	w.SetLayout(vboxLayout)

}
