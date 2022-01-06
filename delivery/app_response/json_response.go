package appresponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponse struct {
	c *gin.Context
}

func (j *JsonResponse) SendData(message *ResponseMessage) {
	j.c.JSON(http.StatusOK, message)
}

func (j *JsonResponse) SendError(errMessage *ErrorMessage) {
	j.c.AbortWithStatusJSON(errMessage.HttpCode, errMessage.ErrorDescription)
}

func NewJsonResponse(c *gin.Context) IAppHttpResponse {
	return &JsonResponse{c}
}
