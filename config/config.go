package config

import (
	appMiddleware "enigmacamp.com/completetesting/delivery/middleware"
	lgr "enigmacamp.com/completetesting/util/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type Config struct {
	RouterEngine   *gin.Engine
	DataSourceName string
	ApiBaseUrl     string
	AppLogger      *lgr.AppLogger
}

func NewConfig() *Config {
	config := new(Config)
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	isDebug := os.Getenv("DEBUG")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	config.DataSourceName = dsn

	if isDebug == "Y" || isDebug == "y" {
		gin.SetMode(gin.DebugMode)
		config.AppLogger = lgr.New(true)
	} else {
		gin.SetMode(gin.ReleaseMode)
		config.AppLogger = lgr.New(false)
	}
	r := gin.New()
	r.Use(appMiddleware.NewLogRequestMiddleware(config.AppLogger).Log())
	r.Use(gin.Recovery())
	config.RouterEngine = r

	config.ApiBaseUrl = fmt.Sprintf("%s:%s", apiHost, apiPort)
	return config
}
