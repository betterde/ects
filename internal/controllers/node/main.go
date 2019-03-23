package node

import (
	"github.com/betterde/ects/internal/models"
	"github.com/betterde/ects/internal/services"
	"github.com/betterde/ects/internal/utils/response"
	"github.com/kataras/iris"
	"log"
)

type (
	WorkerController struct {
		Service services.WorkerService
	}

	CreateRequest struct {
		Name string `json:"name"`
		Remark string `json:"remark"`
	}
)

func (instance *WorkerController) Get() {

}

// 创建节点
func (instance *WorkerController) Post(ctx iris.Context) {
	var params CreateRequest
	if err := ctx.ReadJSON(&params); err != nil {
		log.Println(err)
	}

	var worker models.Worker
	worker.Name = params.Name
	worker.Remark = params.Remark
	log.Println(worker)
	if err := worker.Store(); err != nil {
		log.Println(err)
		if _, err = ctx.JSON(response.Send(1062, "创建节点失败", err)); err != nil {
			// TODO
		}
		return
	}

	if _, err := ctx.JSON(response.Success("创建节点成功", worker)); err != nil {
		// TODO
	}
}