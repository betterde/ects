package routes

import (
	"github.com/betterde/ects/controllers/install"
	"github.com/kataras/iris/mvc"
)

func installation(application *mvc.Application) {
	application.Handle(new(install.Controller))
}