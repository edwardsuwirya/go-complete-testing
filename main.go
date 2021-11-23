package main

import (
	"enigmacamp.com/completetesting/config"
	"log"
)

func main() {
	appConfig := config.NewConfig()
	defer func() {
		if err := appConfig.InfraManager.SqlDb().Close(); err != nil {
			panic(err)
		}
	}()
	routeEngine := appConfig.Routes
	err := routeEngine.RouterEngine.Run(appConfig.ApiBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
}
