package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase(config Config) (*sql.DB, error) {
	return sql.Open("mysql", config.DSN)
}
