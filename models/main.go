package models

type Item struct {
	Name string
	Pid  uint32
	Cpu  string
	Mem  string
}

func (i Item) FilterValue() string { return i.Name }
