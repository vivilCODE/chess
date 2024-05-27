import { useContext, useState } from "react";
import { UserContext } from "../../index";

export const PlayOnline = () => {
  const [user] = useContext(UserContext);
  const [waitingForGame, setWaitingForGame] = useState(false);
  const [foundGame, setFoundGame] = useState(false);

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
      const msg = e.data;

      switch (msg) {
        case "start game":
          setWaitingForGame(() => false);
          setFoundGame(() => true);

        case "waiting to find match":
          setWaitingForGame(() => true);
      }
    };
  };

  return (
    <>
      {user && <button onClick={enterQueue}>Play online</button>}
      {waitingForGame && <span>Waiting...</span>}
      {foundGame && <span>Found game!</span>}
    </>
  );
};
