package actuator

import (
	"context"
	"github.com/betterde/ects/internal/scheduler"
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
func (actuator *Shell) Exec(ctx context.Context, result chan *scheduler.Result) {
	cmd := exec.CommandContext(ctx, "/bin/bash", "-c", actuator.Command)
	if actuator.User != "" {
		sysuser, err := user.Lookup(actuator.User)
		if err != nil {
			//
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

// 创建一个 Shell 执行器
func NewShell(username, dir string, env []string) *Shell {
	return &Shell{
		User: username,
		Env:  env,
		Dir:  dir,
	}
}
