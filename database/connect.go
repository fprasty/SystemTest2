package database

import (
	"log"
	"os"

	"systemtest2/models"
	"systemtest2/salesmodels"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DB2 *gorm.DB

func Connect() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Can't load .env file")
	}
	//Database 1
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can't open database")
	}
	//Database 2
	dsn2 := os.Getenv("DSN2")
	database2, err := gorm.Open(mysql.Open(dsn2), &gorm.Config{})
	if err != nil {
		panic("Can't open database 2")
	}
	
	DB = database
	DB2 = database2

	database.AutoMigrate(
		&models.DynamicUser{},
		&models.AdminDyUser{},
		&models.FrontOffice{},
		&models.MeetingRoom{},
		&models.Maruti{},
		&models.RamaSinta{},
		&models.Kiskenda{},
		&models.Maliawan{},
		&models.OCafe{},
		&models.DropHotel{},
		&models.Villa{},
		&models.Karaoke{},
		//////
		&models.PasarBandungan{},
		&models.TahuOmShin{},
		&models.UserHistory{},
	)
	database2.AutoMigrate(
		&salesmodels.Sales{},
	)
}
