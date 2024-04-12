package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/vivilCODE/chess/chessapi/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type myChessApi struct {
	pb.UnimplementedChessApiServer
}

func (c *myChessApi) NewGame(context.Context, *pb.NewGameRequest) (*pb.NewGameResponse, error) {
	var squares = []*pb.Square{
		{
			Pos: &pb.SquarePosition{X: 1, Y: 1},
			Color: pb.Color_black,
			Piece: pb.Piece_wRook,
		}, 
		{
			Pos: &pb.SquarePosition{X: 2, Y: 1},
			Color: pb.Color_white,
			Piece: pb.Piece_wKnight,
		}, 
		{
			Pos: &pb.SquarePosition{X: 3, Y: 1},
			Color: pb.Color_black,
			Piece: pb.Piece_wBishop,
		}, 
		{
			Pos: &pb.SquarePosition{X: 4, Y: 1},
			Color: pb.Color_white,
			Piece: pb.Piece_wQueen,
		}, 
		{
			Pos: &pb.SquarePosition{X: 5, Y: 1},
			Color: pb.Color_black,
			Piece: pb.Piece_wKing,
		}, 
		{
			Pos: &pb.SquarePosition{X: 6, Y: 1},
			Color: pb.Color_white,
			Piece: pb.Piece_wBishop,
		}, 
		{
			Pos: &pb.SquarePosition{X: 7, Y: 1},
			Color: pb.Color_black,
			Piece: pb.Piece_wKnight,
		}, 
		{
			Pos: &pb.SquarePosition{X: 8, Y: 1},
			Color: pb.Color_white,
			Piece: pb.Piece_wRook,
		}, 

		{
			Pos: &pb.SquarePosition{X: 1, Y: 2},
			Color: pb.Color_white,
			Piece: pb.Piece_wPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 2, Y: 2},
			Color: pb.Color_black,
			Piece: pb.Piece_wPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 3, Y: 2},
			Color: pb.Color_white,
			Piece: pb.Piece_wPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 4, Y: 2},
			Color: pb.Color_black,
			Piece: pb.Piece_wPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 5, Y: 2},
			Color: pb.Color_white,
			Piece: pb.Piece_wPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 6, Y: 2},
			Color: pb.Color_black,
			Piece: pb.Piece_wPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 7, Y: 2},
			Color: pb.Color_white,
			Piece: pb.Piece_wPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 8, Y: 2},
			Color: pb.Color_black,
			Piece: pb.Piece_wPawn,
		}, 

		{
			Pos: &pb.SquarePosition{X: 1, Y: 3},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 2, Y: 3},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 3, Y: 3},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 4, Y: 3},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 5, Y: 3},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 6, Y: 3},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 7, Y: 3},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 8, Y: 3},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 1, Y: 4},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 2, Y: 4},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 3, Y: 4},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 4, Y: 4},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 5, Y: 4},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 6, Y: 4},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 7, Y: 4},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 8, Y: 4},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 1, Y: 5},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 2, Y: 5},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 3, Y: 5},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 4, Y: 5},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 5, Y: 5},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 6, Y: 5},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 7, Y: 5},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 8, Y: 5},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 1, Y: 6},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 2, Y: 6},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 3, Y: 6},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 4, Y: 6},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 5, Y: 6},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 6, Y: 6},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 7, Y: 6},
			Color: pb.Color_white,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 8, Y: 6},
			Color: pb.Color_black,
			Piece: pb.Piece_nil,
		}, 
		{
			Pos: &pb.SquarePosition{X: 1, Y: 7},
			Color: pb.Color_black,
			Piece: pb.Piece_bPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 2, Y: 7},
			Color: pb.Color_white,
			Piece: pb.Piece_bPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 3, Y: 7},
			Color: pb.Color_black,
			Piece: pb.Piece_bPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 4, Y: 7},
			Color: pb.Color_white,
			Piece: pb.Piece_bPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 5, Y: 7},
			Color: pb.Color_black,
			Piece: pb.Piece_bPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 6, Y: 7},
			Color: pb.Color_white,
			Piece: pb.Piece_bPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 7, Y: 7},
			Color: pb.Color_black,
			Piece: pb.Piece_bPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 8, Y: 7},
			Color: pb.Color_white,
			Piece: pb.Piece_bPawn,
		}, 
		{
			Pos: &pb.SquarePosition{X: 1, Y: 8},
			Color: pb.Color_white,
			Piece: pb.Piece_bRook,
		}, 
		{
			Pos: &pb.SquarePosition{X: 2, Y: 8},
			Color: pb.Color_black,
			Piece: pb.Piece_bKnight,
		}, 
		{
			Pos: &pb.SquarePosition{X: 3, Y: 8},
			Color: pb.Color_white,
			Piece: pb.Piece_bBishop,
		}, 
		{
			Pos: &pb.SquarePosition{X: 4, Y: 8},
			Color: pb.Color_black,
			Piece: pb.Piece_bQueen,
		}, 
		{
			Pos: &pb.SquarePosition{X: 5, Y: 8},
			Color: pb.Color_white,
			Piece: pb.Piece_bKing,
		}, 
		{
			Pos: &pb.SquarePosition{X: 6, Y: 8},
			Color: pb.Color_black,
			Piece: pb.Piece_bBishop,
		}, 
		{
			Pos: &pb.SquarePosition{X: 7, Y: 8},
			Color: pb.Color_white,
			Piece: pb.Piece_bKnight,
		}, 
		{
			Pos: &pb.SquarePosition{X: 8, Y: 8},
			Color: pb.Color_black,
			Piece: pb.Piece_bRook,
		}, 

	}
	
	var board = &pb.Board{
		Squares: squares,
	}

	
	var game = &pb.Game{
		ID: 123,
		PlayerOne: nil,
		PlayerTwo: nil,
		Board: board,
	}

	return &pb.NewGameResponse{Game: game}, nil
}

func (c *myChessApi) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	fmt.Println("ping request hit server")
	return &pb.PingResponse{Response: "up and running :)"}, nil
}

func main() {
	// Set up listener to a port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}

	// Initialise grpc server
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	// Register chess api
	api := &myChessApi{}
	pb.RegisterChessApiServer(grpcServer, api)

	// Serve grpc
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc: %v", err)
	}
}
