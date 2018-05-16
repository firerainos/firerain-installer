package parted

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Device struct {
	Model string
	Table string
	Disk  string
	Size  string

	Partitions []Partition
}

func ScanDevice() ([]Device, error) {
	var cmd *exec.Cmd
	if os.Getuid() == 0 {
		cmd = exec.Command("parted", "-lm")
	} else {
		cmd = exec.Command("pkexec", "parted", "-lm")
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var devices []Device

	devLists := strings.Split(string(out), "BYT;\n")

	for _, dev := range devLists {
		if dev == "" || strings.Contains(dev,"Error:") {
			continue
		}
		items := strings.Split(dev, ";\n")
		var device Device
		for i, item := range items {
			if i == 0 {
				tmp := strings.Split(item, ":")
				if len(tmp) < 7 {
					continue
				}
				device = Device{tmp[6], tmp[5], tmp[0], tmp[1], nil}
			} else {
				tmps := strings.Split(item, ":")
				if len(tmps) < 6 {
					continue
				}
				num, _ := strconv.Atoi(tmps[0])
				flags := strings.Split(tmps[len(tmps)-1], ",")
				name := tmps[len(tmps)-2]
				if strings.Contains(name, "8bit") {
					name = ""
				}
				partition := NewPartition(device, num, tmps[1], tmps[2], tmps[3], tmps[4], name, flags)
				device.Partitions = append(device.Partitions, partition)
			}
		}
		devices = append(devices, device)
	}

	return devices, nil
}
