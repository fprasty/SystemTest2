package controllers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"systemtest2/database"
	"systemtest2/models"
	"systemtest2/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

//---------------------------- ADMIN AREA------------------------------//

// Get All Users data
func AdminGetAll(c *fiber.Ctx) error {
	var getuser []models.DynamicUser
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Function for admin login
func LoginAdmin(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	var user models.AdminDyUser
	database.DB.Where("name=?", data["name"]).First(&user)
	if user.Name == "" {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Akun tidak ada",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password salah",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(user.Id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	//Get env file
	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	ip := os.Getenv("IP")

	//Set Cookie
	cookie := fiber.Cookie{
		Name:     "DyAdmin-jwt",
		Domain:   ip,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "you have successfully login",
		"user":    user,
	})

}

func UpdateDyuser(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.DynamicUser
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Register>Unable to parse body")
	}
	//Cookie session
	cookie := c.Cookies("DyAdmin-jwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"messsage": "Cookie error",
		})

	}
	id, _ := strconv.Atoi(c.Params("id"))
	uparams := models.DynamicUser{
		Id: uint(id),
	}
	//Params id cek
	if uparams.Id != 0 {
		c.Status(400)
		return c.JSON("Id not match")
	}

	update := models.DynamicUser{
		Status: data["Status"].(string),
	}

	//UpdateUser
	err := database.DB.Model(&userData).Updates(&update)
	if err != nil {
		log.Println(err)
	}

	var UserData = make(map[string]interface{})
	UserData["id"] = update.Id
	UserData["Status"] = update.Status

	c.Status(200)
	return c.JSON(fiber.Map{
		"cookie":  id,
		"userid":  userData.Id,
		"user":    UserData,
		"message": "user updated successfully",
	})
}

// Function Delete Cookie
func DeleteCookie(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable parse body")
	}
	//Cookie Removed
	var user models.DynamicUser

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
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user":    user,
	})
}

// Fungsi Mengubah status
func SetStatus(c *fiber.Ctx) error {
	var user models.DynamicUser

	update := models.DynamicUser{
		Status: "Dijemput",
	}

	//UpdateUser
	database.DB.Model(&user).Find(&user).Updates(&update)

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Fungsi Menghapus Antrean/User
func DeleteUser(c *fiber.Ctx) error {
	var user models.DynamicUser
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

func RegisterAdmin(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Register>Unable to parse body")
	}

	user := models.AdminDyUser{
		Name: data["name"].(string),
	}

	//HashPassword
	user.SetPassword(data["password"].(string))
	//CreateUser
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"massage": "Akun admin berhasil dibuat",
	})
}
