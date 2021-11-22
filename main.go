package main

import (
	"enigmacamp.com/completetesting/delivery"
	"enigmacamp.com/completetesting/manager"
	"fmt"
	"log"
	"os"
)

func main() {
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")
	infraManager := manager.NewInfra()
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManger(repoManager)
	engine := delivery.NewServer(infraManager, useCaseManager).StartEngine()
	defer func() {
		if err := infraManager.SqlDb().Close(); err != nil {
			panic(err)
		}
	}()
	err := engine.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
	}
}
