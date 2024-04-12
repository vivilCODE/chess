import { ChessApiClient } from "../proto/chess_grpc_web_pb";
import { PingRequest, NewGameRequest } from "../proto/chess_pb";
import { parseGame } from "./parser";
const EnvoyURL = "http://localhost:8081";
const client = new ChessApiClient(EnvoyURL);

export const Ping = () => {
  const req = new PingRequest();
  client.ping(req, {}, (err, response) => {
    console.log(response.getResponse());
  });
};

export const NewGame = () => {
  const req = new NewGameRequest();

  return new Promise((resolve, reject) => {
    client.newGame(req, {}, (err, response) => {
      if (err) {
        reject(err);
      } else {
        let game = parseGame(response);
        resolve(game);
      }
    });
  });
};
