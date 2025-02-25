package model

type Folder struct {
	id   int64
	name string
}

func NewFolder(id int64, name string) *Folder {
	return &Folder{
		id:   id,
		name: name,
	}
}

func (f *Folder) ID() int64 {
	return f.id
}

func (f *Folder) Name() string {
	return f.name
}
