package install

import (
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/models"
	"github.com/betterde/ects/internal/utils/response"
	"github.com/betterde/ects/internal/utils/system"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
)

type (
	InstallationController struct {
	}
	Admin struct {
		Name    string `json:"name" validate:"required"`
		Email   string `json:"email" validate:"required"`
		Pass    string `json:"pass" validate:"required"`
		Confirm string `json:"confirm" validate:"required"`
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
	var (
		admim Admin
		err error
	)
	validate := validator.New()

	if err := ctx.ReadJSON(&config.Conf); err != nil {
		// TODO Add logger
	}

	if err := ctx.ReadJSON(&admim); err != nil {
		// TODO Add logger
	}

	if err := validate.Struct(config.Conf); err != nil {
		if _, err := ctx.JSON(response.ValidationError("请检查配置信息是否填写完整")); err != nil {
			// TODO Add logger
		}
		return
	}

	if err := validate.Struct(config.Conf); err != nil {
		if _, err := ctx.JSON(response.ValidationError("请检查管理员用户信息是否完整")); err != nil {
			// TODO Add logger
		}
		return
	}

	if models.Engine == nil {
		// 初始化数据库
		models.Engine, err = models.Connection()

		if err != nil {
			if _, err = ctx.JSON(response.Success("数据库连接错误", err)); err != nil {
				// TODO
			}
			return
		}

		// 迁移数据表时出现错误则返回异常
		if err := models.Migrate(); err != nil {
			if _, err = ctx.JSON(response.Send(1045, "数据库连接错误", err)); err != nil {
				// TODO
			}
			return
		}
	}

	content, err := yaml.Marshal(&config.Conf)

	if err != nil {
		// TODO
	}

	// 将配置写入文件
	config.WriteConfigToFile(config.Path, content)
	system.Info.Installed = true
	if _, err = ctx.JSON(response.Success("安装成功", system.Info)); err != nil {
		// TODO
	}
}
