package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/model"
	"devport/usecase/work"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateWorkAction struct {
	uc work.CreateWorkUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewCreateWorkAction(uc work.CreateWorkUseCase, v validator.Validator, l logger.Logger) *CreateWorkAction {
	return &CreateWorkAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *CreateWorkAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	const logKey = "create_work"

	userContext, isExistsUserContext := c.Get("user")
	if !isExistsUserContext {
		err := errors.New("not found email")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when email: %v", err))
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	currentUser := userContext.(*model.User)

	var input work.CreateWorkInput

	input.UserID = currentUser.ID()

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when decode create work request")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when validate create work request")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when create work")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success create work")
}
