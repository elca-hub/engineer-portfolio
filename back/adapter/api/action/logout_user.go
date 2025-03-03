package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/middleware"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/user"
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

func (a *LogoutUserAction) Execute(w http.ResponseWriter, r *http.Request) {
	var input user.LogoutUserInput
	const logKey = "logout_user"

	userToken, err := middleware.GetToken(r)

	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when get token")
		response.NewError(err, http.StatusBadRequest).Send(w)

		return
	}

	input.Token = userToken.Token()

	defer r.Body.Close()

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("validation error")
		response.NewErrorMessages(a.v.Messages(), http.StatusBadRequest).Send(w)
	}

	output, err := a.uc.Execute(input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when logout user")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	middleware.DeleteToken(w)

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success logout")
}
