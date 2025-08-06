package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/model"
	"devport/usecase/skill"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindByUserSkillAction struct {
	uc skill.FindByUserSkillUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewFindByUserSkillAction(
	uc skill.FindByUserSkillUseCase,
	v validator.Validator, l logger.Logger,
) *FindByUserSkillAction {
	return &FindByUserSkillAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *FindByUserSkillAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input skill.FindByUserSkillInput
	const logKey = "find_by_user_skill"

	userContext, isExistsUserContext := c.Get("user")
	if !isExistsUserContext {
		err := errors.New("not found user")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when user: %v", err))
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
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when find skills by user")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success find skills by user")
}