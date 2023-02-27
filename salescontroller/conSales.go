package salescontroller

import (
	"errors"
	"fmt"
	"log"

	"math/rand"
	//"os"
	"strconv"
	"systemtest2/database"
	"systemtest2/salesmodels"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

//Upload image function

var letters = []rune("abcdefghijklmnopqrsuvwxyz")

func randLetter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RegisterSales(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}

	user := salesmodels.Sales{
		Username:  data["username"].(string),
		Name:      data["name"].(string),
		Telp:      data["notelp"].(string),
		Jabatan:   data["jabatan"].(string),
		Whatsapp:  data["whatsapp"].(string),
	}

	//CreateUser
	crt := database.DB2.Create(&user)
	if crt != nil {
		log.Println(crt)
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		//"user": userData,
		"massage": "Akun sales berhasil dibuat",
	})
}

func UpdateSales(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData salesmodels.Sales
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	update := salesmodels.Sales{
		Username:  data["username"].(string),
		Name:      data["name"].(string),
		Telp:      data["notelp"].(string),
		Jabatan:     data["jabatan"].(string),
		Whatsapp:  data["whatsapp"].(string),
	}

	id := c.Params("id")

	//UpdateUser
	database.DB2.Model(&userData).Where("id=?", id).Updates(&update)

	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    userData,
		"message": `Success` + userData.Name,
	})
}

func GetSales(c *fiber.Ctx) error {
	var userData []salesmodels.Sales
	var userDatas salesmodels.Sales
	username := c.Params("username")
	database.DB2.Model(&userDatas).Where("username=?", username).First(&userDatas)
	if userDatas.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "User tidak ada",
		})
	}
	database.DB2.Model(&userData).Where("username=?", username).First(&userData)
	return c.JSON(fiber.Map{
		"user": userData,
	})
}

func DeleteSales(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	sales := salesmodels.Sales{
		Id: uint(id),
	}

	deleteQuery := database.DB2.Delete(&sales)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Deleted",
	})
}

func UploadImage(c *fiber.Ctx) error {
	//ip := os.Getenv("IP")
	//port := os.Getenv("PORT")

	if err := godotenv.Load("config.env"); err != nil {
		panic("Error load env file")
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["avatar"]
	fileName := ""

	for _, file := range files {

		fileName = randLetter(15) + "-" + file.Filename
		if err := c.SaveFile(file, "./data/images/avatar/"+fileName); err != nil {
			return nil
		}
	}

	imageurl := "/api/uploads/" + fileName

	update := salesmodels.Sales{
		Avatar: imageurl,
	}
	var userData salesmodels.Sales
	crtupdt := database.DB2.Last(&userData).Updates(&update)
	if crtupdt != nil {
		log.Println(crtupdt)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Succes",
	})
}
