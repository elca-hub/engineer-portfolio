package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/middleware"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/user"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type LoginUserAction struct {
	uc user.LoginUserUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewLoginUserAction(uc user.LoginUserUseCase, v validator.Validator, l logger.Logger) *LoginUserAction {
	return &LoginUserAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *LoginUserAction) Execute(w http.ResponseWriter, r *http.Request) {
	var input user.LoginUserInput

	const logKey = "login_user"

	if _, err := middleware.GetToken(r); err == nil {
		logging.NewInfo(a.l, logKey, http.StatusOK).Log("user already login")
		response.NewError(errors.New("user already login"), http.StatusBadRequest).Send(w)
		return
	}

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

	output, err := a.uc.Execute(input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when login user")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	token, err := middleware.NewCookieToken(output.Token)

	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when create cookie token")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	middleware.SetToken(w, token)

	response.NewSuccess(output, http.StatusOK).Send(w)

	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success login user")
}
