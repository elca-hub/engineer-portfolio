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

type VerifyCookieTokenAction struct {
	uc user.VerifyCookieTokenUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewVerifyCookieTokenAction(uc user.VerifyCookieTokenUseCase, v validator.Validator, l logger.Logger) *VerifyCookieTokenAction {
	return &VerifyCookieTokenAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *VerifyCookieTokenAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input user.VerifyCookieTokenInput
	const logKey = "verify_cookie_token"

	token := r.Header.Get("Authorization")
	if token == "" {
		err := errors.New("not found cookie token")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when login user")
		response.NewError(err, http.StatusInternalServerError).Send(w)

		return
	}

	// tokenの先頭にBearerがついているので取り除く
	input.Token = token[7:]

	output, err := a.uc.Execute(r.Context(), input)

	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when verify cookie token")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	c.Set("user", output.User)

	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success verify cookie token")
}
