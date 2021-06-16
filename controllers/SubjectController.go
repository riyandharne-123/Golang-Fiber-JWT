package controllers

import (
	"../database"
	"../models"
	"github.com/gofiber/fiber"
)

func CreateSubject(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	subject := models.Subject{
		Name:      data["name"],
		StudentId: data["student_id"],
	}

	database.DB.Create(&subject)

	return c.JSON(subject)

}

func GetSubjects(c *fiber.Ctx) error {

	var subjects []models.Subject

	database.DB.Find(&subjects)

	return c.JSON(subjects)
}

func DeleteSubject(c *fiber.Ctx) error {

	var subject models.Subject

	database.DB.Where("id = ?", c.Params("id")).Delete(&subject)

	var subjects []models.Subject

	database.DB.Find(&subjects)

	return c.JSON(subjects)
}
