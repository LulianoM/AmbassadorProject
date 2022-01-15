package routes

import (
	"embassador/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("api")
	admin := api.Group("admin")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	admin.Post("register", controllers.Register)
}
