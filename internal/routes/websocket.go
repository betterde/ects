package routes

import (
	"github.com/betterde/ects/internal/controllers/socket"
	"github.com/betterde/ects/internal/websocket"
	"github.com/kataras/iris/mvc"
)

func registerWebSocket(application *mvc.Application) {
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
	application.Register(ws.Upgrade)
	application.Handle(new(socket.WebSocketController))
}
