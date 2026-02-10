package config

import (
	"fmt"
	"os"

	"github.com/iqbal2604/vehicle-tracking-api/logs"
	models "github.com/iqbal2604/vehicle-tracking-api/models/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db

	DB.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.GPSLocation{}, &logs.Log{})
	return nil

}

func NewDatabase() *gorm.DB {
	return DB
}
