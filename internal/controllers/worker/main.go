package worker

import (
	"github.com/betterde/ects/internal/models"
	"github.com/betterde/ects/internal/services"
	"github.com/betterde/ects/internal/utils/response"
	"github.com/go-xorm/builder"
	"github.com/kataras/iris"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"math/big"
	"strconv"
	"time"
)

type (
	Controller struct {
		Service services.WorkerService
	}

	CreateRequest struct {
		Name   string `json:"name" validate:"required"`
		Remark string `json:"remark"`
	}

	UpdateRequest struct {
		Name   string `json:"name" validate:"required"`
		Remark string `json:"remark"`
	}
)

// 获取节点列表
func (instance *Controller) Get(ctx iris.Context) *response.Response {
	var (
		page  int
		limit int
		start int
		name  string
		total int64
		err   error
	)
	page = 1
	limit = 10
	name = ""
	params := ctx.URLParams()

	if value, exist := params["page"]; exist == true {
		v, err := strconv.Atoi(value)
		if err != nil {

		}
		if v >= 0 {
			page = v
		}
	}

	if value, exist := params["limit"]; exist == true {
		v, err := strconv.Atoi(value)
		if err != nil {

		}
		if v >= 0 {
			limit = v
		}
	}

	if value, exist := params["name"]; exist == true {
		name = value
	}

	start = (page - 1) * limit
	workers := make([]models.Worker, 0)
	if name == "" {
		total, err = models.Engine.Count(&models.Worker{})
		err = models.Engine.Limit(limit, start).Find(&workers)
	} else {
		total, err = models.Engine.Where(builder.Like{"name", name}).Count(&models.Worker{})
		err = models.Engine.Where(builder.Like{"name", name}).Limit(limit, start).Find(&workers)
	}

	if err != nil {
		log.Println(err)
	}

	return response.Success("请求成功", response.Payload{
		"data": workers,
		"meta": &response.Meta{
			Limit: limit,
			Page:  page,
			Total: int(total),
		}})
}

// 创建节点
func (instance *Controller) Post(ctx iris.Context) *response.Response {
	var params CreateRequest
	validate := validator.New()
	if err := ctx.ReadJSON(&params); err != nil {
		log.Println(err)
	}

	err := validate.Struct(params)

	if err != nil {
		ctx.StatusCode(iris.StatusUnprocessableEntity)
		return response.ValidationError("请填写名称")
	}

	var worker models.Worker
	worker.ID = uuid.NewV4().String()
	worker.Name = params.Name
	worker.Remark = params.Remark
	worker.Status = models.STATUS_DISCONNECTED
	worker.IP = big.NewInt(0).Int64()
	worker.CreatedAt = time.Now().Format("2006-1-2 15:04:05")
	worker.UpdatedAt = time.Now().Format("2006-1-2 15:04:05")
	if err := worker.Store(); err != nil {
		log.Println(err)
		return response.Send(1062, "创建节点失败", err)
	}

	return response.Success("创建节点成功", response.Payload{"data": worker})
}

// 更新节点信息
func (instance *Controller) PutBy(id string, ctx iris.Context) *response.Response {
	var params UpdateRequest
	var worker models.Worker
	validate := validator.New()
	if err := ctx.ReadJSON(&params); err != nil {
		log.Println(err)
	}

	err := validate.Struct(params)

	if err != nil {
		ctx.StatusCode(iris.StatusUnprocessableEntity)
		return response.ValidationError("请填写名称")
	}

	result, err := models.Engine.Id(id).Get(&worker)

	if err != nil {

	}

	if result {
		worker.Name = params.Name
		worker.Remark = params.Remark
		if _, err := models.Engine.Id(id).Update(worker); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return response.Send(iris.StatusInternalServerError, "更新失败", make(map[string]interface{}))
		}
	}

	return response.Success("更新节点成功", response.Payload{"data": worker})
}

// 删除指定ID
func (instance *Controller) DeleteBy(id string) *response.Response {
	worker := &models.Worker{
		ID: id,
	}

	_, err := models.Engine.Delete(worker)
	if err != nil {

	}

	return response.Success("删除节点成功", response.Payload{"data": make(map[string]interface{})})
}
