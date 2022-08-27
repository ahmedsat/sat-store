package resources

import "database/sql"

var db *sql.DB

func SetDB(dbAddress *sql.DB) {
	db = dbAddress
}
