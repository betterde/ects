package routes

import (
	"github.com/betterde/ects/controllers/dashboard"
	"github.com/kataras/iris/v12/mvc"
)

func registerDashboard(application *mvc.Application) {
	application.Handle(new(dashboard.Controller))
}
