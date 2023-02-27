package controllers

import (
	"systemtest2/database"
	"systemtest2/models"

	"github.com/gofiber/fiber/v2"
)

func GetHistory(c *fiber.Ctx) error {
	var history []models.UserHistory

	database.DB.Find(&history)
	c.Status(200)
	return c.JSON(fiber.Map{
		"user": history,
	})
}
