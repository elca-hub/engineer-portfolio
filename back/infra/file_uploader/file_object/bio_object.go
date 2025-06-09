package file_object

type BioObject struct {
	fileName    *BioFileName
	fileContent []byte
}

func NewBioObject(fileName *BioFileName, fileContent []byte) *BioObject {
	return &BioObject{
		fileName:    fileName,
		fileContent: fileContent,
	}
}

func (o *BioObject) GetFileContent() []byte {
	return o.fileContent
}

func (o *BioObject) GetFileName() *BioFileName {
	return o.fileName
}
