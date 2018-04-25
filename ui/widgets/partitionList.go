package widgets

import (
	"github.com/firerainos/firerain-installer/core/parted"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/firerainos/firerain-installer/config"
	"strconv"
)

type PartitionList struct {
	widgets.QScrollArea

	hboxLayout  *widgets.QHBoxLayout
	buttonGroup *widgets.QButtonGroup
	listItems   []*PartitionListItem

	parted *parted.Parted

	currendItem *PartitionListItem

	_ func() `constructor:"init"`

	_ func() `signal:"partitionItemChange"`
}

func (p *PartitionList) init() {
	p.SetMinimumSize2(500, 290)
	p.SetWidgetResizable(true)
	p.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)

	frame := widgets.NewQFrame(p, 0)

	p.hboxLayout = widgets.NewQHBoxLayout2(frame)
	p.hboxLayout.SetContentsMargins(40, 40, 40, 20)

	p.buttonGroup = widgets.NewQButtonGroup(p)

	frame.SetLayout(p.hboxLayout)

	p.SetWidget(frame)

	p.parted = parted.NewParted()

	p.initConnect()
	p.ScanPartition()
}

func (p *PartitionList) initConnect() {
	p.buttonGroup.ConnectButtonClicked2(func(id int) {
		p.currendItem = p.listItems[id]
		p.PartitionItemChange()
	})
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

	for _, item := range p.buttonGroup.Buttons() {
		p.buttonGroup.RemoveButton(item)
	}

	devices, err := p.parted.List()
	if err != nil {
		return
	}

	for _, dev := range devices {
		for _, partition := range dev.Partitions {
			if config.Conf.EFIDev == "" && (partition.FileSystem == "fat16" || partition.FileSystem == "fat32") {
				config.Conf.SetEFIDev(partition.Device.Disk + strconv.Itoa(partition.Number))
			}

			if partition.FileSystem != "btrfs" {
				continue
			}

			item := NewPartitionListItem(partition, p)
			p.hboxLayout.AddWidget(item, 0, core.Qt__AlignCenter)
			p.buttonGroup.AddButton(item, len(p.listItems))
			p.listItems = append(p.listItems, item)
			item.SetSelect(config.Conf.InstallDev == item.DevPath())
		}
	}
}

func (p *PartitionList) CurrendItem() *PartitionListItem {
	return p.currendItem
}
