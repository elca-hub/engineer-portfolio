package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	statusCode int
	Errors     []string `json:"errors"`
}

func NewError(err error, statusCode int) *Error {
	return &Error{
		statusCode: statusCode,
		Errors:     []string{err.Error()},
	}
}

func NewErrorMessages(messages []string, statusCode int) *Error {
	return &Error{
		statusCode: statusCode,
		Errors:     messages,
	}
}

func (e *Error) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.statusCode)

	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		panic(err)
	}

	return
}
