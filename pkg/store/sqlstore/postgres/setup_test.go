package postgres

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testDB *sql.DB
)

const (
	postgresDNS = "postgres://root:password@localhost:5432/blog?sslmode=disable"
	driver      = "postgres"
)

func TestMain(m *testing.M) {
	dbConnection, err := sql.Open("postgres", postgresDNS)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConnection.Ping()
	if err != nil {
		log.Fatal(err)
	}

	testDB = dbConnection

	code := m.Run()

	err = testDB.Close()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}
