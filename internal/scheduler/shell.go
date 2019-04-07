package scheduler

import "github.com/betterde/ects/models"

type (
	ShellHandler struct {

	}
)

func (handler *ShellHandler) Run(task models.Task, id string) (result string, err error) {

}
