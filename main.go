package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/MurikiRodgers/simplebank/api"
	db "github.com/MurikiRodgers/simplebank/db/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var serverAddress = "0.0.0.0:8080"

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error! Unable to find .env file")
	}

	dbSource := os.Getenv("DB_SOURCE")

	if dbSource == "" {
		log.Fatal("dbSource not set in .env")
	}

	dbDriver := os.Getenv("DB_DRIVER")

	if dbDriver == "" {
		log.Fatal("dbDriver not set in .env")
	}

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
