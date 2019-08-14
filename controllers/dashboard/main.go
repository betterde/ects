package dashboard

import (
	"context"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/discover"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/models"
	"github.com/coreos/etcd/clientv3"
	"github.com/go-xorm/builder"
	"github.com/kataras/iris/mvc"
)

type (
	Controller struct{}
)

// 获取节点数据
func (instance *Controller) GetNodes() mvc.Response {
	nodes := make([]*models.Node, 0)
	if err := models.Engine.Where(builder.Eq{"status": "online"}).Find(&nodes); err != nil {
		return response.InternalServerError("获取节点信息失败", err)
	}
	worker := 0
	master := 0
	for _, node := range nodes {
		switch node.Mode {
		case "worker":
			worker ++
			break
		case "master":
			master ++
			break
		}
	}

	result := make(map[string]int)
	result["master"] = master
	result["worker"] = worker

	return response.Success("请求成功", response.Payload{"data": result})
}

// 获取正在调度的流水线数量
func (instance *Controller) GetPipelines() mvc.Response {
	resp, err := discover.Client.Get(context.TODO(), config.Conf.Etcd.Pipeline, clientv3.WithPrefix())
	if err != nil {
		return response.InternalServerError("获取流水线信息失败", err)
	}
	return response.Success("请求成功", response.Payload{"data": len(resp.Kvs)})
}

// 获取流水线失败次数
func (instance *Controller) GetFailtures() mvc.Response {
	if count, err := models.Engine.Where(builder.Eq{"status": 0}).Count(&models.PipelineRecords{}); err != nil {
		return response.InternalServerError("获取流水线执行记录失败", err)
	} else {
		return response.Success("请求成功", response.Payload{"data": count})
	}
}
