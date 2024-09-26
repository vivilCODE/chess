package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kouhin/envflag"
	"github.com/vivilCODE/chess/db/dbservice"
	"github.com/vivilCODE/chess/db/log"
	models "github.com/vivilCODE/chess/db/model"
)


func main() {
	dbPort := flag.Int("dbport", 5432, "Postgres address")
	host := flag.String("host", "localhost", "Where to host the http endpoint")


	if err := envflag.Parse(); err != nil {
		log.Logger.Fatal("unable to parse flags", "err", err)
	}

	err := godotenv.Load()
	if err != nil {
		log.Logger.Fatal("unable to load .env file", "err", err)		
	}

	// Initialise database service
	dbUser := os.Getenv("DB_USER")
	dbUserPassword := os.Getenv("DB_USER_PASSWORD")
	dbName := os.Getenv("DB_NAME")


	databaseConfig := dbservice.Config{
		DBPort:   *dbPort,
		DBName:   dbName,
		Host:     *host,
		User:     dbUser,
		Password: *&dbUserPassword,
	}

	service, err := dbservice.New(databaseConfig)
	if err != nil {
		log.Logger.Fatal("unable to initialise database service", "err", err)
		return
	}

	if err = service.Connect(); err != nil {
		log.Logger.Fatal("unable to connect to database", "err", err)
	}
	defer service.Disconnect()


	// Setup REST endpoints
	app := fiber.New()

	app.Get("/db/users", func(c *fiber.Ctx) error {
		return nil
	})
	
	app.Post("/db/users", func(c *fiber.Ctx) error {

		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	
		// Do something with the user object
		log.Logger.Debug("postuser endpoint received user", "user", user)
	
		if err := service.InsertUser(user); err != nil {
			log.Logger.Error("unable to create user", "err", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()}) 
		}

		return c.SendStatus(fiber.StatusOK)
		})

	app.Get("/db/users/:id", func(c *fiber.Ctx) error {
		id:= c.Params("id")	

		log.Logger.Debug("get user request received", "id", id)

		user, err :=service.GetUser(id)
		if err != nil {
			switch err {
			case dbservice.ErrorNoUserFound:
				log.Logger.Debug("no user found with", "id", id)

				c.Status(http.StatusNotFound)				
				return nil
			default:
				log.Logger.Error("getus")

				fmt.Printf("unexpected error in get user endpoint", "err",err) // debuglog

				c.Status(http.StatusInternalServerError)
			}

			return err
		}

		log.Logger.Debug("found user", "user", user)

		c.JSON(user)

		return nil
	})
	app.Get("/db/ping", func(c *fiber.Ctx) error {
		log.Logger.Debug("received ping request")

		c.SendString("pinged :)")
		return nil
	})

	app.Post("/db/games/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		log.Logger.Debug("received request to insert new game", "id", id)
		return nil
	})


	app.Listen(":8082")
}