package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/user"
	"encoding/json"
	"net/http"
)

type CreateUserAction struct {
	uc user.CreateUserUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewCreateUserAction(uc user.CreateUserUseCase, v validator.Validator, l logger.Logger) *CreateUserAction {
	return &CreateUserAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *CreateUserAction) Execute(w http.ResponseWriter, r *http.Request) {
	var input user.CreateUserInput
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

	defer r.Body.Close()

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("validation error")
		response.NewErrorMessages(a.v.Messages(), http.StatusBadRequest).Send(w)
	}

	output, err := a.uc.Execute(input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError)

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)

	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success create user")
}
