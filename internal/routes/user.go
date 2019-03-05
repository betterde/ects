package routes

import (
	"ects/internal/controllers/organization"
	"ects/internal/datasource"
	"ects/internal/repositories"
	"ects/internal/services"
	"github.com/kataras/iris/mvc"
)

func users(application *mvc.Application) {
	repo := repositories.NewUserRepository(datasource.Users)
	userService := services.NewUserService(repo)
	application.Register(userService)
	application.Handle(new(organization.UserController))
}
