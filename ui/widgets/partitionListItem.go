package widgets

import (
	"github.com/firerainos/firerain-installer/core/parted"
	"github.com/therecipe/qt/core"
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

	vboxLayout.AddSpacing(20)
	vboxLayout.AddWidget(p.iconLabel, 0, core.Qt__AlignHCenter)
	vboxLayout.AddWidget(nameLabel, 0, core.Qt__AlignHCenter)
	vboxLayout.AddWidget(sizeLabel, 0, core.Qt__AlignHCenter)
	vboxLayout.AddStretch(1)

	p.SetLayout(vboxLayout)
}

func (p *PartitionListItem) SetSelect(selected bool) {
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
