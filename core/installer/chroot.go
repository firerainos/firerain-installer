package installer

import (
	"os"
	"syscall"
)

var root *os.File

func Chroot() error {
	var err error
	root, err = os.Open("/")
	if err != nil {
		return err
	}

	return syscall.Chroot("/mnt")
}

func ExitChroot() error {
	if err := root.Chdir(); err != nil {
		return err
	}
	defer root.Close()

	return syscall.Chroot(".")
}
