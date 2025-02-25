package action

import (
	"encoding/json"
	"devport/adapter/api/response"
	"devport/adapter/validator"
	"devport/usecase/user"
	"net/http"
)

type CreateUserAction struct {
	uc user.CreateUserUseCase
	v  validator.Validator
}

func NewCreateUserAction(uc user.CreateUserUseCase, v validator.Validator) *CreateUserAction {
	return &CreateUserAction{
		uc: uc,
		v:  v,
	}
}

func (a *CreateUserAction) Execute(w http.ResponseWriter, r *http.Request) {
	var input user.CreateUserInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	defer r.Body.Close()

	if err := a.v.Validate(input); err != nil {
		response.NewErrorMessages(a.v.Messages(), http.StatusBadRequest).Send(w)
	}

	output, err := a.uc.Execute(input)
	if err != nil {
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
}
