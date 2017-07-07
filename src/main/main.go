package main

import (
	"DAO"
	"Network"
)

func main() {

	var database = DAO.Start()

	Network.Start(database)
}
