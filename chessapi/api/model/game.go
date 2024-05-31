package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type Game struct {
	ID          string `json:"id"`
	PlayerWhite User   `json:"playerWhite"`
	PlayerBlack User   `json:"playerBlakc"`
	Board       Board  `json:"board"`
	Started time.Time `json:"started"`
}

type Board struct {
	Squares []Square `json:"squares"`
}

type Square struct {
	Pos     SquarePosition `json:"pos"`
	IsWhite bool           `json:"isWhite"`
	Piece   Piece          `json:"piece"`
}

type SquarePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Piece int

const (
	NilPiece Piece = iota + 1
	WhitePawn
	BlackPawn
	WhiteRook
	BlackRook
	WhiteKnight
	BlackKnight
	WhiteBishop
	BlackBishop
	WhiteQueen
	BlackQueen
	WhiteKing
	BlackKing
)

func (p Piece) String() string {
	names := []string{
		"NilPiece",
		"WhitePawn",
		"BlackPawn",
		"WhiteRook",
		"BlackRook",
		"WhiteKnight",
		"BlackKnight",
		"WhiteBishop",
		"BlackBishop",
		"WhiteQueen",
		"BlackQueen",
		"WhiteKing",
		"BlackKing",
	}
	if p < NilPiece || int(p) > len(names) {
		return "Unknown"
	}
	return names[p-1]
}

func (p Piece) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *Piece) UnmarshalJSON(data []byte) error {
	var pieceName string
	if err := json.Unmarshal(data, &pieceName); err != nil {
		return err
	}

	names := map[string]Piece{
		"NilPiece":    NilPiece,
		"WhitePawn":   WhitePawn,
		"BlackPawn":   BlackPawn,
		"WhiteRook":   WhiteRook,
		"BlackRook":   BlackRook,
		"WhiteKnight": WhiteKnight,
		"BlackKnight": BlackKnight,
		"WhiteBishop": WhiteBishop,
		"BlackBishop": BlackBishop,
		"WhiteQueen":  WhiteQueen,
		"BlackQueen":  BlackQueen,
		"WhiteKing":   WhiteKing,
		"BlackKing":   BlackKing,
	}

	if piece, found := names[pieceName]; found {
		*p = piece
		return nil
	}

	return fmt.Errorf("unknown piece name: %s", pieceName)
}