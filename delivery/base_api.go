package delivery

import (
	appresponse "enigmacamp.com/completetesting/delivery/app_response"
	"enigmacamp.com/completetesting/util/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BaseApi struct {
}

func (api *BaseApi) errLogging(c *gin.Context, err error, errMsg *appresponse.ErrorMessage) {
	logger.ES(err, "DELIVERY", []string{c.Request.Method, c.Request.URL.Path, "CODE", strconv.Itoa(errMsg.HttpCode)}, errMsg.ErrorDescription.Description)
}
