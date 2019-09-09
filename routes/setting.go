package routes

import (
	"github.com/betterde/ects/controllers/setting"
	"github.com/kataras/iris/mvc"
)

func registerSetting(application *mvc.Application) {
	application.Handle(new(setting.Controller))
}
