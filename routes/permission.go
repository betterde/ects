package routes

import (
	"github.com/betterde/ects/controllers/permission"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris/mvc"
)

func registerPermission(application *mvc.Application) {
	application.Register(services.NewPermissionService())
	application.Handle(new(permission.Controller))
}
