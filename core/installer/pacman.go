package installer

import (
	"os/exec"
	"github.com/firerainos/firerain-installer/config"
	"github.com/kr/pty"
	"bufio"
)

func Pacstrap(out chan string) error {
	return installPkg(out,"-r","/mnt","--cachedir=/mnt/var/cache/pacman/pkg")
}

func installPkg (out chan string,arg ...string) error {
	cmdArg := append([]string{"-S","--noconfirm","--force"},arg...)
	cmdArg = append(cmdArg,config.Conf.PkgList...)
	cmd := exec.Command("pacman",arg...)
	pacman,err := pty.Start(cmd)
	if err != nil {
		return err
	}
	defer pacman.Close()

	reader := bufio.NewReader(pacman)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		out <- string(line)
	}

	return nil
}

func SyncDatabase() error {
	cmd := exec.Command("pacman","-Syy")
	return cmd.Run()
}