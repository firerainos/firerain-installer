package installer

import (
	"bufio"
	"github.com/firerainos/firerain-installer/api"
	"github.com/firerainos/firerain-installer/config"
	"github.com/kr/pty"
	"os/exec"
	"strings"
)

func Pacstrap(out chan string) error {
	return installPkg(out, "-r", "/mnt", "--cachedir=/mnt/var/cache/pacman/pkg")
}

func installPkg(out chan string, arg ...string) error {
	cmdArg := append([]string{"-Sy", "--noconfirm"}, arg...)
	cmdArg = append(cmdArg, config.Conf.PkgList...)
	cmd := exec.Command("pacman", cmdArg...)
	pacman, err := pty.Start(cmd)
	if err != nil {
		return err
	}
	defer pacman.Close()

	reader := bufio.NewReaderSize(pacman,1024)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		out <- string(line)
	}

	return cmd.Wait()
}

func SyncDatabase() error {
	cmd := exec.Command("pacman", "-Syy")
	return cmd.Run()
}

func GetGroupPkgInfo(group string) []api.Package {
	cmd := exec.Command("pacman", "-Sg", group)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil
	}

	lines := strings.Split(string(out), "\n")
	pkgs := make([]api.Package, 0)

	pkgInfo := make(chan api.Package)
	done := make(chan bool)
	defer close(pkgInfo)
	defer close(done)

	go func() {
		for {
			if len(pkgs) == len(lines)-1 {
				done <- true
				break
			}
			pkg := <-pkgInfo
			pkgs = append(pkgs, pkg)
		}
	}()

	for _, line := range lines {
		if line == "" {
			continue
		}
		name := strings.Split(line, " ")[1]

		go func() {
			pkgInfo <- api.Package{Name: name, Description: GetPkgDescription(name)}
		}()
	}

	<-done
	return pkgs
}

func GetPkgDescription(pkgName string) string {
	cmd := exec.Command("pacman", "-Si", pkgName)
	cmd.Env = append(cmd.Env, "LANG=C")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Description") {
			return strings.TrimSpace(strings.Split(line, ":")[1])
		}
	}

	return ""
}

func PkgIsExist(pkgName string) bool {
	cmd := exec.Command("pacman", "-Si", pkgName)
	return cmd.Run() == nil
}
