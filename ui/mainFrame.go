package ui

import (
	"github.com/firerainos/firerain-installer/ui/page"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type MainFrame struct {
	*widgets.QFrame

	welcomePage            *page.WelcomePage
	networkPage            *page.NetworkPage
	partitionPage          *page.PartitionPage
	selectDEPage           *page.SelectDEPage
	additionalSoftwarePage *page.AdditionalSoftwarePage
	installPage            *page.InstallPage

	backButton, nextButton *widgets.QPushButton

	stackLayout *widgets.QStackedLayout
}

func NewMainFrame(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *MainFrame {
	frame := &MainFrame{QFrame: widgets.NewQFrame(parent, fo)}

	frame.init()
	frame.initConnect()

	return frame
}

func (m *MainFrame) init() {
	vboxLayout := widgets.NewQVBoxLayout()
	vboxLayout.SetContentsMargins(0, 0, 0, 0)

	m.stackLayout = widgets.NewQStackedLayout()

	m.welcomePage = page.NewWelcomePage(m, 0)
	m.networkPage = page.NewNetworkPage(m, 0)
	m.partitionPage = page.NewPartitionPage(m, 0)
	m.selectDEPage = page.NewSelectDEPage(m, 0)
	m.additionalSoftwarePage = page.NewAdditionalSoftwarePage(m, 0)
	m.installPage = page.NewInstallPage(m, 0)

	m.backButton = widgets.NewQPushButton2("back", m)
	m.nextButton = widgets.NewQPushButton2("next", m)

	m.backButton.SetVisible(false)

	hboxLayout := widgets.NewQHBoxLayout()

	hboxLayout.AddStretch(1)
	hboxLayout.AddWidget(m.backButton, 0, core.Qt__AlignHCenter)
	hboxLayout.AddWidget(m.nextButton, 0, core.Qt__AlignHCenter)
	hboxLayout.AddStretch(1)

	m.stackLayout.AddWidget(m.welcomePage)
	m.stackLayout.AddWidget(m.networkPage)
	m.stackLayout.AddWidget(m.partitionPage)
	m.stackLayout.AddWidget(m.selectDEPage)
	m.stackLayout.AddWidget(m.additionalSoftwarePage)
	m.stackLayout.AddWidget(m.installPage)

	vboxLayout.AddLayout(m.stackLayout, 1)
	vboxLayout.AddLayout(hboxLayout,1)
	vboxLayout.AddSpacing(20)

	m.SetLayout(vboxLayout)
}

func (m *MainFrame) initConnect() {
	m.stackLayout.ConnectCurrentChanged(func(index int) {
		if index == 0 {
			m.backButton.SetVisible(false)
		} else if index == m.stackLayout.Count()-1 {
			m.nextButton.SetVisible(false)
			m.backButton.SetVisible(false)
		} else if index > 0 {
			m.backButton.SetVisible(true)
		}
	})

	m.backButton.ConnectClicked(func(checked bool) {
		m.stackLayout.SetCurrentIndex(m.stackLayout.CurrentIndex() - 1)
	})

	m.nextButton.ConnectClicked(func(checked bool) {
		m.stackLayout.SetCurrentIndex(m.stackLayout.CurrentIndex() + 1)
	})
}
