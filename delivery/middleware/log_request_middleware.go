package appMiddleware

import (
	"enigmacamp.com/completetesting/util/logger"
	"github.com/gin-gonic/gin"
)

type LogRequestMiddleware struct {
	logger *logger.AppLogger
}

func NewLogRequestMiddleware(logger *logger.AppLogger) *LogRequestMiddleware {
	return &LogRequestMiddleware{
		logger: logger,
	}
}

func (v *LogRequestMiddleware) Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		v.logger.Log.Info().Strs("ROUTE", []string{c.Request.Method, c.Request.URL.Path}).Msg("New Request")
	}
}
