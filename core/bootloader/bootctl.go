package bootloader

import (
	"text/template"
	"os"
	"os/exec"
)

type Bootctl struct {
	bootPath string
	rootDevice string
}

func NewBootctl(bootPath,rootDevice string) *Bootctl {
	return &Bootctl{bootPath,rootDevice}
}

func (b *Bootctl) Install() error {
	cmd := exec.Command("bootctl","--path="+b.bootPath,"install")
	_,err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (b *Bootctl) CreateEntries() error {
	temp := `title   FireRain Linux
linux   /vmlinuz-linux
initrd  /initramfs-linux.img
options root=PARTUUID={{ . }} rw`

	t,err:=template.New("entries").Parse(temp)
	if err != nil {
		return err
	}

	file,err:=os.OpenFile(b.bootPath+	"/loader/entries/firerain.conf",os.O_CREATE | os.O_RDWR | os.O_TRUNC,755)
	if err != nil {
		return err
	}

	partuuid,err :=b.getPARTUUID(b.rootDevice)
	if err != nil {
		return err
	}

	err = t.Execute(file,partuuid)
	file.Close()
	return err
}

func (b *Bootctl) getPARTUUID(device string) (string,error) {
	cmd :=exec.Command("blkid","-s","PARTUUID","-o","value",device)
	out,err := cmd.CombinedOutput()
	if err != nil {
		return "",err
	}
	return string(out),nil
}