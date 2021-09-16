package routes

import "github.com/gofiber/fiber/v2"

func PublicRouter(app fiber.Router) {
	app.Post("/public/login", login())
	app.Post("/public/revoke", revoke())
}

func login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"success": true,
		})
	}
}

func revoke() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"success": true,
		})
	}
}
