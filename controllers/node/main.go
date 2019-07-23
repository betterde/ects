package node

import (
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/services"
	"github.com/go-xorm/builder"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"strconv"
	"time"
)

type (
	Controller struct {
		Service services.NodeService
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
func (instance *Controller) Get(ctx iris.Context) mvc.Result {
	var (
		page   int
		limit  int
		start  int
		search string
		total  int64
		err    error
	)
	page = 1
	limit = 10
	search = ""
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

	if value, exist := params["search"]; exist == true {
		search = value
	}

	start = (page - 1) * limit
	workers := make([]models.Node, 0)
	if search == "" {
		total, err = models.Engine.Count(&models.Node{})
		err = models.Engine.Limit(limit, start).Find(&workers)
	} else {
		total, err = models.Engine.Where(builder.Like{"name", search}).Count(&models.Node{})
		err = models.Engine.Where(builder.Like{"name", search}).Limit(limit, start).Find(&workers)
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
func (instance *Controller) Post(ctx iris.Context) mvc.Result {
	var params CreateRequest
	validate := validator.New()
	if err := ctx.ReadJSON(&params); err != nil {
		return response.InternalServerError("解析参数失败", err)
	}

	if err := validate.Struct(params); err != nil {
		return response.ValidationError("请填写名称")
	}

	worker := models.Node{
		Id: uuid.NewV4().String(),
		Name: params.Name,
		Description: params.Remark,
		Status: models.ONLINE,
		Host: "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := worker.Store(); err != nil {
		return response.InternalServerError("创建节点失败", err)
	}

	return response.Success("创建节点成功", response.Payload{"data": worker})
}

// 更新节点信息
func (instance *Controller) PutBy(id string, ctx iris.Context) mvc.Result {
	var params UpdateRequest
	var worker models.Node
	validate := validator.New()
	if err := ctx.ReadJSON(&params); err != nil {
		log.Println(err)
	}

	err := validate.Struct(params)

	if err != nil {
		return response.ValidationError("请填写名称")
	}

	result, err := models.Engine.Id(id).Get(&worker)

	if err != nil {

	}

	if result {
		worker.Name = params.Name
		worker.Description = params.Remark
		if _, err := models.Engine.Id(id).Update(worker); err != nil {
			return response.Send(iris.StatusInternalServerError, "更新失败", make(map[string]interface{}))
		}
	}

	return response.Success("更新节点成功", response.Payload{"data": worker})
}

// 删除指定ID
func (instance *Controller) DeleteBy(id string) mvc.Result {
	worker := &models.Node{
		Id: id,
	}

	_, err := models.Engine.Delete(worker)
	if err != nil {

	}

	return response.Success("删除节点成功", response.Payload{"data": make(map[string]interface{})})
}
