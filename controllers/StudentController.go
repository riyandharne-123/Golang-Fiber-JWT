package controllers

import (
	"../database"
	"../models"
	"github.com/gofiber/fiber"
)

func CreateStudent(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	student := models.Student{
		Name:  data["name"],
		Email: data["email"],
	}

	database.DB.Create(&student)

	return c.JSON(student)

}

func GetStudents(c *fiber.Ctx) error {

	var students []models.Student

	database.DB.Find(&students)

	return c.JSON(students)
}

func GetStudent(c *fiber.Ctx) error {

	var student models.Student

	database.DB.Where("id = ?", c.Params("id")).First(&student)

	return c.JSON(student)
}

func UpdateStudent(c *fiber.Ctx) error {

	var student models.Student

	database.DB.Where("id = ?", c.Params("id")).First(&student)

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	student.Name = data["name"]
	student.Email = data["email"]

	database.DB.Save(&student)

	return c.JSON(student)
}

func DeleteStudent(c *fiber.Ctx) error {

	var student models.Student

	database.DB.Where("id = ?", c.Params("id")).Delete(&student)

	var students []models.Student

	database.DB.Find(&students)

	return c.JSON(students)
}
