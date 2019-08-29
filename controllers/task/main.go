package task

import (
	"github.com/betterde/ects/internal/message"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/services"
	"github.com/go-xorm/builder"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"time"
)

type (
	Controller struct {
		Service services.TaskService
	}

	UpdateRequest struct {
		Name        string `json:"name" validate:"required"`
		Content     string `json:"content" validate:"required"`
		Description string `json:"description"`
	}
)

var (
	validate = validator.New()
)

// 获取任务列表
func (instance *Controller) Get(ctx iris.Context) mvc.Result {
	var (
		total int64
		err   error
	)
	scene := ctx.Params().GetStringDefault("scene", "table")
	tasks := make([]models.Task, 0)

	switch scene {
	case "table":
		search := ctx.URLParamDefault("search", "")
		page, limit, start := utils.Pagination(ctx)

		if search == "" {
			total, err = models.Engine.Limit(limit, start).FindAndCount(&tasks)
		} else {
			total, err = models.Engine.Where(builder.Like{"id", search}.Or(builder.Like{"name", search})).Limit(limit, start).FindAndCount(&tasks)
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
	case "selector":
		if err := models.Engine.Find(&tasks); err != nil {
			return response.InternalServerError("查询任务列表失败", err)
		}
	}

	return response.Success("数据使用场景有误", response.Payload{"data": make([]interface{}, 0)})
}

// 创建任务
func (instance *Controller) Post(ctx iris.Context) mvc.Result {
	task := models.Task{}

	if err := ctx.ReadJSON(&task); err != nil {
		return response.InternalServerError("参数解析失败", err)
	}

	if err := validate.Struct(task); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("task", validationErrors))
	}

	task.Id = uuid.NewV4().String()

	if err := task.Store(); err != nil {
		return response.InternalServerError("Failed to create taks", err)
	}

	return response.Success("Created successful", response.Payload{"data": task})
}

// 更新任务
func (instance *Controller) PutBy(id string, ctx iris.Context) mvc.Result {
	var params UpdateRequest
	validate := validator.New()

	if err := ctx.ReadJSON(&params); err != nil {
		return response.InternalServerError("参数解析失败", err)
	}

	if err := validate.Struct(params); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("task", validationErrors))
	}

	task := &models.Task{
		Id:          id,
		Name:        params.Name,
		Content:     params.Content,
		Description: params.Description,
		UpdatedAt:   utils.Time(time.Now()),
	}

	if err := task.Update(); err != err {
		return response.InternalServerError("更新失败", err)
	}

	return response.Success("更新成功", response.Payload{"data": task})
}

// 删除任务
func (instance *Controller) DeleteBy(id string) mvc.Result {
	task := &models.Task{
		Id: id,
	}

	if err := task.Destroy(); err != nil {
		return response.InternalServerError("Failed to deleted task", err)
	}

	return response.Success("Deleted successful", response.Payload{"data": make(map[string]interface{})})
}
