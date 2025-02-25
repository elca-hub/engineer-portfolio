package user_presenter

import (
	"devport/presenter"
	"devport/usecase/user"
	"net/http"
)

type CreateUserResponse struct {
	Email string `json:"email"`
}

type CreatePresenter struct{}

func NewCreatePresenter() *CreatePresenter {
	return &CreatePresenter{}
}

func (p *CreatePresenter) Success(output user.CreateUserOutput) *presenter.OriginalResponse {
	return &presenter.OriginalResponse{
		Error:      nil,
		Data:       CreateUserResponse{Email: output.Email},
		StatusCode: http.StatusOK,
	}
}

func (p *CreatePresenter) Error(err error, message string) *presenter.OriginalResponse {
	return &presenter.OriginalResponse{
		Error: &presenter.OriginalErrorResponseObj{
			Error:        err,
			ErrorMessage: message,
		},
		Data:       nil,
		StatusCode: http.StatusInternalServerError,
	}
}
