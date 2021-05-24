package routes

import (
	"../controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {

	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
	app.Get("/user", controllers.User)
}
