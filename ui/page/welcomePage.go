package page

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/multimedia"
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

	mediaPlayer := multimedia.NewQMediaPlayer(w, 0)
	videoWidget := multimedia.NewQVideoWidget(w)
	playerlist := multimedia.NewQMediaPlaylist(w)
	playerlist.AddMedia(multimedia.NewQMediaContent2(core.QUrl_FromLocalFile("/home/linux/go/src/github.com/firerainos/firerain-installer/resources/hello.mp4")))
	//playerlist.AddMedia(multimedia.NewQMediaContent2(core.QUrl_FromLocalFile("resources/hello.mp4")))
	playerlist.SetCurrentIndex(1)
	playerlist.SetPlaybackMode(multimedia.QMediaPlaylist__CurrentItemInLoop)
	mediaPlayer.SetPlaylist(playerlist)
	mediaPlayer.SetVideoOutput(videoWidget)
	mediaPlayer.Play()

	videoWidget.SetFixedSize2(900, 505)

	vboxLayout.AddWidget(videoWidget, 0, core.Qt__AlignCenter)
	vboxLayout.AddStretch(1)

	w.SetLayout(vboxLayout)

}
