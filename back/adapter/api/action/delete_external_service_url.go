package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/model"
	"devport/usecase/external_service_url"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteExternalServiceUrlAction struct {
	uc external_service_url.DeleteExternalServiceUrlUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewDeleteExternalServiceUrlAction(
	uc external_service_url.DeleteExternalServiceUrlUseCase,
	v validator.Validator, l logger.Logger,
) *DeleteExternalServiceUrlAction {
	return &DeleteExternalServiceUrlAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *DeleteExternalServiceUrlAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input external_service_url.DeleteExternalServiceUrlInput
	const logKey = "delete_external_service_url"

	userContext, isExistsUserContext := c.Get("user")
	if !isExistsUserContext {
		err := errors.New("not found email")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when email: %v", err))
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	currentUser := userContext.(*model.User)

	input.UserId = currentUser.ID()
	input.ExternalServiceUrlId = c.Param("externalServiceUrlId")

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when delete external service url")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)

	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success delete external service url")
}
