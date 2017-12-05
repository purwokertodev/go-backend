package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var mssqlDB *sql.DB

// GetMsSQLDB function for creating database connection from MSSQL Database
func GetMsSQLDB() *sql.DB {
	var port int
	port = 1433

	if mssqlDB == nil {
		mssqlDB = CreateMsSQLDBConnection(fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
			os.Getenv("MSSQL_DB_HOST"), os.Getenv("MSSQL_DB_USER"), os.Getenv("MSSQL_DB_PASSWORD"), port, os.Getenv("MSSQL_DB_NAME")))
	}
	return mssqlDB
}

// CreateMSSQLDBConnection function for creating  database connection
func CreateMsSQLDBConnection(descriptor string) *sql.DB {
	db, err := sql.Open("mssql", descriptor)
	if err != nil {
		defer db.Close()
		return db
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}
