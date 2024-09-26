package main

import (
	"flag"
	"os"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kouhin/envflag"
	"github.com/vivilCODE/chess/db/dbservice"
	"github.com/vivilCODE/chess/db/handler"
	"github.com/vivilCODE/chess/db/log"
	"github.com/vivilCODE/chess/db/router"
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

	dbUser := os.Getenv("DB_USER")
	dbUserPassword := os.Getenv("DB_USER_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	
	
	databaseConfig := dbservice.Config{
		DBPort:   *dbPort,
		DBName:   dbName,
		Host:     *host,
		User:     dbUser,
		Password: dbUserPassword,
	}
	
	// Initialise database service
	dbs, err := dbservice.New(databaseConfig)
	if err != nil {
		log.Logger.Fatal("unable to initialise database service", "err", err)
		return
	}

	// Connect to the postgres database
	if err = dbs.Connect(); err != nil {
		log.Logger.Fatal("unable to connect to database", "err", err)
	}
	defer dbs.Disconnect()

	handler := handler.New(dbs)

	app := fiber.New()

	// Setup REST endpoints
	router.SetupRoutes(app, handler)

	app.Listen(":8082")
}