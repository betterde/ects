package organization

import (
	"github.com/betterde/ects/internal/message"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/services"
	"github.com/go-xorm/builder"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
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
		Manager bool   `json:"manager"`
	}

	UpdateRequest struct {
		Name    string `json:"name" validate:"required"`
		Email   string `json:"email" validate:"required,email"`
		Manager bool   `json:"manager"`
	}
)

// 获取用户列表
func (instance *UserController) Get(ctx iris.Context) mvc.Response {
	var (
		total int64
		err   error
	)
	scene := ctx.URLParamDefault("scene", "table")
	search := ctx.URLParamDefault("search", "")
	users := make([]models.User, 0)

	switch scene {
	case "table":
		page, limit, start := utils.Pagination(ctx)

		if search == "" {
			total, err = models.Engine.Limit(limit, start).FindAndCount(&users)
		} else {
			total, err = models.Engine.Where(builder.Like{"name", search}).Limit(limit, start).FindAndCount(&users)
		}

		if err != nil {
			return response.InternalServerError("获取用户列表失败", err)
		}

		return response.Success("请求成功", response.Payload{
			"data": users,
			"meta": &response.Meta{
				Limit: limit,
				Page:  page,
				Total: int(total),
			},
		})
	case "selector":
		if search == "" {
			err = models.Engine.Find(&users)
		} else {
			err = models.Engine.Where(builder.Like{"name", search}.Or(builder.Like{"email", search})).Find(&users)
		}

		if err != nil {
			return response.InternalServerError("获取用户列表失败", err)
		}

		return response.Success("请求成功", response.Payload{"data": users})
	}

	return response.Success("请求成功", response.Payload{"data": make([]interface{}, 0)})
}

// 创建用户
func (instance *UserController) Post(ctx iris.Context) mvc.Response {
	var (
		params CreateRequest
	)

	validate := validator.New()

	if err := ctx.ReadJSON(&params); err != nil {
		return response.InternalServerError("参数解析失败", err)
	}

	if err := validate.Struct(params); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("user", validationErrors))
	}

	pass, err := models.GeneratePassword(params.Pass)
	if err != nil {
		return response.InternalServerError("Failed to encryption user password", err)
	}

	if user := instance.Service.FindByEmail(params.Email); user != nil {
		return response.Send(400, "用户邮箱已存在", make(map[string]interface{}))
	}

	user := &models.User{
		Id:        uuid.NewV4().String(),
		Name:      params.Name,
		Email:     params.Email,
		Password:  string(pass),
		Manager:   params.Manager,
		CreatedAt: utils.Time(time.Now()),
		UpdatedAt: utils.Time(time.Now()),
	}

	if err := user.Store(); err != nil {
		return response.InternalServerError("Failed to create user", err)
	}

	return response.Success("创建成功", response.Payload{"data": user})
}

// 修改用户信息
func (instance *UserController) PutBy(id string, ctx iris.Context) mvc.Response {
	var params UpdateRequest
	var user models.User

	validate := validator.New()

	if err := ctx.ReadJSON(&params); err != nil {
		log.Println(err)
	}

	if err := validate.Struct(params); err != nil {
		return response.ValidationError("Failed to validate params")
	}

	result, err := models.Engine.Id(id).Get(&user)
	if err != nil {
		return response.InternalServerError("User does not exist", err)
	}

	if result {
		user.Name = params.Name
		user.Email = params.Email
		user.Manager = params.Manager
		if _, err := models.Engine.Id(id).Update(&user); err != nil {
			return response.Send(iris.StatusInternalServerError, "Failed to update user", err)
		}
	}

	return response.Success("Updated successful", response.Payload{"data": user})
}

// 删除用户
func (instance *UserController) DeleteBy(id string) mvc.Response {
	user := &models.User{
		Id: id,
	}

	_, err := models.Engine.Delete(user)
	if err != nil {

	}

	return response.Success("Deleted successful", response.Payload{"data": make(map[string]interface{})})
}
