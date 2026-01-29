package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iqbal2604/vehicle-tracking-api/helper"
)

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/vehicle-tracking")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)

	log.Println(db, "Database Successfully CONNECTED")
	return db

}
