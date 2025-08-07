package file_object

import "fmt"

type WorkSentenceFileName struct {
	fileName   string
	objectName string
}

func NewWorkSentenceFileName(userId string, workId string) *WorkSentenceFileName {
	return &WorkSentenceFileName{
		fileName:   "work.md",
		objectName: fmt.Sprintf("work_content/%s/%s/work.md", userId, workId),
	}
}

func (f *WorkSentenceFileName) GetFileName() string {
	return f.fileName
}

func (f *WorkSentenceFileName) GetObjectName() string {
	return f.objectName
}