package main

import (
	"log"
	"github.com/gofiber/fiber/v2"

	"github.com/Kunal4now/logstorm/database"
	"github.com/Kunal4now/logstorm/router"
	"github.com/Kunal4now/logstorm/utils"
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

	app := fiber.New()

	router.SetupRouter(app)

	log.Fatal(app.Listen("127.0.0.1:3000"))

	sqlDB, _ := database.DB.DB()

	defer sqlDB.Close()
}
