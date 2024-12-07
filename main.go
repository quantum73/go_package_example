package main

import (
	"fmt"
	"log"

	"github.com/quantum73/go_package_example/formatter"
	"github.com/quantum73/go_package_example/math"
	pg "github.com/quantum73/go_package_example/postgres_client"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)

	db, err := pg.ConnectToPostgres(pg.HOST, pg.PORT, pg.USER, pg.PASSWORD, pg.DBNAME, pg.SSLMODE)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT * FROM account_emailaddress LIMIT 1")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var id, userId int
		var email string
		var verified, primary bool
		if err := rows.Scan(&id, &email, &verified, &primary, &userId); err != nil {
			log.Fatal(err)
		}

		obj := pg.AccountEmailAddressObject{
			Id:       id,
			UserId:   userId,
			Email:    email,
			Verified: verified,
			Primary:  primary,
		}
		fmt.Println(obj)
	}

	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}
}
