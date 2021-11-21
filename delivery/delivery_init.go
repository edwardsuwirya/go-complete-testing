package delivery

import (
	"fmt"
	"log"
	"os"

	"enigmacamp.com/completetesting/db"
	"github.com/gin-gonic/gin"
)

type Routes struct {
}

func (app Routes) StartGin() {
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")

	mdb, err := db.InitResource()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = mdb.Db.Close(); err != nil {
			panic(err)
		}
	}()
	r := gin.Default()
	publicRoute := r.Group("/api")
	NewStudentApi(publicRoute, mdb)

	apiBaseUrl := fmt.Sprintf("%s:%s", host, port)
	err = r.Run(apiBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
}
