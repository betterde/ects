package pipeline

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/discover"
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
)

type (
	Controller struct {
		Service services.PipelineService
	}
	BindNodeRequest struct {
		PipelineId string   `json:"pipeline_id" validate:"required,uuid4"`
		NodesId    []string `json:"nodes_id" validate:"required"`
	}
)

var (
	validate = validator.New()
)

// Get pipelines list
func (instance *Controller) Get(ctx iris.Context) mvc.Result {
	var (
		page   = 1
		limit  = 10
		start  int
		search = ctx.URLParam("name")
		total  int64
		err    error
	)

	pipelines := make([]models.Pipeline, 0)

	start = (page - 1) * limit

	if search != "" {
		total, err = models.Engine.Where(builder.Like{"name", search}).Limit(limit, start).FindAndCount(&pipelines)
	} else {
		total, err = models.Engine.Limit(limit, start).FindAndCount(&pipelines)
	}

	if err != nil {
		return response.InternalServerError("Failed to query pipelines list", err)
	}

	return response.Success("Successful", response.Payload{
		"data": pipelines,
		"meta": &response.Meta{
			Limit: limit,
			Page:  page,
			Total: int(total),
		},
	})
}

// Create a pipeline
func (instance *Controller) Post(ctx iris.Context) mvc.Result {
	pipeline := models.Pipeline{}

	if err := ctx.ReadJSON(&pipeline); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(pipeline); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("pipeline", validationErrors))
	}

	pipeline.Id = uuid.NewV4().String()
	err := pipeline.Store()
	if err != nil {
		return response.InternalServerError("Failed to create pipeline", err)
	}

	key := fmt.Sprintf("%s/%s", config.Conf.Etcd.Pipeline, pipeline.Id)

	bytes, err := json.Marshal(&pipeline)
	if err != nil {
		log.Println(err)
	}

	if _, err := discover.Client.Put(context.TODO(), key, string(bytes)); err != nil {
		log.Println(err)
	}

	if err := models.CreateLog(pipeline, ctx, "CREATE PIPELINE"); err != nil {
		return response.InternalServerError("Failed to create log", err)
	}

	return response.Success("Created successfully", response.Payload{"data": pipeline})
}

// Update pipeline attributes by id
func (instance *Controller) PutBy(id string, ctx iris.Context) mvc.Result {
	pipeline := models.Pipeline{}

	if err := ctx.ReadJSON(&pipeline); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(pipeline); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("pipeline", validationErrors))
	}
	pipeline.Id = id
	err := pipeline.Update()
	if err != nil {
		return response.InternalServerError("Failed to update pipeline", err)
	}

	key := fmt.Sprintf("%s/%s", config.Conf.Etcd.Pipeline, pipeline.Id)

	bytes, err := json.Marshal(&pipeline)
	if err != nil {
		log.Println(err)
	}

	if _, err := discover.Client.Put(context.TODO(), key, string(bytes)); err != nil {
		return response.InternalServerError("Failed to delete pipeline", err)
	}

	return response.Success("Updated successfully", response.Payload{"data": pipeline})
}

// Delete pipeline by id
func (instance *Controller) DeleteBy(id string, ctx iris.Context) mvc.Result {
	pipeline := models.Pipeline{
		Id: id,
	}

	key := fmt.Sprintf("%s/%s", config.Conf.Etcd.Pipeline, pipeline.Id)

	if _, err := discover.Client.Delete(context.TODO(), key); err != nil {
		return response.InternalServerError("Failed to delete pipeline", err)
	}

	if err := pipeline.Destroy(); err != nil {
		return response.InternalServerError("Failed to delete pipeline", err)
	}
	return response.Success("Deleted successfully", response.Payload{"data": make(map[string]interface{})})
}

// Get pipeline binding nodes
func (instance *Controller) GetNodes(ctx iris.Context) mvc.Result {
	id := ctx.URLParam("pipeline_id")

	if id == "" {
		return response.ValidationError("pipeline id is required")
	}

	relations := make([]models.PipelineNodePivot, 0)

	if err := models.Engine.Where(builder.Eq{"pipeline_id": id}).Find(&relations); err != nil {
		return response.InternalServerError("Failed to query relations", err)
	}

	ids := make([]string, 0)

	for _, pivot := range relations{
		ids = append(ids, pivot.NodeId)
	}

	nodes := make([]models.Node, 0)

	if err := models.Engine.Where(builder.Eq{"id": ids}).Find(&nodes); err != nil {
		return response.InternalServerError("Failed to query relations", err)
	}

	return response.Success("Successful", response.Payload{"data": nodes})
}

// Bind pipeline to node
func (instance *Controller) PostNodes(ctx iris.Context) mvc.Result {
	params := BindNodeRequest{}

	if err := ctx.ReadJSON(&params); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(params); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("pipeline", validationErrors))
	}

	if _, err := models.Engine.Where(builder.Eq{"pipeline_id": params.PipelineId}).Delete(&models.PipelineNodePivot{}); err != nil {
		return response.InternalServerError("Failed to delete pipeline and node relations", err)
	}

	relations := make([]*models.PipelineNodePivot, 0)

	for _, id := range params.NodesId {
		relations = append(relations, &models.PipelineNodePivot{
			PipelineId: params.PipelineId,
			NodeId: id,
		})
	}

	_, err := models.Engine.Insert(relations)
	if err != nil {
		return response.InternalServerError("Failed to bind pipeline to node", err)
	}

	return response.Success("Bind successfully", response.Payload{"data": relations})
}

// Get pipeline tasks
func (instance * Controller) GetTasks(ctx iris.Context) mvc.Result {
	id := ctx.URLParam("pipeline_id")

	if id == "" {
		return response.ValidationError("pipeline id is required")
	}

	relations := make([]models.PipelineTaskPivot, 0)

	if err := models.Engine.Where(builder.Eq{"pipeline_id": id}).Find(&relations); err != nil {
		return response.InternalServerError("Failed to query relations", err)
	}

	ids := make([]string, 0)

	for _, pivot := range relations{
		ids = append(ids, pivot.TaskId)
	}

	tasks := make([]models.Task, 0)

	if err := models.Engine.Where(builder.Eq{"id": ids}).Find(&tasks); err != nil {
		return response.InternalServerError("Failed to query relations", err)
	}

	return response.Success("Successful", response.Payload{"data": tasks})
}

// Bind the task to pipeline
func (instance * Controller) PostTask(ctx iris.Context) mvc.Result {
	pivot := models.PipelineTaskPivot{}

	if err := ctx.ReadJSON(&pivot); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(pivot); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("pipeline", validationErrors))
	}

	err := pivot.Store()
	if err != nil {
		return response.InternalServerError("Failed to bind pipeline to node", err)
	}

	return response.Success("Bind successfully", response.Payload{"data": pivot})
}

// Get pipeline detail by id
func (instance *Controller) GetBy(id string) mvc.Result {
	pipeline := models.Pipeline{
		Id: id,
	}

	_, err := models.Engine.Get(&pipeline)

	if err != nil {
		return response.InternalServerError("Query pipeline on error", err)
	}

	_, err = pipeline.Build()

	if err != nil {
		return response.InternalServerError("Failed to build pipeline to string", err)
	}

	return response.Success("", response.Payload{"data": pipeline})
}
