package routes

import (
	"github.com/betterde/ects/internal/controllers/account"
	"github.com/betterde/ects/internal/services"
	"github.com/kataras/iris/mvc"
)

func profile(application *mvc.Application) {
	userService := services.NewUserService()
	application.Register(userService)
	application.Handle(new(account.ProfileController))
}
