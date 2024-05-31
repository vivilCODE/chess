import { useEffect, useRef } from "react";

import "./chessboardsquare.css";

const pieceStringToImgPath = {
  nilPiece: "",
  whitePawn: "/pieces/white-pawn.svg",
  blackPawn: "/pieces/black-pawn.svg",
  whiteRook: "/pieces/white-rook.svg",
  blackRook: "/pieces/black-rook.svg",
  whiteKnight: "/pieces/white-knight.svg",
  blackKnight: "/pieces/black-knight.svg",
  whiteBishop: "/pieces/white-bishop.svg",
  blackBishop: "/pieces/black-bishop.svg",
  whiteKing: "/pieces/white-king.svg",
  blackKing: "/pieces/black-king.svg",
  whiteQueen: "/pieces/white-queen.svg",
  blackQueen: "/pieces/black-queen.svg",
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
  return thisSquare.piece !== "nilPiece";
};

export const ChessBoardSquare = (props) => {
  const { square, selectedSquares, setSelectedSquares, active } = props;
  const thisSquare = useRef();

  const className =
    square.isWhite === true
      ? "chess-board__square--white"
      : "chess-board__square--black";

  // Every time selectedSquares is updated, this checks if the current square should be
  // visually marked as selected, and adds / removes the "selected" class accordingly
  useEffect(() => {
    if (!active) return; // If this board is only a display board, don't do anything on click

    if (thisSquareIsSelected(selectedSquares, square)) {
      thisSquare.current.classList.add("square--selected");
      return;
    }
    thisSquare.current.classList.remove("square--selected");
  }, [selectedSquares, square]);

  // Controls what happens to the selectedSquares state when this square is clicked
  const select = (e) => {
    if (!active) return; // If this board is only a display board, don't do anything on click

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
          alt={square.piece}
          className="chess-board__square__piece-img"
          src={pieceStringToImgPath[square.piece]}
        />
      )}
    </li>
  );
};
