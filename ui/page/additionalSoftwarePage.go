package page

import (
	"github.com/firerainos/firerain-installer/api"
	"github.com/firerainos/firerain-installer/config"
	"github.com/firerainos/firerain-installer/core/installer"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	widgets2 "github.com/firerainos/firerain-installer/ui/widgets"
	"log"
	"strings"
	"github.com/firerainos/firerain-installer/styles"
)

type AdditionalSoftwarePage struct {
	*widgets.QFrame

	itemList                 *widgets.QListWidget
	packageList, installList *widgets.QListView
	leftButton, rightButton  *widgets.QPushButton

	packageLine *widgets2.LineEdit

	account *api.Account

	packageModel map[string]*gui.QStandardItemModel
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
	a.packageLine = widgets2.NewLineEdit(":/resources/package.svg",a)

	a.itemList.SetStyleSheet(styles.ItemList)
	a.packageList.SetStyleSheet(styles.PackageList)
	a.installList.SetStyleSheet(styles.InstallList)

	a.packageList.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)

	a.packageList.SetAlternatingRowColors(true)
	a.installList.SetAlternatingRowColors(true)

	a.packageLine.SetPlaceholderText("输入包名并回车以添加")

	a.itemList.SetFixedSize2(150, 500)
	a.packageList.SetFixedSize2(530, 500)
	a.installList.SetFixedWidth(250)
	a.packageLine.SetFixedWidth(250)

	installLayout := widgets.NewQVBoxLayout2(a)

	installLayout.AddWidget(a.packageLine,0,core.Qt__AlignHCenter)
	installLayout.AddSpacing(5)
	installLayout.AddWidget(a.installList,1,core.Qt__AlignHCenter)

	buttonLayout := widgets.NewQVBoxLayout2(a)
	buttonLayout.SetSpacing(16)

	a.leftButton = widgets.NewQPushButton(a)
	a.rightButton = widgets.NewQPushButton(a)

	a.leftButton.SetIcon(widgets.QApplication_Style().StandardIcon(widgets.QStyle__SP_ArrowLeft, nil, nil))
	a.rightButton.SetIcon(widgets.QApplication_Style().StandardIcon(widgets.QStyle__SP_ArrowRight, nil, nil))

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
	hboxLayout.AddLayout(installLayout,0)

	vboxLayout.AddSpacing(30)
	vboxLayout.AddWidget(welcomeLabel, 0, core.Qt__AlignCenter)
	vboxLayout.AddSpacing(50)
	vboxLayout.AddLayout(hboxLayout, 1)
	vboxLayout.AddStretch(1)

	a.SetLayout(vboxLayout)
}

func (a *AdditionalSoftwarePage) initConnect() {
	a.itemList.ConnectCurrentChanged(func(current *core.QModelIndex, previous *core.QModelIndex) {
		item := current.Data(int(core.Qt__DisplayRole)).ToString()
		if item == "" {
			return
		}
		if len(a.packageModel) != 0 {
			a.packageList.SetModel(a.packageModel[item])
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

	a.packageLine.ConnectKeyReleaseEvent(a.packageLineKeyEvent)

	a.leftButton.ConnectClicked(a.leftButtonClicked)
	a.rightButton.ConnectClicked(a.rightButtonClicked)
}

func (a *AdditionalSoftwarePage) LoadData() {
	if a.itemList.Count() != 0 {
		a.itemList.Clear()
	}

	widgets.NewQListWidgetItem2(config.Conf.DEApplication,a.itemList,0).SetTextAlignment(int(core.Qt__AlignCenter))

	items, err := a.account.GetItem()
	if err != nil {
		log.Println(err)
	}
	a.packageModel = make(map[string]*gui.QStandardItemModel, len(items)+3)

	for _, group := range []string{"kde-applications", "deepin-extra", "gnome-extra"} {
		pkgs := installer.GetGroupPkgInfo(group)
		model := gui.NewQStandardItemModel(a)
		for _, pkg := range pkgs {
			item := gui.NewQStandardItem2(pkg.Name + "\n" + pkg.Description)
			item.SetToolTip(pkg.Description)
			item.SetEditable(false)
			model.AppendRow2(item)
		}
		a.packageModel[group] = model
	}

	for _, item := range items {
		widgets.NewQListWidgetItem2(item.Title,a.itemList,0).SetTextAlignment(int(core.Qt__AlignCenter))
		model := gui.NewQStandardItemModel(a)
		for _, pkg := range item.Packages {
			item := gui.NewQStandardItem2(pkg.Name + "\n" + pkg.Description)
			item.SetEditable(false)
			model.AppendRow2(item)
		}
		a.packageModel[item.Title] = model
	}

	a.itemList.SetCurrentRow(0)
}

func (a *AdditionalSoftwarePage) LoadInstallList() {
	a.itemList.Item(0).SetText(config.Conf.DEApplication)
	if a.itemList.CurrentRow() == 0 {
		a.packageList.SetModel(a.packageModel[config.Conf.DEApplication])
	}

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

func (a *AdditionalSoftwarePage) packageLineKeyEvent(event *gui.QKeyEvent) {
	if event.Key() == int(core.Qt__Key_Enter) || event.Key() == int(core.Qt__Key_Return){
		pkg := a.packageLine.Text()
		if pkg == "" {
			return
		}

		if installer.PkgIsExist(pkg) {
			if !config.Conf.SearchPackage(pkg) {
				config.Conf.AddPackage(pkg)
				a.installModel.SetStringList(config.Conf.PkgList)
			}
			a.packageLine.SetText("")
		} else {
			widgets.NewQMessageBox2(widgets.QMessageBox__NoIcon,"提示",pkg+" 不存在",widgets.QMessageBox__Ok,a,0).Show()
		}
	}
}