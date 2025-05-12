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
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadBioAction struct {
	uc user.UploadBioUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewUploadBioAction(uc user.UploadBioUseCase, v validator.Validator, l logger.Logger) *UploadBioAction {
	return &UploadBioAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *UploadBioAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	const logKey = "upload_bio"

	userContext, isExistsUserContext := c.Get("user")
	if !isExistsUserContext {
		err := errors.New("not found email")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when email: %v", err))
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	currentUser := userContext.(*model.User)

	var input user.UploadBioInput

	input.UserId = currentUser.ID()

	input.IsDeleteImage = c.DefaultQuery("is_delete_image", "false") == "true"

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when decode upload bio request")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when validate upload bio request")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when upload bio")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success upload bio")
}
