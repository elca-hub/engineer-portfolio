package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/model"
	"devport/usecase/user"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserAction struct {
	uc user.UpdateUserUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewUpdateUserAction(uc user.UpdateUserUseCase, v validator.Validator, l logger.Logger) *UpdateUserAction {
	return &UpdateUserAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

type UpdateUserRequest struct {
	Name                string `json:"name"`
	Birthday            string `json:"birthday"`
	Email               string `json:"email"`
	BioPath             string `json:"bio_path"`
	Skills              []*model.Skill
	ExternalServiceURLs []*model.ExternalServiceUrl
}

func (a *UpdateUserAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	const logKey = "update_user"

	userContext, isExistsUserContext := c.Get("user")
	if !isExistsUserContext {
		err := errors.New("not found email")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when email: %v", err))
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	currentUser := userContext.(*model.User)
	input := user.UpdateUserInput{
		UserId:              currentUser.ID(),
		Name:                currentUser.Name(),
		Birthday:            currentUser.Birthday().Format("2006-01-02"),
		Email:               currentUser.Email().Email(),
		BioPath:             currentUser.BioPath(),
		Skills:              currentUser.Skills(),
		ExternalServiceURLs: currentUser.ExternalServiceURLs(),
	}

	// マルチパートフォームの処理
	if err := r.ParseMultipartForm(32 << 20); err != nil && err != http.ErrNotMultipart {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error while parsing multipart form")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	// user_dataの処理
	if userData := r.FormValue("user_data"); userData != "" {
		var jsonRequest UpdateUserRequest
		if err := json.Unmarshal([]byte(userData), &jsonRequest); err != nil {
			logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error while parsing user_data")
			response.NewError(err, http.StatusBadRequest).Send(w)
			return
		}

		if jsonRequest.Name != "" {
			input.Name = jsonRequest.Name
		}
		if jsonRequest.Birthday != "" {
			input.Birthday = jsonRequest.Birthday
		}
		if jsonRequest.Email != "" {
			input.Email = jsonRequest.Email
		}
		if jsonRequest.BioPath != "" {
			input.BioPath = jsonRequest.BioPath
		}
		if jsonRequest.Skills != nil {
			input.Skills = jsonRequest.Skills
		}
		if jsonRequest.ExternalServiceURLs != nil {
			input.ExternalServiceURLs = jsonRequest.ExternalServiceURLs
		}
	}

	// ファイルの取得と処理を行うヘルパー関数
	processFile := func(fieldName string) (multipart.File, *multipart.FileHeader, error) {
		file, header, err := r.FormFile(fieldName)
		if err != nil && errors.Is(err, http.ErrMissingFile) {
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
