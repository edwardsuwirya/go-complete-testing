package delivery

import (
	"enigmacamp.com/completetesting/manager"
	"fmt"
	"github.com/gin-gonic/gin"
)

type IDelivery interface {
	InitRouter(publicRoute *gin.RouterGroup)
}

type Routes struct {
	apiBaseUrl   string
	infraManager manager.Infra
	routers      []IDelivery
	routerEngine *gin.Engine
	publicRoute  *gin.RouterGroup
}

func NewServer(host string, port string) *Routes {
	apiBaseUrl := fmt.Sprintf("%s:%s", host, port)
	newServer := new(Routes)
	infraManager := manager.NewInfra()
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManger(repoManager)
	r := gin.Default()
	publicRoute := r.Group("/api")
	routers := []IDelivery{
		NewStudentApi(useCaseManager.StudentUseCase()),
	}
	newServer.infraManager = infraManager
	newServer.routers = routers
	newServer.apiBaseUrl = apiBaseUrl
	newServer.routerEngine = r
	newServer.publicRoute = publicRoute
	return newServer
}
func (app *Routes) StartEngine() (e error) {
	defer func() {
		if err := app.infraManager.SqlDb().Close(); err != nil {
			panic(err)
		}
	}()

	for _, rt := range app.routers {
		rt.InitRouter(app.publicRoute)
	}
	err := app.routerEngine.Run(app.apiBaseUrl)
	return err
}
