package delivery

import (
	"enigmacamp.com/completetesting/manager"
	"github.com/gin-gonic/gin"
)

func NewServer(engine *gin.Engine, useCaseManager manager.UseCaseManager) error {
	publicRoute := engine.Group("/api")
	_, err := NewStudentApi(publicRoute, useCaseManager.StudentUseCase())
	return err
}
