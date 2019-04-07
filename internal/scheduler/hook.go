package scheduler

import "github.com/betterde/ects/models"

type (
	HookHandler struct {

	}
)

func (handler *HookHandler) Run(task models.Task, id string) (result string, err error) {

}
