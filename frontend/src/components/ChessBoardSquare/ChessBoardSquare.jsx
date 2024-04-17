import { useEffect, useRef } from "react";

import "./chessboardsquare.css";

const pieceEnumToString = {
  0: "",
  1: "White Pawn",
  2: "Black Pawn",
  3: "White Rook",
  4: "Black Rook",
  5: "White Knight",
  6: "Black Knight",
  7: "White Bishop",
  8: "Black Bishop",
  9: "White King",
  10: "Black King",
  11: "White Queen",
  12: "Black Queen",
};

const pieceEnumToImgPath = {
  0: "",
  1: "./pieces/white-pawn.jpg",
  2: "./pieces/black-pawn.jpg",
  3: "./pieces/white-rook.jpg",
  4: "./pieces/black-rook.jpg",
  5: "./pieces/white-knight.jpg",
  6: "./pieces/black-knight.png",
  7: "./pieces/white-bishop.jpg",
  8: "./pieces/black-bishop.png",
  9: "./pieces/white-king.jpg",
  10: "./pieces/black-king.png",
  11: "./pieces/white-queen.webp",
  12: "./pieces/black-queen.jpg",
};

const thisSquareIsSelected = (selectedSquares, thisSquare) => {
  return (
    selectedSquares.from === thisSquare || selectedSquares.to === thisSquare
  );
};

const thisSquareIsFirstSelection = (selectedSquares, thisSquare) => {
  return selectedSquares.from === thisSquare && selectedSquares.to === null;
};

const bothSelectionsAreNull = (selectedSquares) => {
  return selectedSquares.from === null && selectedSquares.to === null;
};

const secondSelectionIsNull = (selectedSquares) => {
  return selectedSquares.from != null && selectedSquares.to === null;
};

const hasPiece = (thisSquare) => {
  return thisSquare.piece !== 0;
};

export const ChessBoardSquare = (props) => {
  const { square, selectedSquares, setSelectedSquares } = props;
  const thisSquare = useRef();

  const className =
    square.color === 0
      ? "chess-board__square--white"
      : "chess-board__square--black";

  // Every time selectedSquares is updated, this checks if the current square should be
  // visually marked as selected by adding / removing a class
  useEffect(() => {
    if (thisSquareIsSelected(selectedSquares, square)) {
      thisSquare.current.classList.add("square--selected");
      return;
    }
    thisSquare.current.classList.remove("square--selected");
  }, [selectedSquares, square]);

  const select = (e) => {
    // If the same square is clicked twice, reset the selection
    if (thisSquareIsFirstSelection(selectedSquares, square)) {
      setSelectedSquares(() => {
        return { from: null, to: null };
      });
      return;
    }

    // If first and second are null, that means none are currently selected
    // and this square is the first selected square
    if (bothSelectionsAreNull(selectedSquares) && hasPiece(square)) {
      setSelectedSquares(() => {
        return { from: square, to: null };
      });
      return;
    }
    // If first is not null but the second is null, this square will be the second selection
    if (secondSelectionIsNull(selectedSquares)) {
      setSelectedSquares((prev) => {
        return { from: prev.from, to: square };
      });
      return;
    }

    // When two selections are made, the backend should either make the move,
    // or evaluate that it is an illegal move which will reset the selection.
    // That means this if statement should never evaluate to true, but I'm keeping it here
    // until I've implemented the logic to make moves and reset selections.
    if (selectedSquares.from != null && selectedSquares.to != null) {
      console.log("two squares are selected, this shouldn't be able to happen");
      return;
    }
  };

  return (
    <li className={className} onClick={select} ref={thisSquare}>
      {hasPiece(square) && (
        <img
          alt={pieceEnumToString[square.piece]}
          className="chess-board__square__piece-img"
          src={pieceEnumToImgPath[square.piece]}
        />
      )}
    </li>
  );
};
