package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/model"
	"devport/usecase/work"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadWorkImageAction struct {
	uc work.UploadWorkImageUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewUploadWorkImageAction(uc work.UploadWorkImageUseCase, v validator.Validator, l logger.Logger) *UploadWorkImageAction {
	return &UploadWorkImageAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *UploadWorkImageAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	const logKey = "upload_work_image"

	userContext, isExistsUserContext := c.Get("user")
	if !isExistsUserContext {
		err := errors.New("not found user")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when user: %v", err))
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	currentUser := userContext.(*model.User)

	var input work.UploadWorkImageInput

	input.UserId = currentUser.ID()
	input.WorkId = c.Param("work_id")

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

	// work画像の処理
	if imageFile, imageHeader, err := processFile("work_image"); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error while getting image file")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	} else if imageFile != nil {
		defer imageFile.Close()
		input.Image = imageFile
		input.ImageHeader = imageHeader
	}

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when validate upload work image request")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when upload work image")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusCreated).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusCreated).Log("success upload work image")
}