package action

import (
	"devport/adapter/api/middleware"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/user"
	"encoding/json"
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

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
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
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	middleware.SetToken(w, output.Token)

	response.NewSuccess(output, http.StatusOK).Send(w)
}
