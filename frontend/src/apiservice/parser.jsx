import {
  MakeMoveRequest,
  NewGameRequest,
  Game,
  Square,
  SquarePosition,
  Color,
  Board,
  Piece,
  Move,
} from "../proto/chess_pb";

export const objectGameToPb = (game) => {
  let pbGame = new Game();
  console.log("objectGameToPb function, game:", game);
  pbGame.setId(game.id);

  let pbBoard = new Board();

  let squares = game.board.squaresList.map((sq) => {
    let square = new Square();
    let squarePosition = new SquarePosition();
    squarePosition.setX(sq.pos.x);
    squarePosition.setY(sq.pos.y);
    square.setPos(squarePosition);
    square.setColor(colorToEnumString[sq.color]);
    square.setPiece(pieceEnumToString[sq.piece]);
    return square;
  });

  console.log(squares);
  pbBoard.setSquaresList(squares);

  // let pbSquares = game.board.squaresList.map((sq) => {
  //   let pbSquare = new Square();
  //   return pbSquare;
  // });

  // pbGame.board = pbBoard;

  pbGame.setBoard(pbBoard);
  console.log("objectGameToPb function, pbGame:", pbGame.toObject());
  return pbGame;
};

export const createNewGameRequest = () => {
  let req = new NewGameRequest();
  return req;
};

export const createMakeMoveRequest = (game, from, to) => {
  let pbGame = objectGameToPb(game);
  const req = new MakeMoveRequest();
  req.setGame(pbGame);

  //   Initialise move
  let move = new Move();

  //   Create and add "from" to move
  let fromSquare = new Square();

  let fromSquarePosition = new SquarePosition();
  fromSquarePosition.setX(from.pos.x);
  fromSquarePosition.setY(from.pos.y);
  fromSquare.setPos(fromSquarePosition);
  fromSquare.setColor(colorToEnumString[from.color]);
  fromSquare.setPiece(pieceEnumToString[from.piece]);
  move.setFrom(fromSquare);

  //   Create and add "to" to move
  let toSquare = new Square();

  let toSquarePosition = new SquarePosition();
  toSquarePosition.setX(to.pos.x);
  toSquarePosition.setY(to.pos.y);
  toSquare.setPos(toSquarePosition);
  toSquare.setColor(colorToEnumString[to.color]);
  toSquare.setPiece(pieceEnumToString[to.piece]);
  move.setTo(toSquare);

  req.setMove(move);

  return req;
};

const colorToEnumString = {
  0: Color.WHITE,
  1: Color.BLACK,
};

const pieceEnumToString = {
  0: Piece.NIL,
  1: Piece.WPAWN,
  2: Piece.BPAWN,
  3: Piece.WROOK,
  4: Piece.BROOK,
  5: Piece.WKNIGHT,
  6: Piece.BKNIGHT,
  7: Piece.WBISHOP,
  8: Piece.BBISHOP,
  9: Piece.WKING,
  10: Piece.BKING,
  11: Piece.WQUEEN,
  12: Piece.BQUEEN,
};
