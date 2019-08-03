package log

import (
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/services"
	"github.com/go-xorm/builder"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
)

type (
	Controller struct {
		Service services.LogService
	}
)

// Query user operation log
func (instance *Controller) Get(ctx iris.Context) mvc.Response {
	var (
		total  int64
		err    error
	)

	search := ctx.Params().GetStringDefault("search", "")
	page, limit, start := utils.Pagination(ctx)
	logs := make([]models.Log, 0)

	if search == "" {
		total, err = models.Engine.Limit(limit, start).Desc("created_at").FindAndCount(&logs)

	} else {
		total, err = models.Engine.Where(builder.Like{"name", search}).Limit(limit, start).Desc("created_at").FindAndCount(&logs)
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
