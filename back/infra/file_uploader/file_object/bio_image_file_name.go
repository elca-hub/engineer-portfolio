package file_object

import "fmt"

type BioImageFileName struct {
	fileName   string
	objectName string
}

const (
	BioImageUri = "bio_images"
)

func BioImageUserPrefix(userId string) string {
	return fmt.Sprintf("%s/%s/", BioImageUri, userId)
}

func NewBioImageFileName(userId string, fileName string) *BioImageFileName {
	return &BioImageFileName{
		fileName:   fileName,
		objectName: fmt.Sprintf("%s%s", BioImageUserPrefix(userId), fileName),
	}
}

func (f *BioImageFileName) GetFileName() string {
	return f.fileName
}

func (f *BioImageFileName) GetObjectName() string {
	return f.objectName
}
