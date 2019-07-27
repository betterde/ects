package node

import (
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

// Get nodes list
func (instance *Controller) Get(ctx iris.Context) mvc.Result {
	var (
		total  int64
		err    error
	)
	search := ctx.Params().GetStringDefault("search", "")
	page, limit, start := utils.Pagination(ctx)
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

	return response.Success("Success", response.Payload{
		"data": workers,
		"meta": &response.Meta{
			Limit: limit,
			Page:  page,
			Total: int(total),
		}})
}

// Create node
func (instance *Controller) Post(ctx iris.Context) mvc.Result {
	var params CreateRequest
	validate := validator.New()
	if err := ctx.ReadJSON(&params); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(params); err != nil {
		return response.ValidationError("Please fill in the name")
	}

	worker := models.Node{
		Id: uuid.NewV4().String(),
		Name: params.Name,
		Description: params.Remark,
		Status: models.ONLINE,
		Host: "",
		CreatedAt: utils.Time(time.Now()),
		UpdatedAt: utils.Time(time.Now()),
	}
	if err := worker.Store(); err != nil {
		return response.InternalServerError("Failed to create node", err)
	}

	return response.Success("Created successful", response.Payload{"data": worker})
}

// Modify node attribute
func (instance *Controller) PutBy(id string, ctx iris.Context) mvc.Result {
	var params UpdateRequest
	var worker models.Node
	validate := validator.New()
	if err := ctx.ReadJSON(&params); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(params); err != nil {
		return response.ValidationError("Failed to validate params")
	}

	result, err := models.Engine.Id(id).Get(&worker)

	if err != nil {
		return response.NotFound("Node does not exist")
	}

	if result {
		worker.Name = params.Name
		worker.Description = params.Remark
		if _, err := models.Engine.Id(id).Update(worker); err != nil {
			return response.Send(iris.StatusInternalServerError, "Failed to update node", make(map[string]interface{}))
		}
	}

	return response.Success("Updated successful", response.Payload{"data": worker})
}

// Delete node
func (instance *Controller) DeleteBy(id string) mvc.Result {
	worker := &models.Node{
		Id: id,
	}

	_, err := models.Engine.Delete(worker)
	if err != nil {

	}

	return response.Success("Deleted successful", response.Payload{"data": make(map[string]interface{})})
}
