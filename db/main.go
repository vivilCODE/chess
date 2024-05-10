package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kouhin/envflag"
	"github.com/vivilCODE/chess/db/dbservice"
	"github.com/vivilCODE/chess/db/models"
)


func main() {
	dbPort := flag.Int("dbport", 5432, "Postgres address")
	host := flag.String("host", "localhost", "Where to host the http endpoint")


	if err := envflag.Parse(); err != nil {
		log.Fatalf("unable to parse flags, %v", err)
	}

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Failed to load .env file: %v", err)
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
		log.Fatalf("unable to initialise database service: %v", err)
		return
	}

	if err = service.Connect(); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
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
		fmt.Println("POSTUSER ENDPOINT: Received user:", user)

		if err := service.CreateUser(user); err != nil {
			fmt.Printf("POSTUSER ENDPOINT: error creating user: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()}) 
		}

		return c.SendStatus(fiber.StatusOK)
		})

	app.Get("/db/users/:id", func(c *fiber.Ctx) error {
		id:= c.Params("id")	

		fmt.Println("GETUSER ENDPOINT: get user request received, id:", id)
		

		user, err :=service.GetUser(id)
		if err != nil {
			switch err {
			case dbservice.ErrorNoUserFound:
				fmt.Printf("GETUSER ENDPOINT: no user found with id: %s\n", id) // debuglog

				c.Status(http.StatusNotFound)				
				return nil
			default:
				fmt.Printf("GETUSER ENDPOINT: unexpected error: %v\n", err) // debuglog

				c.Status(http.StatusInternalServerError)
			}

			return err
		}

		fmt.Printf("GETUSER ENDPOINT: user found: %v\n", user)

		c.JSON(user)

		return nil
	})
	app.Get("/db/ping", func(c *fiber.Ctx) error {
		c.SendString("pinged :)")
		return nil
	})

	app.Listen(":8082")
}