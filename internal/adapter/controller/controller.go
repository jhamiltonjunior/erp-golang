package controller

type Response struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorMessage string `json:"error_message"`
	Data         any    `json:"data"`
}
