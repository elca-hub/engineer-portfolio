package model

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/google/uuid"
)

const (
	MAX_FILE_SIZE = 50 * 1024 * 1024
)

type ImageFile struct {
	file     []byte
	fileName string
}

func NewImageFile(file multipart.File, fileHeader *multipart.FileHeader, path int) (*ImageFile, error) {
	if fileHeader.Size > MAX_FILE_SIZE {
		return nil, fmt.Errorf("ファイルサイズが%dMBを超えています", MAX_FILE_SIZE/1024/1024)
	}

	if fileHeader.Header.Get("Content-Type") != "image/png" && fileHeader.Header.Get("Content-Type") != "image/jpeg" {
		return nil, fmt.Errorf("不正なファイル形式です。Content-Type: %s", fileHeader.Header.Get("Content-Type"))
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

	if err != nil {
		return nil, err
	}

	return &ImageFile{
		file:     buf.Bytes(),
		fileName: fileName,
	}, nil
}

func (f *ImageFile) GetFile() []byte {
	return f.file
}

func (f *ImageFile) GetFileName() string {
	return f.fileName
}
