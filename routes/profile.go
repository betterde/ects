package routes

import (
	"github.com/betterde/ects/controllers/account"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris/v12/mvc"
)

func registerProfile(application *mvc.Application) {
	application.Register(services.NewUserService())
	application.Handle(new(account.Controller))
}
