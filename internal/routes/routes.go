package routes

import (
	"go-projects/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	userHandler := handler.NewUserHandler()

	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUser)
}
