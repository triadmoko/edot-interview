package helpers

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func ResponseApi(message string, code int, status string, data interface{}) Response {
	response := Response{
		Message: message,
		Code:    code,
		Status:  status,
		Data:    data,
	}
	return response
}
