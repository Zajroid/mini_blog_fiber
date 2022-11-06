package main

import (
	"github.com/gofiber/fiber/v2"
	"learnFiber2/internal/router"
	"log"
)

var ()

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	router.SetupRouter(app)

	log.Fatal(app.Listen(":3000"))
}
