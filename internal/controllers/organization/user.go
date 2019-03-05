package organization

import (
	"ects/internal/models"
	"ects/internal/services"
	"github.com/kataras/iris"
)

type UserController struct {
	Service services.UserService
}

func (instance * UserController) Get(ctx iris.Context) (results []models.User) {
	return instance.Service.Users()
}