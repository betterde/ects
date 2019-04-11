package log

import (
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type (
	Controller struct {
		Service services.LogService
	}
)

func (instance *Controller) Get(ctx iris.Context) mvc.Result {
	return response.Success("", response.Payload{"data": make(map[string]interface{})})
}
