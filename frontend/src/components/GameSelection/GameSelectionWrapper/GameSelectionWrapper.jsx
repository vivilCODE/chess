import { ChessBoard } from "../../Game/ChessBoard/ChessBoard";
import { UserContext } from "../../../index";
import { useContext, useState } from "react";

import "./gameselectionwrapper.css";

export const GameSelectionWrapper = () => {
  const [user] = useContext(UserContext);
  const [playHuman, setPlayHuman] = useState(true);
  const [playComputer, setPlayComputer] = useState(false);
  const [waitingForOpponent, setWaitingForOpponent] = useState(false);
  const [foundOpponent, setFoundOpponent] = useState(false);

  const togglePlayHuman = () => {
    setPlayHuman(() => true);
    setPlayComputer(() => false);
  };

  const togglePlayComputer = () => {
    setPlayComputer(() => true);
    setPlayHuman(() => false);
  };

  // Placeholder function for future implementation
  const startBotGame = () => {
    alert("Not implemented");
  };

  // Opens a websocket connection with the server and send user data to it.
  // This places the user in a queue where they will be until an opponent is found.
  const enterQueue = () => {
    let socket = new WebSocket("ws://localhost:8080/chessapi/ws/newgame");

    socket.onopen = () => {
      console.log(
        "game queue websocket connection established, sending user: ",
        user
      );

      const jsonUser = JSON.stringify({
        id: user.user.id,
        name: user.user.name,
        email: user.user.email,
        signed_up: user.user.signed_up,
      });

      socket.send(jsonUser);
    };

    socket.onmessage = (e) => {
      const msg = JSON.parse(e.data);

      switch (msg.type) {
        case "opponentFound":
          setWaitingForOpponent(() => false);
          setFoundOpponent(() => true);
        // redirect to a url like /game/5910919dke9d390209
        // the id is the id of the game

        case "waiting":
          setWaitingForOpponent(() => true);
        default:
          console.log(msg);
      }
    };
  };

  return (
    <div className="game-selection-wrapper">
      <ChessBoard />

      <div className="game-selection-wrapper__selection-column">
        <div className="game-selection-wrapper__selection-column__player-or-computer-btns">
          <button
            onClick={togglePlayHuman}
            className="game-selection-wrapper__selection-column__player-or-computer-btns--human"
          >
            Play Human
          </button>
          <button
            onClick={togglePlayComputer}
            className="game-selection-wrapper__selection-column__player-or-computer-btns--computer"
          >
            Play Computer
          </button>
        </div>

        {!user && playHuman && <p>Please log in to play online</p>}
        {user && playHuman && <button onClick={enterQueue}>Play</button>}
        {playComputer && <button onClick={startBotGame}>Play</button>}
        {waitingForOpponent && <span>waiting to find opponent...</span>}
      </div>

      {/* {!user && <span>Please log in to play online</span>} */}
      {/* <PlayOnline /> */}
    </div>
  );
};
