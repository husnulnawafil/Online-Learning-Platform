package helpers

type APIResponse struct {
	Meta *Meta       `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func ReponseWithData(code int, message string, data interface{}) *APIResponse {
	return &APIResponse{
		&Meta{
			StatusCode: code,
			Message:    message,
		},
		data,
	}
}

func ResponseWithoutData(code int, message string) *APIResponse {
	return &APIResponse{
		&Meta{
			StatusCode: code,
			Message:    message,
		},
		nil,
	}
}
