package log

import (
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	"github.com/go-xorm/builder"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
)

type (
	Controller struct{}
)

// 获取日志
func (instance *Controller) Get(ctx iris.Context) mvc.Response {
	var (
		total int64
		err   error
	)

	scene := ctx.URLParamDefault("scene", "pipeline")
	search := ctx.URLParamDefault("search", "")
	page, limit, start := utils.Pagination(ctx)

	switch scene {
	case "pipeline":
		logs := make([]models.PipelineRecords, 0)
		if search != "" {
			total, err = models.Engine.Where(builder.Like{"pipeline_id", search}).Limit(limit, start).Desc("created_at").FindAndCount(&logs)
		} else {
			total, err = models.Engine.Limit(limit, start).Desc("created_at").FindAndCount(&logs)
		}

		if err != nil {
			return response.InternalServerError("获取流水线日志失败", err)
		}

		for index, _ := range logs {
			logs[index].Steps = make([]*models.TaskRecords, 0)
		}

		return response.Success("请求成功", response.Payload{
			"data": logs,
			"meta": &response.Meta{
				Limit: limit,
				Page:  page,
				Total: int(total),
			},
		})
	case "task":
		logs := make([]models.TaskRecords, 0)
		if search != "" {
			field := ctx.URLParamDefault("field", "pipeline_record_id")
			total, err = models.Engine.Where(builder.Eq{field: search}).Limit(limit, start).Desc("created_at").FindAndCount(&logs)
		} else {
			total, err = models.Engine.Limit(limit, start).Desc("created_at").FindAndCount(&logs)
		}

		if err != nil {
			return response.InternalServerError("获取任务日志失败", err)
		}

		return response.Success("请求成功", response.Payload{
			"data": logs,
			"meta": &response.Meta{
				Limit: limit,
				Page:  page,
				Total: int(total),
			},
		})
	case "user":
		logs := make([]models.Log, 0)
		if search == "" {
			total, err = models.Engine.Limit(limit, start).Desc("created_at").FindAndCount(&logs)
		} else {
			total, err = models.Engine.Where(builder.Eq{"user_id": search}).Limit(limit, start).Desc("created_at").FindAndCount(&logs)
		}

		if err != nil {
			log.Println(err)
		}
		return response.Success("Success", response.Payload{
			"data": logs,
			"meta": &response.Meta{
				Limit: limit,
				Page:  page,
				Total: int(total),
			}})
	}

	return response.Send(400, "请选择日志类型", make([]interface{}, 0))
}
