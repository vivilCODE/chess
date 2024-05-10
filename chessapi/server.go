package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kouhin/envflag"
	_ "github.com/lib/pq"

	"github.com/vivilCODE/chess/chessapi/api"
	"github.com/vivilCODE/chess/chessapi/dbhandler"
	pb "github.com/vivilCODE/chess/chessapi/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	
	if err := envflag.Parse(); err != nil {
		log.Fatalf("unable to parse flags, %v", err)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	if err := pingDatabase(*dbAddress); err != nil {
		log.Fatalf("unable to reach database service: %v\n", err)
	}
	fmt.Println("successfully contacted database")

	dbHandlerConfig := dbhandler.Config{
		Address: *dbAddress,
	}

	dbHandler := dbhandler.DBHandler{
		Config: dbHandlerConfig,
	}



	gapiClientID:= 	os.Getenv("GAPI_CLIENT_ID")
	gapiClientSecret := os.Getenv("GAPI_CLIENT_SECRET")

	// Set up listener to a port
	lis, err := net.Listen("tcp", *httpPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", *httpPort, err)
	}

	// Initialise grpc server
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	// Register chess api
	api := &api.ChessApi{
		DBHandler: dbHandler,
		GapiClientID: gapiClientID,
		GapiClientSecret: gapiClientSecret,
	}
	pb.RegisterChessApiServer(grpcServer, api)

	// Serve grpc
	fmt.Printf("started serving grpc on port %s\n", *httpPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc: %v", err)
	}
}
