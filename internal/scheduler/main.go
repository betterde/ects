package scheduler

import (
	"context"
	"github.com/betterde/ects/internal/pipeline"
	"github.com/betterde/ects/models"
	"time"
)

type Scheduler struct {
	EventsChan chan *pipeline.Event
	Plan       map[string]*models.Pipeline
	Running    map[string]*models.Pipeline
}

var Instance *Scheduler

// 运行调度器
func (scheduler *Scheduler) Run(ctx context.Context) {
	afterTimer := scheduler.TryExecute()
	scheduleTimer := time.NewTimer(afterTimer)
	
	for {
		select {
		case event := <-scheduler.EventsChan:
			scheduler.eventHandler(event)
		case <-scheduleTimer.C:

		}

		after := scheduler.TryExecute()
		scheduleTimer.Reset(after)
	}
}

// 尝试执行Pipeline
func (scheduler *Scheduler) TryExecute() (after time.Duration) {
	var nearTime time.Time
	if len(scheduler.Plan) == 0 {
		after = 1 * time.Second
		return
	}

	now := time.Now()

	for _, pipe := range scheduler.Plan {
		if pipe.NextTime.Before(now) || pipe.NextTime.Equal(now) {
			// TODO Execute Pipeline
			pipe.NextTime = pipe.Expression.Next(now)
		}

		if nearTime.IsZero() || pipe.NextTime.Before(nearTime) {
			nearTime = pipe.NextTime
		}
	}

	after = nearTime.Sub(now)
	return
}

// ETCD事件处理
func (scheduler *Scheduler) eventHandler(event *pipeline.Event) {
	switch event.Type {
	case pipeline.PUT:
		scheduler.Plan[event.Pipeline.Id] = event.Pipeline
	case pipeline.DEL:
		delete(scheduler.Plan, event.Pipeline.Id)
	case pipeline.KILL:
		// TODO KILL handler
	}
}

// 加入队列
func (scheduler *Scheduler) PushEvent(event *pipeline.Event) {
	scheduler.EventsChan <- event
}

// 创建调度器
func New() {
	Instance = &Scheduler{
		EventsChan: make(chan *pipeline.Event, 100),
		Plan:       make(map[string]*models.Pipeline),
		Running:    make(map[string]*models.Pipeline),
	}
}
