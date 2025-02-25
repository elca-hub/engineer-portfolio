package response

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	result     interface{}
	statusCode int
}

func NewSuccess(result interface{}, statusCode int) *Success {
	return &Success{
		result:     result,
		statusCode: statusCode,
	}
}

func (s *Success) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.statusCode)

	if err := json.NewEncoder(w).Encode(s.result); err != nil {
		panic(err)
	}
}
