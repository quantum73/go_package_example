package main

import (
	"github.com/joho/godotenv"
	"github.com/quantum73/go_package_example/env"
	pg "github.com/quantum73/go_package_example/postgres_client"
	"log"
)

const envPath string = ".env"

func main() {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf("Error during loading .env file: %v\n", err)
		return
	}

	pgHost, err := env.GetRequiredEnvValue("DB_HOST")
	if err != nil {
		log.Fatalln(err)
	}
	dbPort, err := env.GetRequiredEnvValue("DB_PORT")
	if err != nil {
		log.Fatalln(err)
	}
	dbPortAsInt, err := env.ParseInt(dbPort)
	if err != nil {
		log.Fatalln(err)
	}
	dbUser, err := env.GetRequiredEnvValue("DB_USER")
	if err != nil {
		log.Fatalln(err)
	}
	dbPassword, err := env.GetRequiredEnvValue("DB_PASS")
	if err != nil {
		log.Fatalln(err)
	}
	dbName, err := env.GetRequiredEnvValue("DB_NAME")
	if err != nil {
		log.Fatalln(err)
	}
	dbSSLMode, err := env.GetRequiredEnvValue("DB_SSLMODE")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := pg.ConnectToPostgres(pgHost, dbPortAsInt, dbUser, dbPassword, dbName, dbSSLMode)
	defer db.Close()
	if err != nil {
		log.Printf("Error connecting to postgres: %v\n", err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Bad postgres ping: %v\n", err)
		return
	}

	log.Println("Successfully connected to postgres")
}
