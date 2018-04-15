package parted

type Partition struct {
	Device     Device
	Number     int
	Start      string
	End        string
	Size       string
	FileSystem string
	Name       string
	Flags      []string
}

func NewPartition(Device Device, Number int, Start, End, Size, FileSystem, Name string, Flags []string) Partition {
	return Partition{Device, Number, Start, End, Size, FileSystem, Name, Flags}
}
