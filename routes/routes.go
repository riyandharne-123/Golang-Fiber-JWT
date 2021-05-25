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
}
