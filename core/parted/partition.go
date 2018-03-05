package parted

type Partition struct {
	Number int
	Start string
	End string
	Size string
	FileSystem string
	Name string
	Flags []string
}

func NewPartition(Number int,Start,End,Size,FileSystem,Name string,Flags []string) Partition {
	return Partition{Number,Start,End,Size,FileSystem,Name,Flags}
}
