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
		logger.Log.Info().Strs("ROUTE", []string{c.Request.Method, c.Request.URL.Path}).Msg("New Request")
	}
}
