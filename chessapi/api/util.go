package api

import "github.com/vivilCODE/chess/chessapi/pb"

func isOdd(n int) bool {
	return n%2 != 0
}

func generateBoard() *pb.Board {
	squares := []*pb.Square{}

	for y := 1; y <= 8; y++ {
		for x := 1; x <= 8; x++ {
			piece := pb.Piece_pieceNil
			color := pb.Color_white

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
				x == 7 && y == 1:
				piece = pb.Piece_wKnight
			case x == 3 && y == 1 ||
				x == 6 && y == 1:
				piece = pb.Piece_wBishop
			case x == 4 && y == 1:
				piece = pb.Piece_wQueen
			case x == 5 && y == 1:
				piece = pb.Piece_wKing
			case y == 2:
				piece = pb.Piece_wPawn

			// Black side of the board
			case y == 7:
				piece = pb.Piece_bPawn
			case x == 1 && y == 8 ||
				x == 8 && y == 8:
				piece = pb.Piece_bRook
			case x == 2 && y == 8 ||
				x == 7 && y == 8:
				piece = pb.Piece_bKnight
			case x == 3 && y == 8 ||
				x == 6 && y == 8:
				piece = pb.Piece_bBishop
			case x == 4 && y == 8:
				piece = pb.Piece_bQueen
			case x == 5 && y == 8:
				piece = pb.Piece_bKing
			}

			squares = append(squares, &pb.Square{
				Pos:   &pb.SquarePosition{X: uint32(x), Y: uint32(y)},
				Color: color,
				Piece: piece,
			})
		}
	}
	return &pb.Board{
		Squares: squares,
	}
}