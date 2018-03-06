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

func (b *Bootctl) Deploy() error {
	err := b.Install()
	if err != nil {
		return err
	}

	partuuid,err :=b.getPARTUUID(b.rootDevice)
	if err != nil {
		return err
	}

	b.CreateEntries("firerain.conf",partuuid,false)
	b.CreateEntries("firerain-fallback.conf",partuuid,true)

	b.SetDefaultEntries("firerain")

	return nil
}

func (b *Bootctl) SetDefaultEntries(entriesName string) error {
	file,err:=os.OpenFile(b.bootPath+"/loader/loader.conf",os.O_CREATE | os.O_RDWR | os.O_TRUNC,755)
	if err != nil {
		return err
	}

	_,err =file.Write([]byte("timeout 3\ndefault "+entriesName))
	if err != nil {
		return err
	}

	file.Close()

	return nil
}

func (b *Bootctl) CreateEntries(confName,partuuid string,fallback bool) error {
	temp := `title   FireRain Linux{{ if .fallback }}(fallback){{ endif }}
linux   /vmlinuz-linux
initrd  /initramfs-linux{{ if .fallback }}-fallback{{ endif }}.img
options root=PARTUUID={{ .partuuid }} rw`

	t,err:=template.New("entries").Parse(temp)
	if err != nil {
		return err
	}

	file,err:=os.OpenFile(b.bootPath+"/loader/entries/"+confName,os.O_CREATE | os.O_RDWR | os.O_TRUNC,755)
	if err != nil {
		return err
	}

	err = t.Execute(file,template.FuncMap{"partuuid":partuuid,"fallback":false})
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