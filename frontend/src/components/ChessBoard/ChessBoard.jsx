import "./chessboard.css";
import { Color } from "../../proto/chess_pb";
import { useEffect, useState } from "react";

const squareToLi = (square) => {
  let className =
    square.color === 0
      ? "chess-board__square--white"
      : "chess-board__square--black";
  let x = square.pos.x;
  let y = square.pos.y;
  return (
    <li className={className}>
      x: {x}, y: {y}
    </li>
  );
};

const getRow = (squares, number) => {
  return squares.map((sq) => {
    if (sq.pos.y === number) {
      return squareToLi(sq);
    }
  });
};

const getSquares = (squares) => {
  let rows = [];
  for (let i = 1; i <= 8; i++) {
    let row = getRow(squares, i);
    rows.unshift(row);
  }
  return rows;
};

export default function ChessBoard(props) {
  const [squares, setSquares] = useState();
  const { game } = props;

  const renderBoard = () => {
    if (game === null) {
      return;
    }

    let squares = getSquares(game.board.squares);
    console.log(squares);
    setSquares(() => squares);
  };

  useEffect(() => {
    renderBoard();
  }, [game]);

  return <ul className="chess-board">{squares}</ul>;
}
