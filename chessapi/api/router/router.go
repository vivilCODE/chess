package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/vivilCODE/chess/chessapi/api/config"
	"github.com/vivilCODE/chess/chessapi/api/handler"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(app *fiber.App, cfg config.Config) {
	api := app.Group("/chessapi")

	// Ping route
	ping := api.Group("/ping")
	ping.Get("/", handler.Ping)

	// Auth routes
	auth := 	api.Group("/auth")
	auth.Post("/signin", handler.NewSignInHandler(cfg.GapiClientID, cfg.GapiClientSecret, cfg.DBHandler))

	// User routes
	user := api.Group("/users")
	user.Get("/:id", handler.NewGetUserHandler(cfg.DBHandler))
	user.Post("/", handler.CreateUser)


	var chatroom = handler.NewChatroom()

	var gameQueue = handler.NewGameQueue()

	// Chatroom route
	chat := api.Group("/ws")
	chat.Use(RequireWebSocketConnection)
	chat.Get("/chatroom", websocket.New(chatroom.ChatroomHandler))
	chat.Get("/newgame", websocket.New(gameQueue.QueueHandler))

}



func RequireWebSocketConnection(c *fiber.Ctx) error{
	if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}

	return fiber.ErrUpgradeRequired
}