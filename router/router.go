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

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error": "Route not found",
		})
	})
}