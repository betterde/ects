package routes

import (
	"github.com/betterde/ects/controllers/setting"
	"github.com/kataras/iris/v12/mvc"
)

func registerSetting(application *mvc.Application) {
	application.Handle(new(setting.Controller))
}
