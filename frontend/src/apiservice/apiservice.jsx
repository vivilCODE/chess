// export const NewGame = () => {
//   const req = parser.createNewGameRequest();

//   return new Promise((resolve, reject) => {
//     client.newGame(req, {}, (err, response) => {
//       if (err) {
//         reject(err);
//       } else {
//         resolve(response.toObject());
//       }
//     });
//   });
// };

// export const MakeMove = (game, from, to) => {
//   let req = parser.createMakeMoveRequest(game, from, to);

//   return new Promise((resolve, reject) => {
//     client.makeMove(req, {}, (err, response) => {
//       if (err) {
//         reject(err);
//       } else {
//         resolve(response.toObject());
//       }
//     });
//   });
// };

export const SignIn = async (code) => {
  console.log("json code: ", JSON.stringify({ code }));

  const res = await fetch("http://localhost:8080/chessapi/auth/signin", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ code }),
  });

  const data = await res.json();

  if (res.ok) {
    console.log("sign in request successful");
  } else {
    console.log("sign in request failed, err: ", data.error);
    return data.error;
  }

  return data;
};
