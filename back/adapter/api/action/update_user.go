package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/model"
	"devport/usecase/user"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
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

func (a *UpdateUserAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input user.UpdateUserInput
	const logKey = "upload_user_icon"

	userContext, isExistsUserContext := c.Get("user")

	if !isExistsUserContext {
		err := errors.New("not found email")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when get email")
		response.NewError(err, http.StatusBadRequest).Send(w)

		return
	}

	input.UserId = userContext.(*model.User).ID()

	file, fileHeader, err := r.FormFile("icon")

	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error while getting file from form")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when close body")
			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}(file)

	input.Icon = file

	input.IconHeader = fileHeader

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when logout user")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success logout")
}
