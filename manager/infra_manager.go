package manager

import (
	"enigmacamp.com/completetesting/config"
	"enigmacamp.com/completetesting/util/logger"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Infra interface {
	SqlDb() *sqlx.DB
	Logger() *logger.AppLogger
}

type infra struct {
	db     *sqlx.DB
	logger *logger.AppLogger
}

func NewInfra(config *config.Config) Infra {
	logger := config.AppLogger
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Database Failed To Start")
	}
	return &infra{
		db:     resource,
		logger: logger,
	}
}

func (i *infra) SqlDb() *sqlx.DB {
	return i.db
}
func (i *infra) Logger() *logger.AppLogger {
	return i.logger
}
func initDbResource(dataSourceName string) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
