package dashboard

import (
	"github.com/betterde/ects/internal/response"
	"github.com/kataras/iris/mvc"
)

type (
	Controller struct{}
)

// 获取节点数据
func (instance *Controller) GetNodes() mvc.Response {
	return response.Success("", response.Payload{"data": make([]interface{}, 0)})
}
