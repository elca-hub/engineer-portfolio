package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/model"
	"devport/usecase/external_service_url"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateExternalServiceUrlAction struct {
	uc external_service_url.UpdateExternalServiceUrlUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewUpdateExternalServiceUrlAction(uc external_service_url.UpdateExternalServiceUrlUseCase, v validator.Validator, l logger.Logger) *UpdateExternalServiceUrlAction {
	return &UpdateExternalServiceUrlAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *UpdateExternalServiceUrlAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input external_service_url.UpdateExternalServiceUrlInput

	const logKey = "update_external_service_url"

	userContext, isExistsUserContext := c.Get("user")
	if !isExistsUserContext {
		err := errors.New("not found email")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when email: %v", err))
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	currentUser := userContext.(*model.User)

	input.UserId = currentUser.ID()
	input.Id = c.Param("externalServiceUrlId")

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when decode update external service url request")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when validate update external service url request")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when update external service url")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success update external service url")
}
