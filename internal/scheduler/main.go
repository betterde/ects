package scheduler

import (
	"context"
	"github.com/betterde/ects/internal/actuator"
	"github.com/betterde/ects/models"
	"github.com/gorhill/cronexpr"
	"log"
	"time"
)

const (
	PUT  = 1 // 新增或更新事件
	DEL  = 2 // 删除事件
	KILL = 3 // 强行终止进程事件
)

type (
	Event struct {
		Type     int              // 事件类型
		Pipeline *models.Pipeline // 流水线
	}
	Contract interface {
		Run(ctx context.Context)             // 运行调度器
		DispatchEvent(event *Event)          // 分发事件
		eventHandler(event *Event)           // 事件处理
		ResultHandler(result *models.Result) // 调度结果处理
	}
)

type Scheduler struct {
	EventsChan chan *Event                 // 事件通道
	ResultChan chan *models.Result         // 执行结果通道
	Plan       map[string]*models.Pipeline // 调度计划
	Running    map[string]*models.Pipeline // 正在运行的流水线
}

var Instance *Scheduler

// 运行调度器
func (scheduler *Scheduler) Run(ctx context.Context) {
	afterTimer := scheduler.TryExecute(ctx)
	scheduleTimer := time.NewTimer(afterTimer)

	for {
		select {
		case event := <-scheduler.EventsChan:
			scheduler.eventHandler(event)
		case <-scheduleTimer.C:
		case result := <-scheduler.ResultChan:
			//log.Printf("%#v", result)
			if err := result.Pipeline.Store(); err != nil {
				log.Fatal(err)
			}
			for _, step := range result.Steps {
				if err := step.Store(); err != nil {
					log.Fatal(err)
				}
			}
		}

		after := scheduler.TryExecute(ctx)
		scheduleTimer.Reset(after)
	}
}

// 尝试执行Pipeline
func (scheduler *Scheduler) TryExecute(ctx context.Context) (after time.Duration) {
	var nearTime time.Time
	if len(scheduler.Plan) == 0 {
		after = 1 * time.Second
		return
	}

	scheduler.ResultChan = make(chan *models.Result, 1000)

	now := time.Now()

	for _, pipe := range scheduler.Plan {
		if pipe.NextTime.Before(now) || pipe.NextTime.Equal(now) {
			actuator.RunPipeline(ctx, pipe, scheduler.ResultChan)
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
func (scheduler *Scheduler) eventHandler(event *Event) {
	switch event.Type {
	case PUT:
		event.Pipeline.Expression = cronexpr.MustParse(event.Pipeline.Spec)
		scheduler.Plan[event.Pipeline.Id] = event.Pipeline
	case DEL:
		delete(scheduler.Plan, event.Pipeline.Id)
	case KILL:
		// TODO KILL handler
	}
}

// 加入队列
func (scheduler *Scheduler) DispatchEvent(event *Event) {
	scheduler.EventsChan <- event
}

// 创建调度器
func New() {
	Instance = &Scheduler{
		EventsChan: make(chan *Event, 100),
		ResultChan: make(chan *models.Result, 100),
		Plan:       make(map[string]*models.Pipeline),
		Running:    make(map[string]*models.Pipeline),
	}
}
