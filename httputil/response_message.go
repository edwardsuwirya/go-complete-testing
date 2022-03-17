package appresponse

type ResponseMessage struct {
	Status      string      `json:"status"`
	Code        string      `json:"code"`
	Description string      `json:"message"`
	Data        interface{} `json:"data"`
}

type ErrorMessage struct {
	HttpCode         int
	ErrorDescription ErrorDescription
}
type ErrorDescription struct {
	Code        string `json:"errorCode"`
	Description string `json:"message"`
}

func NewResponseMessage(code string, description string, data interface{}) *ResponseMessage {
	return &ResponseMessage{
		"Success", code, description, data,
	}
}

func NewErrorMessage(httpCode int, errCode string, message string) *ErrorMessage {
	em := &ErrorMessage{
		HttpCode: httpCode,
		ErrorDescription: ErrorDescription{
			Code:        errCode,
			Description: message,
		},
	}
	return em
}
