package delivery

import (
	"enigmacamp.com/completetesting/manager"
	"github.com/gin-gonic/gin"
)

type IDelivery interface {
	InitRouter(publicRoute *gin.RouterGroup)
}

type Routes struct {
	infraManager manager.Infra
	routers      []IDelivery
	routerEngine *gin.Engine
	publicRoute  *gin.RouterGroup
}

func NewServer(infraManager manager.Infra, useCaseManager manager.UseCaseManager) *Routes {
	newServer := new(Routes)

	r := gin.Default()
	publicRoute := r.Group("/api")
	routers := []IDelivery{
		NewStudentApi(useCaseManager.StudentUseCase()),
	}
	newServer.infraManager = infraManager
	newServer.routers = routers
	newServer.routerEngine = r
	newServer.publicRoute = publicRoute
	return newServer
}
func (app *Routes) StartEngine() (rt *gin.Engine) {
	for _, rt := range app.routers {
		rt.InitRouter(app.publicRoute)
	}
	return app.routerEngine
}
