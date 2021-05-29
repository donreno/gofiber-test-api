package routes

import (
	"github.com/donreno/gofiber-test-api/auth"
	"github.com/donreno/gofiber-test-api/wire"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(
		logger.New(),
		compress.New(),
		basicauth.New(basicauth.Config{Authorizer: auth.DoAuth}))

	api := app.Group("/products")

	handler := wire.MakeHandler()

	api.Get("/", handler.GetAll)
	api.Get("/:id", handler.Get)
	api.Post("/", handler.Update)
	api.Put("/", handler.Create)
	api.Delete("/", handler.Delete)
}
