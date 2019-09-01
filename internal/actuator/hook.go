package actuator

import (
	"context"
	"github.com/betterde/ects/models"
)

type (
	Hook struct {

	}
)

func (actuator *Hook) Exec(ctx context.Context, result chan *models.Result) {

}
