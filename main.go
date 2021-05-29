package main

import (
	"github.com/donreno/gofiber-test-api/routes"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
