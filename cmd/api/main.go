package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	postgresDNS = "postgres://root:password@localhost:5432/blog?sslmode=disable"
	driver      = "postgres"
)

func main() {

	srv, err := setup()
	if err != nil {
		log.Fatal(err)
	}

	_, err = connectToDB(driver)
	if err != nil {
		log.Fatal(err)
	}

	if err = srv.run(":8000"); err != nil {
		log.Fatal(err)
	}
}

func setup() (*server, error) {
	srv := newServer()

	srv.setupRoutes()

	return srv, nil
}

func connectToDB(driver string) (*sql.DB, error) {
	db, err := sql.Open(driver, postgresDNS)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
