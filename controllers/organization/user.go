package organization

import (
	"github.com/betterde/ects/internal/services"
	"github.com/betterde/ects/models"
	"github.com/kataras/iris"
)

type UserController struct {
	Service services.UserService
}

func (instance * UserController) Get(ctx iris.Context) (results []models.User) {
	return instance.Service.Users()
}