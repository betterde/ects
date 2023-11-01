package initialize

import (
	"context"
	"encoding/json"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/controllers/auth"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/service"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/services"
	"github.com/coreos/etcd/clientv3"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"time"
)

type (
	Controller struct {
	}
	PostRequest struct {
		Etcd     config.Etcd     `json:"etcd"`
		Database config.Database `json:"database"`
		User     config.User     `json:"user"`
		Auth     config.Auth     `json:"auth"`
	}
)

// 获取系统信息
func (instance *Controller) Get(ctx iris.Context) mvc.Response {
	return response.Success("请求成功", response.Payload{"data": service.Runtime})
}

// 创建服务配置
func (instance *Controller) Post(ctx iris.Context) mvc.Response {
	var (
		params = &PostRequest{}
		err    error
	)
	validate := validator.New()

	if err := ctx.ReadJSON(params); err != nil {
		return response.InternalServerError("参数解析失败", err)
	}

	if err := validate.Struct(params); err != nil {
		return response.ValidationError("配置参数有误")
	}

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   params.Etcd.EndPoints,
		DialTimeout: 10 * time.Second,
	})

	if err != nil {
		return response.InternalServerError("无法连接ETCD", err)
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}()

	config.Conf.Etcd = params.Etcd
	config.Conf.Database = params.Database
	config.Conf.Auth = params.Auth

	buf, err := json.Marshal(config.Conf)
	if err != nil {
		return response.InternalServerError("序列化配置信息失败", err)
	}

	if _, err = client.Put(context.TODO(), params.Etcd.Config, string(buf)); err != nil {
		return response.InternalServerError("保存配置信息失败", err)
	}

	if err := utils.CreateDatabase(); err != nil {
		return response.InternalServerError("创建数据库失败", err)
	}

	if models.Engine == nil {
		// Create database engine
		models.Engine, err = models.Connection()
		if err != nil {
			return response.InternalServerError("链接数据库失败", err)
		}
	}

	if err := models.Migrate(); err != nil {
		return response.InternalServerError("迁移数据表失败", err)
	}

	pass, err := models.GeneratePassword(params.User.Pass)

	user := &models.User{
		Id:        uuid.NewV4().String(),
		Name:      params.User.Name,
		Email:     params.User.Email,
		Password:  string(pass),
		Manager:   true,
		CreatedAt: utils.Time(time.Now()),
		UpdatedAt: utils.Time(time.Now()),
	}

	if _, err := models.Engine.Insert(user); err != nil {
		return response.InternalServerError("创建系统管理员失败", err)
	}

	token, err := services.IssueToken(user)

	return response.Success("初始化成功", response.Payload{"data": auth.SignInSuccess{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   time.Now().Add(time.Duration(config.Conf.Auth.TTL) * time.Second).Unix(),
	}})
}

// 生成 JWT Secret
func (instance *Controller) GetSecret(ctx iris.Context) mvc.Response {
	return response.Success("Success", response.Payload{"data": utils.Random(64)})
}

// 验证数据库是否存在
func (instance *Controller) GetDatabase(ctx iris.Context) mvc.Response {
	config.Conf.Database.User = ctx.URLParam("user")
	config.Conf.Database.Pass = ctx.URLParam("pass")
	config.Conf.Database.Host = ctx.URLParam("host")
	config.Conf.Database.Port = ctx.Params().GetIntDefault("port", 3306)
	config.Conf.Database.Char = ctx.URLParam("char")
	config.Conf.Database.Name = ctx.URLParam("name")
	return response.Success("Success", response.Payload{"data": map[string]bool{"exist": utils.IsDatabaseExist()}})
}
