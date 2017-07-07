package Network

import (
	"Authorization"
	"log"
	"net/http"
	"database/sql"
)

var authExecutor *Authorization.Executor

func Start(database *sql.DB) {

	authExecutor = &Authorization.Executor{database}

	var router = newRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
