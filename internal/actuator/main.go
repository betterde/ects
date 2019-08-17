package actuator

import (
	"context"
	"github.com/betterde/ects/internal/scheduler"
)

type (
	Contract interface {
		Exec(ctx context.Context, result chan *scheduler.Result)
	}
)