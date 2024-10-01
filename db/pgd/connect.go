package pgd

import (
	"context"
	"fmt"
	"log"

	"github.com/Rajendro1/Talenzen/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ReadConnectPostgresDB() (*pgxpool.Pool, error) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.ReadPostgresUser, config.ReadPostgresPassword, config.ReadPostgresHost, config.ReadPostgresPort, config.ReadPostgresDatabaseName)

	configConn, err := pgxpool.ParseConfig(psqlInfo)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
	}

	configConn.MaxConns = 1000

	// Establish a connection pool
	poolDb, err := pgxpool.NewWithConfig(context.Background(), configConn)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	// defer poolDb.Close()

	return poolDb, nil
}
func WriteConnectPostgresDB() (*pgxpool.Pool, error) {
	err := createPostgreSQLDatabase()
	if err != nil {
		fmt.Println("ConnectPostgresDB ", err)
	}
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.WritePostgresUser, config.WritePostgresPassword, config.WritePostgresHost, config.WritePostgresPort, config.WritePostgresDatabaseName)

	configConn, err := pgxpool.ParseConfig(psqlInfo)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
	}

	configConn.MaxConns = 1000

	// Establish a connection pool
	poolDb, err := pgxpool.NewWithConfig(context.Background(), configConn)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	// defer poolDb.Close()

	return poolDb, nil
}
func createPostgreSQLDatabase() error {
	// Connection string for the default "postgres" database
	defaultDBConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres", config.WritePostgresUser, config.WritePostgresPassword, config.WritePostgresHost, config.WritePostgresPort)

	// Establish a connection to the default "postgres" database
	cn, err := pgxpool.ParseConfig(defaultDBConnStr)
	if err != nil {
		return fmt.Errorf("unable to parse default database URL: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), cn)
	if err != nil {
		return fmt.Errorf("unable to create connection pool for default database: %v", err)
	}
	defer pool.Close()

	// Check if the target database exists
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s')", config.WritePostgresDatabaseName)
	err = pool.QueryRow(context.Background(), query).Scan(&exists)
	if err != nil {
		return fmt.Errorf("unable to check if database exists: %v", err)
	}

	// If the database does not exist, create it
	if !exists {
		_, err = pool.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %s", config.WritePostgresDatabaseName))
		if err != nil {
			return fmt.Errorf("unable to create database: %v", err)
		}
		fmt.Println("Database created successfully")
	} else {
		fmt.Println("Database already exists")
	}

	return nil
}
