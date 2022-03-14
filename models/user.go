package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//Id       uint `gorm:"PrimaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password []byte `json:"-"`
}
