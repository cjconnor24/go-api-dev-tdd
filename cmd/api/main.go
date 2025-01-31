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
	log.Println("Hello World")

	_, err := connectToDB(driver)
	if err != nil {
		log.Fatal(err)
	}
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
