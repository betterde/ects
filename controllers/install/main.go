package install

import (
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/controllers/auth"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/system"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/seeds"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
	"log"
	"time"
)

type (
	Controller struct {
	}
)

func (instance *Controller) Get(ctx iris.Context) {
	_, err := ctx.JSON(response.Success("请求成功", response.Payload{"data": system.Info}))

	if err != nil {
		// TODO
	}
}

// 根据用户填写配置数据生成配置文件
func (instance *Controller) Post(ctx iris.Context) mvc.Result {
	var err error
	validate := validator.New()

	// 如果已经安装则返回错误
	if system.Info.Installed {
		return response.Send(400, "系统已安装", make(map[string]interface{}))
	}

	if err := ctx.ReadJSON(&config.Conf); err != nil {
		// TODO Add logger
	}

	if err := validate.Struct(config.Conf); err != nil {
		return response.ValidationError("请检查配置信息是否填写完整")
	}

	if models.Engine == nil {
		// 初始化数据库
		models.Engine, err = models.Connection()

		if err != nil {
			return response.Success("数据库连接错误", response.Payload{"data": err})
		}

		// 迁移数据表时出现错误则返回异常
		if err := models.Migrate(); err != nil {
			return response.Send(1045, "数据库连接错误", err)
		}

		// 填充系统初始数据
		seedService := seeds.Seeder{}
		if err := seedService.Run(); err != nil {
			return response.Send(1045, "填出系统数据失败", err)
		}
	}

	pass, err := models.GeneratePassword(config.Conf.User.Pass)
	id := uuid.NewV4()

	user := &models.User{
		Id:        id.String(),
		Name:      config.Conf.User.Name,
		Email:     config.Conf.User.Email,
		Password:  string(pass),
		Manager:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := models.Engine.Insert(user); err != nil {
		return response.Send(1062, "创建用户失败", err)
	}

	content, err := yaml.Marshal(&config.Conf)

	if err != nil {
		log.Println(err)
	}

	// 将配置写入文件
	config.WriteConfigToFile(config.Path, content)
	system.Info.Installed = true

	token, err := services.IssueToken(user)

	return response.Success("安装成功", response.Payload{"data": auth.SignInSuccess{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   time.Now().Add(time.Duration(config.Conf.Auth.TTL) * time.Second).Unix(),
	}})
}
