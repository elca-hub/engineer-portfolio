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
	"regexp"
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

	uuidRegex := regexp.MustCompile(`^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}$`)

	rawToken := token[7:]

	if !uuidRegex.MatchString(rawToken) {
		err := errors.New("invalid cookie token")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when login user")
		response.NewError(err, http.StatusInternalServerError).Send(w)

		return
	}

	input.Token = rawToken

	output, err := a.uc.Execute(r.Context(), input)

	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when verify cookie token")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	c.Set("user", output.User)

	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success verify cookie token")
}
