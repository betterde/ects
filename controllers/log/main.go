package log

import (
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/services"
	"github.com/go-xorm/builder"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"strconv"
)

type (
	Controller struct {
		Service services.LogService
	}
)

// Query user operation log
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

	logs := make([]models.Log, 0)

	if search == "" {
		total, err = models.Engine.Count(&models.Log{})
		err = models.Engine.Limit(limit, start).Find(&logs)
	} else {
		total, err = models.Engine.Where(builder.Like{"name", search}).Count(&models.Log{})
		err = models.Engine.Where(builder.Like{"name", search}).Limit(limit, start).Find(&logs)
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
