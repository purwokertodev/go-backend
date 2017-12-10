package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

//postgresDB global var
var postgresDB *sql.DB

// GetPostgresDB function for creating database connection from Postgres Database
func GetPostgresDB() *sql.DB {
	if postgresDB == nil {
		postgresDB = CreatePostgresDBConnection(fmt.Sprintf("host=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_DB_HOST"), os.Getenv("POSTGRES_DB_USER"), os.Getenv("POSTGRES_DB_PASSWORD"), os.Getenv("POSTGRES_DB_NAME")))

	}
	return postgresDB
}

// CreatePostgresDBConnection function for creating database connection
func CreatePostgresDBConnection(descriptor string) *sql.DB {
	db, err := sql.Open("postgres", descriptor)
	if err != nil {
		defer db.Close()
		return db
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}
