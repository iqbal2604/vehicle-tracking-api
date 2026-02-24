package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/iqbal2604/vehicle-tracking-api/logs"
	"github.com/iqbal2604/vehicle-tracking-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	var db *gorm.DB
	var err error

	maxAttempts := 12
	wait := 5 * time.Second
	for i := 0; i < maxAttempts; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			DB = db
			DB.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.GPSLocation{}, &logs.Log{}, &models.TokenBlacklist{}, &models.Geofence{})
			return nil
		}
		log.Printf("Database not ready (attempt %d/%d): %v", i+1, maxAttempts, err)
		time.Sleep(wait)
	}

	return err

}

func NewDatabase() *gorm.DB {
	return DB
}
