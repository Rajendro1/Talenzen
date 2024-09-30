package pgd

import (
	"database/sql"
	"log"

	"github.com/Rajendro1/Talenzen/config"
)

func SetupReadDatabase() *sql.DB {
	db, err := sql.Open("postgres", config.GetReadDatabaseConnectionString())
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Ensure the database is reachable
	if err := db.Ping(); err != nil {
		log.Fatalf("Could not ping the database: %v", err)
	}

	log.Println("Read Database connected!")
	return db
}
func SetupWriteDatabase() *sql.DB {
	db, err := sql.Open("postgres", config.GetWriteDatabaseConnectionString())
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Ensure the database is reachable
	if err := db.Ping(); err != nil {
		log.Fatalf("Could not ping the database: %v", err)
	}

	log.Println("WriteDatabase connected!")
	return db
}
