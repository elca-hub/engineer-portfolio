package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/model"
	"devport/usecase/certification"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteCertificationAction struct {
	uc certification.DeleteCertificationUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewDeleteCertificationAction(
	uc certification.DeleteCertificationUseCase,
	v validator.Validator, l logger.Logger,
) *DeleteCertificationAction {
	return &DeleteCertificationAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *DeleteCertificationAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input certification.DeleteCertificationInput
	const logKey = "delete_certification"

	userContext, isExistsUserContext := c.Get("user")
	if !isExistsUserContext {
		err := errors.New("not found user")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when user: %v", err))
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	currentUser := userContext.(*model.User)
	input.UserId = currentUser.ID()
	input.ID = c.Param("certificationId")

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("validation error")
		response.NewErrorMessages(a.v.Messages(), http.StatusBadRequest).Send(w)
		return
	}

	err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when delete certification")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(nil, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success delete certification")
}