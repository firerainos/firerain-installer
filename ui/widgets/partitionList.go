package widgets

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	"github.com/firerainos/firerain-installer/core/parted"
)

type PartitionList struct {
	*widgets.QScrollArea

	hboxLayout *widgets.QHBoxLayout

	parted *parted.Parted
}

func NewPartitionList(parent widgets.QWidget_ITF) *PartitionList {
	widget := widgets.NewQScrollArea(parent)

	partitionList := &PartitionList{QScrollArea: widget,parted:parted.NewParted()}
	partitionList.init()
	partitionList.ScanPartition()

	return partitionList
}

func (p *PartitionList) init() {
	p.SetMinimumSize2(500,290)
	p.SetWidgetResizable(true)
	p.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)

	frame := widgets.NewQFrame(p,0)

	p.hboxLayout = widgets.NewQHBoxLayout2(frame)
	p.hboxLayout.SetContentsMargins(40,40,40,20)

	frame.SetLayout(p.hboxLayout)

	p.SetWidget(frame)
}

func (p *PartitionList) ScanPartition() {
	for {
		item := p.hboxLayout.TakeAt(0)
		if item.Pointer() == nil {
			break
		}

		if item.Widget().Pointer() != nil {
			item.Widget().DestroyQWidget()
		}

		if item.Layout().Pointer() != nil {
			item.Layout().DestroyQLayoutItem()
		}
	}

	devices,err := p.parted.List()
	if err != nil {
		return
	}

	for _,dev := range devices {
		for _,partition := range dev.Partitions{
			if partition.FileSystem!="btrfs" {
				continue
			}
			item := NewPartitionListItem(partition,p)
			p.hboxLayout.AddWidget(item,0,core.Qt__AlignCenter)
		}
	}
}