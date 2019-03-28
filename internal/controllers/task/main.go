package task

import (
	"github.com/betterde/ects/internal/models"
	"github.com/betterde/ects/internal/utils/response"
	"github.com/kataras/iris"
)

type (
	Controller struct {
	}

	CreateRequest struct {
	}

	UpdateRequest struct {
	}
)

func (instance *Controller) Get(ctx iris.Context) *response.Response {
	tasks := new(models.Task)
	return response.Success("", response.Payload{
		"data": tasks,
		"meta": &response.Meta{

		}})
}

func (instance *Controller) Post(ctx iris.Context) *response.Response {
	var task models.Task
	return response.Success("创建任务成功", response.Payload{"data": task})
}

func (instance *Controller) PutBy(id string, ctx iris.Context) *response.Response {
	var task models.Task
	return response.Success("更新任务成功", response.Payload{"data": task})
}

func (instance *Controller) DeleteBy(id string) *response.Response {
	return response.Success("删除任务成功", response.Payload{"data": make(map[string]interface{})})
}
