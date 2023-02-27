package controllers

import (
	"systemtest2/database"
	"systemtest2/models"

	"github.com/gofiber/fiber/v2"
)

//-----------------CREATE USERCEK FUNCTION-------------------------------->
//-----------------10 Inside Location------------------->

// Front Office
func FOUserCek(c *fiber.Ctx) error {
	var getuser []models.FrontOffice
	var user models.FrontOffice

	cookies := c.Cookies("FOjwt")
	database.DB.Where("id=?", 1).First(&user)

	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Tidak ada antrian",
		})
	}

	if cookies != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Anda sudah antri",
		})
	}

	if user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sedang ada antrian",
		})
	}

	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Meeting Room
func MRUserCek(c *fiber.Ctx) error {
	var getuser []models.MeetingRoom
	var user models.MeetingRoom

	cookies := c.Cookies("MRjwt")
	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Tidak ada antrian",
		})
	}

	if cookies != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Anda sudah antri",
		})
	}

	database.DB.Where("id=?", 1).First(&user)
	if user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sedang ada antrian",
		})
	}
	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Maruti
func MRTIUserCek(c *fiber.Ctx) error {
	var getuser []models.Maruti
	var user models.Maruti

	cookies := c.Cookies("MRTIjwt")
	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Tidak ada antrian",
		})
	}

	if cookies != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Anda sudah antri",
		})
	}

	database.DB.Where("id=?", 1).First(&user)
	if user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sedang ada antrian",
		})
	}
	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Rama Shinta
func RamaSintaUserCek(c *fiber.Ctx) error {
	var getuser []models.RamaSinta
	var user models.RamaSinta

	cookies := c.Cookies("RSjwt")
	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Tidak ada antrian",
		})
	}

	if cookies != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Anda sudah antri",
		})
	}

	database.DB.Where("id=?", 1).First(&user)
	if user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sedang ada antrian",
		})
	}
	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Kiskenda
func KiskendaUserCek(c *fiber.Ctx) error {
	var getuser []models.Kiskenda
	var user models.Kiskenda

	cookies := c.Cookies("KKDjwt")
	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Tidak ada antrian",
		})
	}

	if cookies != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Anda sudah antri",
		})
	}

	database.DB.Where("id=?", 1).First(&user)
	if user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sedang ada antrian",
		})
	}
	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Maliawan
func MaliawanUserCek(c *fiber.Ctx) error {
	var getuser []models.Maliawan
	var user models.Maliawan

	cookies := c.Cookies("MLIjwt")
	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Tidak ada antrian",
		})
	}

	if cookies != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Anda sudah antri",
		})
	}

	database.DB.Where("id=?", 1).First(&user)
	if user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sedang ada antrian",
		})
	}
	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// O Cafe
func OCafeUserCek(c *fiber.Ctx) error {
	var getuser []models.OCafe
	var user models.OCafe

	cookies := c.Cookies("OCjwt")
	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Tidak ada antrian",
		})
	}

	if cookies != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Anda sudah antri",
		})
	}

	database.DB.Where("id=?", 1).First(&user)
	if user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sedang ada antrian",
		})
	}
	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Drop Off Hotel
func DOUserCek(c *fiber.Ctx) error {
	var getuser []models.DropHotel
	var user models.DropHotel

	cookies := c.Cookies("DOjwt")
	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Tidak ada antrian",
		})
	}

	if cookies != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Anda sudah antri",
		})
	}

	database.DB.Where("id=?", 1).First(&user)
	if user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sedang ada antrian",
		})
	}
	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Villa
func VillaUserCek(c *fiber.Ctx) error {
	var getuser []models.Villa
	var user models.Villa

	cookies := c.Cookies("Villajwt")
	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Tidak ada antrian",
		})
	}

	if cookies != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Anda sudah antri",
		})
	}

	database.DB.Where("id=?", 1).First(&user)
	if user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sedang ada antrian",
		})
	}
	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Karouke
func KaraokeUserCek(c *fiber.Ctx) error {
	var getuser []models.Karaoke
	var user models.Karaoke

	cookies := c.Cookies("KRjwt")
	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Tidak ada antrian",
		})
	}

	if cookies != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Anda sudah antri",
		})
	}

	database.DB.Where("id=?", 1).First(&user)
	if user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sedang ada antrian",
		})
	}
	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Pasar Bandungan
func PBUserCek(c *fiber.Ctx) error {
	var getuser []models.PasarBandungan
	var user models.PasarBandungan

	cookies := c.Cookies("PBjwt")
	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Tidak ada antrian",
		})
	}

	if cookies != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Anda sudah antri",
		})
	}

	database.DB.Where("id=?", 1).First(&user)
	if user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sedang ada antrian",
		})
	}
	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Tahu OmShin
func TOSUserCek(c *fiber.Ctx) error {
	var getuser []models.TahuOmShin
	var user models.TahuOmShin

	cookies := c.Cookies("TOSjwt")
	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Tidak ada antrian",
		})
	}

	if cookies != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Anda sudah antri",
		})
	}

	database.DB.Where("id=?", 1).First(&user)
	if user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sedang ada antrian",
		})
	}
	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}
