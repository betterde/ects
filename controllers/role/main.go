package role

import (
	"github.com/betterde/ects/internal/message"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/services"
	"github.com/go-xorm/builder"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
)

type (
	Controller struct {
		Service services.RoleService
	}
	BindPermissionsRequest struct {
		RoleId string   `json:"role_id" validate:"required,uuid4"`
		Codes  []string `json:"codes" validate:"required"`
	}
)

var (
	validate = validator.New()
)

// Get roles list
func (instance *Controller) Get(ctx iris.Context) mvc.Result {
	var (
		page   = 1
		limit  = 10
		start  int
		search = ctx.URLParam("name")
		total  int64
		err    error
	)

	roles := make([]models.Role, 0)

	start = (page - 1) * limit

	if search != "" {
		total, err = models.Engine.Where(builder.Like{"name", search}).Limit(limit, start).FindAndCount(&roles)
	} else {
		total, err = models.Engine.Limit(limit, start).FindAndCount(&roles)
	}

	if err != nil {
		return response.InternalServerError("Failed to query roles list", err)
	}

	return response.Success("Success", response.Payload{
		"data": roles,
		"meta": &response.Meta{
			Limit: limit,
			Page:  page,
			Total: int(total),
		},
	})
}

// Create role
func (instance *Controller) Post(ctx iris.Context) mvc.Result {
	role := models.Role{}
	if err := ctx.ReadJSON(&role); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(role); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("role", validationErrors))
	}
	role.Id = uuid.NewV4().String()
	role.System = 0
	err := role.Store()
	if err != nil {
		return response.InternalServerError("Failed to create role", err)
	}

	return response.Success("Role created successfully", response.Payload{"data": role})
}

// Update role attributes by id
func (instance *Controller) PutBy(id string, ctx iris.Context) mvc.Result {
	role := models.Role{}
	if err := ctx.ReadJSON(&role); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(role); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("role", validationErrors))
	}

	role.Id = id
	if err := role.Update(); err != err {
		return response.InternalServerError("Failed to update role", err)
	}

	return response.Success("Role updated successfully", response.Payload{"data": role})
}

// Delete role by id
func (instance *Controller) DeleteBy(id string) mvc.Result {
	role := &models.Role{
		Id: id,
	}

	if err := role.Destroy(); err != nil {
		return response.InternalServerError("Failed to delete role", err)
	}
	return response.Success("Role deleted successfully", response.Payload{"data": make(map[string]interface{})})
}

// Get role permissions
func (instance *Controller) GetPermissions(ctx iris.Context) mvc.Result {
	id := ctx.URLParam("role_id")

	if id == "" {
		return response.ValidationError("pipeline id is required")
	}

	relations := make([]models.RolePermissionPivot, 0)

	if err := models.Engine.Where(builder.Eq{"role_id": id}).Find(&relations); err != nil {
		return response.InternalServerError("Failed to query relations", err)
	}

	codes := make([]string, 0)

	for _, pivot := range relations {
		codes = append(codes, pivot.PermissionCode)
	}

	permissions := make([]models.Permission, 0)

	if err := models.Engine.Where(builder.Eq{"code": codes}).Find(&permissions); err != nil {
		return response.InternalServerError("Failed to query relations", err)
	}

	return response.Success("Success", response.Payload{"data": permissions})
}

// Bind the permissions to role
func (instance *Controller) PostPermissions(ctx iris.Context) mvc.Result {
	params := BindPermissionsRequest{}

	if err := ctx.ReadJSON(&params); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(params); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("pipeline", validationErrors))
	}

	if _, err := models.Engine.Where(builder.Eq{"pipeline_id": params.RoleId}).Delete(&models.PipelineNodePivot{}); err != nil {
		return response.InternalServerError("Failed to delete pipeline and node relations", err)
	}

	relations := make([]*models.RolePermissionPivot, 0)

	for _, code := range params.Codes {
		relations = append(relations, &models.RolePermissionPivot{
			RoleId: params.RoleId,
			PermissionCode: code,
		})
	}

	_, err := models.Engine.Insert(relations)
	if err != nil {
		return response.InternalServerError("Failed to bind pipeline to node", err)
	}

	return response.Success("Bind successfully", response.Payload{"data": relations})
}
