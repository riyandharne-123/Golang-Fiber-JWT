package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}
