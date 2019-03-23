package install

import (
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/controllers/auth"
	"github.com/betterde/ects/internal/models"
	"github.com/betterde/ects/internal/services"
	"github.com/betterde/ects/internal/utils/response"
	"github.com/betterde/ects/internal/utils/system"
	"github.com/kataras/iris"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
	"log"
	"time"
)

type (
	InstallationController struct {
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
	var err error
	validate := validator.New()

	// 如果已经安装则返回错误
	if system.Info.Installed {
		if _, err = ctx.JSON(response.Send(400, "系统已安装", make(map[string]interface{}))); err != nil {
			// TODO
		}
		return
	}

	if err := ctx.ReadJSON(&config.Conf); err != nil {
		// TODO Add logger
	}

	if err := validate.Struct(config.Conf); err != nil {
		if _, err := ctx.JSON(response.ValidationError("请检查配置信息是否填写完整")); err != nil {
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

	pass, err := models.GeneratePassword(config.Conf.User.Pass)
	id := uuid.NewV4()

	user := &models.User{
		ID:        id.String(),
		Name:      config.Conf.User.Name,
		Email:     config.Conf.User.Email,
		Password:  string(pass),
		Manager:   true,
		CreatedAt: time.Now().Format("2006-1-2 15:04:05"),
		UpdatedAt: time.Now().Format("2006-1-2 15:04:05"),
	}

	if _, err := models.Engine.Insert(user); err != nil {
		if _, err = ctx.JSON(response.Send(1062, "创建用户失败", err)); err != nil {
			// TODO
		}
		return
	}

	content, err := yaml.Marshal(&config.Conf)

	if err != nil {
		log.Println(err)
	}

	// 将配置写入文件
	config.WriteConfigToFile(config.Path, content)
	system.Info.Installed = true

	token, err := services.IssueToken(user)

	if _, err := ctx.JSON(response.Success("安装成功", auth.SignInSuccess{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   time.Now().Add(time.Duration(config.Conf.Auth.TTL) * time.Second).Unix(),
	})); err != nil {
		// TODO Add logger
	}
}
