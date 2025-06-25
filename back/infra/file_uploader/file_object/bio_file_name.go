package file_object

import "fmt"

type BioFileName struct {
	fileName   string
	objectName string
}

func NewBioFileName(userId string) *BioFileName {
	return &BioFileName{
		fileName:   "bio.md",
		objectName: fmt.Sprintf("bio/%s/bio.md", userId),
	}
}

func (f *BioFileName) GetFileName() string {
	return f.fileName
}

func (f *BioFileName) GetObjectName() string {
	return f.objectName
}
