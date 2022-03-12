package routes

import (
	"go-auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	app.Post("/register", controllers.Register)
}
