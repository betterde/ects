package services

type (
	PermissionInterface interface {
	}
	PermissionService struct {
	}
)

func NewPermissionService() PermissionInterface {
	return &PermissionService{}
}
