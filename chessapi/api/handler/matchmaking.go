package handler

import (
	"encoding/json"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/vivilCODE/chess/chessapi/api/model"
	"github.com/vivilCODE/chess/chessapi/log"
)

type gameClient struct {
	conn *websocket.Conn
	user model.User
}

type GameQueue struct {
	queue []*gameClient

	queueIsFull chan bool

	sync.Mutex
}

func NewGameQueue() *GameQueue {
	q :=  &GameQueue{
		queue: make([]*gameClient, 0),	
		queueIsFull: make(chan bool),
	}

	go q.matchMaker()

	return q
}

// func matchMaker waits for a notification saying that the queue is full,
// then it calls start game with the two first players in the queue.
func (q *GameQueue) matchMaker() {
	for {
		<- q.queueIsFull
		q.Lock()

		if len(q.queue) < 2 {
			q.Unlock()
			continue
		}

		playerOne := q.queue[0]
		playerTwo := q.queue[1]
		q.queue = q.queue[2:]

		q.Unlock()


		ge := NewGameEngine(playerOne, playerTwo)
		ge.StartGame()
	}
}

func (q *GameQueue) addClient(client *gameClient) {
	q.Lock()
	defer q.Unlock()
	q.queue = append(q.queue, client)

	log.Logger.Debug("Added player to queue", "name", client.user.Name, "current queue length", len(q.queue))

	
	if len(q.queue) >=2 {
		log.Logger.Debug("queue is full, notifying handler", "queue length", len(q.queue))		
		q.queueIsFull <- true
	}
}

// func QueueHandler accepts incoming requests and adds clients to the queue which will 
// take them to a game once two people have joined it
func (q *GameQueue) QueueHandler(c *websocket.Conn) {
	log.Logger.Debug("received connection request to queue handler")
	
	defer func() {
		if err := c.Close(); err != nil {
			log.Logger.Error("error closing websocket connection", "err", err)
		}
	}()

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

	if err := c.WriteMessage(websocket.TextMessage, []byte("waiting to find match")); err != nil {
		log.Logger.Error("unable to message client", "player", user.Name, "err", err)
	}
	
	q.addClient(&gameClient{
		conn: c,
		user: user,
	})

	// This loop is just to keep the connection alive by stopping the handler from exiting
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Logger.Error("error reading message", "player", user.Name, "err", err)
		}
	}
}
