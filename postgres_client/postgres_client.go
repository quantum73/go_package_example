package postgres_client

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type AccountEmailAddressObject struct {
	Id, UserId        int
	Email             string
	Verified, Primary bool
}

func (emailAddr AccountEmailAddressObject) String() string {
	return fmt.Sprintf(
		"AccountEmailAddressObject[id=%d, email=%s, verified=%t, primary=%t, user_id=%d]",
		emailAddr.Id, emailAddr.Email, emailAddr.Verified, emailAddr.Primary, emailAddr.UserId,
	)
}

func ConnectToPostgres(
	host string, port int, user, password, dbname, sslmode string,
) (*sql.DB, error) {
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
