import { useEffect, useState } from "react";
import * as apiService from "../apiservice/apiservice";
import { ChessBoard } from "../components/ChessBoard/ChessBoard";

const twoSquaresSelected = (selectedSquares) => {
  return selectedSquares.from != null && selectedSquares.to != null;
};

export const GameWindow = () => {
  const [game, setGame] = useState(null);
  const [selectedSquares, setSelectedSquares] = useState({
    from: null,
    to: null,
  });

  // When two squares have been selected, a move is ready to be made.
  // Use the apiService to make a MakeMove requesst to the backend,
  // then update the game state and reset selectedSquares state
  useEffect(() => {
    const MakeMove = async () => {
      let updatedGame = await apiService.MakeMove(
        game,
        selectedSquares.from,
        selectedSquares.to
      );
      setGame(() => updatedGame.game);
      setSelectedSquares(() => {
        return { from: null, to: null };
      });
    };

    if (twoSquaresSelected(selectedSquares)) {
      MakeMove();
    }
  }, [selectedSquares, game]);

  const NewGame = async () => {
    let game = await apiService.NewGame();
    console.log("game: ", game);
    setGame(() => game.game);
  };

  return (
    <div className="game-window">
      <button onClick={NewGame}>NewGame</button>

      {game && (
        <ChessBoard
          game={game}
          selectedSquares={selectedSquares}
          setSelectedSquares={setSelectedSquares}
        />
      )}
    </div>
  );
};
