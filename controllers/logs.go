package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Kunal4now/logstorm/database"
	"github.com/Kunal4now/logstorm/model"
	"time"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetLogs(c *fiber.Ctx) error {
	level, tag := c.Query("level"), c.Query("tag")

	logs, err := database.GetLogs(level, tag)

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
	newLog := new(model.Log)
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