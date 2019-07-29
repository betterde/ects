package organization

import (
	"github.com/betterde/ects/internal/message"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/services"
	"github.com/go-xorm/builder"
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
		Manager bool   `json:"manager" validate:"required"`
	}
	UpdateRequest struct {
		Name    string `json:"name" validate:"required"`
		Email   string `json:"email" validate:"required,email"`
		Manager bool   `json:"manager" validate:"required"`
	}
)

// Get users list
func (instance *UserController) Get(ctx iris.Context) mvc.Result {
	var (
		total int64
		err   error
	)
	search := ctx.Params().GetStringDefault("search", "")
	page, limit, start := utils.Pagination(ctx)
	users := make([]models.User, 0)

	if search == "" {
		total, err = models.Engine.Limit(limit, start).FindAndCount(&users)
	} else {
		total, err = models.Engine.Where(builder.Like{"name", search}).Limit(limit, start).FindAndCount(&users)
	}

	if err != nil {
		return response.InternalServerError("Failed to query pipelines list", err)
	}

	return response.Success("Successful", response.Payload{
		"data": users,
		"meta": &response.Meta{
			Limit: limit,
			Page:  page,
			Total: int(total),
		},
	})
}

// Create user
func (instance *UserController) Post(ctx iris.Context) mvc.Result {
	var (
		params CreateRequest
	)

	validate := validator.New()

	if err := ctx.ReadJSON(&params); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(params); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("user", validationErrors))
	}

	pass, err := models.GeneratePassword(params.Pass)
	if err != nil {
		return response.InternalServerError("Failed to encryption user password", err)
	}

	if user, _ := instance.Service.FindByEmail(params.Email); user != nil {
		return response.Send(400, "This mail address already exists", make(map[string]interface{}))
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

	return response.Success("Created successful", response.Payload{"data": user})
}

// Modify user attribute
func (instance *UserController) PutBy(id string, ctx iris.Context) mvc.Result {
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
		if _, err := models.Engine.Id(id).Update(user); err != nil {
			return response.Send(iris.StatusInternalServerError, "Failed to update user", make(map[string]interface{}))
		}
	}

	return response.Success("Updated successful", response.Payload{"data": user})
}

// Delete user
func (instance *UserController) DeleteBy(id string) mvc.Result {
	user := &models.User{
		Id: id,
	}

	_, err := models.Engine.Delete(user)
	if err != nil {

	}

	return response.Success("Deleted successful", response.Payload{"data": make(map[string]interface{})})
}
