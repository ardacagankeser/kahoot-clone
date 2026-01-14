package handler

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"quiz.com/quiz/internal/service"
)

type WebsocketHandler struct {
	netService *service.NetService
}

func NewWebsocketHandler(netService *service.NetService) WebsocketHandler {
	return WebsocketHandler{
		netService: netService,
	}
}

func (h WebsocketHandler) Ws(con *websocket.Conn) {
	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = con.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}

		h.netService.OnIncomingMessage(con, mt, msg)
	}
}
