package pipeline

import (
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type (
	Controller struct {
		Service services.PipelineService
	}
)

func (instance *Controller) Get(ctx iris.Context) mvc.Result {
	return response.Success("", response.Payload{})
}