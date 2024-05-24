package handler

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/vivilCODE/chess/chessapi/log"

	"github.com/gofiber/fiber/v2"
	"github.com/vivilCODE/chess/chessapi/api/model"
	"github.com/vivilCODE/chess/chessapi/dbhandler"
	"github.com/vivilCODE/chess/db/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type SignInHandler struct {
	GapiClientID     string
	GapiClientSecret string
	DBHandler *dbhandler.DBHandler
}

func NewSignInHandler(gapiID string, gapiSecret string, dbh *dbhandler.DBHandler) func(c *fiber.Ctx) error {
	h := &SignInHandler{
		GapiClientID:     gapiID,
		GapiClientSecret: gapiSecret,
		DBHandler: dbh,
	}

	return h.SignIn
}



func (h *SignInHandler) SignIn(c *fiber.Ctx) error {
	log.Logger.Debug("sign in requested")
	
	var req model.SignInRequest
	if err := c.BodyParser(&req); err != nil {
		log.Logger.Error("unable to parse sign in request body", "err", err)
	
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	context := context.Background()

	// Configuration for the oauth request, scopes decide what data we are asking for about the user
	cfg := &oauth2.Config{
		ClientID:     h.GapiClientID,
		ClientSecret: h.GapiClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		RedirectURL: "http://localhost:3000/signin",
		Endpoint:    google.Endpoint,
	}

	// Exchange the authorization code into a token
	token, err := cfg.Exchange(context, req.Code)
	if err != nil {
		log.Logger.Error("unable to exchange code", "err", err)
	
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	// Initialise oauth client with the configuration and the exchanged token
	client := cfg.Client(context, token)
	
	// Request user info
	res, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Logger.Error("unable to get google user info", "err", err)
	
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Internal server error"})
	}
	defer res.Body.Close()

	// Should define a new type called gapiUserInfo or something for parsing the gapi data,
	// becuase the user model is going to increase in size the more this application develops.
	// user will include friends, past games, rank, maybe messages etc.
	var userInfo models.User	

	// Decode json response into user struct
	if err := json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		log.Logger.Error("unable to decode user info", "err", err)
	
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Internal server error"}) 
	}


	userDoesNotExist := false

	// Try to fetch the same user from database
	dbUser, err := h.DBHandler.GetUser(userInfo.Id)
	if err != nil {
		log.Logger.Error("unable to fetch user from database", "err", err)
	
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Internal server error"}) 
	}

	// User has never signed in before, create new user in the db with this userinfo
	if dbUser == nil {
		userDoesNotExist = true
	}
	
	
	if userDoesNotExist {
		log.Logger.Debug("user has not signed in before, create new user")
		
		userInfo.SignedUp = time.Now()
		if err := h.DBHandler.PostUser(&userInfo); err != nil {
			log.Logger.Error("unable to create user", "err", err)
	
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Internal server error"}) 
		}
		
	}


	user := model.User{
		ID: userInfo.Id,
		Name: strings.Split(userInfo.Email, "@")[0],
		Email: userInfo.Email,
	}


	c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user})

	return nil
}