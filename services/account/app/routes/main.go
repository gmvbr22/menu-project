package routes

import "github.com/gofiber/fiber/v2"

func MainRouter(app fiber.Router) {
	app.Get("/", hello())
}

func hello() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"success": true,
		})
	}
}
