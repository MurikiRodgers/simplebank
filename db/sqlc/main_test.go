package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")

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
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
