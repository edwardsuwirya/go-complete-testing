package delivery

import (
	"enigmacamp.com/completetesting/manager"
	"github.com/gin-gonic/gin"
)

func NewServer(engine *gin.Engine, useCaseManager manager.UseCaseManager) {
	publicRoute := engine.Group("/api")
	NewStudentApi(publicRoute, useCaseManager.StudentUseCase())
}
