package services

type (
	LogInterface interface {

	}

	LogService struct {

	}
)

func NewLogService() LogInterface {
	return &LogService{}
}
