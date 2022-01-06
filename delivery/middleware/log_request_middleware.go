package appMiddleware

import (
	"enigmacamp.com/completetesting/util/logger"
	"github.com/gin-gonic/gin"
)

type LogRequestMiddleware struct {
}

func NewLogRequestMiddleware() *LogRequestMiddleware {
	return &LogRequestMiddleware{}
}

func (v *LogRequestMiddleware) Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.IS("DELIVERY", []string{c.Request.Method, c.Request.URL.Path}, "New Request")
	}
}
