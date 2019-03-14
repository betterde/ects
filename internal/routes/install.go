package routes

import (
	"github.com/betterde/ects/internal/controllers/install"
	"github.com/betterde/ects/internal/services"
	"github.com/kataras/iris/mvc"
)

func installation(application *mvc.Application) {
	installService := services.NewInstallService()
	application.Register(installService)
	application.Handle(new(install.InstallationController))
}