package actuator

import (
	"context"
	"github.com/betterde/ects/models"
)

type (
	Http struct {}
)

func (actuator *Http) Exec(ctx context.Context, result chan *models.Result) {

}
