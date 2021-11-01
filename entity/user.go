package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"name"`
	Email    string `gorm:"index,unique"`
	Password string `json:"password"`
}
