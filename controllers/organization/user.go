package organization

import (
	"github.com/betterde/ects/internal/message"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"time"
)

type (
	UserController struct {
		Service services.UserService
	}

	CreateRequest struct {
		Name    string `json:"name" validate:"required"`
		Email   string `json:"email" validate:"required,email"`
		Pass    string `json:"pass" validate:"required"`
		Confirm string `json:"confirm" validate:"eqfield=Pass"`
		TeamId  string `json:"team_id" validate:"omitempty,uuid4"`
		Manager bool   `json:"manager" validate:"required"`
	}
	UpdateRequest struct {
		Name    string `json:"name" validate:"required"`
		Email   string `json:"email" validate:"required,email"`
		TeamId  string `json:"team_id" validate:"omitempty,uuid4"`
		Manager bool   `json:"manager" validate:"required"`
	}
)

// 获取用户列表
func (instance *UserController) Get(ctx iris.Context) mvc.Result {
	params := ctx.URLParams()
	users, meta := instance.Service.Users(params)
	return response.Success("请求成功", response.Payload{"data": users, "meta": meta})
}

// 创建用户
func (instance *UserController) Post(ctx iris.Context) mvc.Result {
	var (
		params CreateRequest
	)

	validate := validator.New()

	if err := ctx.ReadJSON(&params); err != nil {
		return response.InternalServerError("解析参数失败", err)
	}

	if err := validate.Struct(params); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("user", validationErrors))
	}

	pass, err := models.GeneratePassword(params.Pass)
	if err != nil {
		return response.InternalServerError("用户密码加密失败", err)
	}

	if user, _ := instance.Service.FindByEmail(params.Email); user != nil {
		return response.Send(400, "用户邮箱已经存在", make(map[string]interface{}))
	}

	user := &models.User{
		Id:        uuid.NewV4().String(),
		Name:      params.Name,
		Email:     params.Email,
		Password:  string(pass),
		TeamId:    params.TeamId,
		Manager:   params.Manager,
		CreatedAt: time.Now().Format("2016-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2016-01-02 15:04:05"),
	}

	if err := user.Store(); err != nil {
		return response.InternalServerError("创建用户失败", err)
	}

	return response.Success("创建用户成功", response.Payload{"data": user})
}

// 更新用户信息
func (instance *UserController) PutBy(id string, ctx iris.Context) mvc.Result {
	var params UpdateRequest
	var user models.User

	validate := validator.New()

	if err := ctx.ReadJSON(&params); err != nil {
		log.Println(err)
	}

	if err := validate.Struct(params); err != nil {
		return response.ValidationError("请填写名称")
	}

	result, err := models.Engine.Id(id).Get(&user)
	if err != nil {
		return response.InternalServerError("用户ID不存在", err)
	}

	if result {
		user.Name = params.Name
		user.Email = params.Email
		user.TeamId = params.TeamId
		user.Manager = params.Manager
		if _, err := models.Engine.Id(id).Update(user); err != nil {
			return response.Send(iris.StatusInternalServerError, "更新失败", make(map[string]interface{}))
		}
	}

	return response.Success("更新用户信息成功", response.Payload{"data": user})
}

// 删除用户
func (instance *UserController) DeleteBy(id string) mvc.Result {
	user := &models.User{
		Id: id,
	}

	_, err := models.Engine.Delete(user)
	if err != nil {

	}

	return response.Success("删除用户成功", response.Payload{"data": make(map[string]interface{})})
}
