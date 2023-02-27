package controllers

import (
	"systemtest2/database"
	"systemtest2/models"

	"github.com/gofiber/fiber/v2"
)

//---------------------------- ADMIN AREA------------------------------//

// Get All Users data
// Front Office
func FOGetAll(c *fiber.Ctx) error {
	var getuser []models.FrontOffice
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Meeting Room
func MRGetAll(c *fiber.Ctx) error {
	var getuser []models.MeetingRoom
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Maruti
func MRTIGetAll(c *fiber.Ctx) error {
	var getuser []models.Maruti
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Rama Shinta
func RMGetAll(c *fiber.Ctx) error {
	var getuser []models.RamaSinta
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Kiskenda
func KKDGetAll(c *fiber.Ctx) error {
	var getuser []models.Kiskenda
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Maliawan
func MLGetAll(c *fiber.Ctx) error {
	var getuser []models.Maliawan
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// OCafe
func OCGetAll(c *fiber.Ctx) error {
	var getuser []models.OCafe
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Drop off hotel
func DOGetAll(c *fiber.Ctx) error {
	var getuser []models.DropHotel
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Villa
func VillaGetAll(c *fiber.Ctx) error {
	var getuser []models.Villa
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Karouke
func KRKGetAll(c *fiber.Ctx) error {
	var getuser []models.Karaoke
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Pasar Bandungan
func PBGetAll(c *fiber.Ctx) error {
	var getuser []models.PasarBandungan
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Tahu Om Shin
func OMGetAll(c *fiber.Ctx) error {
	var getuser []models.TahuOmShin
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}
