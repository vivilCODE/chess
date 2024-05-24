package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vivilCODE/chess/chessapi/log"
)

// Ping returns a 200 status code and nothing else
func Ping(c *fiber.Ctx) error {

	log.Logger.Debug("chess api received ping request")

	c.SendStatus(http.StatusOK)

	return nil
}