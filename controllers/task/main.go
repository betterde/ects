package task

import (
	"github.com/betterde/ects/internal/message"
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
		Service services.TaskService
	}

	CreateRequest struct {
		Name        string `json:"name" validate:"required"`
		Content     string `json:"content" validate:"required"`
		Event       string `json:"event" validate:"required"`
		Mode        string `json:"mode" validate:"required"`
		Overlap     bool   `json:"overlap" validate:"required"`
		Timeout     int    `json:"timeout" validate:"gte=0"`
		Interval    int    `json:"interval" validate:"gte=0"`
		Retries     int    `json:"retries" validate:"gte=0"`
		Status      string `json:"status" validate:"required"`
		Description string `json:"description"`
	}

	UpdateRequest struct {
		Name        string `json:"name" validate:"required"`
		Content     string `json:"content" validate:"required"`
		Event       string `json:"event" validate:"required"`
		Mode        string `json:"mode" validate:"required"`
		Overlap     bool   `json:"overlap" validate:"required"`
		Timeout     int    `json:"timeout" validate:"gte=0"`
		Interval    int    `json:"interval" validate:"gte=0"`
		Retries     int    `json:"retries" validate:"gte=0"`
		Status      string `json:"status" validate:"required"`
		Description string `json:"description"`
	}
)

var (
	validate = validator.New()
)

// 获取任务俩表
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

	start = (page - 1) * limit

	tasks := make([]models.Task, 0)

	if search == "" {
		total, err = models.Engine.Count(&models.Task{})
		err = models.Engine.Limit(limit, start).Find(&tasks)
	} else {
		total, err = models.Engine.Where(builder.Like{"name", search}).Count(&models.Task{})
		err = models.Engine.Where(builder.Like{"name", search}).Limit(limit, start).Find(&tasks)
	}

	if err != nil {
		log.Println(err)
	}

	return response.Success("请求成功", response.Payload{
		"data": tasks,
		"meta": &response.Meta{
			Limit: limit,
			Page:  page,
			Total: int(total),
		}})
}

// 创建任务
func (instance *Controller) Post(ctx iris.Context) mvc.Result {
	var (
		params CreateRequest
	)

	if err := ctx.ReadJSON(&params); err != nil {
		return response.InternalServerError("解析参数失败", err)
	}

	if err := validate.Struct(params); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("task", validationErrors))
	}

	task := &models.Task{
		Id:          uuid.NewV4().String(),
		Name:        params.Name,
		Content:     params.Content,
		Mode:        params.Mode,
		Description: params.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := task.Store(); err != nil {
		return response.InternalServerError("创建任务失败", err)
	}

	return response.Success("创建任务成功", response.Payload{"data": task})
}

// 更新任务
func (instance *Controller) PutBy(id string, ctx iris.Context) mvc.Result {
	var params UpdateRequest
	validate := validator.New()

	if err := ctx.ReadJSON(&params); err != nil {
		return response.InternalServerError("解析参数失败", err)
	}

	if err := validate.Struct(params); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("task", validationErrors))
	}

	task := &models.Task{
		Id:          id,
		Name:        params.Name,
		Content:     params.Content,
		Mode:        params.Mode,
		Description: params.Description,
		UpdatedAt:   time.Now(),
	}

	if err := task.Update(); err != err {
		return response.InternalServerError("创建任务失败", err)
	}

	return response.Success("更新任务成功", response.Payload{"data": task})
}

// 删除任务
func (instance *Controller) DeleteBy(id string) mvc.Result {
	task := &models.Task{
		Id: id,
	}

	if err := task.Destroy(); err != nil {
		return response.InternalServerError("删除任务失败", err)
	}

	return response.Success("删除任务成功", response.Payload{"data": make(map[string]interface{})})
}
