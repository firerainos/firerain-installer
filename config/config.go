package config

import "os"

var Conf *Config

type Config struct {
	*Account
	*Package
	DEApplication string
	InstallDev string
	IsUEFI bool
	EFIDev string
}

func InitConfig() {
	Conf = &Config{Account: &Account{}, Package: &Package{},IsUEFI:true}
	Conf.AddPackage("base")
	Conf.AddPackage("base-devel")
	Conf.AddPackage("vim")
	Conf.AddPackage("sudo")
	Conf.AddPackage("networkmanager")
	Conf.AddPackage("firerain-fristboot")

	Conf.DEApplication="kde-applications"

	if _,err:=os.Stat("/sys/firmware/efi/efivars");err!=nil{
		if os.IsNotExist(err) {
			Conf.IsUEFI = false
		}
	}
}

func (config *Config) SetInstallDev(dev string) {
	config.InstallDev = dev
}
func (config *Config) SetEFIDev(dev string) {
	config.EFIDev = dev
}
