package file_object

type BioImageObject struct {
	fileName    *BioImageFileName
	fileContent []byte
}

func NewBioImageObject(fileName *BioImageFileName, fileContent []byte) *BioImageObject {
	return &BioImageObject{
		fileName:    fileName,
		fileContent: fileContent,
	}
}

func (o *BioImageObject) GetFileContent() []byte {
	return o.fileContent
}

func (o *BioImageObject) GetFileName() *BioImageFileName {
	return o.fileName
}
