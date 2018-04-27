package installer

import (
	"github.com/firerainos/firerain-installer/config"
	"syscall"
)

func MountMnt() error {
	return syscall.Mount(config.Conf.InstallDev, "/mnt", "", syscall.MS_BIND, "")
}

func BindMnt() {
	syscall.Mount("proc", "/mnt/proc", "proc", syscall.MS_NOSUID|syscall.MS_NOEXEC|syscall.MS_NODEV, "")
	syscall.Mount("sys", "/mnt/sys", "sysfs", syscall.MS_NOSUID|syscall.MS_NOEXEC|syscall.MS_NODEV, "ro")
	syscall.Mount("efivarfs", "/mnt/sys/firmware/efi/efivars", "efivarfs",
		syscall.MS_NOSUID|syscall.MS_NOEXEC|syscall.MS_NODEV, "")
	syscall.Mount("udev","/mnt/dev","devtmpfs",syscall.MS_NOSUID,"mode=0755")
	syscall.Mount("devpts","/mnt/dev/pts","devpts",
		syscall.MS_NOSUID | syscall.MS_NOEXEC,"mode=0620,gid=5")
	syscall.Mount("shm","/mnt/dev/shm","tmpfs",
		syscall.MS_NOSUID | syscall.MS_NODEV,"mode=1777")
	syscall.Mount("run","/mnt/run","tmpfs",
		syscall.MS_NOSUID | syscall.MS_NODEV,"mode=0755")
	syscall.Mount("tmp","/mnt/tmp","tmpfs",
		syscall.MS_NOSUID | syscall.MS_NODEV | syscall.MS_STRICTATIME,"mode=1777")
}

func UnBindMnt() {
	syscall.Unmount("/mnt/proc",0)
	syscall.Unmount("/mnt/sys/firmware/efi/efivars",0)
	syscall.Unmount("/mnt/sys",0)
	syscall.Unmount("/mnt/dev/pts",0)
	syscall.Unmount("/mnt/dev/shm",0)
	syscall.Unmount("/mnt/dev",0)
	syscall.Unmount("/mnt/run",0)
	syscall.Unmount("/mnt/tmp",0)
}
