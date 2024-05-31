import { useContext, useEffect, useState } from "react";
import { UserContext } from "../../index";

import "./chatroom.css";

export const Chatroom = () => {
  const [user] = useContext(UserContext);
  const [socket, setSocket] = useState(null);
  const [messages, setMessages] = useState([]);

  //   Once a user has logged in, call enterChatroom to open a websocket connection
  useEffect(() => {
    if (user === null) {
      return;
    }

    enterChatroom(user);
  }, [user]);

  //   Initialise the websocket connection and set the socket state
  const enterChatroom = async (u) => {
    let s = new WebSocket("ws://localhost:8080/chessapi/ws/chatroom");

    setSocket(() => s);
  };

  //   Read text value from input field and send it to the socket
  const sendMessage = () => {
    const messageField = document.querySelector(".chatbox__input");

    const message = messageField.value;
    messageField.value = "";
    socket.send(message);
  };

  useEffect(() => {
    if (socket !== null) {
      // When socket connection opens, log that a connection has been established and send
      // some user data as the first message. The server uses this to keep track of clients connected to the chatroom.
      socket.onopen = function () {
        console.log("websocket connection established, sending user: ", user);

        const jsonUser = JSON.stringify({
          id: user.user.id,
          name: user.user.name,
          email: user.user.email,
          signed_up: user.user.signed_up,
        });

        console.log("jsonuser: ", jsonUser);

        socket.send(jsonUser);
      };

      // When a message is received from the socket, log it and append it to the messages state
      socket.onmessage = (e) => {
        console.log("message from server: ", e.data);
        setMessages((prev) => [...prev, e.data]);
      };

      socket.onclose = function () {
        console.log("WebSocket connection closed");
      };
    }
  }, [socket]);

  return (
    <div className="chatbox">
      <h3>Chatroom</h3>
      <ul className="chatbox__chat-log">
        {messages.map((message) => {
          return (
            <li className="chatbox__message">
              <span>{message}</span>
            </li>
          );
        })}
      </ul>
      <input type="text" className="chatbox__input" />
      <button className="chatbox__send-btn" onClick={sendMessage}>
        Send
      </button>
    </div>
  );
};

// let socket;

// document.getElementById("connect").onclick = function () {
//   const name = document.getElementById("name").value;
//   const email = document.getElementById("email").value;

//   socket = new WebSocket("ws://localhost:3000/ws");

//   socket.onopen = function () {
//     console.log("WebSocket connection established");
//     document.getElementById("messages").textContent +=
//       "Connected to chatroom\n";

//     const user = JSON.stringify({ id: "1", name: name, email: email });
//     socket.send(user);
//   };

//   socket.onmessage = function (event) {
//     console.log("Message from server: ", event.data);
//     document.getElementById("messages").textContent +=
//       "Message from server: " + event.data + "\n";
//   };

//   socket.onclose = function () {
//     console.log("WebSocket connection closed");
//     document.getElementById("messages").textContent +=
//       "Disconnected from chatroom\n";
//   };
// };

// document.getElementById("send").onclick = function () {
//   if (socket && socket.readyState === WebSocket.OPEN) {
//     const message = document.getElementById("message").value;
//     socket.send(message);
//   }
// };
