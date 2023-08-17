package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber API!")
	})

	todoGroup := app.Group("/todos")
	SetupTodoRoutes(todoGroup)
}
