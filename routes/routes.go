package routes

import (
	"../controllers"
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

func Setup(app *fiber.App) {

	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	app.Get("/user", controllers.User)

	//student Crud
	app.Get("/students", controllers.GetStudents)
	app.Post("/students/create", controllers.CreateStudent)
	app.Get("/students/:id", controllers.GetStudent)
	app.Put("/students/update/:id", controllers.UpdateStudent)
	app.Delete("/students/delete/:id", controllers.DeleteStudent)

	//subject Crud
	app.Post("/subjects/create", controllers.CreateSubject)
	app.Delete("/subjects/delete/:id", controllers.DeleteSubject)
}
