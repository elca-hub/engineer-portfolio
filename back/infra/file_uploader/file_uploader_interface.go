//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package file_uploader

type FileUploader interface {
	UploadFile(fileBytes []byte, bucketName string, objectName string) (string, error)
	CrateBucket(bucketName string) error
}
