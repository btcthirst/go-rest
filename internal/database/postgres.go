package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Init(conn string) {

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to the database")
}
