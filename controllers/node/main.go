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

// 获取节点列表
func (instance *Controller) Get(ctx iris.Context) mvc.Response {
	var (
		total int64
		err   error
	)
	search := ctx.Params().GetStringDefault("search", "")
	page, limit, start := utils.Pagination(ctx)
	nodes := make([]models.Node, 0)

	if search == "" {
		total, err = models.Engine.Limit(limit, start).FindAndCount(&nodes)
	} else {
		total, err = models.Engine.Where(builder.Like{"name", search}).Limit(limit, start).FindAndCount(&models.Node{})
	}

	if err != nil {
		log.Println(err)
	}

	nids := make([]string, 0)

	for _, node := range nodes {
		nids = append(nids, node.Id)
	}

	relations := make([]models.PipelineNodePivot, 0)
	if err := models.Engine.Where(builder.In("node_id", nids)).Find(&relations); err != nil {
		return response.InternalServerError("查询绑定关系失败", err)
	}

	pids := make([]string, 0)

	for _, relation := range relations {
		pids = append(pids, relation.PipelineId)
	}

	pipelines := make(map[string]models.Pipeline)

	if err := models.Engine.Where(builder.In("id", pids)).Find(&pipelines); err != nil {
		return response.InternalServerError("获取关联的流水线信息失败", err)
	}

	for rindex, relation := range relations {
		pipeline := pipelines[relation.PipelineId]
		relations[rindex].Pipeline = &pipeline
	}

	for nindex, node := range nodes {
		nodes[nindex].Pipelines = make([]*models.PipelineNodePivot, 0)
		for rindex, relation := range relations {
			if node.Id == relation.NodeId {
				nodes[nindex].Pipelines = append(nodes[nindex].Pipelines, &relations[rindex])
			}
		}
	}

	return response.Success("请求成功", response.Payload{
		"data": nodes,
		"meta": &response.Meta{
			Limit: limit,
			Page:  page,
			Total: int(total),
		}})
}

// 创建节点
func (instance *Controller) Post(ctx iris.Context) mvc.Response {
	var params CreateRequest
	validate := validator.New()
	if err := ctx.ReadJSON(&params); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(params); err != nil {
		return response.ValidationError("Please fill in the name")
	}

	worker := models.Node{
		Id:          uuid.NewV4().String(),
		Name:        params.Name,
		Description: params.Remark,
		Status:      models.ONLINE,
		Host:        "",
		CreatedAt:   utils.Time(time.Now()),
		UpdatedAt:   utils.Time(time.Now()),
	}
	if err := worker.Store(); err != nil {
		return response.InternalServerError("Failed to create node", err)
	}

	return response.Success("Created successful", response.Payload{"data": worker})
}

// 修改节点信息
func (instance *Controller) PutBy(id string, ctx iris.Context) mvc.Response {
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
		if _, err := models.Engine.Id(id).Update(&worker); err != nil {
			return response.Send(iris.StatusInternalServerError, "Failed to update node", err)
		}
	}

	return response.Success("更新成功", response.Payload{"data": worker})
}

// 删除节点
func (instance *Controller) DeleteBy(id string) mvc.Response {
	worker := &models.Node{
		Id: id,
	}

	_, err := models.Engine.Delete(worker)
	if err != nil {

	}

	return response.Success("删除成功", response.Payload{"data": make(map[string]interface{})})
}

// 获取节点关联的流水线
func (instance *Controller) GetPipelines(ctx iris.Context) mvc.Response {
	id := ctx.URLParamDefault("node_id", "")
	if id == "" {
		return response.ValidationError("请选择节点")
	}

	relations := make([]models.PipelineNodePivot, 0)

	if err := models.Engine.Where(builder.Eq{"node_id": id}).Find(&relations); err != nil {
		return response.InternalServerError("查询节点和流水线关系失败", err)
	}

	ids := make([]string, 0)

	// 获取流水线ID数组
	for _, relation := range relations {
		ids = append(ids, relation.PipelineId)
	}

	pipelines := make(map[string]models.Pipeline)

	if err := models.Engine.Where(builder.Eq{"id": ids}).Find(&pipelines); err != nil {
		return response.InternalServerError("Failed to query relations", err)
	}

	// 将流水线关联到节点关系
	for index, relation := range relations {
		task := pipelines[relation.PipelineId]
		relations[index].Pipeline = &task
	}

	return response.Success("获取关联关系成功", response.Payload{"data": relations})
}

// 关联流水线
func (instance *Controller) PostPipeline(ctx iris.Context) mvc.Response {
	relation := models.PipelineNodePivot{
		Id: uuid.NewV4().String(),
	}

	if err := ctx.ReadJSON(&relation); err != nil {
		return response.InternalServerError("参数解析失败", err)
	}

	if err := relation.Store(); err != nil {
		return response.InternalServerError("关联失败", err)
	}

	relation.Pipeline = &models.Pipeline{}

	if _, err := models.Engine.Id(relation.PipelineId).Get(relation.Pipeline); err != nil {
		return response.InternalServerError("查询流水线信息失败", err)
	}

	// 查询流水线关联的任务数量，如果没有任务则无法关联
	if count, err := models.Engine.Where(builder.Eq{"pipeline_id": relation.PipelineId}).Count(&models.PipelineTaskPivot{}); err != nil {
		return response.InternalServerError("获取流水线关联的任务数量失败", err)
	} else if count < 1 {
		return response.Send(400, "当前流水线没有关联任何任务", make([]interface{}, 0))
	}

	return response.Success("关联成功", response.Payload{"data": relation})
}

// 解绑流水线
func (instance *Controller) DeletePipelineBy(id string, ctx iris.Context) mvc.Response {
	if id == "" {
		return response.ValidationError("请选择要解绑的关联ID")
	}

	relation := models.PipelineNodePivot{
		Id: id,
	}

	if _, err := models.Engine.Delete(&relation); err != nil {
		return response.InternalServerError("解绑流水线失败", err)
	}

	return response.Success("解绑成功", response.Payload{"data": make([]interface{}, 0)})
}
