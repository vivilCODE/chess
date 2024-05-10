package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/vivilCODE/chess/chessapi/dbhandler"
	"github.com/vivilCODE/chess/chessapi/pb"
	"github.com/vivilCODE/chess/db/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


type ChessApi struct {
	pb.UnimplementedChessApiServer
	DBHandler dbhandler.DBHandler
	GapiClientID     string
	GapiClientSecret string
}

func (c *ChessApi) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	context := context.Background()
	
	fmt.Println("API SIGNIN FUNC: sign in requested") // debuglog

	// Configuration for the oauth request, scopes decide what data we are asking for about the user
	conf := &oauth2.Config{
		ClientID:     c.GapiClientID,
		ClientSecret: c.GapiClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		RedirectURL: "http://localhost:3000/signin",
		Endpoint:    google.Endpoint,
	}
	
	// Exchange the authorization code into a token
	token, err := conf.Exchange(context, req.GetCode())
	if err != nil {
		log.Fatal("unable to exchange code:", err)
	}
	
	// Initialise oauth client with the configuration and the exchanged token
	client := conf.Client(context, token)
	
	// Request user info
	res, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("unable to get google userinfo: %v", err)
	}
	defer res.Body.Close()
	
	var userInfo models.User
	
	// Decode json response into user struct
	if err := json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		fmt.Printf("unable to decode userinfo: %v", err)
		return nil, nil
	}
	
	userDoesNotExist := false

	// Try to fetch the same user from database
	dbUser, err := c.DBHandler.GetUser(userInfo.Id)
	if err != nil {
		return nil, fmt.Errorf("error fetching user from database: %v", err)
	}

	// User has never signed in before, create new user in the db with this userinfo
	if dbUser == nil {
		fmt.Println("CHESSAPI SIGN IN FUNC: user and err nil, user has not signed in before") // debuglog
		userDoesNotExist = true
	}


	if userDoesNotExist {
		fmt.Println("CHESSAPI SIGN IN FUNC: user logging in for the first time, post user info to database") // debuglog
		userInfo.SignedUp = time.Now()
		if err := c.DBHandler.PostUser(&userInfo); err != nil {
			return &pb.SignInResponse{}, fmt.Errorf("unable to create user: %v", err)
		}
		
	}
		
	var pbUser = pb.User{
		ID: userInfo.Id,
		Name: strings.Split(userInfo.Email, "@")[0],
		Email: userInfo.Email,
	}
		
	return &pb.SignInResponse{User: &pbUser}, nil
}

func (c *ChessApi) NewGame(context.Context, *pb.NewGameRequest) (*pb.NewGameResponse, error) {
	board := generateBoard()

	var game = &pb.Game{
		ID:        123,
		PlayerOne: nil,
		PlayerTwo: nil,
		Board:     board,
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