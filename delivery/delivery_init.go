package delivery

import (
	"enigmacamp.com/completetesting/manager"
	"enigmacamp.com/completetesting/util/logger"
	"github.com/gin-gonic/gin"
)

func NewServer(engine *gin.Engine, useCaseManager manager.UseCaseManager, logger *logger.AppLogger) error {
	publicRoute := engine.Group("/api")
	_, err := NewStudentApi(publicRoute, useCaseManager.StudentUseCase(), logger)
	return err
}
