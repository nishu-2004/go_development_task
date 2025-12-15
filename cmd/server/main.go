package main

import (
	"log"

	"go-projects/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Register all routes
	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
