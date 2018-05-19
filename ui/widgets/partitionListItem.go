package widgets

import (
	"github.com/firerainos/firerain-installer/core/parted"
	"github.com/firerainos/firerain-installer/styles"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"strconv"
	"strings"
)

type PartitionListItem struct {
	*widgets.QPushButton

	iconLabel *widgets.QLabel

	partition parted.Partition
}

func NewPartitionListItem(partition parted.Partition, parent widgets.QWidget_ITF) *PartitionListItem {
	btn := widgets.NewQPushButton(parent)

	partitionListItem := &PartitionListItem{QPushButton: btn, partition: partition}
	partitionListItem.init()

	return partitionListItem
}

func (p *PartitionListItem) init() {
	p.SetFixedSize2(180, 200)
	p.SetCheckable(true)
	p.SetStyleSheet(styles.PartitionListItem)

	vboxLayout := widgets.NewQVBoxLayout2(p)

	p.iconLabel = widgets.NewQLabel(p, 0)
	device := p.partition.Device.Disk + strconv.Itoa(p.partition.Number)
	device = strings.Replace(device, "/dev/", "", 1)
	name := p.partition.Name
	if name == "" {
		name = "未命名"
	}
	nameLabel := widgets.NewQLabel2(name+"("+device+")", p, 0)
	sizeLabel := widgets.NewQLabel2("总共 "+p.partition.Size, p, 0)

	p.iconLabel.SetFixedSize2(120, 120)
	p.iconLabel.SetPixmap(gui.NewQPixmap5(":/resources/harddisk.svg", "", 0).Scaled2(120, 120, core.Qt__KeepAspectRatioByExpanding, 0))

	vboxLayout.AddSpacing(10)
	vboxLayout.AddWidget(p.iconLabel, 0, core.Qt__AlignHCenter)
	vboxLayout.AddSpacing(10)
	vboxLayout.AddWidget(nameLabel, 0, core.Qt__AlignHCenter)
	vboxLayout.AddWidget(sizeLabel, 0, core.Qt__AlignHCenter)
	vboxLayout.AddStretch(1)

	p.SetLayout(vboxLayout)
}

func (p *PartitionListItem) SetSelect(selected bool) {
	p.SetChecked(selected)
	if selected {
	} else {

	}
}

func (p *PartitionListItem) DevPath() string {
	return p.partition.Device.Disk + strconv.Itoa(p.partition.Number)
}

func (p *PartitionListItem) Size() string {
	return p.partition.Size
}
