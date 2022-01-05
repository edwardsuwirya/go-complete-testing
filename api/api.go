package api

import (
	"context"
	"database/sql"
	"enigmacamp.com/completetesting/config"
	"enigmacamp.com/completetesting/delivery"
	"enigmacamp.com/completetesting/manager"
	"enigmacamp.com/completetesting/util/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
			logger.Log.Fatal().Msg("Database Failed To Close")
		}
	}(db)
	err := delivery.NewServer(s.config.RouterEngine, s.usecase)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Server Failed To Run")
	}
	logger.Log.Info().Msg(fmt.Sprintf("Server Runs on %s", s.config.ApiBaseUrl))
	srv := &http.Server{
		Addr:    s.config.ApiBaseUrl,
		Handler: s.config.RouterEngine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal().Err(err).Msg("Server Failed To Run")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal().Err(err).Msg("Server Failed To Shutdown")
	}

	logger.Log.Info().Msg("Server Is Exiting")
}
