package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type Config struct {
	RouterEngine   *gin.Engine
	DataSourceName string
	ApiBaseUrl     string
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

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	config.DataSourceName = dsn

	r := gin.Default()
	config.RouterEngine = r

	config.ApiBaseUrl = fmt.Sprintf("%s:%s", apiHost, apiPort)
	return config
}
