package installer

import (
	"github.com/firerainos/firerain-installer/config"
	"os/exec"
)

func EnableServices() error {
	if err := Chroot(); err != nil {
		return err
	}

	for _, pkg := range config.Conf.PkgList {
		var err error
		switch pkg {
		case "sddm":
			err = EnableService("sddm")
		case "lightdm":
			err = EnableService("lightdm")
		case "networkmanager":
			err = EnableService("NetworkManager")
		case "teamviewer":
			fallthrough
		case "teamviewer-beta":
			err = EnableService("teamviewerd")
		case "bumblebee":
			err = EnableService("bumblebeed")
		case "firerain-fristboot":
			err = EnableService("firerain-fristboot")
		case "tlp":
			if err = EnableService("tlp");err!=nil {
				break
			}
			err = EnableService("tlp-sleep")
		}
		if err != nil {
			return err
		}
	}

	return ExitChroot()
}

func EnableService(service string) error {
	cmd := exec.Command("systemctl", "enable", service)
	return cmd.Run()
}
