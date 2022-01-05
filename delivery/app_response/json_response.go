package appresponse

import (
	"enigmacamp.com/completetesting/util/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponse struct {
	c      *gin.Context
	logger *logger.AppLogger
}

func (j *JsonResponse) SendData(message *ResponseMessage) {
	j.c.JSON(http.StatusOK, message)
}

func (j *JsonResponse) SendError(httpCode int, errMessage *ErrorMessage, err error) {
	j.logger.Log.Error().Err(err).Strs("ROUTE", []string{j.c.Request.Method, j.c.Request.URL.Path}).Int("CODE", httpCode).Msg(errMessage.ErrorDescription)
	j.c.AbortWithStatusJSON(httpCode, errMessage)
}

func NewJsonResponse(c *gin.Context, logger *logger.AppLogger) IAppHttpResponse {
	return &JsonResponse{c, logger}
}
