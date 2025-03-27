package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/usecase/user"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type IsExistsUserAction struct {
	uc user.IsExistsUserUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewIsExistsUserAction(uc user.IsExistsUserUseCase, v validator.Validator, l logger.Logger) *IsExistsUserAction {
	return &IsExistsUserAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *IsExistsUserAction) Execute(w http.ResponseWriter, r *http.Request, ctx *gin.Context) {
	var input user.IsExistsUserInput
	const logKey = "is_exists_user"

	input.Email = ctx.Query("email")

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
