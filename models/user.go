package models

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Vehicles []Vehicle `gorm:"foreignKey:userID;references:ID"`
}
