package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/Kunal4now/logstorm/controllers"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Get("/", controllers.Hello)

	logs := api.Group("/log")

	logs.Get("/", controllers.GetLogs)
	logs.Get("/:id", controllers.GetLog)
	logs.Post("/", controllers.CreateLog)
	logs.Delete("/:id", controllers.DeleteLog)
}