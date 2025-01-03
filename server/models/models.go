package models

type User struct {
	USERID   uint   `gorm:"primaryKey"`
	USERNAME string `gorm:"unique"`
	PASSWORD string
}
