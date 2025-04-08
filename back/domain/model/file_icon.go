package model

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/google/uuid"
)

const (
	MAX_FILE_SIZE = 50 * 1024 * 1024
)

type FileIcon struct {
	file     []byte
	fileName *FileIconName
}

func NewFileIcon(file multipart.File, fileHeader *multipart.FileHeader, path int) (*FileIcon, error) {
	if fileHeader.Size > MAX_FILE_SIZE {
		return nil, errors.New("file size exceeds the limit")
	}

	if fileHeader.Header.Get("Content-Type") != "image/png" && fileHeader.Header.Get("Content-Type") != "image/jpeg" {
		return nil, errors.New("invalid file type")
	}

	extensionTmp := fileHeader.Filename

	extension := ""

	for i := len(extensionTmp) - 1; i >= 0; i-- {
		if extensionTmp[i] == '.' {
			extension = extensionTmp[i+1:]
			break
		}
	}

	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}

	fileNameUUID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%s.%s", fileNameUUID.String(), extension)

	iconName, err := NewFileIconName(fileName, path)

	if err != nil {
		return nil, err
	}

	return &FileIcon{
		file:     buf.Bytes(),
		fileName: iconName,
	}, nil
}

func (f *FileIcon) GetFile() []byte {
	return f.file
}

func (f *FileIcon) GetFileName() *FileIconName {
	return f.fileName
}
