package response

type ErrorResponse struct {
	ErrorMessage string
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{ErrorMessage: err.Error()}
}
