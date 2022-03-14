package routes

import (
	"go-auth/controllers"

	jwtware "github.com/gofiber/jwt/v3"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	private := app.Group("/private")
	private.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	private.Get("/", controllers.Private)

	public := app.Group("/public")
	public.Get("/", controllers.Public)

	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}
