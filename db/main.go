package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kouhin/envflag"
	"github.com/vivilCODE/chess/db/dbservice"
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
	
	app.Get("/db/ping", func(c *fiber.Ctx) error {
		c.SendString("pinged :)")
		return nil
	})



	app.Listen(":8082")
}