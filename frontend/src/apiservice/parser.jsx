import { Game } from "../proto/chess_pb";

export const parseGame = (newGameResponse) => {
  let game = newGameResponse.getGame();
  let gameID = game.getId();
  let board = game.getBoard();
  let squares = board.getSquaresList();

  let squaresArray = squares.map((square) => {
    let position = square.getPos();
    let x = position.getX();
    let y = position.getY();
    return {
      pos: { x: x, y: y },
      color: square.getColor(),
      piece: square.getPiece(),
    };
  });

  return {
    ID: gameID,
    board: {
      squares: squaresArray,
    },
  };
};
