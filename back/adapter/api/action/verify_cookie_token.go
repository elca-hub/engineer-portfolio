package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/middleware"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/user"
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

	userToken, err := middleware.GetToken(r)

	if err != nil {
		errMsg := "error when get token"

		if err.Error() == "empty token" {
			errMsg = "token is empty"
		}

		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(errMsg)
		response.NewError(err, http.StatusBadRequest).Send(w)

		return
	}

	input.Token = userToken.Token()

	defer r.Body.Close()

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("validation error")
		response.NewErrorMessages(a.v.Messages(), http.StatusBadRequest).Send(w)
	}

	output, err := a.uc.Execute(input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when verify cookie token")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	cookie, err := middleware.NewCookieToken(output.Token)

	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when set token")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	middleware.SetToken(w, cookie)
	c.Set("email", output.Email)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success verify cookie token")
}
