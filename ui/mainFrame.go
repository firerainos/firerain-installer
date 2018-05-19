package ui

import (
	"github.com/firerainos/firerain-installer/api"
	"github.com/firerainos/firerain-installer/config"
	"github.com/firerainos/firerain-installer/core/installer"
	"github.com/firerainos/firerain-installer/styles"
	"github.com/firerainos/firerain-installer/ui/page"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
	"os/exec"
	"strings"
)

type MainFrame struct {
	*widgets.QFrame

	welcomePage            *page.WelcomePage
	networkPage            *page.NetworkPage
	accountPage            *page.AccountPage
	partitionPage          *page.PartitionPage
	selectDEPage           *page.SelectDEPage
	additionalSoftwarePage *page.AdditionalSoftwarePage
	installPage            *page.InstallPage
	endPage                *page.EndPage

	backButton, nextButton *widgets.QPushButton

	stackLayout *widgets.QStackedLayout

	account *api.Account

	mediaPlayer *multimedia.QMediaPlayer
}

func NewMainFrame(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *MainFrame {
	frame := &MainFrame{QFrame: widgets.NewQFrame(parent, fo), account: api.NewAccount()}

	frame.init()
	frame.initPlayer()
	frame.initConnect()

	return frame
}

func (m *MainFrame) init() {
	vboxLayout := widgets.NewQVBoxLayout()
	vboxLayout.SetContentsMargins(0, 0, 0, 0)

	m.stackLayout = widgets.NewQStackedLayout()

	m.welcomePage = page.NewWelcomePage(m, 0)
	m.networkPage = page.NewNetworkPage(m, 0)
	m.accountPage = page.NewAccountPage(m, 0)
	m.partitionPage = page.NewPartitionPage(m, 0)
	m.selectDEPage = page.NewSelectDEPage(m, 0)
	m.additionalSoftwarePage = page.NewAdditionalSoftwarePage(m.account, m, 0)
	m.installPage = page.NewInstallPage(m, 0)
	m.endPage = page.NewEndPage(m, 0)

	m.backButton = widgets.NewQPushButton2("返回", m)
	m.nextButton = widgets.NewQPushButton2("继续", m)

	m.backButton.SetMinimumWidth(60)
	m.nextButton.SetMinimumWidth(60)

	m.backButton.SetStyleSheet(styles.BackButton)
	m.nextButton.SetStyleSheet(styles.NextButton)

	m.backButton.SetVisible(false)

	hboxLayout := widgets.NewQHBoxLayout()
	hboxLayout.SetSpacing(40)

	hboxLayout.AddStretch(1)
	hboxLayout.AddWidget(m.backButton, 0, core.Qt__AlignHCenter)
	hboxLayout.AddWidget(m.nextButton, 0, core.Qt__AlignHCenter)
	hboxLayout.AddStretch(1)

	m.stackLayout.AddWidget(m.welcomePage)
	m.stackLayout.AddWidget(m.networkPage)
	m.stackLayout.AddWidget(m.accountPage)
	m.stackLayout.AddWidget(m.partitionPage)
	m.stackLayout.AddWidget(m.selectDEPage)
	m.stackLayout.AddWidget(m.additionalSoftwarePage)
	m.stackLayout.AddWidget(m.installPage)
	m.stackLayout.AddWidget(m.endPage)

	vboxLayout.AddLayout(m.stackLayout, 1)
	vboxLayout.AddLayout(hboxLayout, 1)
	vboxLayout.AddSpacing(20)

	m.SetLayout(vboxLayout)
}

func (m *MainFrame) initPlayer() {
	m.mediaPlayer = multimedia.NewQMediaPlayer(m, 0)
	playList := multimedia.NewQMediaPlaylist(m)
	playList.AddMedia(multimedia.NewQMediaContent2(core.NewQUrl3("qrc:/resources/background.mp3", core.QUrl__TolerantMode)))
	playList.SetPlaybackMode(multimedia.QMediaPlaylist__CurrentItemInLoop)
	m.mediaPlayer.SetPlaylist(playList)
}

func (m *MainFrame) initConnect() {
	m.stackLayout.ConnectCurrentChanged(func(index int) {
		if index == 0 {
			m.backButton.SetVisible(false)
		} else if index == m.stackLayout.Count()-2 || index == 1 {
			m.nextButton.SetVisible(false)
			m.backButton.SetVisible(false)
		} else if index > 0 && index < 7 {
			m.backButton.SetVisible(true)
		}

		if index == 1 {
			go m.checkNetwork()
		}
	})

	m.backButton.ConnectClicked(func(checked bool) {
		m.stackLayout.SetCurrentIndex(m.stackLayout.CurrentIndex() - 1)
	})

	m.nextButton.ConnectClicked(func(checked bool) {
		index := m.stackLayout.CurrentIndex()
		switch index {
		case 1:
			m.networkPage.SetTips("正在检查网络...")
			go m.checkNetwork()
			return
		case 2:
			if config.Conf.Username == "" || config.Conf.Password == "" {
				m.accountPage.SetTips("请输入用户名或密码")
				return
			}

			m.setButtonVisible(false)
			m.accountPage.SetEnableLogin(false)
			m.accountPage.SetTips("登陆中...")
			m.accountPage.Repaint()

			if err := m.account.Login(config.Conf.Username, config.Conf.Password); err != nil {
				if strings.Contains(err.Error(), "username or password errors") {
					m.accountPage.SetTips("用户名或密码错误")
				} else {
					m.accountPage.SetTips("登陆失败")
				}
				m.setButtonVisible(true)
				m.accountPage.SetEnableLogin(true)
				return
			}
			m.accountPage.SetTips("加载数据中...")
			m.accountPage.Repaint()
			installer.SyncDatabase()
			m.additionalSoftwarePage.LoadData()
			m.nextButton.SetVisible(true)
			m.accountPage.SetEnableLogin(true)
			m.accountPage.SetTips("")
		case 3:
			if config.Conf.InstallDev == "" {
				return
			}
		case 4:
			m.additionalSoftwarePage.LoadInstallList()
		case 5:
			m.mediaPlayer.Play()
			go m.install()
		case 7:
			m.reboot()
		}
		m.stackLayout.SetCurrentIndex(index + 1)
	})
}

func (m *MainFrame) checkNetwork() {
	cmd := exec.Command("ping", "-c", "3", "www.baidu.com")
	if err := cmd.Run(); err != nil {
		m.networkPage.ConnectNetwork()
	} else {
		m.stackLayout.SetCurrentIndex(2)
	}
	m.setButtonVisible(true)
}

func (m *MainFrame) install() {
	message := make(chan string)

	go func() {
		for {
			msg := <-message
			if msg == "" {
				break
			}
			if strings.HasPrefix(msg, "message:") {
				m.installPage.SetTips(strings.Split(msg, ":")[1])
			} else if strings.HasPrefix(msg, "action:") {
				if strings.HasSuffix(msg, "closeMessage") {
					m.installPage.SetMessageVisible()
				}
			} else {
				m.installPage.AddMessage(msg)
			}
		}
	}()

	err := installer.Install(message)
	if err != nil {
		m.endPage.SetTips("安装失败\n错误:" + err.Error())
	}

	m.stackLayout.SetCurrentIndex(7)
	m.nextButton.SetText("重启")
	m.nextButton.SetVisible(true)
}

func (m *MainFrame) setButtonVisible(enable bool) {
	m.backButton.SetVisible(enable)
	m.nextButton.SetVisible(enable)
}

func (m *MainFrame) reboot() {
	exec.Command("reboot").Run()
}
