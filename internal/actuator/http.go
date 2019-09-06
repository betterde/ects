package actuator

import (
	"context"
	"github.com/betterde/ects/models"
	"io/ioutil"
	"net/http"
	"strings"
)

type (
	Http struct {
		Url     string
		Method  string
		Content string
	}
)

func (actuator *Http) Exec(ctx context.Context) *models.TaskRecords {
	client := &http.Client{}
	record := &models.TaskRecords{}
	req, err := http.NewRequest(actuator.Method, actuator.Url, strings.NewReader(actuator.Content))
	if err != nil {
		record.Result = err.Error()
		record.Status = "failed"
		return record
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		record.Result = err.Error()
		record.Status = "failed"
		return record
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		record.Result = err.Error()
		record.Status = "failed"
		return record
	}

	record.Result = string(body)
	record.Status = "finished"
	return record
}
