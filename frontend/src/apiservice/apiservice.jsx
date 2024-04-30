import { ChessApiClient } from "../proto/chess_grpc_web_pb";
import { PingRequest } from "../proto/chess_pb";
import * as parser from "./parser";
const EnvoyURL = "http://localhost:8081";
const client = new ChessApiClient(EnvoyURL);

export const Ping = () => {
  const req = new PingRequest();
  client.ping(req, {}, (err, response) => {
    console.log(response.getResponse());
  });
};

export const NewGame = () => {
  const req = parser.createNewGameRequest();

  return new Promise((resolve, reject) => {
    client.newGame(req, {}, (err, response) => {
      if (err) {
        reject(err);
      } else {
        resolve(response.toObject());
      }
    });
  });
};

export const MakeMove = (game, from, to) => {
  let req = parser.createMakeMoveRequest(game, from, to);

  return new Promise((resolve, reject) => {
    client.makeMove(req, {}, (err, response) => {
      if (err) {
        reject(err);
      } else {
        resolve(response.toObject());
      }
    });
  });
};

export const SignIn = (code) => {
  let req = parser.createSignInRequest(code);
  return new Promise((resolve, reject) => {
    client.signIn(req, {}, (err, response) => {
      if (err) {
        reject(err);
      } else {
        resolve(response.toObject());
      }
    });
  });
};
