package services

type (
	PipelineInterface interface {
	}

	PipelineService struct {
	}
)

func NewPipelineService() PipelineInterface {
	return &PipelineService{}
}
