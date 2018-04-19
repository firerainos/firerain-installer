package widgets

import (
	"github.com/firerainos/firerain-installer/styles"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type LineEdit struct {
	*widgets.QLineEdit

	iconLabel *widgets.QLabel
}

func NewLineEdit(icon string, parent widgets.QWidget_ITF) *LineEdit {
	widget := widgets.NewQLineEdit(parent)

	lineEdit := &LineEdit{QLineEdit: widget}
	lineEdit.init(icon)

	return lineEdit
}

func (lineEdit *LineEdit) init(icon string) {
	lineEdit.iconLabel = widgets.NewQLabel(lineEdit, 0)

	lineEdit.iconLabel.SetPixmap(gui.NewQPixmap5(icon, "", 0))
	lineEdit.iconLabel.SetFixedSize2(12, 12)
	lineEdit.iconLabel.Move2(12, 12)

	lineEdit.SetStyleSheet(styles.LineEdit)

	lineEdit.SetFixedSize2(310, 36)

	lineEdit.SetContextMenuPolicy(core.Qt__NoContextMenu)
}
