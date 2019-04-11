package routes

import (
	"github.com/betterde/ects/controllers/log"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris/mvc"
)

func registerLog(application *mvc.Application) {
	application.Register(services.NewLogService())
	application.Handle(new(log.Controller))
}