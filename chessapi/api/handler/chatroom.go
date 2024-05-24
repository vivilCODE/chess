package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/vivilCODE/chess/chessapi/api/model"
)

type Chatroom struct {
	// clients hollds all current clients in this room
	Clients map[*websocket.Conn]model.User

	// join is a channel for clients wishing to join the room
	// join chan *client

	// leave chan *client

	// incoming messages that should be sent to other users
	// forward chan []byte

	Mu sync.Mutex
}

func NewChatroom() *Chatroom {
	return &Chatroom{
		Clients: make(map[*websocket.Conn]model.User),
	}
}

func (room *Chatroom) ChatroomHandler(c *websocket.Conn) {

	fmt.Printf("received connection request to /chatroom")
		
	_, message, err := c.ReadMessage()
	if err != nil {
		fmt.Printf("unable to read socket message, error: %v\n", err)
		return
	}

	var user model.User

	if err := json.Unmarshal(message, &user); err != nil {
		fmt.Printf("unable to unmarshal user in first socket message, error: %v\n", err)
		return
	}

	room.Mu.Lock()
	room.Clients[c] = user

	room.Mu.Unlock()

	
	fmt.Printf("new websocket connection established for user %v\n", user)

	defer func() {
		// Remove the client from the chatroom on disconnect
		room.Mu.Lock()
		delete(room.Clients, c)
		room.Mu.Unlock()
		c.Close()
	}()
	
	// Message handling loop
	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Printf("Received from %s: %s", user.Name, message)
	
		// Broadcast the message to all clients in the chatroom
		room.Mu.Lock()
		for client, _ := range room.Clients {
			if err := client.WriteMessage(messageType, []byte(user.Name+": "+string(message))); err != nil {
				fmt.Println("Write error:", err)
			}
		}

		room.Mu.Unlock()
	}
	
}
