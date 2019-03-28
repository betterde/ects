package response

import (
	"github.com/kataras/iris"
	"reflect"
)

type (
	Meta struct {
		Page int `json:"page"`
		Limit int `json:"limit"`
		Total int `json:"total"`
	}
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta *Meta `json:"meta,omitempty"`
	}
	Payload map[string]interface{}
)

// 发送成功响应
func Success(message string, payload map[string]interface{}) *Response {
	meta, ok := payload["meta"]
	if ok {
		return &Response{
			Code: iris.StatusOK,
			Message: message,
			Data: payload["data"],
			Meta: reflect.ValueOf(meta).Interface().(*Meta),
		}
	}

	return &Response{
		Code: iris.StatusOK,
		Message: message,
		Data: payload["data"],
	}
}

// 认证失败响应
func UnAuthenticated(message string) *Response {
	return &Response{
		Code: iris.StatusUnauthorized,
		Message: message,
		Data: make(map[string]interface{}),
	}
}

func NotFound(message string) *Response {
	return &Response{
		Code: iris.StatusNotFound,
		Message: message,
		Data: make(map[string]interface{}),
	}
}

func ValidationError(message string) *Response {
	return &Response{
		Code: iris.StatusUnprocessableEntity,
		Message: message,
		Data: make(map[string]interface{}),
	}
}

func InternalServerError(message string, data interface{}) *Response {
	return &Response{
		Code: iris.StatusInternalServerError,
		Message: message,
		Data: data,
	}
}

func Send(code int, message string, data interface{}) *Response {
	return &Response{
		Code: code,
		Message: message,
		Data: data,
	}
}