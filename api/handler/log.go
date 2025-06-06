package handler

import (
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	"github.com/go-xorm/builder"
	"github.com/gofiber/fiber/v2"
	"log"
)

func GetUserLog(ctx *fiber.Ctx) error {
	var err error

	search := ctx.Query("search", "")
	meta := utils.Pagination(ctx)
	logs := make([]models.Log, 0)
	if search == "" {
		meta.Total, err = models.Engine.Limit(meta.Limit, meta.Start).Desc("created_at").FindAndCount(&logs)
	} else {
		meta.Total, err = models.Engine.Where(builder.Eq{"user_id": search}).Limit(meta.Limit, meta.Start).Desc("created_at").FindAndCount(&logs)
	}

	if err != nil {
		log.Println(err)
	}

	return ctx.JSON(response.Success("", logs, &meta))
}

func GetPipelineLog(ctx *fiber.Ctx) error {
	var err error

	search := ctx.Query("search", "")
	meta := utils.Pagination(ctx)
	logs := make([]models.TaskRecords, 0)

	if search != "" {
		field := ctx.Query("field", "pipeline_record_id")
		meta.Total, err = models.Engine.Where(builder.Eq{field: search}).Limit(meta.Limit, meta.Start).Desc("created_at").FindAndCount(&logs)
	} else {
		meta.Total, err = models.Engine.Limit(meta.Limit, meta.Start).Desc("created_at").FindAndCount(&logs)
	}

	if err != nil {
		return ctx.JSON(response.InternalServerError("获取任务日志失败", err))
	}

	return ctx.JSON(response.Success("Success", logs, &meta))
}

func GetTaskLog(ctx *fiber.Ctx) error {
	var err error

	search := ctx.Query("search", "")
	meta := utils.Pagination(ctx)
	logs := make([]models.PipelineRecords, 0)

	if search != "" {
		meta.Total, err = models.Engine.Where(builder.Like{"pipeline_id", search}).Limit(meta.Limit, meta.Start).Desc("created_at").FindAndCount(&logs)
	} else {
		meta.Total, err = models.Engine.Limit(meta.Limit, meta.Start).Desc("created_at").FindAndCount(&logs)
	}

	if err != nil {
		return ctx.JSON(response.InternalServerError("获取流水线日志失败", err))
	}

	for index, _ := range logs {
		logs[index].Steps = make([]*models.TaskRecords, 0)
	}

	return ctx.JSON(response.Success("Success", logs, &meta))
}
