package routes

import (
	"github.com/betterde/ects/controllers/organization"
	"github.com/betterde/ects/internal/services"
	"github.com/kataras/iris/mvc"
)

func users(application *mvc.Application) {
	userService := services.NewUserService()
	application.Register(userService)
	application.Handle(new(organization.UserController))
}
