package services

type (
	RoleInterface interface {
	}
	RoleService struct {
	}
)

func NewRoleService() RoleInterface {
	return &RoleService{}
}
