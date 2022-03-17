package main

import (
	"enigmacamp.com/completetesting/api"
)

// @title           Student API
// @version         1.0
// @description     This is a sample Go API Documentation.
// @termsOfService  http://enigmacamp.com/terms/

// @contact.name   Edo
// @contact.url    http://enigmacamp.com
// @contact.email  support@enigmacamp.com

// @host      localhost:8080
// @BasePath  /api
func main() {
	api.NewApiServer().Run()
}
