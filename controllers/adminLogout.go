package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// Admin Logout 
func AdminLogout(c *fiber.Ctx) error {
	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "DyAdmin-jwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Succes Logout",
	})
}
