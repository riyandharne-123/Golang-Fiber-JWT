package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID       uint
	Name     string    `json:"name"`
	Email    string    `gorm:"unique" json:"email"`
	Subjects []Subject `gorm:"foreignKey:StudentId"`
}
