package main

import (
	"log"
	"github.com/Kunal4now/logstorm/packages/database"
	"github.com/Kunal4now/logstorm/packages/utils"
)

func main() {
	_, err := utils.LoadConfig(".")
	dbErr := database.InitDB()

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	if err != nil {
		log.Fatal(err)
	}

	//TODO: Add routes and start server
}