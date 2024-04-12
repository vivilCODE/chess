import { useState } from "react";
import * as apiService from "../apiservice/apiservice";
import ChessBoard from "../components/ChessBoard/ChessBoard";

export default function LandingPage() {
  const [game, setGame] = useState(null);

  const NewGame = async () => {
    let game = await apiService.NewGame();
    setGame(() => game);
  };

  return (
    <div className="landing-page">
      <button onClick={apiService.Ping}>Ping</button>
      <button onClick={NewGame}>NewGame</button>

      <ChessBoard game={game} />
    </div>
  );
}
