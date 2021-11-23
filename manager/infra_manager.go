package manager

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

type Infra interface {
	SqlDb() *sqlx.DB
}

type infra struct {
	db *sqlx.DB
}

func NewInfra(dataSourceName string) Infra {
	resource, err := initDbResource(dataSourceName)
	if err != nil {
		log.Panicln(err)
	}
	return &infra{
		db: resource,
	}
}

func (i *infra) SqlDb() *sqlx.DB {
	return i.db
}

func initDbResource(dataSourceName string) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
