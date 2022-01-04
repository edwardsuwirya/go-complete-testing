package api

import (
	"database/sql"
	"enigmacamp.com/completetesting/config"
	"enigmacamp.com/completetesting/delivery"
	"enigmacamp.com/completetesting/manager"
	"log"
)

type Server interface {
	Run()
}

type server struct {
	config  *config.Config
	infra   manager.Infra
	usecase manager.UseCaseManager
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
	}
}

func (s *server) Run() {
	db := s.infra.SqlDb().DB
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	err := delivery.NewServer(s.config.RouterEngine, s.usecase)
	err = s.config.RouterEngine.Run(s.config.ApiBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
}
