package responses

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(statusCode int, message string, data interface{}) Response {
	return Response{
		Code:    statusCode,
		Status:  "Success",
		Message: message,
		Data:    data,
	}
}

func Error(statusCode int, message string, err interface{}) Response {
	return Response{
		Code:    statusCode,
		Status:  "Error",
		Message: message,
		Error:   err,
	}
}
