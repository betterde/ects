package actuator

import (
	"context"
	"github.com/betterde/ects/models"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

type (
	Shell struct {
		User string
		Env  []string
		Dir  string
		Command string
	}
)

// 执行 Shell 任务
func (actuator *Shell) Exec(ctx context.Context) *models.TaskRecords {
	cmd := exec.CommandContext(ctx, "/bin/bash", "-c", actuator.Command)
	record :=  &models.TaskRecords{}
	if actuator.User != "" {
		credential, err := getCredential(actuator.User)
		if err != nil {
			record.Status = "failed"
			record.Result = err.Error()
			return record
		}
		cmd.SysProcAttr.Credential = credential
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
	res := <-resChan
	record.Result = string(res.output)
	if res.err != nil {
		record.Status = "failed"
	} else {
		record.Status = "finished"
	}
	return record
}

// 获取执行证书
func getCredential(username string) (*syscall.Credential, error) {
	sysuser, err := user.Lookup(username)
	if err != nil {
		return nil, err
	}
	uid, err := strconv.Atoi(sysuser.Uid)
	gid, err := strconv.Atoi(sysuser.Gid)
	return &syscall.Credential{
		Uid:         uint32(uid),
		Gid:         uint32(gid),
		Groups:      nil,
		NoSetGroups: false,
	}, nil
}
