package initialize

import (
	"context"
	"encoding/json"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/controllers/auth"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/system"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/services"
	"github.com/coreos/etcd/clientv3"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
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

// Get system info
func (instance *Controller) Get(ctx iris.Context) mvc.Result {
	return response.Success("请求成功", response.Payload{"data": system.Info})
}

// Initialize system config
func (instance *Controller) Post(ctx iris.Context) mvc.Result {
	var (
		params = &PostRequest{}
		err    error
	)
	validate := validator.New()

	if err := ctx.ReadJSON(params); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(params); err != nil {
		return response.ValidationError("Please check whether the configuration information is complete")
	}

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   params.Etcd.EndPoints,
		DialTimeout: 10 * time.Second,
	})

	if err != nil {
		return response.InternalServerError("Failed to connect to etcd endpoints", err)
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
		return response.InternalServerError("Failed to Marshal JSON", err)
	}

	if _, err = client.Put(context.TODO(), params.Etcd.Config, string(buf)); err != nil {
		return response.InternalServerError("Failed to put config to etcd", err)
	}

	if err := utils.CreateDatabase(); err != nil {
		return response.InternalServerError("Failed to create database", err)
	}

	if models.Engine == nil {
		// Create database engine
		models.Engine, err = models.Connection()
		if err != nil {
			return response.InternalServerError("Failed to connect to database", err)
		}
	}

	if err := models.Migrate(); err != nil {
		return response.InternalServerError("Failed to migrate the table", err)
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
		return response.InternalServerError("Failed to create system manager", err)
	}

	token, err := services.IssueToken(user)

	return response.Success("Congratulations! successful installation", response.Payload{"data": auth.SignInSuccess{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   time.Now().Add(time.Duration(config.Conf.Auth.TTL) * time.Second).Unix(),
	}})
}

// Get JWT Secret
func (instance *Controller) GetSecret(ctx iris.Context) mvc.Result {
	return response.Success("Success", response.Payload{"data": utils.Random(64)})
}

// Validate database exist
func (instance *Controller) GetDatabase(ctx iris.Context) mvc.Result {
	config.Conf.Database.User = ctx.URLParam("user")
	config.Conf.Database.Pass = ctx.URLParam("pass")
	config.Conf.Database.Host = ctx.URLParam("host")
	config.Conf.Database.Port = ctx.Params().GetIntDefault("port", 3306)
	config.Conf.Database.Char = ctx.URLParam("char")
	config.Conf.Database.Name = ctx.URLParam("name")
	return response.Success("Success", response.Payload{"data": map[string]bool{"exist": utils.IsDatabaseExist()}})
}
