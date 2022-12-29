package response

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"reflect"
)

type (
	Meta struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
		Total int `json:"total"`
	}
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta    *Meta       `json:"meta,omitempty"`
	}
	Payload map[string]interface{}
)

// 发送成功响应
func Success(message string, payload Payload) mvc.Response {
	data := payload["data"]
	meta, ok := payload["meta"]
	if ok {
		return mvc.Response{
			Code: iris.StatusOK,
			Object: Response{
				Code:    iris.StatusOK,
				Message: message,
				Data:    data,
				Meta:    reflect.ValueOf(meta).Interface().(*Meta),
			},
		}
	}

	return mvc.Response{
		Code: iris.StatusOK,
		Object: Response{
			Code:    iris.StatusOK,
			Message: message,
			Data:    data,
		},
	}
}

// 认证失败响应
func UnAuthenticated(message string) mvc.Response {
	return mvc.Response{
		Code: iris.StatusUnauthorized,
		Object: Response{
			Code:    iris.StatusUnauthorized,
			Message: message,
			Data:    make(map[string]interface{}),
		},
	}
}

func NotFound(message string) mvc.Response {
	return mvc.Response{
		Code: iris.StatusNotFound,
		Object: Response{
			Code:    iris.StatusNotFound,
			Message: message,
			Data:    make(map[string]interface{}),
		},
	}
}

func ValidationError(message string) mvc.Response {
	return mvc.Response{
		Code: iris.StatusUnprocessableEntity,
		Object: Response{
			Code:    iris.StatusUnprocessableEntity,
			Message: message,
			Data:    make(map[string]interface{}),
		},
	}
}

func InternalServerError(message string, err error) mvc.Response {
	return mvc.Response{
		Code: iris.StatusInternalServerError,
		Object: Response{
			Code:    iris.StatusInternalServerError,
			Message: message,
			Data:    err,
		},
	}
}

func Send(code int, message string, data interface{}) mvc.Response {
	return mvc.Response{
		Code: code,
		Object: Response{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}
}
