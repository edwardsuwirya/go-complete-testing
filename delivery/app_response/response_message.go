package appresponse

import (
	"enigmacamp.com/completetesting/util/app_status"
)

type ResponseMessage struct {
	Status      string      `json:"status"`
	Code        string      `json:"code"`
	Description string      `json:"message"`
	Data        interface{} `json:"data"`
}
type ErrorMessage struct {
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"message"`
}

func NewResponseMessage(code string, description string, data interface{}) *ResponseMessage {
	return &ResponseMessage{
		app_status.StatusText(app_status.Success), code, description, data,
	}
}

func NewErrorMessage(errCode string, message string) *ErrorMessage {
	em := &ErrorMessage{
		ErrorCode:        errCode,
		ErrorDescription: message,
	}
	return em
}
