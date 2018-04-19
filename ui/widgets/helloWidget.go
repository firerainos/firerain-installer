package widgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type HelloWidget struct {
	*widgets.QWidget

	helloStr []string
	x        []int
}

func NewHelloWidget(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *HelloWidget {
	widget := widgets.NewQWidget(parent, fo)

	helloWidget := &HelloWidget{QWidget: widget}
	helloWidget.init()

	helloWidget.StartTimer(20, core.Qt__PreciseTimer)
	helloWidget.ConnectTimerEvent(func(event *core.QTimerEvent) {
		helloWidget.roller()
	})

	return helloWidget
}

func (helloWidget *HelloWidget) init() {
	helloWidget.helloStr = []string{"Olá", "Здравствуйте", "Salve", "여보세요", "Hallo", "ສະບາຍດີ", "你好", "Hello", "xin chào", "Halo", "saluton", "سلام", "Bonjour", "Kamusta"}
	helloWidget.x = []int{100, 600, 200, 600, 350, 350, 420, 420, 380, 480, 750, 300, 250, 680}

	helloWidget.SetFixedSize2(900, len(helloWidget.helloStr)*30+50)

	helloWidget.ConnectPaintEvent(helloWidget.paintEvent)
}

func (helloWidget *HelloWidget) Start() {
	go helloWidget.roller()
}

func (helloWidget *HelloWidget) roller() {
	for i := 0; i < len(helloWidget.helloStr); i++ {
		if helloWidget.x[i] < -200 {
			helloWidget.x[i] = 900
		} else {
			helloWidget.x[i] -= 5
		}
		helloWidget.Repaint()
	}
}

func (helloWidget *HelloWidget) paintEvent(event *gui.QPaintEvent) {
	painter := gui.NewQPainter2(helloWidget)
	defer painter.DestroyQPainter()
	painter.SetRenderHint(gui.QPainter__Antialiasing, true)

	for i := 0; i < len(helloWidget.helloStr); i++ {
		painter.DrawText3(helloWidget.x[i], i*30+20, helloWidget.helloStr[i])
	}
}
