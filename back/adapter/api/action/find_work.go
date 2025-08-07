package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/work"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindWorkAction struct {
	uc work.FindWorkUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewFindWorkAction(
	uc work.FindWorkUseCase,
	v validator.Validator, l logger.Logger,
) *FindWorkAction {
	return &FindWorkAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *FindWorkAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input work.FindWorkInput
	const logKey = "find_work"

	input.UserId = c.Param("userId")
	input.WorkId = c.Param("workId")

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("validation error")
		response.NewErrorMessages(a.v.Messages(), http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), input)

	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when find work")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success find work")
}
