package models

type Vehicle struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"index"`
	Name   string
	Plate  string `gorm:"unique"`
	Type   string
	Model  string
	Color  string

	User User `gorm:"foreignKey:UserID"`
}
