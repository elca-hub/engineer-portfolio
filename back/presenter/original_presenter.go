package presenter

type OriginalErrorResponseObj struct {
	Error        error  `json:"detail"`
	ErrorMessage string `json:"error_message"`
}

type OriginalResponse struct {
	Error      *OriginalErrorResponseObj `json:"error"`
	Data       interface{}               `json:"data"`
	StatusCode int                       `json:"-"`
}

type OriginalPresenter interface {
	Success(output interface{}) *OriginalResponse
	Error(err error, message string) *OriginalResponse
}
