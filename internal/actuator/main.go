package actuator

import (
	"context"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	uuid "github.com/satori/go.uuid"
	"log"
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
			NodeId:     "",
			WorkerName: "",
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
			env := strings.Split(pivot.Environment, " ")
			taskRecord, err := pivot.Task.Exec(pctx, pivot.User, pivot.Directory, env)
			if err != nil {
				record.Status = 0
				goto END
			}

			taskRecord.NodeId = ""
			taskRecord.Timeout = pivot.Timeout
			taskRecord.Retries = pivot.Retries
			taskRecord.WorkerName = ""
			taskRecord.CreatedAt = utils.Time(time.Now())
			taskRecord.UpdatedAt = utils.Time(time.Now())

			result.Steps = append(result.Steps, taskRecord)

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
		if err := record.Store(); err != nil {
			// TODO
		}

		// 当流水线成功时触发
		if record.Status == 1 && pipeline.Finished != "" {
			if _, err := pipeline.FinishedTask.Exec(context.Background(), "", "", []string{}); err != nil {
				log.Println(err)
			}
		}

		// 当流水线失败时触发
		if record.Status == 0 && pipeline.Failed != "" {
			if _, err := pipeline.FailedTask.Exec(context.Background(), "", "", []string{}); err != nil {
				log.Println(err)
			}
		}

		result.Pipeline = record
		resChan <- result
	}
}
