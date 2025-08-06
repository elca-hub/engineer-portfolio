package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/skill"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindSkillByUserIdAction struct {
	uc skill.FindByUserSkillUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewFindSkillByUserIdAction(
	uc skill.FindByUserSkillUseCase,
	v validator.Validator, l logger.Logger,
) *FindSkillByUserIdAction {
	return &FindSkillByUserIdAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *FindSkillByUserIdAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input skill.FindByUserSkillInput
	const logKey = "find_skill_by_user_id"

	input.UserId = c.Param("userId")

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("validation error")
		response.NewErrorMessages(a.v.Messages(), http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when find skills by user id")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success find skills by user id")
}