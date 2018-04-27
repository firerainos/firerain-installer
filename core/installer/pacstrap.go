package installer

import (
	"os"
)

func Pacstrap() error {
	createDir()
	BindMnt()
}

func createDir() {
	os.MkdirAll("/mnt/var/cache/pacman/pkg",os.FileMode(0755))
	os.MkdirAll("/mnt/var/lib/pacman",os.FileMode(0755))
	os.MkdirAll("/mnt/var/log",os.FileMode(0755))
	os.MkdirAll("/mnt/dev",os.FileMode(0755))
	os.MkdirAll("/mnt/run",os.FileMode(0755))
	os.MkdirAll("/mnt/etc",os.FileMode(0755))
	os.MkdirAll("/mnt/tmp",os.FileMode(1777))
	os.MkdirAll("/mnt/sys",os.FileMode(0555))
	os.MkdirAll("/mnt/proc",os.FileMode(1777))
}
