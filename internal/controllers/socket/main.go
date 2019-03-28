package socket

import (
	"github.com/betterde/ects/internal/websocket"
	"github.com/kataras/iris"
	"log"
	"sync/atomic"
)

type (
	Controller struct {
		Conn websocket.Connection
	}
)

var visits uint64

func increment() uint64 {
	return atomic.AddUint64(&visits, 1)
}

func decrement() uint64 {
	return atomic.AddUint64(&visits, ^uint64(0))
}

func (c *Controller) Get(ctx iris.Context) {
	if c.Conn.Err() != nil {
		log.Println(c.Conn.Err())
	}
	//c.Conn.OnLeave(c.onLeave)
	//c.Conn.On("visit", c.update)
	//在所有事件回调注册后调用它
	//c.Conn.Wait()
}

func (c *Controller) update() {
	newCount := increment()
	c.Conn.To(websocket.All).Emit("visit", newCount)
}

func (c *Controller) onLeave(roomName string) {
	newCount := decrement()
	c.Conn.To(websocket.Broadcast).Emit("visit", newCount)
}