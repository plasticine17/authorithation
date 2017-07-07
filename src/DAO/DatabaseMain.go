package DAO

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Database *sql.DB

func Start() *sql.DB {
	Database = createDatabase()
	return Database
}

func Stop() {
	Database.Close()
}
