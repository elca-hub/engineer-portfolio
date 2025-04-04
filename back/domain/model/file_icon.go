package model

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
)

const (
	MAX_FILE_SIZE = 500 * 1024 * 1024
)

type FileIcon struct {
	file      *bytes.Buffer
	extension string
}

func NewFileIcon(file multipart.File, fileHeader *multipart.FileHeader) (*FileIcon, error) {
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

	return &FileIcon{
		file:      buf,
		extension: extension,
	}, nil
}

func (f *FileIcon) GetFile() *bytes.Buffer {
	return f.file
}

func (f *FileIcon) GetExtension() string {
	return f.extension
}
