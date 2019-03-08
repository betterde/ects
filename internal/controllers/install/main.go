package install

import (
	"github.com/betterde/ects/internal/utils/response"
	"github.com/betterde/ects/internal/utils/system"
	"github.com/kataras/iris"
)

type (
	InstallationController struct{}
	SystemStatus struct {
		Installed bool   `json:"installed"`
		Version   string `json:"version"`
	}
)

func (instance *InstallationController) Get(ctx iris.Context) {
	_, err := ctx.JSON(response.Success("请求成功", &SystemStatus{
		Installed: system.Installed,
		Version: system.Version,
	}))

	if err != nil {
		// TODO
	}
}

func (instance *InstallationController) Post(ctx iris.Context) {

}
