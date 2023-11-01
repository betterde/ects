package routes

import (
	"github.com/betterde/ects/controllers/initialize"
	"github.com/kataras/iris/v12/mvc"
)

func Initialize(application *mvc.Application) {
	application.Handle(new(initialize.Controller))
}
