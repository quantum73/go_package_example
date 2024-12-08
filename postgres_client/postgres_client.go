package postgres_client

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectToPostgres(host string, port int, user, password, dbname, sslmode string) (*sql.DB, error) {
	pgInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)
	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
