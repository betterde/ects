package routes

import (
	"github.com/betterde/ects/internal/controllers/organization"
	"github.com/betterde/ects/internal/datasource"
	"github.com/betterde/ects/internal/repositories"
	"github.com/betterde/ects/internal/services"
	"github.com/kataras/iris/mvc"
)

func users(application *mvc.Application) {
	repo := repositories.NewUserRepository(datasource.Users)
	userService := services.NewUserService(repo)
	application.Register(userService)
	application.Handle(new(organization.UserController))
}
