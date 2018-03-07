package networkmanager

import (
	"os/exec"
	"strings"
	"strconv"
)

type NetworkManager struct {
	wifiList WifiList
}

type WifiInfo struct {
	Ssid string
	Signal int
	Security string
	InUse bool
}

type WifiList map[string]WifiInfo

func NewNetworkManager() *NetworkManager {
	return &NetworkManager{}
}

func (n *NetworkManager) SetWifiStatus(status bool) {
	var cmd *exec.Cmd
	if status {
		cmd = exec.Command("nmcli","r","wifi","on")
	}else {
		cmd = exec.Command("nmcli","r","wifi","off")
	}
	cmd.Start()
}

func (n *NetworkManager) WifiScan() {
	tmpList := WifiList{}

	cmd := exec.Command("nmcli", "-t","dev", "wifi","list")

	out,err := cmd.Output()
	if err != nil {
		return
	}

	for _,wifi := range strings.Split(string(out),"\n"){
		item := strings.Split(wifi,":")
		if len(item) < 2 {
			continue
		}
		if item[1]== "" {
			continue
		}

		signal,_:=strconv.Atoi(item[5])
		info := WifiInfo{item[1],signal,item[7],item[0]=="*"}

		tmpList[info.Ssid] = info
	}

	n.wifiList = tmpList
}

func (n *NetworkManager) WifiList() *WifiList {
	return &n.wifiList
}
