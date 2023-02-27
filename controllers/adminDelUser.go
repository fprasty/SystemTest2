package controllers

import (
	"os"
	"systemtest2/database"
	"systemtest2/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// Front Office
func FODeleteUser(c *fiber.Ctx) error {
	var user models.FrontOffice
	var history models.UserHistory

	database.DB.Where("Status=?", "Dijemput").Find(&user)
	migrate := models.UserHistory{
		Loc:       user.Loc,
		Status:    "Selesai",
		Cookie:    user.Cookie,
		Antre:     user.Antre,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}

	database.DB.Create(&history).Updates(&migrate)

	database.DB.Where("Status=?", "Dijemput").Find(&user)
	database.DB.Delete(&user)
	//Cookie Removed

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "Dyjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}
	c.Cookie(&cookie)
	c.ClearCookie("Dyjwt")

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Meeting Room
func MRDeleteUser(c *fiber.Ctx) error {
	var user models.MeetingRoom
	var history models.UserHistory

	database.DB.Where("Status=?", "Dijemput").Find(&user)
	migrate := models.UserHistory{
		Loc:       user.Loc,
		Status:    "Selesai",
		Cookie:    user.Cookie,
		Antre:     user.Antre,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}

	database.DB.Create(&history).Updates(&migrate)
	database.DB.Where("Status=?", "Dijemput").Find(&user)
	database.DB.Delete(&user)
	//Cookie Removed

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "Dyjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}
	c.Cookie(&cookie)
	c.ClearCookie("Dyjwt")

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Maruti
func MRTIDeleteUser(c *fiber.Ctx) error {
	var user models.Maruti
	var history models.UserHistory

	database.DB.Where("Status=?", "Dijemput").Find(&user)
	migrate := models.UserHistory{
		Loc:       user.Loc,
		Status:    "Selesai",
		Cookie:    user.Cookie,
		Antre:     user.Antre,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}

	database.DB.Create(&history).Updates(&migrate)
	database.DB.Where("Status=?", "Dijemput").Find(&user)
	database.DB.Delete(&user)
	//Cookie Removed

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "Dyjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}
	c.Cookie(&cookie)
	c.ClearCookie("Dyjwt")

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Rama Shinta
func RSDeleteUser(c *fiber.Ctx) error {
	var user models.RamaSinta
	var history models.UserHistory

	database.DB.Where("Status=?", "Dijemput").Find(&user)
	migrate := models.UserHistory{
		Loc:       user.Loc,
		Status:    "Selesai",
		Cookie:    user.Cookie,
		Antre:     user.Antre,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}

	database.DB.Create(&history).Updates(&migrate)
	database.DB.Where("Status=?", "Dijemput").Find(&user)
	database.DB.Delete(&user)
	//Cookie Removed

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "Dyjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}
	c.Cookie(&cookie)
	c.ClearCookie("Dyjwt")

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Kiskenda
func KKDDeleteUser(c *fiber.Ctx) error {
	var user models.Kiskenda
	var history models.UserHistory

	database.DB.Where("Status=?", "Dijemput").Find(&user)
	migrate := models.UserHistory{
		Loc:       user.Loc,
		Status:    "Selesai",
		Cookie:    user.Cookie,
		Antre:     user.Antre,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}

	database.DB.Create(&history).Updates(&migrate)
	database.DB.Where("Status=?", "Dijemput").Find(&user)
	database.DB.Delete(&user)
	//Cookie Removed

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "Dyjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}
	c.Cookie(&cookie)
	c.ClearCookie("Dyjwt")

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Maliawan
func MLIDeleteUser(c *fiber.Ctx) error {
	var user models.Maliawan
	var history models.UserHistory

	database.DB.Where("Status=?", "Dijemput").Find(&user)
	migrate := models.UserHistory{
		Loc:       user.Loc,
		Status:    "Selesai",
		Cookie:    user.Cookie,
		Antre:     user.Antre,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}

	database.DB.Create(&history).Updates(&migrate)
	database.DB.Where("Status=?", "Dijemput").Find(&user)
	database.DB.Delete(&user)
	//Cookie Removed

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "Dyjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}
	c.Cookie(&cookie)
	c.ClearCookie("Dyjwt")

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// OCafe
func OCDeleteUser(c *fiber.Ctx) error {
	var user models.OCafe
	var history models.UserHistory

	database.DB.Where("Status=?", "Dijemput").Find(&user)
	migrate := models.UserHistory{
		Loc:       user.Loc,
		Status:    "Selesai",
		Cookie:    user.Cookie,
		Antre:     user.Antre,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}

	database.DB.Create(&history).Updates(&migrate)
	database.DB.Where("Status=?", "Dijemput").Find(&user)
	database.DB.Delete(&user)
	//Cookie Removed

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "Dyjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}
	c.Cookie(&cookie)
	c.ClearCookie("Dyjwt")

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Drop office
func DODeleteUser(c *fiber.Ctx) error {
	var user models.DropHotel
	var history models.UserHistory

	database.DB.Where("Status=?", "Dijemput").Find(&user)
	migrate := models.UserHistory{
		Loc:       user.Loc,
		Status:    "Selesai",
		Cookie:    user.Cookie,
		Antre:     user.Antre,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}

	database.DB.Create(&history).Updates(&migrate)
	database.DB.Where("Status=?", "Dijemput").Find(&user)
	database.DB.Delete(&user)
	//Cookie Removed

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "Dyjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}
	c.Cookie(&cookie)
	c.ClearCookie("Dyjwt")

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Villa
func VLADeleteUser(c *fiber.Ctx) error {
	var user models.Villa
	var history models.UserHistory

	database.DB.Where("Status=?", "Dijemput").Find(&user)
	migrate := models.UserHistory{
		Loc:       user.Loc,
		Status:    "Selesai",
		Cookie:    user.Cookie,
		Antre:     user.Antre,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}

	database.DB.Create(&history).Updates(&migrate)
	database.DB.Where("Status=?", "Dijemput").Find(&user)
	database.DB.Delete(&user)
	//Cookie Removed

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "Dyjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}
	c.Cookie(&cookie)
	c.ClearCookie("Dyjwt")

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Karouke
func KRKDeleteUser(c *fiber.Ctx) error {
	var user models.Karaoke
	var history models.UserHistory

	database.DB.Where("Status=?", "Dijemput").Find(&user)
	migrate := models.UserHistory{
		Loc:       user.Loc,
		Status:    "Selesai",
		Cookie:    user.Cookie,
		Antre:     user.Antre,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}

	database.DB.Create(&history).Updates(&migrate)
	database.DB.Where("Status=?", "Dijemput").Find(&user)
	database.DB.Delete(&user)
	//Cookie Removed

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "Dyjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}
	c.Cookie(&cookie)
	c.ClearCookie("Dyjwt")

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Pasar Bandungan
func PBDeleteUser(c *fiber.Ctx) error {
	var user models.PasarBandungan
	var history models.UserHistory

	database.DB.Where("Status=?", "Dijemput").Find(&user)
	migrate := models.UserHistory{
		Loc:       user.Loc,
		Status:    "Selesai",
		Cookie:    user.Cookie,
		Antre:     user.Antre,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}

	database.DB.Create(&history).Updates(&migrate)
	database.DB.Where("Status=?", "Dijemput").Find(&user)
	database.DB.Delete(&user)
	//Cookie Removed

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "Dyjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}
	c.Cookie(&cookie)
	c.ClearCookie("Dyjwt")

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Tahu Om Shin
func TOSDeleteUser(c *fiber.Ctx) error {
	var user models.TahuOmShin
	var history models.UserHistory

	database.DB.Where("Status=?", "Dijemput").Find(&user)
	migrate := models.UserHistory{
		Loc:       user.Loc,
		Status:    "Selesai",
		Cookie:    user.Cookie,
		Antre:     user.Antre,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}

	database.DB.Create(&history).Updates(&migrate)
	database.DB.Where("Status=?", "Dijemput").Find(&user)
	database.DB.Delete(&user)
	//Cookie Removed

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	cookie := fiber.Cookie{
		Name:     "Dyjwt",
		Domain:   ip,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}
	c.Cookie(&cookie)
	c.ClearCookie("Dyjwt")

	return c.JSON(fiber.Map{
		"user": user,
	})
}
