package routes

import (
	"github.com/betterde/ects/controllers/node"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris/mvc"
)

func registerWorker(application *mvc.Application) {
	application.Register(services.NewNodeService())
	application.Handle(new(node.Controller))
}