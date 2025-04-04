//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package file_uploader

type FileUploader interface {
	UploadFile(fileBytes []byte, objectName string) error
	DeleteFile(objectName string) error
}
