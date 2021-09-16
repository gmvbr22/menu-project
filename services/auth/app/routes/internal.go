package routes

import "github.com/gofiber/fiber/v2"

func InternalRouter(app fiber.Router) {
	app.Post("/introspection", introspection())
}

func introspection() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"success": true,
		})
	}
}
