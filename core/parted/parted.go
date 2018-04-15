package parted

type Parted struct {
}

func NewParted() *Parted {
	return &Parted{}
}

func (p *Parted) List() ([]Device, error) {
	return ScanDevice()
}
