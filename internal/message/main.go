package message

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

var (
	modules = map[string]map[string]map[string]string{
		"task":     taskMessage(),
		"user":     userMessage(),
		"role":     roleMessage(),
		"team":     teamMessage(),
		"pipeline": pipelineMessage(),
	}
)

// 获取制定模块表单验证的单条消息
func Get(module string, validationErrors validator.ValidationErrors) string {
	first := validationErrors[0]
	//field := first.Field()
	//condition := first.Tag()
	//return modules[module][field][condition]
	return fmt.Sprintf("Key: '%s' Error:Field validation for '%s' failed on the '%s' tag", first.Namespace(), first.Field(), first.Tag())
}
