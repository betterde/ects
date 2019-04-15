package initialize

import (
	"context"
	"encoding/json"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/controllers/auth"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/system"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/seeds"
	"github.com/betterde/ects/services"
	"github.com/coreos/etcd/clientv3"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"os"
	"time"
)

type (
	Controller struct {
	}
	PostRequest struct {
		Etcd     config.Etcd     `json:"etcd"`
		Service  config.Service  `json:"service"`
		Database config.Database `json:"database"`
		User     config.User     `json:"user"`
		Auth     config.Auth     `json:"auth"`
	}
)

// 获取系统信息
func (instance *Controller) Get(ctx iris.Context) mvc.Result {
	return response.Success("请求成功", response.Payload{"data": system.Info})
}

// 根据用户填写配置数据生成配置文件
func (instance *Controller) Post(ctx iris.Context) mvc.Result {
	var (
		params = &PostRequest{}
		err    error
	)
	validate := validator.New()

	if err := ctx.ReadJSON(params); err != nil {
		return response.InternalServerError("解析JSON失败", err)
	}

	if err := validate.Struct(params); err != nil {
		return response.ValidationError("请检查配置信息是否填写完整")
	}

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   params.Etcd.EndPoints,
		DialTimeout: 10 * time.Second,
	})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}()

	requestctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	config.Conf.Etcd = params.Etcd
	config.Conf.Service = params.Service
	config.Conf.Database = params.Database
	config.Conf.Auth = params.Auth

	bites, err := json.Marshal(config.Conf)
	if err != nil {
		log.Println(err)
	}

	_, err = client.Put(requestctx, params.Etcd.Config, string(bites))
	cancel()
	if err != nil {
		log.Println(err)
	}

	// 如果已经安装则返回错误
	if system.Info.Installed {
		return response.Send(400, "系统已安装", make(map[string]interface{}))
	}

	if models.Engine == nil {
		config.Conf.Database = params.Database
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

	pass, err := models.GeneratePassword(params.User.Pass)

	user := &models.User{
		Id:        uuid.NewV4().String(),
		Name:      params.User.Name,
		Email:     params.User.Email,
		Password:  string(pass),
		Manager:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := models.Engine.Insert(user); err != nil {
		return response.Send(1062, "创建用户失败", err)
	}

	token, err := services.IssueToken(user)

	return response.Success("安装成功", response.Payload{"data": auth.SignInSuccess{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   time.Now().Add(time.Duration(config.Conf.Auth.TTL) * time.Second).Unix(),
	}})
}
