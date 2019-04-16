package response

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
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
func Success(message string, payload map[string]interface{}) mvc.Result {
	meta, ok := payload["meta"]
	if ok {
		return mvc.Response{
			Code: iris.StatusOK,
			Object: Response{
				Code: iris.StatusOK,
				Message: message,
				Data: payload["data"],
				Meta: reflect.ValueOf(meta).Interface().(*Meta),
			},
		}
	}

	return mvc.Response{
		Code: iris.StatusOK,
		Object: Response{
			Code: iris.StatusOK,
			Message: message,
			Data: payload["data"],
		},
	}
}

// 认证失败响应
func UnAuthenticated(message string) mvc.Result {
	return mvc.Response{
		Code: iris.StatusUnauthorized,
		Object: Response{
			Code: iris.StatusUnauthorized,
			Message: message,
			Data: make(map[string]interface{}),
		},
	}
}

func NotFound(message string) mvc.Result {
	return mvc.Response{
		Code: iris.StatusNotFound,
		Object: Response{
			Code: iris.StatusNotFound,
			Message: message,
			Data: make(map[string]interface{}),
		},
	}
}

func ValidationError(message string) mvc.Result {
	return mvc.Response{
		Code: iris.StatusUnprocessableEntity,
		Object: Response{
			Code: iris.StatusUnprocessableEntity,
			Message: message,
			Data: make(map[string]interface{}),
		},
	}
}

func InternalServerError(message string, err error) mvc.Result {
	return mvc.Response{
		Code: iris.StatusInternalServerError,
		Object: Response{
			Code: iris.StatusInternalServerError,
			Message: message,
			Data: err,
		},
	}
}

func Send(code int, message string, data interface{}) mvc.Result {
	return mvc.Response{
		Code: code,
		Object: Response{
			Code: code,
			Message: message,
			Data: data,
		},
	}
}