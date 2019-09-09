package actuator

import (
	"context"
	"github.com/betterde/ects/internal/notify"
	"github.com/betterde/ects/models"
)

type (
	Mail struct {
		Mail *notify.Mail
	}
)

func (actuator *Mail) Exec(ctx context.Context) *models.TaskRecords {
	err := actuator.Mail.Generator("info").Send()
	if err != nil {
		return &models.TaskRecords{
			Status: "failed",
			Result: err.Error(),
		}
	}

	return &models.TaskRecords{
		Status: "finished",
	}
}
