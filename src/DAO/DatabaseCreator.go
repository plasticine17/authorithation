package DAO

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"Utils"
)


func createDatabase() *sql.DB {
	var dbInfo = createDBInfo()

	var database, err = sql.Open(Utils.DriverName, dbInfo)

	if err != nil {
		log.Fatal(err)
	}

	return database
}

func createDBInfo() string {
	var format = "user=%s password=%s dbname=%s sslmode=disable"

	return fmt.Sprintf(format, Utils.DatabaseUser, Utils.DatabasePassword, Utils.DatabaseName)
}
