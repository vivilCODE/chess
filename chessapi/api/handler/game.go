package handler

import (
	"encoding/json"
	"math/rand"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/vivilCODE/chess/chessapi/api/model"
	"github.com/vivilCODE/chess/chessapi/log"
)

type Client struct {
	Conn *websocket.Conn
	User model.User
	IsWhite bool
}

type GameQueue struct {
	Queue []*Client

	queueIsFull chan bool

	sync.Mutex
}

func NewGameQueue() *GameQueue {
	q :=  &GameQueue{
		Queue: make([]*Client, 0),	
		queueIsFull: make(chan bool),
	}

	go q.matchMaker()

	return q
}

// func matchMaker waits for a notification saying that the queue is full,
// then is calls start game with the two first players in the queue.
func (q *GameQueue) matchMaker() {
	for {
		<- q.queueIsFull
		q.Lock()

		if len(q.Queue) < 2 {
			q.Unlock()
			continue
		}

		playerOne := q.Queue[0]
		playerTwo := q.Queue[1]
		q.Queue = q.Queue[2:]

		q.Unlock()

		q.StartGame(playerOne, playerTwo)
	}
}

func (q *GameQueue) StartGame(playerOne, playerTwo *Client) {
	log.Logger.Debug("two players in queue, start game between", "player1", playerOne.User.Name, "player2", playerTwo.User.Name)

	startMessage := []byte("start game")

	p1, p2:=	assignBlackWhite(playerOne, playerTwo)

	log.Logger.Debug("colors", p1.User.Name, p1.IsWhite, p2.User.Name, p2.IsWhite)

	// Send the start message to both players
	if err := p1.Conn.WriteMessage(websocket.TextMessage, startMessage); err != nil {
		log.Logger.Error("unable to start game for player1", "err", err)
	} else {
		log.Logger.Debug("game start message sent to player1", "name", p1.User.Name)
	}

	if err := p2.Conn.WriteMessage(websocket.TextMessage, startMessage); err != nil {
		log.Logger.Error("unable to start game for player2", "err", err)
	} else {
		log.Logger.Debug("game start message sent to player2", "name", p2.User.Name)
	}
}


func (q *GameQueue) AddClient(client *Client) {
	q.Lock()
	defer q.Unlock()
	q.Queue = append(q.Queue, client)

	log.Logger.Debug("Added player to queue", "name", client.User.Name, "current queue length", len(q.Queue))

	
	if len(q.Queue) >=2 {
		log.Logger.Debug("queue is full, notifying handler", "queue length", len(q.Queue))		
		q.queueIsFull <- true
	}
}

// func QueueHandler accepts incoming requests and adds clients to the queue which will 
// take them to a gaee once two people have joined it
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
	
	q.AddClient(&Client{
		Conn: c,
		User: user,
	})

	// This loop is just to keep the connection alive by stopping the handler for exiting
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Logger.Error("error reading message", "player", user.Name, "err", err)
		}
	}
}


// func assingBlackWhite assigns black or white pieces to the two players
// based on a random int
func assignBlackWhite(playerOne, playerTwo *Client) (*Client, *Client) {
	p1 := playerOne
	p2 := playerTwo
	
	if rand.Intn(2) == 1 {
		p1.IsWhite = true
		p2.IsWhite = false
	} else {
		p1.IsWhite = false
		p2.IsWhite = true
	}

	return p1, p2
} 