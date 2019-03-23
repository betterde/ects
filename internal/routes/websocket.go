package routes

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/websocket"
)

func registerWebSocket(application *mvc.Application) {
	ws := websocket.New(websocket.Config{})
	application.Register(ws.Upgrade)
	//application.Handle(new())
}
