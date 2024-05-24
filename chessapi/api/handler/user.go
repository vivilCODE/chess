package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vivilCODE/chess/chessapi/dbhandler"
)


type GetUserHandler struct {
	DBHandler *dbhandler.DBHandler
}

func NewGetUserHandler(dbh *dbhandler.DBHandler) func(c *fiber.Ctx) error {
	h := &GetUserHandler{
		DBHandler: dbh,
	}

	return h.GetUser
}

// GetUser retrieves a user by ID
func (h *GetUserHandler) GetUser(c *fiber.Ctx) error {
	// id := c.Params("id")

	// In a real application, you would fetch the user from the database
	// user := model.User{
	// 	ID:    id,
	// 	Name:  "John Doe",
	// 	Email: "johndoe@example.com",
	// }

	// return c.JSON(user)
		return nil
}

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
	// user := new(model.User)
	// if err := c.BodyParser(user); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	// }

	// In a real application, you would save the user to the database
	// user.ID = "123456"

	// return c.Status(fiber.StatusCreated).JSON(user)
		return nil
}