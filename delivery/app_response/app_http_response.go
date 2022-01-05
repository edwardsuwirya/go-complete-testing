package appresponse

type IAppHttpResponse interface {
	SendData(message *ResponseMessage)
	SendError(httpCode int, errMessage *ErrorMessage, err error)
}
