package actuator

import (
	"context"
	"github.com/betterde/ects/internal/scheduler"
)

type (
	Mail struct {

	}
)

func (actuator *Mail) Exec(ctx context.Context, result chan *scheduler.Result) {

}
