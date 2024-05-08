package controller

type ResponseSuccessMessage struct {
	Message string `json:"message"`
}

type ResponseErrorMessage struct {
	Error string `json:"error"`
}

func NewResponseSuccessMessage(message string) ResponseSuccessMessage {
	return ResponseSuccessMessage{
		Message: message,
	}
}

func NewResponseErrorMessage(err string) ResponseErrorMessage {
	return ResponseErrorMessage{
		Error: err,
	}
}
