package installer

import (
	"os/exec"
	"github.com/firerainos/firerain-installer/config"
	"github.com/kr/pty"
	"bufio"
	"strings"
	"github.com/firerainos/firerain-installer/api"
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

func GetGroupPkgInfo(group string) []api.Package {
	cmd := exec.Command("pacman","-Sg",group)
	out,err:=cmd.CombinedOutput()
	if err != nil {
		return nil
	}

	lines := strings.Split(string(out),"\n")
	pkgs := make([]api.Package,len(lines))

	for i,line:= range lines {
		if line=="" {
			continue
		}
		name := strings.Split(line," ")[1]
		pkg := api.Package{Name:name,Description:GetPkgDescription(name)}
		pkgs[i] = pkg
	}

	return pkgs
}

func GetPkgDescription(pkgName string) string {
	cmd := exec.Command("pacman","-Si",pkgName)
	cmd.Env = append(cmd.Env, "LANG=C")
	out,err:=cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(out),"\n")
	for _,line := range lines {
		if strings.Contains(line,"Description") {
			return strings.TrimSpace(strings.Split(line,":")[1])
		}
	}

	return ""
}