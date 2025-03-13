package global

import "context"

var (
	Ctx        context.Context
	CancelFunc context.CancelFunc
)

func init() {
	Ctx, CancelFunc = context.WithCancel(context.Background())
}
