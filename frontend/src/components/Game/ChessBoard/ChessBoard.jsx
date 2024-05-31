import "./chessboard.css";
import { useEffect, useState } from "react";
import { ChessBoardSquare } from "./ChessBoardSquare";

// The board gets passed a 'game' object, which contains all information needed to render the board.
// If game is nil, which it will be in the game selection screen until a game is initiated,
// the chessboard still renders, but no pieces can be selected/moved.
export const ChessBoard = (props) => {
  const [squares, setSquares] = useState();
  const { game, selectedSquares, setSelectedSquares, isWhitePerspective } =
    props;

  // Re-renders the board every time a game update occurs or selected squares changes.
  // useEffect(() => {
  // getSquareElements takes an array of 64 squares and returns 64 ChessBoardSquare elements
  //   const getSquareElements = (squareList) => {
  //     let squares = [];
  //     for (let i = 1; i <= 8; i++) {
  //       let rawRow = squareList.filter((sq) => sq.pos.y === i);
  //       let row = rawRow.map((sq) => {
  //         return (
  //           <ChessBoardSquare
  //             key={"x:" + sq.pos.x + "y:" + sq.pos.y}
  //             square={sq}
  //             selectedSquares={selectedSquares}
  //             setSelectedSquares={setSelectedSquares}
  //           />
  //         );
  //       });
  //       squares.unshift(row);
  //     }
  //     return squares;
  //   };

  //   let squares = getSquareElements(game.board.squaresList);
  //   setSquares(() => squares);
  // }, [game, selectedSquares, setSelectedSquares]);

  // If the client is currently not in a game, display the board as it would look
  // before any moves are made, from whites perspective
  useEffect(() => {
    if (!game) {
      const squares = generateBoardWhite();
      setSquares(() => squares);
    }
  }, [game]);

  return <ul className="chess-board">{squares}</ul>;
};

// Returns an array of all squares with their pieces in the starting position
// x1 y1... x8 y1, x1 y2... x8 y2 etc
const generateSquares = () => {
  const isOdd = (n) => n % 2 === 0;

  let squares = [];

  for (let y = 1; y <= 8; y++) {
    for (let x = 1; x <= 8; x++) {
      let piece;
      let isWhite = true;

      if ((isOdd(y) && isOdd(x)) || (!isOdd(y) && !isOdd(x))) {
        isWhite = false;
      }

      switch (true) {
        // White side of the board
        case (x === 1 && y === 1) || (x === 8 && y === 1):
          piece = "whiteRook";
          break;
        case (x === 2 && y === 1) || (x === 7 && y === 1):
          piece = "whiteKnight";
          break;
        case (x === 3 && y === 1) || (x === 6 && y === 1):
          piece = "whiteBishop";
          break;
        case x === 4 && y === 1:
          piece = "whiteQueen";
          break;
        case x === 5 && y === 1:
          piece = "whiteKing";
          break;
        case y === 2:
          piece = "whitePawn";
          break;

        // Black side of the board
        case y === 7:
          piece = "blackPawn";
          break;
        case (x === 1 && y === 8) || (x === 8 && y === 8):
          piece = "blackRook";
          break;
        case (x === 2 && y === 8) || (x === 7 && y === 8):
          piece = "blackKnight";
          break;
        case (x === 3 && y === 8) || (x === 6 && y === 8):
          piece = "blackBishop";
          break;
        case x === 4 && y === 8:
          piece = "blackQueen";
          break;
        case x === 5 && y === 8:
          piece = "blackKing";
          break;

        default:
          piece = "nilPiece"; // Assuming you want to set a default piece
          break;
      }

      squares.push({
        pos: { x: x, y: y },
        piece: piece,
        isWhite: isWhite,
        active: false,
      });
    }
  }

  return squares;
};

// Takes an array of squares (objects) and returns them as ChessBoardSquare elemtns
// in the order that they should appear  from whites perspective
const generateBoardWhite = () => {
  const objectSquares = generateSquares();

  let elementSquares = [];
  for (let y = 1; y <= 8; y++) {
    let objectRow = objectSquares.filter((sq) => sq.pos.y === y);
    let squareElementRow = objectRow.map((sq) => {
      return (
        <ChessBoardSquare key={"x:" + sq.pos.x + "y:" + sq.pos.y} square={sq} />
      );
    });
    elementSquares.unshift(squareElementRow);
  }
  return elementSquares;
};

const generateBoardBlack = () => {
  const objectSquares = generateSquares();

  let elementSquares = [];

  for (let y = 1; y <= 8; y++) {
    let objectRow = objectSquares.filter((sq) => sq.pos.y === y).reverse();
    let squareElementRow = objectRow.map((sq) => {
      return (
        <ChessBoardSquare key={"x:" + sq.pos.x + "y:" + sq.pos.y} square={sq} />
      );
    });
    elementSquares.push(squareElementRow);
  }

  return elementSquares;
};
