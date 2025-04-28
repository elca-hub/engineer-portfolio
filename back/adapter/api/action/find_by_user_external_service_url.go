package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/external_service_url"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type FindByUserExternalServiceUrlAction struct {
	uc external_service_url.FindByUserExternalServiceUrlUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewFindByUserExternalServiceUrlAction(uc external_service_url.FindByUserExternalServiceUrlUseCase, v validator.Validator, l logger.Logger) *FindByUserExternalServiceUrlAction {
	return &FindByUserExternalServiceUrlAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *FindByUserExternalServiceUrlAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input external_service_url.FindByUserExternalServiceUrlInput
	const logKey = "find_by_user_external_service_url"

	input.UserId = c.Param("userId")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log(fmt.Sprintf("error when close body because: %v", err))
			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}(r.Body)

	output, err := a.uc.Execute(r.Context(), input)

	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when checking if user exists: %v", err))
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success fetch user info")
}
