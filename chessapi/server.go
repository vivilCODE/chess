package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/kouhin/envflag"
	_ "github.com/lib/pq"
	"github.com/vivilCODE/chess/chessapi/dbcontroller"
	pb "github.com/vivilCODE/chess/chessapi/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func isOdd(n int) bool {
	return n % 2 != 0
}

func generateBoard() (*pb.Board) {
	squares := []*pb.Square{}
	

	for y:=1; y<=8; y++ {
		for x:= 1; x<=8; x++ {
			piece := pb.Piece_nil
			color:= pb.Color_white
			
			if isOdd(y) && isOdd(x) || 
			!isOdd(y) && !isOdd(x) {
				color = pb.Color_black
			}

			switch {
			// White side of the board
			case x == 1 && y == 1 || 
			x == 8 && y == 1:
				piece = pb.Piece_wRook
			case x == 2 && y == 1 ||
			x== 7 && y==1:
				piece = pb.Piece_wKnight
			case x == 3 && y == 1 ||
			x == 6 && y==1:
				piece = pb.Piece_wBishop
			case x == 4 && y == 1:
				piece = pb.Piece_wQueen
			case x == 5 && y == 1:
				piece = pb.Piece_wKing
			case y == 2:
				piece = pb.Piece_wPawn
	
			// Black side of the board
			case y==7:
				piece = pb.Piece_bPawn
			case x == 1 && y == 8 ||
			x== 8 && y==8:
				piece = pb.Piece_bRook
				case x == 2 && y == 8 ||
			x == 7 && y==8:
				piece = pb.Piece_bKnight
				case x == 3 && y == 8 ||
			x == 6 && y==8:
				piece = pb.Piece_bBishop
			case x == 4 && y == 8:
				piece = pb.Piece_bQueen			
			case x == 5 && y == 8:
				piece = pb.Piece_bKing
			}

			squares = append(squares, &pb.Square{
				Pos: &pb.SquarePosition{X:uint32(x), Y: uint32(y)},
				Color: color,
				Piece: piece,
			})
		}
	}
	return &pb.Board{
		Squares: squares,
	}
}


type ChessApi struct {
	pb.UnimplementedChessApiServer
	dbcontroller *dbcontroller.DBController
}

func (c *ChessApi) NewGame(context.Context, *pb.NewGameRequest) (*pb.NewGameResponse, error) {
	board := generateBoard();
	
	var game = &pb.Game{
		ID: 123,
		PlayerOne: nil,
		PlayerTwo: nil,
		Board: board,
	}

	return &pb.NewGameResponse{Game: game}, nil
}

func (c *ChessApi) MakeMove(ctx context.Context, req *pb.MakeMoveRequest) (*pb.MakeMoveResponse, error) {
	fmt.Println("make move request hit the server")
	
	game:= req.Game

	fromPosition := req.Move.From.Pos
	toPosition := req.Move.To.Pos

	pieceOnFromSquare := req.Move.From.Piece
	pieceOnToSquare := req.Move.To.Piece

	fmt.Printf("%v from %v to %v, capturing %v\n", pieceOnFromSquare, fromPosition, toPosition, pieceOnToSquare)

	for _, sq := range game.Board.Squares {
		if sq.Pos.GetX() == fromPosition.GetX() && sq.Pos.GetY() == fromPosition.GetY() {
			fmt.Printf("found fromposition: %v\n", sq)
			sq.Piece = pb.Piece_nil
		}
		if sq.Pos.GetX() == toPosition.GetX() && sq.Pos.GetY() == toPosition.GetY() {
			fmt.Printf("found toosition: %v\n", sq)
			sq.Piece = pieceOnFromSquare
		}
	}

	
	for _, sq := range game.Board.Squares {
		if sq.Pos.GetX() == fromPosition.GetX() && sq.Pos.GetY() == fromPosition.GetY() {
			fmt.Printf("found fromposition2: %v\n", sq)
		}
		if sq.Pos.GetX() == toPosition.GetX() && sq.Pos.GetY() == toPosition.GetY() {
			fmt.Printf("found toosition2: %v\n", sq)
		}
	}

	

	return &pb.MakeMoveResponse{Game: game}, nil
}

func (c *ChessApi) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	fmt.Println("ping request hit server")
	return &pb.PingResponse{Response: "up and running :)"}, nil
}

func main() {
	httpPort := flag.String("port", ":8080", "Address for listening to http requests")
	dbPort := flag.Int("dbport", 5432, "Postgres address")
	host := flag.String("host", "localhost", "Where to host the http endpoint")



	if err := envflag.Parse(); err != nil {
		log.Fatalf("unable to parse flags, %v", err)
	}

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Failed to load .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbUserPassword := os.Getenv("DB_USER_PASSWORD")
	dbName := os.Getenv("DB_NAME")


	databaseConfig := dbcontroller.Config{
		DBPort:   *dbPort,
		DBName:   dbName,
		Host:     *host,
		User:     dbUser,
		Password: *&dbUserPassword,
	}

	dbc, err := dbcontroller.New(databaseConfig)
	if err != nil {
		log.Fatalf("unable to initialise database controller: %v", err)
		return
	}

	if err = dbc.Connect(); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer dbc.Disconnect()

	// Set up listener to a port
	lis, err := net.Listen("tcp", *httpPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", *httpPort, err)
	}

	// Initialise grpc server
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	// Register chess api
	api := &ChessApi{
		dbcontroller: dbc,
	}
	pb.RegisterChessApiServer(grpcServer, api)

	// Serve grpc
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc: %v", err)
	}
}
