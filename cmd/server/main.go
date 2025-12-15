package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	sqlc "go-projects/db/sqlc"
	dbconn "go-projects/internal/db"
	"go-projects/internal/handler"
	"go-projects/internal/logger"
	"go-projects/internal/middleware"
	"go-projects/internal/repository"
	"go-projects/internal/routes"
	"go-projects/internal/service"
)

func main() {
	// ✅ INIT LOGGER
	logger.Init()
	logger.Log.Info("🚀 server starting")

	app := fiber.New()

	// ✅ REGISTER MIDDLEWARE (BEFORE ROUTES)
	app.Use(middleware.RequestID())
	app.Use(middleware.Logger())

	// DB connection
	pool := dbconn.NewPostgresPool()
	queries := sqlc.New(pool)

	// Dependency wiring
	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	routes.RegisterRoutes(app, userHandler)

	log.Fatal(app.Listen(":8080"))
}
