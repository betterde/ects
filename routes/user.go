package routes

import (
	"github.com/betterde/ects/controllers/organization"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris/v12/mvc"
)

func registerUser(application *mvc.Application) {
	application.Register(services.NewUserService())
	application.Handle(new(organization.UserController))
}
