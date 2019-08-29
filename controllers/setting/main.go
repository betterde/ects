package setting

import (
	"context"
	"encoding/json"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/discover"
	"github.com/betterde/ects/internal/response"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type (
	Controller struct{}
)

// 获取通知配置信息
func (instance *Controller) GetNotification(ctx iris.Context) mvc.Response {
	resp, err := discover.Client.Get(context.TODO(), config.Conf.Etcd.Config)
	if err != nil {
		return response.InternalServerError("获取配置信息失败", err)
	}

	conf := &config.Config{}
	if err := json.Unmarshal(resp.Kvs[0].Value, &conf); err != nil {
		return response.InternalServerError("反序列化失败", err)
	}

	return response.Success("请求成功", response.Payload{"data": conf.Notification})
}

// 测试发送邮件功能
func (instance *Controller) PostMail(ctx iris.Context) mvc.Response {
	return response.Success("请求成功", response.Payload{"data": make([]interface{}, 0)})
}

// 更新服务配置
func (instance *Controller) PutNotification(ctx iris.Context) mvc.Response {
	params := config.Notification{}
	if err := ctx.ReadJSON(&params); err != nil {
		return response.Send(400, "解析参数失败", err)
	}

	config.Conf.Notification = params

	bytes, err := json.Marshal(config.Conf)
	if err != nil {
		return response.InternalServerError("序列化失败", err)
	}

	if _, err := discover.Client.Put(context.TODO(), config.Conf.Etcd.Config, string(bytes)); err != nil {
		return response.InternalServerError("更新配置失败", err)
	}

	return response.Success("请求成功", response.Payload{"data": params})
}
