package db

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Resource struct {
	Db *sqlx.DB
}

func InitResource() (*Resource, error) {
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
	return &Resource{
		Db: conn,
	}, nil
}
