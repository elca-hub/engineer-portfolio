package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/model"
	"devport/usecase/user"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadBioImageAction struct {
	uc user.UploadBioImageUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewUploadBioImageAction(uc user.UploadBioImageUseCase, v validator.Validator, l logger.Logger) *UploadBioImageAction {
	return &UploadBioImageAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *UploadBioImageAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	const logKey = "upload_bio_image"

	userContext, isExistsUserContext := c.Get("user")
	if !isExistsUserContext {
		err := errors.New("not found email")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when email: %v", err))
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	currentUser := userContext.(*model.User)

	var input user.UploadBioImageInput

	input.UserId = currentUser.ID()

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
	if imageFile, imageHeader, err := processFile("bio_image"); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error while getting image file")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	} else if imageFile != nil {
		defer imageFile.Close()
		input.Image = imageFile
		input.ImageHeader = imageHeader
	}

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when validate upload bio image request")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when upload bio image")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success upload bio image")
}
