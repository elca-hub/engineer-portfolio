package file_object

type FileNameInterface interface {
	GetFileName() string
	GetObjectName() string
}

type ObjectNameInterface interface {
	GetFileName() string
	GetObjectName() string
	GetFileContent() []byte
	GetFilePattern() uint
}
