package routes

import (
	"github.com/betterde/ects/controllers/log"
	"github.com/kataras/iris/v12/mvc"
)

func registerLog(application *mvc.Application) {
	application.Handle(new(log.Controller))
}
