package controllers

import (
	"systemtest2/database"
	"systemtest2/models"

	"github.com/gofiber/fiber/v2"
)

//-----------------CREATE AllUser FUNCTION-------------------------------->
//-----------------10 Inside Location------------------->

// Front Office
func FOAllUser(c *fiber.Ctx) error {
	var getuser []models.FrontOffice
	var user models.FrontOffice

	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Waiting",
		})
	}

	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Meeting Room
func MRAllUser(c *fiber.Ctx) error {
	var getuser []models.MeetingRoom
	var user models.MeetingRoom

	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Waiting",
		})
	}

	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Maruti
func MarutiAllUser(c *fiber.Ctx) error {
	var getuser []models.Maruti
	var user models.Maruti

	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Waiting",
		})
	}

	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Rama Shinta
func RSAllUser(c *fiber.Ctx) error {
	var getuser []models.RamaSinta
	var user models.RamaSinta

	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Waiting",
		})
	}

	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Kiskenda
func KiskendaAllUser(c *fiber.Ctx) error {
	var getuser []models.Kiskenda
	var user models.Kiskenda

	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Waiting",
		})
	}

	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Maliawan
func MaliawanAllUser(c *fiber.Ctx) error {
	var getuser []models.Maliawan
	var user models.Maliawan

	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Waiting",
		})
	}

	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// OCafe
func OCafeAllUser(c *fiber.Ctx) error {
	var getuser []models.OCafe
	var user models.OCafe

	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Waiting",
		})
	}

	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Drop User
func DOAllUser(c *fiber.Ctx) error {
	var getuser []models.DropHotel
	var user models.DropHotel

	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Waiting",
		})
	}

	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Villa
func VillaAllUser(c *fiber.Ctx) error {
	var getuser []models.Villa
	var user models.Villa

	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Waiting",
		})
	}

	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Karouke
func KaraokeAllUser(c *fiber.Ctx) error {
	var getuser []models.Karaoke
	var user models.Karaoke

	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Waiting",
		})
	}

	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

//-----------------10 Outside Location------------------->

// Pasar Bandungan
func PBAllUser(c *fiber.Ctx) error {
	var getuser []models.PasarBandungan
	var user models.PasarBandungan

	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Waiting",
		})
	}

	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Tahu OmShin
func TOSAllUser(c *fiber.Ctx) error {
	var getuser []models.TahuOmShin
	var user models.TahuOmShin

	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Waiting",
		})
	}

	database.DB.Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}
