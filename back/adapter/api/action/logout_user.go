package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/user"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutUserAction struct {
	uc user.LogoutUserUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewLogoutUserAction(uc user.LogoutUserUseCase, v validator.Validator, l logger.Logger) *LogoutUserAction {
	return &LogoutUserAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *LogoutUserAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input user.LogoutUserInput
	const logKey = "logout_user"

	contextToken, isExists := c.Get("token")

	if !isExists {
		errObj := errors.New("not found cookie token")
		logging.NewError(a.l, errObj, logKey, http.StatusBadRequest).Log("error when get token")
		response.NewError(errObj, http.StatusBadRequest).Send(w)

		return
	}

	input.Token = contextToken.(string)

	output, err := a.uc.Execute(input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when logout user")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success logout")
}
