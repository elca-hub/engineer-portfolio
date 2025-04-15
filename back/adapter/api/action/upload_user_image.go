package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/user"
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadUserImageAction struct {
	uc user.UploadUserImageUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewUploadUserImageAction(uc user.UploadUserImageUseCase, v validator.Validator, l logger.Logger) *UploadUserImageAction {
	return &UploadUserImageAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *UploadUserImageAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	const logKey = "upload_user_image"

	var input user.UploadUserImageInput

	// マルチパートフォームの処理
	if err := r.ParseMultipartForm(32 << 20); err != nil && err != http.ErrNotMultipart {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error while parsing multipart form")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	// ファイルの取得と処理を行うヘルパー関数
	processFile := func(fieldName string) (multipart.File, *multipart.FileHeader, error) {
		file, header, err := r.FormFile(fieldName)
		if err != nil && errors.Is(err, http.ErrMissingFile) {
			if errors.Is(err, http.ErrMissingFile) {
				return nil, nil, nil
			}

			return nil, nil, err
		}
		return file, header, nil
	}

	// アイコン画像の処理
	if iconFile, iconHeader, err := processFile("icon"); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error while getting icon file")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	} else if iconFile != nil {
		defer iconFile.Close()
		input.Icon = iconFile
		input.IconHeader = iconHeader
	}

	// ヘッダー画像の処理
	if headerFile, headerHeader, err := processFile("header_icon"); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error while getting header file")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	} else if headerFile != nil {
		defer headerFile.Close()
		input.Header = headerFile
		input.HeaderHeader = headerHeader
	}

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when update user")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success update user")
}
