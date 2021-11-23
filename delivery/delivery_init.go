package delivery

import (
	"enigmacamp.com/completetesting/manager"
	"github.com/gin-gonic/gin"
)

type IDelivery interface {
	InitRouter(publicRoute *gin.RouterGroup)
}

type Routes struct {
	routers      []IDelivery
	RouterEngine *gin.Engine
	publicRoute  *gin.RouterGroup
}

func NewServer(useCaseManager manager.UseCaseManager) *Routes {
	newServer := new(Routes)

	r := gin.Default()
	publicRoute := r.Group("/api")
	routers := []IDelivery{
		NewStudentApi(useCaseManager.StudentUseCase()),
	}
	newServer.routers = routers
	newServer.RouterEngine = r
	newServer.publicRoute = publicRoute
	newServer.initAppRoutes()
	return newServer
}
func (app *Routes) initAppRoutes() {
	for _, rt := range app.routers {
		rt.InitRouter(app.publicRoute)
	}
}
