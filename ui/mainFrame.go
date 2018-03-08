package ui

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	"github.com/firerainos/firerain-installer/ui/page"
)

type MainFrame struct {
	frame *widgets.QFrame
}

func NewMainFrame(parent widgets.QWidget_ITF,fo core.Qt__WindowType) *MainFrame {
	frame := widgets.NewQFrame(parent,fo)

	vboxLayout := widgets.NewQVBoxLayout()

	stackLayout := widgets.NewQStackedLayout()

	welcomePage := page.NewWelcomePage(frame,0)
	networkPage := page.NewNetworkPage(frame,0)
	partitionPage := page.NewPartitionPage(frame, 0)
	selectDMPage := page.NewSelectDMPage(frame, 0)
	additionalSoftwarePage := page.NewAdditionalSoftwarePage(frame, 0)
	installPage := page.NewInstallPage(frame,0)

	backButton := widgets.NewQPushButton2("back",frame)
	nextButton := widgets.NewQPushButton2("next",frame)

	backButton.SetVisible(false)

	stackLayout.AddWidget(welcomePage.Frame)
	stackLayout.AddWidget(networkPage.Frame)
	stackLayout.AddWidget(partitionPage.Frame)
	stackLayout.AddWidget(selectDMPage.Frame)
	stackLayout.AddWidget(additionalSoftwarePage.Frame)
	stackLayout.AddWidget(installPage.Frame)

	vboxLayout.AddWidget(backButton,0,core.Qt__AlignLeft)
	vboxLayout.AddLayout(stackLayout,1)
	vboxLayout.AddWidget(nextButton,0,core.Qt__AlignHCenter)

	frame.SetLayout(vboxLayout)

	stackLayout.ConnectCurrentChanged(func(index int) {
		if index == 0 {
			backButton.SetVisible(false)
		} else if index == stackLayout.Count() - 1 {
			nextButton.SetVisible(false)
			backButton.SetVisible(false)
		} else if index > 0 {
			backButton.SetVisible(true)
		}
	})

	backButton.ConnectClicked(func(checked bool) {
		stackLayout.SetCurrentIndex(stackLayout.CurrentIndex()-1)
	})

	nextButton.ConnectClicked(func(checked bool) {
		stackLayout.SetCurrentIndex(stackLayout.CurrentIndex()+1)
	})

	return &MainFrame{frame}
}