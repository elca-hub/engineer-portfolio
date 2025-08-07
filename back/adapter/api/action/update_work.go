package action

import (
	"devport/adapter/api/logging"
	"devport/adapter/api/response"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/model"
	"devport/usecase/work"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateWorkAction struct {
	uc work.UpdateWorkUseCase
	v  validator.Validator
	l  logger.Logger
}

func NewUpdateWorkAction(uc work.UpdateWorkUseCase, v validator.Validator, l logger.Logger) *UpdateWorkAction {
	return &UpdateWorkAction{
		uc: uc,
		v:  v,
		l:  l,
	}
}

func (a *UpdateWorkAction) Execute(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var input work.UpdateWorkInput
	const logKey = "update_work"

	userContext, isExistsUserContext := c.Get("user")
	if !isExistsUserContext {
		err := errors.New("not found user")
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log(fmt.Sprintf("error when user: %v", err))
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	currentUser := userContext.(*model.User)
	input.UserId = currentUser.ID()

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logging.NewError(
			a.l,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("error while decoding request body")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when close body")
			response.NewError(err, http.StatusInternalServerError).Send(w)

			return
		}
	}(r.Body)

	// クエリパラメータからis_delete_imageを取得
	isDeleteImageStr := r.URL.Query().Get("is_delete_image")
	if isDeleteImageStr != "" {
		isDeleteImage, err := strconv.ParseBool(isDeleteImageStr)
		if err != nil {
			logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("error parsing is_delete_image parameter")
			response.NewError(err, http.StatusBadRequest).Send(w)
			return
		}
		input.IsDeleteImage = isDeleteImage
	}

	if err := a.v.Validate(input); err != nil {
		logging.NewError(a.l, err, logKey, http.StatusBadRequest).Log("validation error")
		response.NewErrorMessages(a.v.Messages(), http.StatusBadRequest).Send(w)

		return
	}

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(a.l, err, logKey, http.StatusInternalServerError).Log("error when update work")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)

	logging.NewInfo(a.l, logKey, http.StatusOK).Log("success update work")
}