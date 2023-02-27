package controllers

import (
	"fmt"
	"log"
	"os"

	"strconv"
	"time"

	"systemtest2/database"
	"systemtest2/models"
	"systemtest2/util"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// Membuat user untuk antrean
func CreateDyuser(c *fiber.Ctx) error {
	var data map[string]interface{}
	//cookied := c.Cookies("Dyjwt")
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	user := models.DynamicUser{
		Status: data["Status"].(string),
	}

	//Step check  1
	//Check apakah Cokkied kosong?
	/*
		if cookied != "" {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "Mohon antre",
			})
		}
	*/

	//Step check  2
	//Check apakah Status "Menunggu" || "Dijemput"?
	var usercheck models.DynamicUser
	/*
		if err := database.DB.Where("id=?", 1).First(&usercheck); err != nil {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "Ada id",
			})
		}
	*/
	//Step check  3
	//Check apakah Status "Menunggu" || "Dijemput"?
	database.DB.Where("status=?", "Menunggu").First(&usercheck)
	database.DB.Where("status=?", "Dijemput").First(&usercheck)
	if usercheck.Status == "Menunggu" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	database.DB.Where("status=?", "Dijemput").First(&user)
	if usercheck.Status == "Dijemput" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	//Step check 4
	//Check apakah nama token kosong?
	/*
		if err := c.Cookies("Dyjwt"); err != "" {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "Mohon antre",
			})
		}
	*/

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	//Set Cookie

	cookie := fiber.Cookie{
		Name:     "Dyjwt",
		Domain:   "192.168.2.119",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//Create user baru
	db := database.DB.Create(&user)
	update := models.DynamicUser{
		Cookie: token,
	}

	//Change id to 1 if > 1 || < 1
	updateid := models.DynamicUser{
		Id: 1,
	}
	if usercheck.Id > 1 || usercheck.Id < 1 {
		database.DB.Model(&user).Updates(&updateid)
		println("Perubahahan berhasil, id adalah", usercheck.Id)
	}

	//Menambahkan token ke database "Cookie"
	database.DB.Model(&user).Updates(&update)
	if db != nil {
		log.Println(err)
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Create antrean succes",
		"user":    user,
	})

}

// Get All User datas
func AllDyuser(c *fiber.Ctx) error {
	var getuser []models.DynamicUser
	var user models.DynamicUser

	database.DB.Where("id=?", 1).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Waiting",
		})
	}

	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

// Get All User datas
func UserCek(c *fiber.Ctx) error {
	var getuser []models.DynamicUser
	var user models.DynamicUser

	cookies := c.Cookies("Dyjwt")
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

	/*
		cookie := fiber.Cookie{}
		if cookie.Value == "" {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "Tidak ada antrian",
			})
		}
	*/
	database.DB.Preload("User").Find(&getuser)
	return c.JSON(fiber.Map{
		"user": getuser,
	})
}

//----------------->
//----------------------------->
//--------------------------------------->

//-----------------CREATE ANTREAN FUNCTION-------------------------------->
//-----------------10 Inside Location------------------->

// Front Office
func FOCreate(c *fiber.Ctx) error {
	var data map[string]interface{}
	var usercheck models.FrontOffice
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	user := models.FrontOffice{
		Loc:       data["lokasi"].(string),
		Status:    data["status"].(string),
		Antre:     data["antre"].(string),
		CreatedAt: data["created_at"].(string),
	}

	//Step check  2
	//Check apakah Status "Menunggu" || "Dijemput"?

	//Step check  3
	//Check apakah Status "Menunggu" || "Dijemput"?
	database.DB.Where("status=?", "Menunggu").First(&usercheck)
	if usercheck.Status == "Menunggu" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}
	database.DB.Where("status=?", "Dijemput").First(&usercheck)
	if usercheck.Status == "Dijemput" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
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
		Name:     "FOjwt",
		Domain:   ip,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//Create user baru
	db := database.DB.Create(&user)
	update := models.FrontOffice{
		Cookie: token,
	}

	//Change id to 1 if > 1 || < 1
	updateid := models.FrontOffice{
		Id: 1,
	}
	if usercheck.Id > 1 || usercheck.Id < 1 {
		database.DB.Model(&user).Updates(&updateid)
		println("Perubahahan berhasil, id adalah", usercheck.Id)
	}

	//Menambahkan token ke database "Cookie"
	database.DB.Model(&user).Updates(&update)
	if db != nil {
		log.Println(err)
	}
	c.Cookie(&cookie)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Create antrean succes",
		"user":    user,
	})

}

// Meeting Room
func MRCreate(c *fiber.Ctx) error {
	var data map[string]interface{}
	var usercheck models.MeetingRoom

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	user := models.MeetingRoom{
		Loc:       data["lokasi"].(string),
		Status:    data["status"].(string),
		Antre:     data["antre"].(string),
		CreatedAt: data["created_at"].(string),
	}

	//Step check  2
	//Check apakah Status "Menunggu" || "Dijemput"?

	//Step check  3
	//Check apakah Status "Menunggu" || "Dijemput"?
	database.DB.Where("status=?", "Menunggu").First(&usercheck)
	if usercheck.Status == "Menunggu" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}
	database.DB.Where("status=?", "Dijemput").First(&usercheck)
	if usercheck.Status == "Dijemput" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
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
		Name:     "MRjwt",
		Domain:   ip,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//Create user baru
	db := database.DB.Create(&user)
	update := models.MeetingRoom{
		Cookie: token,
	}

	//Change id to 1 if > 1 || < 1
	updateid := models.MeetingRoom{
		Id: 1,
	}
	if usercheck.Id > 1 || usercheck.Id < 1 {
		database.DB.Model(&user).Updates(&updateid)
		println("Perubahahan berhasil, id adalah", usercheck.Id)
	}

	//Menambahkan token ke database "Cookie"
	database.DB.Model(&user).Updates(&update)
	if db != nil {
		log.Println(err)
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Create antrean succes",
		"user":    user,
	})

}

// Maruti
func MRTICreate(c *fiber.Ctx) error {
	var data map[string]interface{}
	var usercheck models.Maruti
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	user := models.Maruti{
		Loc:       data["lokasi"].(string),
		Status:    data["status"].(string),
		Antre:     data["antre"].(string),
		CreatedAt: data["created_at"].(string),
	}

	//Step check  2
	//Check apakah Status "Menunggu" || "Dijemput"?

	//Step check  3
	//Check apakah Status "Menunggu" || "Dijemput"?
	database.DB.Where("status=?", "Menunggu").First(&usercheck)
	if usercheck.Status == "Menunggu" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}
	database.DB.Where("status=?", "Dijemput").First(&usercheck)
	if usercheck.Status == "Dijemput" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
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
		Name:     "MRTIjwt",
		Domain:   ip,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//Create user baru
	db := database.DB.Create(&user)
	update := models.Maruti{
		Cookie: token,
	}

	//Change id to 1 if > 1 || < 1
	updateid := models.Maruti{
		Id: 1,
	}
	if usercheck.Id > 1 || usercheck.Id < 1 {
		database.DB.Model(&user).Updates(&updateid)
		println("Perubahahan berhasil, id adalah", usercheck.Id)
	}

	//Menambahkan token ke database "Cookie"
	database.DB.Model(&user).Updates(&update)
	if db != nil {
		log.Println(err)
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Create antrean succes",
		"user":    user,
	})

}

// Rama Shinta
func RSCreate(c *fiber.Ctx) error {
	var data map[string]interface{}
	var usercheck models.RamaSinta
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	user := models.RamaSinta{
		Loc:       data["lokasi"].(string),
		Status:    data["status"].(string),
		Antre:     data["antre"].(string),
		CreatedAt: data["created_at"].(string),
	}

	//Step check  2
	//Check apakah Status "Menunggu" || "Dijemput"?

	//Step check  3
	//Check apakah Status "Menunggu" || "Dijemput"?
	database.DB.Where("status=?", "Menunggu").First(&usercheck)
	if usercheck.Status == "Menunggu" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}
	database.DB.Where("status=?", "Dijemput").First(&usercheck)
	if usercheck.Status == "Dijemput" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
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
		Name:     "RSjwt",
		Domain:   ip,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//Create user baru
	db := database.DB.Create(&user)
	update := models.RamaSinta{
		Cookie: token,
	}

	//Change id to 1 if > 1 || < 1
	updateid := models.RamaSinta{
		Id: 1,
	}
	if usercheck.Id > 1 || usercheck.Id < 1 {
		database.DB.Model(&user).Updates(&updateid)
		println("Perubahahan berhasil, id adalah", usercheck.Id)
	}

	//Menambahkan token ke database "Cookie"
	database.DB.Model(&user).Updates(&update)
	if db != nil {
		log.Println(err)
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Create antrean succes",
		"user":    user,
	})

}

// Kiskenda
func KiskendaCreate(c *fiber.Ctx) error {
	var data map[string]interface{}
	var usercheck models.Kiskenda
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	user := models.Kiskenda{
		Loc:       data["lokasi"].(string),
		Status:    data["status"].(string),
		Antre:     data["antre"].(string),
		CreatedAt: data["created_at"].(string),
	}

	//Step check  2
	//Check apakah Status "Menunggu" || "Dijemput"?

	//Step check  3
	//Check apakah Status "Menunggu" || "Dijemput"?
	database.DB.Where("status=?", "Menunggu").First(&usercheck)
	if usercheck.Status == "Menunggu" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}
	database.DB.Where("status=?", "Dijemput").First(&usercheck)
	if usercheck.Status == "Dijemput" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
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
		Name:     "KKDjwt",
		Domain:   ip,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//Create user baru
	db := database.DB.Create(&user)
	update := models.Kiskenda{
		Cookie: token,
	}

	//Change id to 1 if > 1 || < 1
	updateid := models.Kiskenda{
		Id: 1,
	}
	if usercheck.Id > 1 || usercheck.Id < 1 {
		database.DB.Model(&user).Updates(&updateid)
		println("Perubahahan berhasil, id adalah", usercheck.Id)
	}

	//Menambahkan token ke database "Cookie"
	database.DB.Model(&user).Updates(&update)
	if db != nil {
		log.Println(err)
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Create antrean succes",
		"user":    user,
	})

}

// Maliawan
func MaliawanCreate(c *fiber.Ctx) error {
	var data map[string]interface{}
	var usercheck models.Maliawan
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	user := models.Maliawan{
		Loc:       data["lokasi"].(string),
		Status:    data["status"].(string),
		Antre:     data["antre"].(string),
		CreatedAt: data["created_at"].(string),
	}

	//Step check  2
	//Check apakah Status "Menunggu" || "Dijemput"?

	//Step check  3
	//Check apakah Status "Menunggu" || "Dijemput"?
	database.DB.Where("status=?", "Menunggu").First(&usercheck)
	if usercheck.Status == "Menunggu" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}
	database.DB.Where("status=?", "Dijemput").First(&usercheck)
	if usercheck.Status == "Dijemput" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
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
		Name:     "MLIjwt",
		Domain:   ip,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//Create user baru
	db := database.DB.Create(&user)
	update := models.Maliawan{
		Cookie: token,
	}

	//Change id to 1 if > 1 || < 1
	updateid := models.Maliawan{
		Id: 1,
	}
	if usercheck.Id > 1 || usercheck.Id < 1 {
		database.DB.Model(&user).Updates(&updateid)
		println("Perubahahan berhasil, id adalah", usercheck.Id)
	}

	//Menambahkan token ke database "Cookie"
	database.DB.Model(&user).Updates(&update)
	if db != nil {
		log.Println(err)
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Create antrean succes",
		"user":    user,
	})

}

// OCafe
func OCafeCreate(c *fiber.Ctx) error {
	var data map[string]interface{}
	var usercheck models.OCafe
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	user := models.OCafe{
		Loc:       data["lokasi"].(string),
		Status:    data["status"].(string),
		Antre:     data["antre"].(string),
		CreatedAt: data["created_at"].(string),
	}

	//Step check  2
	//Check apakah Status "Menunggu" || "Dijemput"?

	//Step check  3
	//Check apakah Status "Menunggu" || "Dijemput"?
	database.DB.Where("status=?", "Menunggu").First(&usercheck)
	if usercheck.Status == "Menunggu" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}
	database.DB.Where("status=?", "Dijemput").First(&usercheck)
	if usercheck.Status == "Dijemput" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
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
		Name:     "OCjwt",
		Domain:   ip,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//Create user baru
	db := database.DB.Create(&user)
	update := models.OCafe{
		Cookie: token,
	}

	//Change id to 1 if > 1 || < 1
	updateid := models.OCafe{
		Id: 1,
	}
	if usercheck.Id > 1 || usercheck.Id < 1 {
		database.DB.Model(&user).Updates(&updateid)
		println("Perubahahan berhasil, id adalah", usercheck.Id)
	}

	//Menambahkan token ke database "Cookie"
	database.DB.Model(&user).Updates(&update)
	if db != nil {
		log.Println(err)
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Create antrean succes",
		"user":    user,
	})

}

// Drop Off Hotel
func DOCreate(c *fiber.Ctx) error {
	var data map[string]interface{}
	var usercheck models.DropHotel
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	user := models.DropHotel{
		Loc:       data["lokasi"].(string),
		Status:    data["status"].(string),
		Antre:     data["antre"].(string),
		CreatedAt: data["created_at"].(string),
	}

	//Step check  2
	//Check apakah Status "Menunggu" || "Dijemput"?

	//Step check  3
	//Check apakah Status "Menunggu" || "Dijemput"?
	database.DB.Where("status=?", "Menunggu").First(&usercheck)
	if usercheck.Status == "Menunggu" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}
	database.DB.Where("status=?", "Dijemput").First(&usercheck)
	if usercheck.Status == "Dijemput" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
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
		Name:     "DOjwt",
		Domain:   ip,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//Create user baru
	db := database.DB.Create(&user)
	update := models.DropHotel{
		Cookie: token,
	}

	//Change id to 1 if > 1 || < 1
	updateid := models.DropHotel{
		Id: 1,
	}
	if usercheck.Id > 1 || usercheck.Id < 1 {
		database.DB.Model(&user).Updates(&updateid)
		println("Perubahahan berhasil, id adalah", usercheck.Id)
	}

	//Menambahkan token ke database "Cookie"
	database.DB.Model(&user).Updates(&update)
	if db != nil {
		log.Println(err)
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Create antrean succes",
		"user":    user,
	})

}

// Villa
func VillaCreate(c *fiber.Ctx) error {
	var data map[string]interface{}
	var usercheck models.Villa
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	user := models.Villa{
		Loc:       data["lokasi"].(string),
		Status:    data["status"].(string),
		Antre:     data["antre"].(string),
		CreatedAt: data["created_at"].(string),
	}

	//Step check  2
	//Check apakah Status "Menunggu" || "Dijemput"?

	//Step check  3
	//Check apakah Status "Menunggu" || "Dijemput"?
	database.DB.Where("status=?", "Menunggu").First(&usercheck)
	if usercheck.Status == "Menunggu" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}
	database.DB.Where("status=?", "Dijemput").First(&usercheck)
	if usercheck.Status == "Dijemput" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
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
		Name:     "Villajwt",
		Domain:   ip,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//Create user baru
	db := database.DB.Create(&user)
	update := models.Villa{
		Cookie: token,
	}

	//Change id to 1 if > 1 || < 1
	updateid := models.Villa{
		Id: 1,
	}
	if usercheck.Id > 1 || usercheck.Id < 1 {
		database.DB.Model(&user).Updates(&updateid)
		println("Perubahahan berhasil, id adalah", usercheck.Id)
	}

	//Menambahkan token ke database "Cookie"
	database.DB.Model(&user).Updates(&update)
	if db != nil {
		log.Println(err)
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Create antrean succes",
		"user":    user,
	})

}

// Karaoke
func KaraokeCreate(c *fiber.Ctx) error {
	var data map[string]interface{}
	var usercheck models.Karaoke
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	user := models.Karaoke{
		Loc:       data["lokasi"].(string),
		Status:    data["status"].(string),
		Antre:     data["antre"].(string),
		CreatedAt: data["created_at"].(string),
	}

	//Step check  2
	//Check apakah Status "Menunggu" || "Dijemput"?

	//Step check  3
	//Check apakah Status "Menunggu" || "Dijemput"?
	database.DB.Where("status=?", "Menunggu").First(&usercheck)
	if usercheck.Status == "Menunggu" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}
	database.DB.Where("status=?", "Dijemput").First(&usercheck)
	if usercheck.Status == "Dijemput" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
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
		Name:     "KRjwt",
		Domain:   ip,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//Create user baru
	db := database.DB.Create(&user)
	update := models.Karaoke{
		Cookie: token,
	}

	//Change id to 1 if > 1 || < 1
	updateid := models.Karaoke{
		Id: 1,
	}
	if usercheck.Id > 1 || usercheck.Id < 1 {
		database.DB.Model(&user).Updates(&updateid)
		println("Perubahahan berhasil, id adalah", usercheck.Id)
	}

	//Menambahkan token ke database "Cookie"
	database.DB.Model(&user).Updates(&update)
	if db != nil {
		log.Println(err)
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Create antrean succes",
		"user":    user,
	})

}

// -----------------2 Outside Location------------------->
// Pasar Bandungan
func PBCreate(c *fiber.Ctx) error {
	var data map[string]interface{}
	var usercheck models.PasarBandungan
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	user := models.PasarBandungan{
		Loc:       data["lokasi"].(string),
		Status:    data["status"].(string),
		Antre:     data["antre"].(string),
		CreatedAt: data["created_at"].(string),
	}

	//Step check  2
	//Check apakah Status "Menunggu" || "Dijemput"?

	//Step check  3
	//Check apakah Status "Menunggu" || "Dijemput"?
	database.DB.Where("status=?", "Menunggu").First(&usercheck)
	if usercheck.Status == "Menunggu" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}
	database.DB.Where("status=?", "Dijemput").First(&usercheck)
	if usercheck.Status == "Dijemput" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
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
		Name:     "PBjwt",
		Domain:   ip,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//Create user baru
	db := database.DB.Create(&user)
	update := models.PasarBandungan{
		Cookie: token,
	}

	//Change id to 1 if > 1 || < 1
	updateid := models.PasarBandungan{
		Id: 1,
	}
	if usercheck.Id > 1 || usercheck.Id < 1 {
		database.DB.Model(&user).Updates(&updateid)
		println("Perubahahan berhasil, id adalah", usercheck.Id)
	}

	//Menambahkan token ke database "Cookie"
	database.DB.Model(&user).Updates(&update)
	if db != nil {
		log.Println(err)
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Create antrean succes",
		"user":    user,
	})

}

// Tahu Om Shin
func TOCreate(c *fiber.Ctx) error {
	var data map[string]interface{}
	var usercheck models.TahuOmShin
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	user := models.TahuOmShin{
		Loc:       data["lokasi"].(string),
		Status:    data["status"].(string),
		Antre:     data["antre"].(string),
		CreatedAt: data["created_at"].(string),
	}

	//Step check  2
	//Check apakah Status "Menunggu" || "Dijemput"?

	//Step check  3
	//Check apakah Status "Menunggu" || "Dijemput"?
	database.DB.Where("status=?", "Menunggu").First(&usercheck)
	if usercheck.Status == "Menunggu" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}
	database.DB.Where("status=?", "Dijemput").First(&usercheck)
	if usercheck.Status == "Dijemput" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon antre",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
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
		Name:     "TOSjwt",
		Domain:   ip,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//Create user baru
	db := database.DB.Create(&user)
	update := models.TahuOmShin{
		Cookie: token,
	}

	//Change id to 1 if > 1 || < 1
	updateid := models.TahuOmShin{
		Id: 1,
	}
	if usercheck.Id > 1 || usercheck.Id < 1 {
		database.DB.Model(&user).Updates(&updateid)
		println("Perubahahan berhasil, id adalah", usercheck.Id)
	}

	//Menambahkan token ke database "Cookie"
	database.DB.Model(&user).Updates(&update)
	if db != nil {
		log.Println(err)
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Create antrean succes",
		"user":    user,
	})

}
