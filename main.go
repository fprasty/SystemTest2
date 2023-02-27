package main

import (
	"log"
	"os"

	"systemtest2/database"
	"systemtest2/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Can't load .env file")
	}
	port := os.Getenv("PORT")
	ip := os.Getenv("IP")

	app := fiber.New()
	routes.Setup(app)
	app.Listen(ip + ":" + port)
}
