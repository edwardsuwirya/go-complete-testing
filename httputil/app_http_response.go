package appresponse

type IAppHttpResponse interface {
	SendData(message *ResponseMessage)
	SendError(errMessage *ErrorMessage)
}
