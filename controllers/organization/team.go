package organization

import (
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type TeamController struct {
	Service services.TeamService
}

func (instance *TeamController) Get(ctx iris.Context) mvc.Result {
	return response.Success("请求成功", response.Payload{})
}

func (instance *TeamController) Post(ctx iris.Context) mvc.Result {
	return response.Success("请求成功", response.Payload{})
}

func (instance *TeamController) PutBy(id string, ctx iris.Context) mvc.Result {
	return response.Success("请求成功", response.Payload{})
}

func (instance *TeamController) DeleteBy(id string) mvc.Result {
	return response.Success("请求成功", response.Payload{})
}
