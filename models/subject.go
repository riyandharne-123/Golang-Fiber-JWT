package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	ID        uint
	StudentId string `json:"student_id"`
	Name      string `json:"name"`
}
