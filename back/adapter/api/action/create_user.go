package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/user"
	"encoding/json"
	"io"
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
		)
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// TODO: ログの追加
			return
		}
	}(r.Body)

	if err := a.v.Validate(input); err != nil {
		response.NewErrorMessages(a.v.Messages(), http.StatusBadRequest).Send(w)
	}

	output, err := a.uc.Execute(input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError)

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
}
