package models

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/service"
	"github.com/betterde/ects/internal/utils"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	MODESHELL      = "shell"
	MODEHTTP       = "http"
	MODEMAIL       = "mail"
	MODEHOOK       = "hook"
)

// 任务模型
type Task struct {
	Id          string     `json:"id" validate:"-" xorm:"not null pk comment('用户ID') CHAR(36)"`
	Name        string     `json:"name" validate:"required" xorm:"not null comment('名称') VARCHAR(255)"`
	Mode        string     `json:"mode" validate:"required" xorm:"not null default('shell') comment('任务模式') VARCHAR(32)"`
	Url         string     `json:"url" validate:"omitempty" xorm:"null comment('请求URL') VARCHAR(255)"`
	Method      string     `json:"method" validate:"omitempty" xorm:"null comment('任务模式') VARCHAR(255)"`
	Content     string     `json:"content" validate:"omitempty" xorm:"null comment('内容') TEXT"`
	Description string     `json:"description" validate:"-" xorm:"null comment('描述') VARCHAR(255)"`
	CreatedAt   utils.Time `json:"created_at" validate:"-" xorm:"not null created comment('创建于') DATETIME"`
	UpdatedAt   utils.Time `json:"updated_at" validate:"-" xorm:"not null updated comment('更新于') DATETIME"`
}

// 定义模型的数据表名称
func (task *Task) TableName() string {
	return "tasks"
}

// 创建任务
func (task *Task) Store() error {
	_, err := Engine.Insert(task)
	return err
}

// 更新任务
func (task *Task) Update() error {
	_, err := Engine.Id(task.Id).Update(task)
	return err
}

// 删除任务
func (task *Task) Destroy() error {
	_, err := Engine.Delete(task)
	return err
}

// 序列化
func (task *Task) ToString() (string, error) {
	result, err := json.Marshal(task)
	return string(result), err
}

// 执行任务
func (task *Task) Exec(ctx context.Context, username string, dir string, env []string) (record *TaskRecords, err error) {
	record = &TaskRecords{
		TaskId:     task.Id,
		NodeId:     service.Runtime.Id,
		WorkerName: service.Runtime.Name,
		TaskName:   task.Name,
		Mode:       task.Mode,
		Content:    task.Content,
	}
	beginWith := time.Now()

	switch task.Mode {
	case MODESHELL:
		cmd := exec.CommandContext(ctx, "/bin/bash", "-c", task.Content)
		cmd.Env = env
		cmd.Dir = dir
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setpgid: true,
		}
		// 如果要以指定用户运行命令
		if username != "" {
			sysuser, err := user.Lookup(username)
			if err != nil {
				goto END
			}
			uid, err := strconv.Atoi(sysuser.Uid)
			gid, err := strconv.Atoi(sysuser.Gid)
			cmd.SysProcAttr.Credential = &syscall.Credential{
				Uid:         uint32(uid),
				Gid:         uint32(gid),
				Groups:      nil,
				NoSetGroups: false,
			}
		}

		resChan := make(chan struct {
			output []byte
			err    error
		})

		go func() {
			output, err := cmd.CombinedOutput()
			resChan <- struct {
				output []byte
				err    error
			}{output: output, err: err}
		}()

		select {
		case result := <-resChan:
			record.Result = string(result.output)
			if result.err != nil {
				record.Status = "failed"
			} else {
				record.Status = "finished"
			}
			break
		}
		break
	case MODEHOOK:
		var resp *http.Response
		resp, err = http.Post(task.Url, "application/json", strings.NewReader(task.Content))
		if err != nil {
			goto END
		}
		defer func() {
			if err := resp.Body.Close(); err != nil {
				log.Println(err)
			}
		}()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		record.Result = string(body)
		break
	case MODEHTTP:
		client := &http.Client{}
		req, err := http.NewRequest(task.Method, task.Url, strings.NewReader(task.Content))
		if err != nil {
			return record, err
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			return record, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return record, err
		}
		record.Result = string(body)
		break
	case MODEMAIL:
		mailer := gomail.NewDialer(config.Conf.Notification.Host, config.Conf.Notification.Port, config.Conf.Notification.User, config.Conf.Notification.Pass)
		mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		message := gomail.NewMessage()
		message.SetHeader("From", fmt.Sprintf("%s<%s>", config.Conf.Notification.Name, config.Conf.Notification.User))
		message.SetHeader("To", task.Content)
		message.SetHeader("Subject", "Notification")
		message.SetBody("text/html", "Hello Goerge")
		resChan := make(chan error)
		go func(resChan chan error) {
			if err := mailer.DialAndSend(message); err != nil {
				resChan <- err
			}
			resChan <- nil
		}(resChan)
		err = <-resChan
	}

END:

	finishWith := time.Now()

	record.BeginWith = utils.Time(beginWith)
	record.FinishWith = utils.Time(finishWith)
	record.FinishWith = utils.Time(time.Now())
	record.Duration = int64(finishWith.Sub(beginWith).Seconds())
	return record, err
}
