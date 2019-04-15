package routes

import (
	"github.com/betterde/ects/controllers/initialize"
	"github.com/kataras/iris/mvc"
)

func Initialize(application *mvc.Application) {
	application.Handle(new(initialize.Controller))
}