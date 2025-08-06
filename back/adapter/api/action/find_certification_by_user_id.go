package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/certification"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindCertificationByUserIdAction struct {
	uc certification.FindByUserCertificationUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewFindCertificationByUserIdAction(
	uc certification.FindByUserCertificationUseCase,
	v validator.Validator, l logger.Logger,
) *FindCertificationByUserIdAction {
	return &FindCertificationByUserIdAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *FindCertificationByUserIdAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input certification.FindByUserCertificationInput
	const logKey = "find_certification_by_user_id"

	input.UserId = c.Param("userId")

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("validation error")
		response.NewErrorMessages(a.v.Messages(), http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when find certifications by user id")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success find certifications by user id")
}