package router

import "github.com/gofiber/fiber/v2"


func NotFoundRoute(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "error",
				"message":   "sorry, endpoint is not found",
			})
		},
	)
}