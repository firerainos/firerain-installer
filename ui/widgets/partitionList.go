package widgets

import (
	"github.com/firerainos/firerain-installer/config"
	"github.com/firerainos/firerain-installer/core/parted"
	"github.com/firerainos/firerain-installer/styles"
	"github.com/fsnotify/fsnotify"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os/exec"
	"strconv"
	"strings"
)

type PartitionList struct {
	widgets.QScrollArea

	hboxLayout  *widgets.QHBoxLayout
	buttonGroup *widgets.QButtonGroup
	listItems   []*PartitionListItem

	openButton *widgets.QPushButton

	parted *parted.Parted

	watcher *fsnotify.Watcher
	devList []string

	currendItem *PartitionListItem

	_ func() `constructor:"init"`

	_ func() `signal:"partitionItemChange"`
	_ func() `signal:"deviceChange"`
}

func (p *PartitionList) init() {
	p.SetMinimumSize2(500, 290)
	p.SetWidgetResizable(true)
	p.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	p.SetStyleSheet(styles.PartitionList)

	frame := widgets.NewQFrame(p, 0)

	p.hboxLayout = widgets.NewQHBoxLayout2(frame)
	p.hboxLayout.SetContentsMargins(40, 40, 40, 20)

	p.buttonGroup = widgets.NewQButtonGroup(p)

	p.openButton = widgets.NewQPushButton2("打开分区管理器", p)
	p.openButton.SetVisible(false)
	p.openButton.SetGeometry2(185, 120, 130, 50)

	frame.SetLayout(p.hboxLayout)

	p.SetWidget(frame)

	p.parted = parted.NewParted()

	p.watcher, _ = fsnotify.NewWatcher()
	p.watcher.Add("/dev")

	go func() {
		for {
			event := <-p.watcher.Events
			name := ""
			if event.Op&fsnotify.Create == fsnotify.Create {
				name = event.Name
			} else if event.Op&fsnotify.Remove == fsnotify.Remove {
				name = event.Name
			}

			if name != "" {
				for _, dev := range p.devList {
					if strings.HasPrefix(name, dev) {
						p.DeviceChange()
						break
					}
				}
			}
		}
	}()

	p.initConnect()
	p.ScanPartition()
}

func (p *PartitionList) initConnect() {
	p.buttonGroup.ConnectButtonClicked2(func(id int) {
		p.currendItem = p.listItems[id]
		p.PartitionItemChange()
	})

	p.openButton.ConnectClicked(func(checked bool) {
		exec.Command("/usr/bin/partitionmanager").Start()
	})

	p.ConnectDeviceChange(func() {
		p.ScanPartition()
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

	p.devList = make([]string, 0)
	for _, dev := range devices {
		p.devList = append(p.devList, dev.Disk)
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

	p.openButton.SetVisible(len(p.listItems) == 0)
}

func (p *PartitionList) CurrendItem() *PartitionListItem {
	return p.currendItem
}
