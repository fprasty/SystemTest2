package controllers

import (
	"systemtest2/database"
	"systemtest2/models"

	"github.com/gofiber/fiber/v2"
)

// Fungsi Mengubah status
// -------------------- 10 Inside Location ---------------------->
// Front Office
func FOSetStatus(c *fiber.Ctx) error {
	//var data map[string]interface{}
	var data map[string]interface{}
	var userData models.FrontOffice
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	update := models.FrontOffice{
		Status:   data["status"].(string),
		UpdateAt: data["update_at"].(string),
	}

	//UpdateUser
	database.DB.Model(&userData).First(&userData).Updates(&update)

	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    userData,
		"message": "FO updated successfully",
	})
}

// Meeting Room
func MRSetStatus(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.MeetingRoom
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	update := models.MeetingRoom{
		Status:   data["status"].(string),
		UpdateAt: data["update_at"].(string),
	}

	//UpdateUser
	database.DB.Model(&userData).First(&userData).Updates(&update)

	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    userData,
		"message": "MO updated successfully",
	})
}

// Maruti
func MRTISetStatus(c *fiber.Ctx) error {
	var data map[string]interface{}
	var user models.Maruti
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	update := models.Maruti{
		Status:   data["status"].(string),
		UpdateAt: data["update_at"].(string),
	}

	//UpdateUser
	database.DB.Model(&user).Find(&user).Updates(&update)

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Rama Shinta
func RSSetStatus(c *fiber.Ctx) error {
	var data map[string]interface{}
	var user models.RamaSinta
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	update := models.RamaSinta{
		Status:   data["status"].(string),
		UpdateAt: data["update_at"].(string),
	}

	//UpdateUser
	database.DB.Model(&user).Find(&user).Updates(&update)

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Kiskenda
func KKDSetStatus(c *fiber.Ctx) error {
	var data map[string]interface{}
	var user models.Kiskenda
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	update := models.Kiskenda{
		Status:   data["status"].(string),
		UpdateAt: data["update_at"].(string),
	}

	//UpdateUser
	database.DB.Model(&user).Find(&user).Updates(&update)

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Maliawan
func MLISetStatus(c *fiber.Ctx) error {
	var data map[string]interface{}
	var user models.Maliawan
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	update := models.Maliawan{
		Status:   data["status"].(string),
		UpdateAt: data["update_at"].(string),
	}

	//UpdateUser
	database.DB.Model(&user).Find(&user).Updates(&update)

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// OCafe
func OCSetStatus(c *fiber.Ctx) error {
	var data map[string]interface{}
	var user models.OCafe
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	update := models.OCafe{
		Status:   data["status"].(string),
		UpdateAt: data["update_at"].(string),
	}

	//UpdateUser
	database.DB.Model(&user).Find(&user).Updates(&update)

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Drop Off Hotel
func DOSetStatus(c *fiber.Ctx) error {
	var data map[string]interface{}
	var user models.DropHotel
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	update := models.DropHotel{
		Status:   data["status"].(string),
		UpdateAt: data["update_at"].(string),
	}

	//UpdateUser
	database.DB.Model(&user).Find(&user).Updates(&update)

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Villa
func VillaSetStatus(c *fiber.Ctx) error {
	var data map[string]interface{}
	var user models.Villa
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	update := models.Villa{
		Status:   data["status"].(string),
		UpdateAt: data["update_at"].(string),
	}

	//UpdateUser
	database.DB.Model(&user).Find(&user).Updates(&update)

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Karouke
func KRKSetStatus(c *fiber.Ctx) error {
	var data map[string]interface{}
	var user models.Karaoke
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	update := models.Karaoke{
		Status:   data["status"].(string),
		UpdateAt: data["update_at"].(string),
	}

	//UpdateUser
	database.DB.Model(&user).Find(&user).Updates(&update)

	return c.JSON(fiber.Map{
		"user": user,
	})
}

//-------------------- 2 Outside Location ---------------------->

// Pasar Bandungan
func PBSetStatus(c *fiber.Ctx) error {
	var data map[string]interface{}
	var user models.PasarBandungan
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	update := models.PasarBandungan{
		Status:   data["status"].(string),
		UpdateAt: data["update_at"].(string),
	}

	//UpdateUser
	database.DB.Model(&user).Find(&user).Updates(&update)

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Tahu Om Shin
func TOSSetStatus(c *fiber.Ctx) error {
	var data map[string]interface{}
	var user models.TahuOmShin
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	update := models.TahuOmShin{
		Status:   data["status"].(string),
		UpdateAt: data["update_at"].(string),
	}

	//UpdateUser
	database.DB.Model(&user).Find(&user).Updates(&update)

	return c.JSON(fiber.Map{
		"user": user,
	})
}
