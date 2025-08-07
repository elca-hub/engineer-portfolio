package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/work"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FetchWorkAction struct {
	uc work.FetchWorkUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewFetchWorkAction(uc work.FetchWorkUseCase, v validator.Validator, l logger.Logger) *FetchWorkAction {
	return &FetchWorkAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *FetchWorkAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input work.FetchWorkInput
	const logKey = "fetch_work"

	input.UserId = c.Param("userId")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log(fmt.Sprintf("error when close body because: %v", err))
			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}(r.Body)

	input.Type = c.Query("type")

	output, err := a.uc.Execute(r.Context(), input)

	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when checking if user exists: %v", err))
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success fetch work")
}
