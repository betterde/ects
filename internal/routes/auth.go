package routes

import (
	"github.com/betterde/ects/internal/controllers/auth"
	"github.com/betterde/ects/internal/datasource"
	"github.com/betterde/ects/internal/repositories"
	"github.com/betterde/ects/internal/services"
	"github.com/kataras/iris/mvc"
)

func authentication(application *mvc.Application) {
	repo := repositories.NewUserRepository(datasource.Users)
	userService := services.NewUserService(repo)
	application.Register(userService)
	application.Handle(new(auth.AuthenticationController))
}
