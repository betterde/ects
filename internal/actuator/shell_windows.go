//go:build windows
// +build windows

package actuator

import (
	"context"
	"github.com/betterde/ects/models"
	"os/exec"
	"strconv"
	"syscall"
)

// Exec Execute shell tasks
func (actuator *Shell) Exec(ctx context.Context) *models.TaskRecords {
	cmd := exec.CommandContext(ctx, "cmd", "/C", actuator.Command)
	record := &models.TaskRecords{}

	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	if actuator.User != "" {
		token, err := syscall.OpenCurrentProcessToken()
		if err != nil {
			record.Status = "failed"
			record.Result = err.Error()
			return record
		}

		cmd.SysProcAttr.Token = token
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
	case <-ctx.Done():
		exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(cmd.Process.Pid)).Run()
		cmd.Process.Kill()
	case res := <-resChan:
		record.Result = string(res.output)
		if res.err != nil {
			record.Status = "failed"
		} else {
			record.Status = "finished"
		}
	}

	return record
}
