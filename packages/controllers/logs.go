package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Kunal4now/logstorm/packages/database"
	"time"
)

func GetLogs(c *fiber.Ctx) error {
	logs, err := database.GetLogs()

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err,
			"success": false,
		})
	}

	c.Status(200).JSON(&fiber.Map{
		"logs": logs,
		"success": true,
	})

	return nil
}

func CreateLog(c *fiber.Ctx) error {
	newLog := new(database.Log)
	err := c.BodyParser(newLog)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err,
			"success": false,
		})
	}

	newLog.Timestamp = time.Now()

	log, err := database.CreateLog(newLog.Level, newLog.Message, newLog.Timestamp, newLog.Tag, newLog.Data)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err,
			"success": false,
		})
	}

	c.Status(200).JSON(&fiber.Map{
		"log": log,
		"success": true,
	})

	return nil
}

func GetLog(c *fiber.Ctx) error {
	id := c.Params("id")

	log, err := database.GetLog(id)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err,
			"success": false,
		})
	}

	c.Status(200).JSON(&fiber.Map{
		"log": log,
		"success": true,
	})

	return nil
}

func DeleteLog(c *fiber.Ctx) error {
	id := c.Params("id")

	err := database.DeleteLog(id)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err,
			"success": false,
		})
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
	})

	return nil
}