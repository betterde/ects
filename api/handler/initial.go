package handler

import (
	"context"
	"encoding/json"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/service"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type (
	PostRequest struct {
		Auth     config.Auth     `json:"auth"`
		User     config.User     `json:"user"`
		Etcd     config.Etcd     `json:"etcd"`
		Database config.Database `json:"database"`
	}
)

func GetSystemInfo(ctx *fiber.Ctx) error {
	return ctx.JSON(response.Success("Success", service.Runtime))
}

func GenSystemSecret(ctx *fiber.Ctx) error {
	return ctx.JSON(response.Success("Success", utils.Random(64)))
}

func SetSystemDatabase(ctx *fiber.Ctx) error {
	err := ctx.BodyParser(&config.Conf.Database)
	if err != nil {
		return err
	}
	return ctx.JSON(response.Success("Success", map[string]bool{"exist": utils.IsDatabaseExist()}))
}

func InitializationSystem(ctx *fiber.Ctx) error {
	var (
		params = PostRequest{}
	)
	validate := validator.New()

	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	if err := validate.Struct(params); err != nil {
		return err
	}

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   params.Etcd.EndPoints,
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		return err
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}()

	config.Conf.Auth = params.Auth
	config.Conf.Etcd = params.Etcd
	config.Conf.Database = params.Database

	buf, err := json.Marshal(config.Conf)
	if err != nil {
		return err
	}

	if _, err = client.Put(context.TODO(), params.Etcd.Config, string(buf)); err != nil {
		return err
	}

	if err = utils.CreateDatabase(); err != nil {
		return err
	}

	if models.Engine == nil {
		// Create database engine
		models.Engine, err = models.Connection()
		if err != nil {
			return err
		}
	}

	if err = models.Migrate(); err != nil {
		return err
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

	if _, err = models.Engine.Insert(user); err != nil {
		return err
	}

	token, err := service.IssueToken(user)

	return ctx.JSON(response.Success("Success", SignInSuccess{
		TokenType:   "Bearer",
		ExpiresIn:   uint64(time.Now().Add(time.Duration(config.Conf.Auth.TTL) * time.Second).Unix()),
		AccessToken: token,
	}))
}
