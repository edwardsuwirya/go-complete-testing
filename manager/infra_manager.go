package manager

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Infra interface {
	SqlDb() *sqlx.DB
}

type infra struct {
	db *sqlx.DB
}

func NewInfra() Infra {
	resource, err := initDbResource()
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

func initDbResource() (*sqlx.DB, error) {
	host := os.Getenv("PSQL_HOST")
	port := os.Getenv("PSQL_PORT")
	dbName := os.Getenv("PSQL_DBNAME")
	dbUser := os.Getenv("PSQL_USER")
	dbPassword := os.Getenv("PSQL_PASSWD")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, host, port, dbName)

	conn, err := sqlx.Connect("pgx", url)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
