package config

var Conf *Config

type Config struct {
	*Account
	*Package
	InstallDev string
}


func InitConfig() {
	Conf = &Config{}
}

func (config *Config) SetInstallDev(dev string) {
	config.InstallDev = dev
}