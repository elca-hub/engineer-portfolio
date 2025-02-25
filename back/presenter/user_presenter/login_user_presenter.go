package user_presenter

import (
	"devport/presenter"
	"devport/usecase/user"
	"net/http"
)

type LoginUserResponse struct {
	Email string `json:"email"`
}

type LoginUserPresenter struct{}

func NewLoginUserPresenter() *LoginUserPresenter {
	return &LoginUserPresenter{}
}

func (p *LoginUserPresenter) Success(output user.LoginUserOutput) *presenter.OriginalResponse {
	return &presenter.OriginalResponse{
		Error:      nil,
		Data:       LoginUserResponse{Email: output.Email},
		StatusCode: http.StatusOK,
	}
}

func (p *LoginUserPresenter) Error(err error, message string) *presenter.OriginalResponse {
	return &presenter.OriginalResponse{
		Error: &presenter.OriginalErrorResponseObj{
			Error:        err,
			ErrorMessage: message,
		},
		Data:       nil,
		StatusCode: http.StatusInternalServerError,
	}
}
