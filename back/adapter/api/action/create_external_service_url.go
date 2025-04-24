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
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateExternalServiceUrlAction struct {
	uc external_service_url.CreateExternalServiceUrlUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewCreateExternalServiceUrlAction(
	uc external_service_url.CreateExternalServiceUrlUseCase,
	v validator.Validator, l logger.Logger,
) *CreateExternalServiceUrlAction {
	return &CreateExternalServiceUrlAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *CreateExternalServiceUrlAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input external_service_url.CreateExternalServiceUrlInput
	const logKey = "create_external_service_url"

	userContext, isExistsUserContext := c.Get("user")
	if !isExistsUserContext {
		err := errors.New("not found email")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when email: %v", err))
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	currentUser := userContext.(*model.User)

	input.UserId = currentUser.ID()

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logging.NewError(
			a.l,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("error while decoding request body")
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
	}(r.Body)
	
	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("validation error")
		response.NewErrorMessages(a.v.Messages(), http.StatusBadRequest).Send(w)

		return
	}

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when create external service url")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)

	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success external service url")
}
