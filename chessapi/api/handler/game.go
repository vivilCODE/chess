package handler

import (
	"github.com/gofiber/websocket/v2"
	"github.com/vivilCODE/chess/chessapi/log"
	"github.com/vivilCODE/chess/db/models"
)

type Game struct {

}

type GameQueue struct {
	clients map[*websocket.Conn]models.User
}

func NewGameQueue() *GameQueue {
	return &GameQueue{
		clients: make(map[*websocket.Conn]models.User),
	}
}


func (q *GameQueue) QueueHandler(c *websocket.Conn) {
	log.Logger.Debug("received connection request to queue handler")



}