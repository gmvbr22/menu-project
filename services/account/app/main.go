package main

import (
	"fmt"

	"github.com/gmvbr/menu-project/services/account/app/env"
	"github.com/gmvbr/menu-project/services/account/app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	port := env.GetEnv("PORT", "3000")

	routes.MainRouter(app)

	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
