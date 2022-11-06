package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouter(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/")
}
