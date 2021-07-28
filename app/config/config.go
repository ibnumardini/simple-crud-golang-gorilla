package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	return
}
