package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/kouhin/envflag"
	_ "github.com/lib/pq"

	pb "github.com/vivilCODE/chess/chessapi/pb"
	"github.com/vivilCODE/chess/db/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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
			piece := pb.Piece_pieceNil
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
	gapiClientID string
	gapiClientSecret string
}

func (c *ChessApi) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	code := req.GetCode()
	context := context.Background()

	conf := &oauth2.Config{
		ClientID: c.gapiClientID,
		ClientSecret: c.gapiClientSecret,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		// Scopes: []string{"./auth/userinfo.email", "./auth/userinfo.profile"},
		RedirectURL: "http://localhost:3000/signin",
		Endpoint: google.Endpoint,
	}


	token, err :=	conf.Exchange(context, code)
	if err != nil {
		log.Fatal("unable to exchange code:", err)
	}

	client := conf.Client(context, token)

	res, err :=	client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		fmt.Printf("unable to get userinfo: %v\n", err)
		return nil, nil
	}
	defer res.Body.Close()

	var user models.User

	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		fmt.Printf("unable to decode userinfo: %v", err)
		return nil,nil
	}

	user.SignedUp = time.Now()
	var pbUser =  pb.User{
		ID: user.Id,
		Name: strings.Split(user.Email, "@")[0],
		Email: user.Email,
	}

	// HTTP CALL TO DB SERVICE TO CREATE USER

	fmt.Println("sign in requested")
	return &pb.SignInResponse{User: &pbUser}, nil
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
	game := req.Game
	fromPosition := req.Move.From.Pos
	toPosition := req.Move.To.Pos

	pieceOnFromSquare := req.Move.From.Piece
	
	for _, sq := range game.Board.Squares {
		if sq.Pos.GetX() == fromPosition.GetX() && sq.Pos.GetY() == fromPosition.GetY() {
			sq.Piece = pb.Piece_pieceNil
		}
		if sq.Pos.GetX() == toPosition.GetX() && sq.Pos.GetY() == toPosition.GetY() {
			sq.Piece = pieceOnFromSquare
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
	api := &ChessApi{
		gapiClientID: gapiClientID,
		gapiClientSecret: gapiClientSecret,
	}
	pb.RegisterChessApiServer(grpcServer, api)

	// Serve grpc
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc: %v", err)
	}
}



func pingDatabase(address string) error {
	res, err := http.Get(address)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("received non 200 response: %v", res.StatusCode)
	}

	return nil
}