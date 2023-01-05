package main

import (
	"log"
	"github.com/Kunal4now/logstorm/packages/database"
	"github.com/Kunal4now/logstorm/packages/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/Kunal4now/logstorm/packages/controllers"
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/logs", controllers.GetLogs)
	app.Get("/logs/:id", controllers.GetLog)
	app.Post("/logs", controllers.CreateLog)
	app.Delete("/logs/:id", controllers.DeleteLog)

	app.Listen(":3000")
}