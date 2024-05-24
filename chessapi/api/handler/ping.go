package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Ping returns a 200 status code and nothing else
func Ping(c *fiber.Ctx) error {
	fmt.Println("chess api received ping")

	c.SendStatus(http.StatusOK)

	return nil
}