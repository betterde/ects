package organization

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

type TeamController struct {
	Service services.TeamService
}

// Get teams list
func (instance *TeamController) Get(ctx iris.Context) mvc.Result {
	var (
		page   = 1
		limit  = 10
		start  int
		search = ctx.URLParam("name")
		total  int64
		err    error
	)

	teams := make([]models.Team, 0)

	start = (page - 1) * limit

	if search != "" {
		total, err = models.Engine.Where(builder.Like{"name", search}).Limit(limit, start).FindAndCount(&teams)
	} else {
		total, err = models.Engine.Limit(limit, start).FindAndCount(&teams)
	}

	if err != nil {
		return response.InternalServerError("Failed to query teams list", err)
	}

	return response.Success("Success", response.Payload{
		"data": teams,
		"meta": &response.Meta{
			Limit: limit,
			Page:  page,
			Total: int(total),
		},
	})
}

// Create team
func (instance *TeamController) Post(ctx iris.Context) mvc.Result {
	team := models.Team{}
	validate := validator.New()

	if err := ctx.ReadJSON(&team); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(team); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("team", validationErrors))
	}

	team.Id = uuid.NewV4().String()

	if err := team.Store(); err != nil {
		return response.InternalServerError("Failed to create team", err)
	}

	// Create operation log
	if err := models.CreateLog(team, ctx, "CREATE TEAM"); err != nil {
		return response.InternalServerError("Failed to create log", err)
	}

	return response.Success("Team created successfully", response.Payload{"data": team})
}

// Update team attributes by id
func (instance *TeamController) PutBy(id string, ctx iris.Context) mvc.Result {
	team := models.Team{}
	validate := validator.New()

	if err := ctx.ReadJSON(&team); err != nil {
		return response.InternalServerError("Failed to Unmarshal JSON", err)
	}

	if err := validate.Struct(team); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ValidationError(message.Get("team", validationErrors))
	}

	team.Id = id
	if err := team.Update(); err != err {
		return response.InternalServerError("Failed to update team", err)
	}

	// Create operation log
	if err := models.CreateLog(team, ctx, "MODIFY TEAM"); err != nil {
		return response.InternalServerError("Failed to create log", err)
	}

	return response.Success("Team updated successfully", response.Payload{"data": team})
}

// Delete team by id
func (instance *TeamController) DeleteBy(id string, ctx iris.Context) mvc.Result {
	team := &models.Team{
		Id: id,
	}

	if _, err := models.Engine.Get(team); err != nil {
		return response.InternalServerError("Failed to delete team", err)
	}

	if err := team.Destroy(); err != nil {
		return response.InternalServerError("Failed to delete team", err)
	}

	// Create operation log
	if err := models.CreateLog(*team, ctx, "DELETE TEAM"); err != nil {
		return response.InternalServerError("Failed to create log", err)
	}

	return response.Success("Team deleted successfully", response.Payload{"data": make(map[string]interface{})})
}
