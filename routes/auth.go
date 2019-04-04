package routes

import (
	"github.com/betterde/ects/controllers/auth"
	"github.com/betterde/ects/internal/services"
	"github.com/kataras/iris/mvc"
)

func authentication(application *mvc.Application) {
	application.Register(services.NewUserService())
	application.Handle(new(auth.Controller))
}
