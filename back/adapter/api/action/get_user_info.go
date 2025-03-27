package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/user"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type GetUserInfoAction struct {
	uc user.GetUserInfoUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewGetUserAction(uc user.GetUserInfoUseCase, v validator.Validator, l logger.Logger) *GetUserInfoAction {
	return &GetUserInfoAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *GetUserInfoAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input user.GetUserInfoInput
	const logKey = "get_user_info"

	contextToken, isExistsToken := c.Get("token")

	if !isExistsToken {
		err := errors.New("not found cookie token")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when get token")
		response.NewError(err, http.StatusBadRequest).Send(w)

		return
	}

	keysEmail, isExistsEmail := c.Get("email")

	if !isExistsEmail {
		err := errors.New("not found email")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when get email")
		response.NewError(err, http.StatusBadRequest).Send(w)

		return
	}

	input.Token = contextToken.(string)
	input.Email = keysEmail.(string)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when close body")
			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}(r.Body)

	output, err := a.uc.Execute(r.Context(), input)

	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error when get user info")
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success get user info")
}
