package routes

import (
	"github.com/betterde/ects/controllers/role"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris/mvc"
)

func registerRole(application *mvc.Application) {
	application.Register(services.NewRoleService())
	application.Handle(new(role.Controller))
}