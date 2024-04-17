import "./chessboard.css";
import { useEffect, useState } from "react";
import { ChessBoardSquare } from "../ChessBoardSquare/ChessBoardSquare";

// The board gets passed a 'game' object, which contains all information needed to render the board.
export const ChessBoard = (props) => {
  const [squares, setSquares] = useState();
  const { game, selectedSquares, setSelectedSquares } = props;

  // Re-renders the board every time a game update occurs or selected squares changes.
  useEffect(() => {
    const getSquares = (squareList) => {
      let squares = [];
      for (let i = 1; i <= 8; i++) {
        let rawRow = squareList.filter((sq) => sq.pos.y === i);
        let row = rawRow.map((sq) => {
          return (
            <ChessBoardSquare
              key={"x:" + sq.pos.x + "y:" + sq.pos.y}
              square={sq}
              selectedSquares={selectedSquares}
              setSelectedSquares={setSelectedSquares}
            />
          );
        });
        squares.unshift(row);
      }
      return squares;
    };

    let squares = getSquares(game.board.squaresList);
    setSquares(() => squares);
  }, [game, selectedSquares, setSelectedSquares]);

  return <ul className="chess-board">{squares}</ul>;
};
