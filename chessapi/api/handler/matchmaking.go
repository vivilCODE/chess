package handler

import (
	"encoding/json"
	"math/rand"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
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

		q.startGame(playerOne, playerTwo)
	}
}

// Create a game, store it in the database, and send the game id to both clients to start the game.
func (q *GameQueue) startGame(p1, p2 *gameClient) {
	gameId := uuid.NewString()
	
	pWhite, pBlack := assignBlackWhite(p1, p2)

	log.Logger.Debug("starting new game", "game id", gameId, "white", pWhite.user.Name, "black", pBlack.user.Name)
	

	// board := generateBoard()

	// game := model.Game{
	// 	ID: gameId,
	// 	PlayerWhite: pWhite.user,
	// 	PlayerBlack: pBlack.user,
	// 	Board: board,
	// 	Started: time.Now(),
	// }


	startMessage := []byte(gameId)
	// messageWithGame := append(startMessage, gameJSON...)

	// Send the start message to both players
	if err := pWhite.conn.WriteMessage(websocket.TextMessage, startMessage); err != nil {
		log.Logger.Error("unable to start game for white", "err", err)
	} else {
		log.Logger.Debug("game start message sent to white", "name", pWhite.user.Name)
	}

	if err := pBlack.conn.WriteMessage(websocket.TextMessage, startMessage); err != nil {
		log.Logger.Error("unable to start game for black", "err", err)
		} else {
		log.Logger.Debug("game start message sent to black", "name", pBlack.user.Name)
	}

}

// func assingBlackWhite assigns black or white pieces to the two players
// based on a random int
func assignBlackWhite(p1, p2 *gameClient) (*gameClient, *gameClient) {
	if rand.Intn(2) == 1 {
		return p1, p2
	} else {
		return p2, p1
	}
}

func generateBoard() model.Board {
	squares := []model.Square{}

	for y := 1; y <= 8; y++ {
		for x := 1; x <= 8; x++ {
			piece := model.NilPiece
			isWhite := true

			if isOdd(y) && isOdd(x) ||
				!isOdd(y) && !isOdd(x) {
				isWhite = false
			}

			switch {
			// White side of the board
			case x == 1 && y == 1 ||
				x == 8 && y == 1:
				piece = model.WhiteRook
			case x == 2 && y == 1 ||
				x == 7 && y == 1:
				piece = model.WhiteKnight
			case x == 3 && y == 1 ||
				x == 6 && y == 1:
				piece = model.WhiteBishop
			case x == 4 && y == 1:
				piece = model.WhiteQueen
			case x == 5 && y == 1:
				piece = model.WhiteKing
			case y == 2:
				piece = model.WhitePawn

			// Black side of the board
			case y == 7:
				piece = model.BlackPawn
			case x == 1 && y == 8 ||
				x == 8 && y == 8:
				piece = model.BlackRook
			case x == 2 && y == 8 ||
				x == 7 && y == 8:
				piece = model.BlackKnight
			case x == 3 && y == 8 ||
				x == 6 && y == 8:
				piece = model.BlackBishop
			case x == 4 && y == 8:
				piece = model.BlackQueen
			case x == 5 && y == 8:
				piece = model.BlackKing
			}

			squares = append(squares, model.Square{
				Pos: model.SquarePosition{X:x, Y: y},
				IsWhite: isWhite,
				Piece: piece,
			})
		}
	}
	return model.Board{
		Squares: squares,
	}
}

func isOdd(n int) bool {
	return n % 2 != 0
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



