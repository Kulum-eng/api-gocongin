package responses

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, err interface{}) Response {
	var errorData string
	if e, ok := err.(error); ok {
		errorData = e.Error()
	} else if err != nil {
		errorData = err.(string)
	}
	return Response{
		Message: message,
		Data:    errorData,
	}
}
