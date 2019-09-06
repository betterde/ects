package actuator

import (
	"context"
	"fmt"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/notify"
	"github.com/betterde/ects/internal/service"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

type (
	Contract interface {
		Exec(ctx context.Context) *models.TaskRecords
	}
)

// 执行流水线
func RunPipeline(ctx context.Context, pipeline *models.Pipeline, resChan chan *models.Result) {
	if len(pipeline.Steps) > 0 {
		record := &models.PipelineRecords{
			Id:         uuid.NewV4().String(),
			PipelineId: pipeline.Id,
			NodeId:     service.Runtime.Id,
			WorkerName: service.Runtime.Name,
			Spec:       pipeline.Spec,
			Status:     1,
			Duration:   0,
		}
		beginWith := time.Now()
		result := &models.Result{}
		// 按照任务的排序，逐个执行
		for _, pivot := range pipeline.Steps {
			var pctx context.Context
			var cancelFunc func()

			// 如果设置了超时时间则
			if pivot.Timeout == 0 {
				pctx, cancelFunc = context.WithCancel(ctx)
			} else {
				pctx, cancelFunc = context.WithTimeout(ctx, time.Duration(pivot.Timeout)*time.Second)
			}
			taskRecord := RunStep(pctx, pivot)

			taskRecord.Timeout = pivot.Timeout
			taskRecord.Retries = pivot.Retries
			taskRecord.PipelineRecordId = record.Id
			taskRecord.CreatedAt = utils.Time(time.Now())

			result.Steps = append(result.Steps, taskRecord)
			if taskRecord.Status == "failed" {
				record.Status = 0
				goto END
			}

			select {
			case <-ctx.Done():
				cancelFunc()
				break
			default:
				continue
			}
		}
	END:
		finishWith := time.Now()
		record.Duration = int64(finishWith.Sub(beginWith).Seconds())
		record.BeginWith = utils.Time(beginWith)
		record.FinishWith = utils.Time(finishWith)
		record.FinishWith = utils.Time(time.Now())
		record.CreatedAt = utils.Time(time.Now())
		record.UpdatedAt = utils.Time(time.Now())

		// 当流水线成功时触发
		if record.Status == 1 && pipeline.Finished != "" {
			switch pipeline.FinishedTask.Mode {
			case models.MODESHELL:
				shell :=&Shell{
					User:    "",
					Env:     nil,
					Dir:     "",
					Command: "",
				}
				shell.Exec(ctx)
				break
			case models.MODEHTTP:
				break
			case models.MODEHOOK:
				break
			case models.MODEMAIL:
				break
			}
		}

		// 当流水线失败时触发
		if record.Status == 0 && pipeline.Failed != "" {
			switch pipeline.FailedTask.Mode {
			case models.MODESHELL:
				shell :=&Shell{
					User:    "",
					Env:     nil,
					Dir:     "",
					Command: "",
				}
				shell.Exec(ctx)
				break
			case models.MODEHTTP:
				break
			case models.MODEHOOK:
				break
			case models.MODEMAIL:
				break
			}
		}

		result.Pipeline = record
		resChan <- result
	}
}

// 运行任务
func RunStep(ctx context.Context, pivot *models.PipelineTaskPivot) *models.TaskRecords {
	record := &models.TaskRecords{}
	beginWith := time.Now()
	if pivot.Retries == 0 {
		record = runActuator(ctx, pivot)
	} else {
		for i := 0; i < pivot.Retries; i ++ {
			record = runActuator(ctx, pivot)
			if record.Status == "finished" {
				break
			}
		}
	}

	record.TaskId = pivot.TaskId
	record.NodeId = service.Runtime.Id
	record.TaskName = pivot.Task.Name
	record.WorkerName = service.Runtime.Name
	record.Content =  pivot.Task.Content
	record.Mode = pivot.Task.Mode
	record.Timeout = pivot.Timeout
	record.Retries = pivot.Retries
	finishWith := time.Now()
	record.BeginWith = utils.Time(beginWith)
	record.FinishWith = utils.Time(time.Now())
	record.Duration = int64(finishWith.Sub(beginWith).Seconds())
	return record
}

func runActuator(ctx context.Context, pivot *models.PipelineTaskPivot) *models.TaskRecords {
	switch pivot.Task.Mode {
	case models.MODESHELL:
		shell :=&Shell{
			User:    pivot.User,
			Env:     strings.Split(pivot.Environment, " "),
			Dir:     pivot.Directory,
			Command: pivot.Task.Content,
		}
		return shell.Exec(ctx)
	case models.MODEMAIL:
		mail := Mail{
			Mail: &notify.Mail{
				From:        fmt.Sprintf("%s<%s>", config.Conf.Notification.Name, config.Conf.Notification.User),
				To:          pivot.Task.Url,
				Subject:     pivot.Task.Name,
				Year:        time.Now().Year(),
				SiteURL:     config.Conf.Notification.Url,
				SiteTitle:   "Elastic Crontab System",
				Greeting:    "Hello",
				Intro:       pivot.Task.Content,
				Outro:       "",
				Salutation:  "Regards",
			},
		}

		return mail.Exec(ctx)
	case models.MODEHTTP:
		request := &Http{
			Url:     pivot.Task.Url,
			Method:  pivot.Task.Method,
			Content: pivot.Task.Content,
		}
		return request.Exec(ctx)
	case models.MODEHOOK:
		request := &Http{
			Url:     pivot.Task.Url,
			Method:  pivot.Task.Method,
			Content: pivot.Task.Content,
		}
		return request.Exec(ctx)
	}

	return nil
}
