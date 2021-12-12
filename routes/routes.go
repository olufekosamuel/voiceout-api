package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/olufekosamuel/voiceout-api/controllers"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"error": false,
			"msg":   "server is up and running",
		})
	})

	app.Post("/v1/admin/register", controllers.RegisterAdmin)
	app.Post("/v1/admin/login", controllers.LoginAdmin)

}
