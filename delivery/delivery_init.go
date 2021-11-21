package delivery

import (
	"enigmacamp.com/completetesting/manager"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Routes struct {
}

func (app Routes) StartGin() {
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")

	infraManager := manager.NewInfra()
	useCaseManager := manager.NewUseCaseManger(infraManager)
	defer func() {
		if err := infraManager.SqlDb().Close(); err != nil {
			panic(err)
		}
	}()
	r := gin.Default()
	publicRoute := r.Group("/api")

	NewStudentApi(publicRoute, useCaseManager.StudentUseCase())

	apiBaseUrl := fmt.Sprintf("%s:%s", host, port)
	err := r.Run(apiBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
}
