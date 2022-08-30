package services

import "database/sql"

var DB *sql.DB

func SetDB(dbAddress *sql.DB) {
	DB = dbAddress
}
