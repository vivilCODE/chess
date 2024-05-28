package handler

import (
	"encoding/json"
	"math/rand"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/vivilCODE/chess/chessapi/api/model"
	"github.com/vivilCODE/chess/chessapi/log"
)

type GameEngine struct {
	playerWhite *gameClient
	playerBlack *gameClient
	game        model.Game
}

func NewGameEngine(p1, p2 *gameClient) *GameEngine{
	gameId := uuid.NewString()
	
	pWhite, pBlack := assignBlackWhite(p1, p2)

	ge := &GameEngine{
		playerWhite: pWhite,
		playerBlack: pBlack,
		game: model.Game{
			ID: gameId,
			PlayerWhite: pWhite.user,
			PlayerBlack: pBlack.user,
			Board: generateBoard(),
		},
	}

	return ge
}


func (ge *GameEngine) StartGame() {
	pWhite := ge.playerWhite
	pBlack := ge.playerBlack

	
	log.Logger.Debug("starting new game", "white", pWhite.user.Name, "black", pBlack.user.Name)
	
	// log.Logger.Debug("game:","game", ge.game)

	gameJSON, err := json.Marshal(ge.game)
	if err != nil {
		log.Logger.Error("unable to marshal game data", "err", err)
	}

	log.Logger.Debug("game json:", "game", gameJSON)
	
	startMessage := []byte("start game")
	messageWithGame := append(startMessage, gameJSON...)

	// Send the start message to both players
	if err := pWhite.conn.WriteMessage(websocket.TextMessage, messageWithGame); err != nil {
		log.Logger.Error("unable to start game for white", "err", err)
	} else {
		log.Logger.Debug("game start message sent to white", "name", pWhite.user.Name)
	}

	if err := pBlack.conn.WriteMessage(websocket.TextMessage, messageWithGame); err != nil {
		log.Logger.Error("unable to start game for black", "err", err)
		} else {
		log.Logger.Debug("game start message sent to black", "name", pBlack.user.Name)
	}
}

// func assingBlackWhite assigns black or white pieces to the two players
// based on a random int
func assignBlackWhite(p1, p2 *gameClient) (*gameClient, *gameClient) {
	if rand.Intn(2) == 1 {
		return p1, p2
	} else {
		return p2, p1
	}
}


func generateBoard() model.Board {
	squares := []model.Square{}

	for y := 1; y <= 8; y++ {
		for x := 1; x <= 8; x++ {
			piece := model.NilPiece
			isWhite := true

			if isOdd(y) && isOdd(x) ||
				!isOdd(y) && !isOdd(x) {
				isWhite = false
			}

			switch {
			// White side of the board
			case x == 1 && y == 1 ||
				x == 8 && y == 1:
				piece = model.WhiteRook
			case x == 2 && y == 1 ||
				x == 7 && y == 1:
				piece = model.WhiteKnight
			case x == 3 && y == 1 ||
				x == 6 && y == 1:
				piece = model.WhiteBishop
			case x == 4 && y == 1:
				piece = model.WhiteQueen
			case x == 5 && y == 1:
				piece = model.WhiteKing
			case y == 2:
				piece = model.WhitePawn

			// Black side of the board
			case y == 7:
				piece = model.BlackPawn
			case x == 1 && y == 8 ||
				x == 8 && y == 8:
				piece = model.BlackRook
			case x == 2 && y == 8 ||
				x == 7 && y == 8:
				piece = model.BlackKnight
			case x == 3 && y == 8 ||
				x == 6 && y == 8:
				piece = model.BlackBishop
			case x == 4 && y == 8:
				piece = model.BlackQueen
			case x == 5 && y == 8:
				piece = model.BlackKing
			}

			squares = append(squares, model.Square{
				Pos: model.SquarePosition{X:x, Y: y},
				IsWhite: isWhite,
				Piece: piece,
			})
		}
	}
	return model.Board{
		Squares: squares,
	}
}


func isOdd(n int) bool {
	return n % 2 != 0
}
