package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kouhin/envflag"
	_ "github.com/lib/pq"

	"github.com/vivilCODE/chess/chessapi/api/config"
	"github.com/vivilCODE/chess/chessapi/api/router"
	"github.com/vivilCODE/chess/chessapi/dbhandler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vivilCODE/chess/chessapi/log"
)

func pingDatabase(address string) error {
	res, err := http.Get(fmt.Sprintf("http://%s/db/ping", address))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("received non 200 response: %v", res.StatusCode)
	}

	return nil
}


func main() {
	httpPort := flag.String("port", ":8080", "Address for listening to http requests")
	dbAddress := flag.String("dbaddress", "localhost:8082", "Address for communicating with db")

	log.Logger.Info("starting chessapi")

	if err := envflag.Parse(); err != nil {
		log.Logger.Fatal("unable to parse flags", "err", err)
	}

	err := godotenv.Load()
	if err != nil {
		log.Logger.Fatal("unable to load .env file", "err", err)
	}

	// Ensure database is running and reachable
	if err := pingDatabase(*dbAddress); err != nil {
		log.Logger.Fatal("unable to reach database service", "err", err)
	}
	log.Logger.Info("successfully contacted database")

	dbHandlerConfig := &dbhandler.Config{
		Address: *dbAddress,
	}

	dbHandler := &dbhandler.DBHandler{
		Config: dbHandlerConfig,
	}

	gapiClientID:= 	os.Getenv("GAPI_CLIENT_ID")
	gapiClientSecret := os.Getenv("GAPI_CLIENT_SECRET")

	chessapiConfig := config.Config{
		DBHandler: dbHandler,
		GapiClientID: gapiClientID,
		GapiClientSecret: gapiClientSecret,
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
        AllowOrigins:     "*",
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

	router.SetupRoutes(app, chessapiConfig)

	if err := app.Listen(*httpPort); err != nil {
		log.Logger.Fatal("unable to serve chessapi on", "port", *httpPort, "err", err)
	}
}
