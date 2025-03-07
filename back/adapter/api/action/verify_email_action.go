package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/middleware"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/user"
	"encoding/json"
	"io"
	"net/http"
)

type VerifyEmailAction struct {
	uc user.VerificationEmailUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewVerifyEmailAction(uc user.VerificationEmailUseCase, v validator.Validator, l logger.Logger) *VerifyEmailAction {
	return &VerifyEmailAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *VerifyEmailAction) Execute(w http.ResponseWriter, r *http.Request) {
	var input user.VerificationEmailInput
	const logKey = "create_user"

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
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when verify email")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	token, err := middleware.NewCookieToken(output.Token)

	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when set token")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	middleware.SetToken(w, token)

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success verify email")
}
