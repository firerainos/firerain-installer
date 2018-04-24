package config

var Conf *Config

type Config struct {
	*Account
	*Package
	InstallDev string
}

func InitConfig() {
	Conf = &Config{Account: &Account{}, Package: &Package{}}
	Conf.AddPackage("base")
	Conf.AddPackage("base-devel")
	Conf.AddPackage("vim")
	Conf.AddPackage("sudo")
	Conf.AddPackage("networkmanager")
}

func (config *Config) SetInstallDev(dev string) {
	config.InstallDev = dev
}
