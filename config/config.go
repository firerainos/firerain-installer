package config

var Conf *Config

type Config struct {
	*Account
	*Package
	InstallDev string
}


func InitConfig() {
	Conf = &Config{Account:&Account{},Package:&Package{}}
}

func (config *Config) SetInstallDev(dev string) {
	config.InstallDev = dev
}