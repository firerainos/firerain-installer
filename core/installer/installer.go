package installer

import (
	"bufio"
	"github.com/firerainos/firerain-installer/config"
	"github.com/firerainos/firerain-installer/core/bootloader"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Install(out chan string) error {
	defer close(out)
	out <- "message:正在初始化"
	if err := MountMnt(); err != nil {
		return err
	}
	defer UnMountMnt()

	createDir()

	BindMnt()
	defer UnBindMnt()

	out <- "message:正在安装"
	if err := Pacstrap(out); err != nil {
		return err
	}
	out <- "action:closeMessage"

	if config.Conf.IsUEFI {
		out <- "message:正在创建引导"
		if err := bootloader.NewBootctl("/mnt/boot", config.Conf.InstallDev).Deploy(); err != nil {
			return err
		}
	}

	out <- "message:正在完成最后操作"
	if err := EnableServices(); err != nil {
		return err
	}

	if err := genfstab(); err != nil {
		return err
	}

	if config.Conf.SearchPackage("sudo") {
		if err := setSudoers(); err != nil {
			return err
		}
	}

	if err := setLocale(); err != nil {
		return err
	}

	if err := AutoSetTimezone(); err != nil {
		return err
	}

	Chroot()
	defer ExitChroot()
	cmd := exec.Command("passwd", "-d")
	return cmd.Run()
}

func genfstab() error {
	cmd := exec.Command("bash", "-c", "genfstab -U /mnt >> /mnt/etc/fstab")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func createDir() {
	os.MkdirAll("/mnt/var/cache/pacman/pkg", 0755)
	os.MkdirAll("/mnt/var/lib/pacman", 0755)
	os.MkdirAll("/mnt/var/log", 0755)
	os.MkdirAll("/mnt/dev", 0755)
	os.MkdirAll("/mnt/run", 0755)
	os.MkdirAll("/mnt/etc", 0755)
	os.MkdirAll("/mnt/tmp", 1777)
	os.MkdirAll("/mnt/sys", 0555)
	os.MkdirAll("/mnt/proc", 1777)
}

func setSudoers() error {
	sudoers, err := os.Open("/mnt/etc/sudoers")
	if err != nil {
		return err
	}

	var data string

	reader := bufio.NewReader(sudoers)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if io.EOF == err {
				break
			} else {
				return err
			}
		}

		lineStr := string(line)
		if strings.Contains(lineStr, "# %wheel ALL=(ALL) ALL") {
			data += "%wheel ALL=(ALL) ALL"
		} else {
			data += lineStr
		}
		data += "\n"
	}

	return ioutil.WriteFile("/mnt/etc/sudoers", []byte(data), 0440)
}

func setLocale() error {
	locale, err := os.Open("/mnt/etc/locale.gen")
	if err != nil {
		return err
	}

	var data string

	reader := bufio.NewReader(locale)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if io.EOF == err {
				break
			} else {
				return err
			}
		}

		lineStr := string(line)
		if strings.Contains(lineStr, "#en_US.UTF-8 UTF-8") {
			data += "en_US.UTF-8 UTF-8"
		} else if strings.Contains(lineStr, "#zh_CN.UTF-8 UTF-8") {
			data += "zh_CN.UTF-8 UTF-8"
		} else if strings.Contains(lineStr, "#zh_TW.UTF-8 UTF-8") {
			data += "zh_TW.UTF-8 UTF-8"
		} else {
			data += lineStr
		}
		data += "\n"
	}

	if err := ioutil.WriteFile("/mnt/etc/locale.gen", []byte(data), 0644); err != nil {
		return err
	}

	cmd := exec.Command("locale-gen")
	return cmd.Run()
}
