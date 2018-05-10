package page

import (
	"github.com/firerainos/firerain-installer/api"
	"github.com/firerainos/firerain-installer/config"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"log"
	"strings"
)

type AdditionalSoftwarePage struct {
	*widgets.QFrame

	itemList                 *widgets.QListWidget
	packageList, installList *widgets.QListView
	leftButton, rightButton  *widgets.QPushButton

	account *api.Account

	packageModel []*gui.QStandardItemModel
	installModel *core.QStringListModel
}

func NewAdditionalSoftwarePage(account *api.Account, parent widgets.QWidget_ITF, fo core.Qt__WindowType) *AdditionalSoftwarePage {
	frame := widgets.NewQFrame(parent, fo)

	page := &AdditionalSoftwarePage{QFrame: frame, account: account}
	page.init()
	page.initConnect()

	return page
}

func (a *AdditionalSoftwarePage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(a)

	welcomeLabel := widgets.NewQLabel2("选择附加软件包", a, 0)

	hboxLayout := widgets.NewQHBoxLayout2(a)
	hboxLayout.SetSpacing(0)

	a.itemList = widgets.NewQListWidget(a)
	a.packageList = widgets.NewQListView(a)
	a.installList = widgets.NewQListView(a)

	a.itemList.SetFixedSize2(150, 500)
	a.packageList.SetFixedSize2(430, 500)
	a.installList.SetFixedSize2(250, 500)

	buttonLayout := widgets.NewQVBoxLayout2(a)
	buttonLayout.SetSpacing(16)

	a.leftButton = widgets.NewQPushButton(a)
	a.rightButton = widgets.NewQPushButton(a)

	a.leftButton.SetIcon(widgets.QApplication_Style().StandardIcon(widgets.QStyle__SP_ArrowLeft,nil,nil))
	a.rightButton.SetIcon(widgets.QApplication_Style().StandardIcon(widgets.QStyle__SP_ArrowRight,nil,nil))

	a.leftButton.SetFixedSize2(32, 32)
	a.rightButton.SetFixedSize2(32, 32)

	a.leftButton.SetEnabled(false)
	a.rightButton.SetEnabled(false)

	buttonLayout.AddStretch(1)
	buttonLayout.AddWidget(a.leftButton, 0, core.Qt__AlignCenter)
	buttonLayout.AddWidget(a.rightButton, 0, core.Qt__AlignCenter)
	buttonLayout.AddStretch(1)

	hboxLayout.AddWidget(a.itemList, 0, core.Qt__AlignLeft)
	hboxLayout.AddWidget(a.packageList, 0, core.Qt__AlignLeft)
	hboxLayout.AddLayout(buttonLayout, 0)
	hboxLayout.AddWidget(a.installList, 0, core.Qt__AlignRight)

	vboxLayout.AddStretch(1)
	vboxLayout.AddWidget(welcomeLabel, 0, core.Qt__AlignCenter)
	vboxLayout.AddSpacing(50)
	vboxLayout.AddLayout(hboxLayout, 1)
	vboxLayout.AddStretch(1)

	a.SetLayout(vboxLayout)
}

func (a *AdditionalSoftwarePage) initConnect() {
	a.itemList.ConnectCurrentChanged(func(current *core.QModelIndex, previous *core.QModelIndex) {
		if len(a.packageModel) != 0 {
			a.packageList.SetModel(a.packageModel[current.Row()])
		}
	})

	a.packageList.ConnectClicked(func(index *core.QModelIndex) {
		currentText := index.Data(int(core.Qt__DisplayRole)).ToString()
		if currentText != "" {
			a.leftButton.SetEnabled(false)
			a.rightButton.SetEnabled(!config.Conf.SearchPackage(strings.Split(currentText, "\n")[0]))
		}
		a.installList.ClearSelection()
	})

	a.installList.ConnectClicked(func(index *core.QModelIndex) {
		currentText := index.Data(int(core.Qt__DisplayRole)).ToString()
		if currentText != "" {
			a.leftButton.SetEnabled(true)
			a.rightButton.SetEnabled(false)
		}
		a.packageList.ClearSelection()
	})

	a.packageList.ConnectDoubleClicked(func(index *core.QModelIndex) {
		currentText := index.Data(int(core.Qt__DisplayRole)).ToString()
		if currentText != "" && !config.Conf.SearchPackage(strings.Split(currentText, "\n")[0]) {
			a.rightButtonClicked(true)
		}
	})

	a.installList.ConnectDoubleClicked(func(index *core.QModelIndex) {
		a.leftButtonClicked(true)
		a.leftButton.SetEnabled(false)
	})

	a.leftButton.ConnectClicked(a.leftButtonClicked)
	a.rightButton.ConnectClicked(a.rightButtonClicked)
}

func (a *AdditionalSoftwarePage) LoadData() {
	if len(a.packageModel) !=0 {
		a.packageModel = make([]*gui.QStandardItemModel,0)
		a.itemList.Clear()
	}

	items, err := a.account.GetItem()
	if err != nil {
		log.Println(err)
	}

	for _, item := range items {
		a.itemList.AddItem(item.Title)
		model := gui.NewQStandardItemModel(a)
		for _, pkg := range item.Packages {
			item := gui.NewQStandardItem2(pkg.Name + "\n" + pkg.Description)
			item.SetEditable(false)
			model.AppendRow2(item)
		}
		a.packageModel = append(a.packageModel, model)
	}

	a.itemList.SetCurrentRow(0)

	a.packageList.SetModel(a.packageModel[0])
}

func (a *AdditionalSoftwarePage) LoadInstallList() {
	a.installModel = core.NewQStringListModel2(config.Conf.PkgList, a)
	a.installList.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	a.installList.SetModel(a.installModel)
}

func (a *AdditionalSoftwarePage) leftButtonClicked(checked bool) {
	if str := a.installList.CurrentIndex().Data(int(core.Qt__DisplayRole)).ToString(); str != "" {
		config.Conf.RemovePackage(strings.Split(str, "\n")[0])
		a.installModel.SetStringList(config.Conf.PkgList)
	}
}

func (a *AdditionalSoftwarePage) rightButtonClicked(checked bool) {
	if str := a.packageList.CurrentIndex().Data(int(core.Qt__DisplayRole)).ToString(); str != "" {
		config.Conf.AddPackage(strings.Split(str, "\n")[0])
		a.installModel.SetStringList(config.Conf.PkgList)
	}
}
