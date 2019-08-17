package actuator

import (
	"context"
	"github.com/betterde/ects/internal/scheduler"
)

type (
	Hook struct {

	}
)

func (actuator *Hook) Exec(ctx context.Context, result chan *scheduler.Result) {

}
