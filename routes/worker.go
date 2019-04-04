package routes

import (
	"github.com/betterde/ects/controllers/worker"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris/mvc"
)

func registerWorker(application *mvc.Application) {
	application.Register(services.NewWorkerService())
	application.Handle(new(worker.Controller))
}