package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	dbDriver := "msql"
	dbUser := "root"
	dbPass := ""
	dbName := "patients"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	return db, err
}
