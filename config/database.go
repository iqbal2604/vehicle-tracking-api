package config

import (
	"github.com/iqbal2604/vehicle-tracking-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := "root@tcp(localhost:3306)/vehicle_tracking?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db

	DB.AutoMigrate(&models.User{}, &models.Vehicle{})
	return nil

}

func NewDatabase() *gorm.DB {
	return DB
}
