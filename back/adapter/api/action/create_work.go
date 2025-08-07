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
	var input work.CreateWorkInput
	const logKey = "create_work"

	userContext, isExistsUserContext := c.Get("user")
	if !isExistsUserContext {
		err := errors.New("not found email")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when email: %v", err))
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	currentUser := userContext.(*model.User)

	input.UserId = currentUser.ID()

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("validation error")
		response.NewErrorMessages(a.v.Messages(), http.StatusBadRequest).Send(w)

		return
	}

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when create work")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusCreated).Send(w)

	logging.NewInfo(a.l, logKey, http.StatusCreated).Log("success create work")
}
