package handler

import (
	"context"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/discover"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/models"
	"github.com/go-xorm/builder"
	"github.com/gofiber/fiber/v2"
	"go.etcd.io/etcd/client/v3"
)

func GetNodes(ctx *fiber.Ctx) error {
	nodes := make([]*models.Node, 0)
	if err := models.Engine.Where(builder.Eq{"status": "online"}).Find(&nodes); err != nil {
		return err
	}

	res := struct {
		Master uint `json:"master"`
		Worker uint `json:"worker"`
	}{
		Master: 0,
		Worker: 0,
	}

	for _, node := range nodes {
		switch node.Mode {
		case "worker":
			res.Worker++
			break
		case "master":
			res.Master++
			break
		}
	}

	return ctx.JSON(response.Success("success", res, nil))
}

func GetPipelines(ctx *fiber.Ctx) error {
	resp, err := discover.Client.Get(context.TODO(), config.Conf.Etcd.Pipeline, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	return ctx.JSON(response.Success("success", len(resp.Kvs), nil))
}

func GetFailtures(ctx *fiber.Ctx) error {
	if count, err := models.Engine.Where(builder.Eq{"status": 0}).Count(&models.PipelineRecords{}); err != nil {
		return err
	} else {
		return ctx.JSON(response.Success("success", count, nil))
	}
}
