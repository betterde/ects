package routes

import (
	"github.com/betterde/ects/internal/controllers/auth"
	"github.com/betterde/ects/internal/services"
	"github.com/kataras/iris/mvc"
)

func authentication(application *mvc.Application) {
	userService := services.NewUserService()
	application.Register(userService)
	application.Handle(new(auth.AuthenticationController))
}
