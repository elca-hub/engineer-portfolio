package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type FetchUserInfoAction struct {
	uc user.FetchUserInfoUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewFetchUserInfoAction(uc user.FetchUserInfoUseCase, v validator.Validator, l logger.Logger) *FetchUserInfoAction {
	return &FetchUserInfoAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *FetchUserInfoAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input user.FetchUserInfoInput
	const logKey = "fetch_user_info"

	input.UserId = c.Param("userId")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log(fmt.Sprintf("error when close body because: %v", err))
			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}(r.Body)

	output, err := a.uc.Execute(r.Context(), input)

	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when checking if user exists: %v", err))
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success fetch user info")
}
