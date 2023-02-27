package controllers

import (
	"fmt"
	"os"
	"systemtest2/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// Function Delete Cookie
// Front Office
func FODeleteCookie(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable parse body")
	}
	//Cookie Removed
	var user models.FrontOffice

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "FOjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user":    user,
	})
}

// Meeting Room
func MRDeleteCookie(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable parse body")
	}
	//Cookie Removed
	var user models.MeetingRoom

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "MRjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user":    user,
	})
}

// Maruti
func MRTIDeleteCookie(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable parse body")
	}
	//Cookie Removed
	var user models.Maruti

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "MRTIjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user":    user,
	})
}

// Rama Shinta
func RSDeleteCookie(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable parse body")
	}
	//Cookie Removed
	var user models.RamaSinta

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "RSjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user":    user,
	})
}

// Kiskenda
func KKDDeleteCookie(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable parse body")
	}
	//Cookie Removed
	var user models.Kiskenda

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "KKDjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user":    user,
	})
}

// Maliawan
func MLIDeleteCookie(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable parse body")
	}
	//Cookie Removed
	var user models.Maliawan

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "MLIjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user":    user,
	})
}

// OCafe
func OCDeleteCookie(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable parse body")
	}
	//Cookie Removed
	var user models.OCafe

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "OCjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user":    user,
	})
}

// Drop Office
func DODeleteCookie(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable parse body")
	}
	//Cookie Removed
	var user models.DropHotel

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "DOjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user":    user,
	})
}

// Villa
func VLADeleteCookie(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable parse body")
	}
	//Cookie Removed
	var user models.Villa

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "Villajwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user":    user,
	})
}

// Karaoke
func KRKDeleteCookie(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable parse body")
	}
	//Cookie Removed
	var user models.Karaoke

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "KRjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user":    user,
	})
}

// Pasar Bandungan
func PBDeleteCookie(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable parse body")
	}
	//Cookie Removed
	var user models.PasarBandungan

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "PBjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user":    user,
	})
}

// Function Delete Cookie
func TOSDeleteCookie(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable parse body")
	}
	//Cookie Removed
	var user models.TahuOmShin

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "TOSjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user":    user,
	})
}
