package routes

import (
	"github.com/betterde/ects/controllers/account"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris/mvc"
)

func registerProfile(application *mvc.Application) {
	userService := services.NewUserService()
	application.Register(userService)
	application.Handle(new(account.Controller))
}
