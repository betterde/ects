package permission

import (
	"github.com/betterde/ects/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type (
	Controller struct {
		Service services.PermissionService
	}
)

func (instance *Controller) Get(ctx iris.Context) mvc.Result {

}

func (instance *Controller) Post(ctx iris.Context) mvc.Result {

}

func (instance *Controller) PutBy(id string, ctx iris.Context) mvc.Result {

}

func (instance *Controller) DeleteBy(id string) mvc.Result {

}