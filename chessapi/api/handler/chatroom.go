package handler

import (
	"encoding/json"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/vivilCODE/chess/chessapi/api/model"
	"github.com/vivilCODE/chess/chessapi/log"
)

type Chatroom struct {
	// clients hollds all current clients in this room
	Clients map[*websocket.Conn]model.User

	Mu sync.Mutex
}

func NewChatroom() *Chatroom {
	return &Chatroom{
		Clients: make(map[*websocket.Conn]model.User),
	}
}

func (room *Chatroom) ChatroomHandler(c *websocket.Conn) {
	log.Logger.Debug("received connection request to /chatroom")
	
		
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Logger.Error("unable to read socket message", "err", err)
	
		return
	}

	var user model.User

	if err := json.Unmarshal(message, &user); err != nil {
		log.Logger.Error("unable to unmarshal user in first socket message", "err", err)
	
		return
	}

	room.Mu.Lock()
	room.Clients[c] = user

	room.Mu.Unlock()

	
	log.Logger.Debug("new websocket connection established for", "user", user)

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
			log.Logger.Error("read error", "err", err)
			
			break
		}

		log.Logger.Debug("received", "message", message, "from", "user", user.Name)
			
		// Broadcast the message to all clients in the chatroom
		room.Mu.Lock()
		for client, _ := range room.Clients {
			if err := client.WriteMessage(messageType, []byte(user.Name+": "+string(message))); err != nil {
				log.Logger.Error("write error", "err", err)
			}
		}

		room.Mu.Unlock()
	}
	
}
