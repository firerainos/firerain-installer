package networkmanager

import (
	"os/exec"
	"strconv"
	"strings"
)

type NetworkManager struct {
	wifiList WifiList
}

type WifiInfo struct {
	Ssid     string
	Signal   int
	Security string
	InUse    bool
}

type WifiList map[string]WifiInfo

func NewNetworkManager() *NetworkManager {
	return &NetworkManager{}
}

func (n *NetworkManager) SetWifiStatus(status bool) {
	var cmd *exec.Cmd
	if status {
		cmd = exec.Command("nmcli", "r", "wifi", "on")
	} else {
		cmd = exec.Command("nmcli", "r", "wifi", "off")
	}
	cmd.Start()
}

func (n *NetworkManager) WifiStatus() bool {
	var cmd *exec.Cmd
	cmd = exec.Command("nmcli", "r", "wifi")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	return string(out) == "enabled"
}

func (n *NetworkManager) CheckHasWifi() bool {
	var cmd *exec.Cmd
	cmd = exec.Command("nmcli", "-t", "dev")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	return strings.Contains(string(out), "wifi")
}

func (n *NetworkManager) WifiScan() {
	tmpList := WifiList{}

	cmd := exec.Command("nmcli", "-t", "dev", "wifi", "list")

	out, err := cmd.Output()
	if err != nil {
		return
	}

	for _, wifi := range strings.Split(string(out), "\n") {
		item := strings.Split(wifi, ":")
		if len(item) < 2 {
			continue
		}
		if item[1] == "" {
			continue
		}

		signal, _ := strconv.Atoi(item[5])
		info := WifiInfo{item[1], signal, item[7], item[0] == "*"}

		tmpList[info.Ssid] = info
	}

	n.wifiList = tmpList
}

func (n *NetworkManager) WifiList() *WifiList {
	return &n.wifiList
}

func (n *NetworkManager) ConnectWifi(ssid, password string) {
	var cmd *exec.Cmd
	if password == "" {
		cmd = exec.Command("nmcli", "-t", "dev", "wifi", "connect", ssid)
	} else {
		cmd = exec.Command("nmcli", "-t", "dev", "wifi", "connect", ssid, "password", password)
	}
	cmd.Run()
}
