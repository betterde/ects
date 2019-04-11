package scheduler

import (
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/services"
	"log"
	"sync"
)

type (
	Handler interface {
		Run(task models.Task, id string) (result string, err error)
	}

	TaskCount struct {
		wg   sync.WaitGroup
		exit chan struct{}
	}

	Instance struct {
		m sync.Map
	}

	TaskResult struct {
		Result     string
		Err        error
		RetryTimes int8
	}
)

var (
	cron        *Cron
	taskCount   TaskCount
	runInstance Instance
)

func Initialize() {
	cron = New()
	cron.Start()
	service := services.NewTaskService()
	tasks, err := service.NormalTasks()

	if err != nil {
		log.Println(err)
	}

	for _, task := range tasks {
		cron.AddFunc(task.Spec, task.Content)
	}
}

func createJob(task models.Task) FuncJob {
	handler := createHandler(task)
	if handler == nil {
		return nil
	}

	taskFunc := func() {
		taskCount.Add()
		defer taskCount.Done()
		log := before(task)
	}

	return taskFunc
}

func createHandler(task models.Task) Handler {
	var handler Handler = nil
	switch task.Mode {
	case models.TASK_MODE_SHELL:
		handler = new(ShellHandler)
		break
	case models.TASK_MODE_REST:
		handler = new(RestHandler)
		break
	case models.TASK_MODE_HOOK:
		handler = new(HookHandler)
		break
	case models.TASK_MODE_MAIL:
		handler = new(MailHandler)
		break
	}

	return handler
}

func before(task models.Task) (log int64) {
	if task.Overlap == false && runInstance.has(task.ID) {
		createTaskLog(taskModel, models.Cancel)
		return
	}
	taskLogId, err := createTaskLog(taskModel, models.Running)
	if err != nil {
		logger.Error("任务开始执行#写入任务日志失败-", err)
		return
	}
	logger.Debugf("任务命令-%s", taskModel.Command)

	return taskLogId
}

func after(task models.Task, result TaskResult, id string)  {

}

func (tc *TaskCount) Add() {
	tc.wg.Add(1)
}

func (tc *TaskCount) Done() {
	tc.wg.Done()
}

func (tc *TaskCount) Exit() {
	tc.wg.Done()
	<-tc.exit
}

func (tc *TaskCount) Wait() {
	tc.Add()
	tc.wg.Wait()
	close(tc.exit)
}

// 是否有任务处于运行中
func (i *Instance) has(id string) bool {
	_, ok := i.m.Load(id)

	return ok
}

func (i *Instance) add(key int) {
	i.m.Store(key, struct{}{})
}

func (i *Instance) done(key int) {
	i.m.Delete(key)
}
