package task

import (
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type (
	Controller struct {
	}

	CreateRequest struct {
		Name        string    `json:"name"`
		ParentID    string    `json:"parent_id"`
		Content     string    `json:"content"`
		Event       string    `json:"event"`
		Mode        string    `json:"mode"`
		Overlap     bool      `json:"overlap"`
		Timeout     int       `json:"timeout"`
		Interval    int       `json:"interval"`
		Retries     int       `json:"retries"`
		Status      string    `json:"status"`
		Triggering  time.Time `json:"triggering"`
		Description string    `json:"description"`
	}

	UpdateRequest struct {
		Name        string    `json:"name"`
		ParentID    string    `json:"parent_id"`
		Content     string    `json:"content"`
		Event       string    `json:"event"`
		Mode        string    `json:"mode"`
		Overlap     bool      `json:"overlap"`
		Timeout     int       `json:"timeout"`
		Interval    int       `json:"interval"`
		Retries     int       `json:"retries"`
		Status      string    `json:"status"`
		Triggering  time.Time `json:"triggering"`
		Description string    `json:"description"`
	}
)

func (instance *Controller) Get(ctx iris.Context) mvc.Result {
	tasks := new(models.Task)
	return response.Success("", response.Payload{
		"data": tasks,
		"meta": &response.Meta{

		}})
}

// 创建任务
func (instance *Controller) Post(ctx iris.Context) mvc.Result {
	var params CreateRequest
	validate := validator.New()
	if err := ctx.ReadJSON(&params); err != nil {
		return response.InternalServerError("解析参数失败", err)
	}

	if err := validate.Struct(params); err != nil {
		return response.ValidationError("请填写名称")
	}

	var task models.Task
	task.ID = uuid.NewV4().String()
	task.Name = params.Name
	task.ParentID = params.ParentID
	task.Content = params.Content
	task.Event = params.Event
	task.Mode = params.Mode
	task.Overlap = params.Overlap
	task.Timeout = params.Timeout
	task.Interval = params.Interval
	task.Retries = params.Retries
	task.Status = params.Status
	task.Triggering = params.Triggering
	task.Description = params.Description
	task.CreatedAt = time.Now().Format("2006-1-2 15:04:05")
	task.UpdatedAt = time.Now().Format("2006-1-2 15:04:05")
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
		return response.ValidationError("请填写名称")
	}

	var task models.Task

	return response.Success("更新任务成功", response.Payload{"data": task})
}

// 删除任务
func (instance *Controller) DeleteBy(id string) mvc.Result {
	task := &models.Task{
		ID: id,
	}

	if err := task.Destroy(); err != nil {
		return response.InternalServerError("删除任务失败", err)
	}

	return response.Success("删除任务成功", response.Payload{"data": make(map[string]interface{})})
}
