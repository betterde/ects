package actuator

import (
	"context"
	"github.com/betterde/ects/internal/scheduler"
)

type (
	Http struct {}
)

func (actuator *Http) Exec(ctx context.Context, result chan *scheduler.Result) {

}
