package install

import (
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/utils/response"
	"github.com/betterde/ects/internal/utils/system"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
)

type (
	InstallationController struct{
	}
)

func (instance *InstallationController) Get(ctx iris.Context) {
	_, err := ctx.JSON(response.Success("请求成功", system.Info))

	if err != nil {
		// TODO
	}
}

// 根据用户填写配置数据生成配置文件
func (instance *InstallationController) Post(ctx iris.Context) {
	var conf config.Config

	validate := validator.New()

	if err := ctx.ReadJSON(&conf); err != nil {
		// TODO Add logger
	}

	if err := validate.Struct(conf); err != nil {
		if _, err := ctx.JSON(response.ValidationError("用户名和密码不能为空")); err != nil {
			// TODO Add logger
		}
		return
	}

	content, err := yaml.Marshal(&conf)

	if err != nil {
		// TODO
	}

	config.WriteConfigToFile(config.Path, content)

	_, err = ctx.JSON(response.Success("请求成功", system.Info))

	if err != nil {
		// TODO
	}
}
