package api

import (
	"database/sql"
	"enigmacamp.com/completetesting/config"
	"enigmacamp.com/completetesting/delivery"
	"enigmacamp.com/completetesting/manager"
	"enigmacamp.com/completetesting/util/logger"
	"fmt"
)

type Server interface {
	Run()
}

type server struct {
	config  *config.Config
	infra   manager.Infra
	usecase manager.UseCaseManager
	logger  *logger.AppLogger
}

func NewApiServer() Server {
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUseCaseManger(repo)
	return &server{
		config:  appConfig,
		infra:   infra,
		usecase: usecase,
		logger:  appConfig.AppLogger,
	}
}

func (s *server) Run() {
	db := s.infra.SqlDb().DB
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			s.logger.Log.Fatal().Msg("Database Failed To Close")
		}
	}(db)
	err := delivery.NewServer(s.config.RouterEngine, s.usecase, s.logger)
	s.logger.Log.Info().Msg(fmt.Sprintf("Server Runs on %s", s.config.ApiBaseUrl))
	err = s.config.RouterEngine.Run(s.config.ApiBaseUrl)
	if err != nil {
		s.logger.Log.Fatal().Err(err).Msg("Server Failed To Run")
	}

}
