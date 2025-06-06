package handler

import (
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/service"
	"github.com/betterde/ects/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

type (
	SignIn struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	SignUp struct {
		Name     string `json:"name" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Confirm  string `json:"confirm" validate:"required"`
	}

	SignInSuccess struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   uint64 `json:"expires_in"`
	}
)

func SignInHandler(ctx *fiber.Ctx) error {
	var params SignIn
	validate := validator.New()
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	if err := validate.Struct(params); err != nil {
		return err
	}

	userService := service.NewUserService()

	token, err := userService.Attempt(params.Username, params.Password)
	if err != nil {
		return err
	}

	return ctx.JSON(SignInSuccess{
		TokenType:   "Bearer",
		ExpiresIn:   config.Conf.Auth.TTL,
		AccessToken: token,
	})
}

func SignUpHandler(ctx *fiber.Ctx) error {
	var params SignUp
	validate := validator.New()
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	if err := validate.Struct(params); err != nil {
		return err
	}

	user := &models.User{
		Name: params.Name,
	}

	hash, err := models.GeneratePassword(params.Password)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	if err := user.Save(); err != nil {
		log.Println(err)
	}
	return ctx.JSON(response.Success("success", user, nil))
}
