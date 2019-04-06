package services

import (
	"errors"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/models"
	"github.com/go-xorm/builder"
	"log"
	"strconv"
)

type (
	TeamInterface interface {
		Teams(params map[string]string) (*[]models.Team, *response.Meta)
		FindByID(id string) (*models.Team, error)
		Update(id string, params *UpdateRequest) (*models.Team, error)
		Delete(id string) error
	}

	TeamService struct {
	}

	CreateRequest struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}

	UpdateRequest struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}
)

func NewTeamService() TeamInterface {
	return &TeamService{}
}

// 获取团队列表
func (service *TeamService) Teams(params map[string]string) (*[]models.Team, *response.Meta) {
	var (
		page  = 1
		limit = 10
		start int
		total int64
		err   error
	)

	if value, exist := params["page"]; exist == true {
		v, err := strconv.Atoi(value)
		if err != nil {

		}
		if v >= 0 {
			page = v
		}
	}

	if value, exist := params["limit"]; exist == true {
		v, err := strconv.Atoi(value)
		if err != nil {

		}
		if v >= 0 {
			limit = v
		}
	}

	start = (page - 1) * limit

	teams := make([]models.Team, 0)

	if search, exist := params["search"]; exist && search != "" {
		total, err = models.Engine.Where(builder.Like{"name", search}).Count(&models.Team{})
		err = models.Engine.Where(builder.Like{"name", search}).Limit(limit, start).Find(&teams)
	} else {
		total, err = models.Engine.Count(&models.Team{})
		err = models.Engine.Limit(limit, start).Find(&teams)
	}

	if err != nil {
		log.Println(err)
	}

	return &teams, &response.Meta{
		Limit: limit,
		Page:  page,
		Total: int(total),
	}
}

// 根据ID获取团队信息
func (service *TeamService) FindByID(id string) (*models.Team, error) {
	var team models.Team
	result, err := models.Engine.Id(id).Get(&team)
	if err != nil {
		log.Println(err)
	}

	if result {
		return &team, nil
	}

	return &team, errors.New("团队不存在")
}

// 更新团队信息
func (service *TeamService) Update(id string, params *UpdateRequest) (*models.Team, error) {

}

// 删除团队
func (service *TeamService) Delete(id string) error {
	_, err := models.Engine.Delete(&models.Worker{
		ID: id,
	})

	return err
}
