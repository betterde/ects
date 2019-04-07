package scheduler

import "github.com/betterde/ects/models"

type (
	RestHandler struct {

	}
)

func (handler *RestHandler) Run(task models.Task, id string) (result string, err error) {

}
