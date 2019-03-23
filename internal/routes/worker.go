package routes

import (
	"github.com/betterde/ects/internal/controllers/node"
	"github.com/betterde/ects/internal/services"
	"github.com/kataras/iris/mvc"
)

func registerWorker(application *mvc.Application) {
	application.Register(services.NewWorkerService())
	application.Handle(new(node.WorkerController))
}