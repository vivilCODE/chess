package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vivilCODE/chess/db/handler"
)

func SetupRoutes(app *fiber.App, h *handler.Handler) {
	// Insert user into database
	app.Post("/db/users", h.InsertUser)

	// Get user by ID
	app.Get("/db/users/:id", h.GetUser)

	// Ping endpoint
	app.Get("/db/ping", h.Ping)

	// Insert new game into database
	app.Post("/db/games", h.InsertGame)
}