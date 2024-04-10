package main

import (
	"context"
	"log"
	"net"

	pb "github.com/vivilCODE/chess/pb"
	"google.golang.org/grpc"
)

type myChessApi struct {
	pb.UnimplementedChessApiServer
}

func (c *myChessApi) MakeMove(context.Context, *pb.MakeMoveRequest) (*pb.MakeMoveResponse, error) {
	return &pb.MakeMoveResponse{Body: "response from server"}, nil
}

func main() {
	// Set up listener to a port
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	// Initialise grpc server
	grpcServer := grpc.NewServer()

	// Register chess api
	api := &myChessApi{}
	pb.RegisterChessApiServer(grpcServer, api)

	// Serve grpc
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc: %v", err)
	}
}
