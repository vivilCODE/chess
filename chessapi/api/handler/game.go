package handler

import (
	"github.com/google/uuid"
	"github.com/vivilCODE/chess/chessapi/api/model"
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


