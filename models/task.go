package models

import (
	"encoding/json"
	"github.com/betterde/ects/internal/utils"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

const (
	STATUSNORMAL   = "normal"
	STATUSDISABLED = "disabled"
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
	Url         string     `json:"url" validate:"omitempty" xorm:"null comment('任务模式') VARCHAR(255)"`
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

func (task *Task) Exec(username string, dir string, env []string) ([]byte, error) {
	switch task.Mode {
	case MODESHELL:
		cmd := exec.Command("/bin/bash", "-c", task.Content)
		cmd.Env = env
		cmd.Dir = dir
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setpgid:    true,
		}
		if username != "" {
			sysuser, err := user.Lookup(username)
			if err != nil {
				return []byte{}, nil
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
		break
	case MODEHOOK:
		break
	case MODEHTTP:
		break
	case MODEMAIL:
		break
	}
	return []byte{1,2,4}, nil
}
