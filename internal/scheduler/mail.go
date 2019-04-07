package scheduler

import "github.com/betterde/ects/models"

type (
	MailHandler struct {

	}
)

func (handler *MailHandler) Run(task models.Task, id string) (result string, err error) {

}
