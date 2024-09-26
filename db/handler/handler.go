package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vivilCODE/chess/db/dbservice"
	"github.com/vivilCODE/chess/db/log"
	"github.com/vivilCODE/chess/db/model"
)

type Handler struct {
	DBService *dbservice.DBService
}

func New(dbs *dbservice.DBService) *Handler{
	h := &Handler{
		DBService: dbs,
	}

	return h
}

func(h *Handler) GetUser(c *fiber.Ctx) error{
	id:= c.Params("id")	

	log.Logger.Debug("get user request received", "id", id)

	user, err :=h.DBService.GetUser(id)
	if err != nil {
		switch err {
		case dbservice.ErrorNoUserFound:
			log.Logger.Debug("no user found with", "id", id)

			c.Status(http.StatusNotFound)				
			return nil
		default:
			log.Logger.Error("getus")
			log.Logger.Debug("unexpected error in get user endpoint", "err", err)

			c.Status(http.StatusInternalServerError)
		}

		return err
	}

	log.Logger.Debug("found user", "user", user)

	c.JSON(user)

	return nil
}

func(h *Handler) InsertUser(c *fiber.Ctx) error{
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	
	// Do something with the user object
	log.Logger.Debug("postuser endpoint received user", "user", user)
	
	if err := h.DBService.InsertUser(user); err != nil {
		log.Logger.Error("unable to create user", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()}) 
	}

	return c.SendStatus(fiber.StatusOK)
}

func(h *Handler) InsertGame(c *fiber.Ctx) error{
		// NOT IMPLEMENTED
		
		id := c.Params("id")
		log.Logger.Debug("received request to insert new game", "id", id)
		return nil
}

func(h *Handler) Ping(c *fiber.Ctx) error{
	log.Logger.Debug("received ping request")

	c.SendString("pinged :)")
	return nil
}