package response

import (
	"github.com/betterde/ects/models"
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
func Success(message string, payload map[string]interface{}) mvc.Response {
	data := payload["data"]
	meta, ok := payload["meta"]
	if ok {
		return mvc.Response{
			Code: iris.StatusOK,
			Object: Response{
				Code: iris.StatusOK,
				Message: message,
				Data: data,
				Meta: reflect.ValueOf(meta).Interface().(*Meta),
			},
		}
	}

	switch reflect.TypeOf(data).String() {
	case "models.Pipeline":
		pipeline := reflect.ValueOf(data).Interface().(models.Pipeline)

		return mvc.Response{
			Code: iris.StatusOK,
			Object: Response{
				Code: iris.StatusOK,
				Message: message,
				Data: &pipeline,
			},
		}
	case "models.Node":
		node := reflect.ValueOf(data).Interface().(models.Node)

		return mvc.Response{
			Code: iris.StatusOK,
			Object: Response{
				Code: iris.StatusOK,
				Message: message,
				Data: &node,
			},
		}
	case "models.User":
		user := reflect.ValueOf(data).Interface().(models.User)

		return mvc.Response{
			Code: iris.StatusOK,
			Object: Response{
				Code: iris.StatusOK,
				Message: message,
				Data: &user,
			},
		}
	case "models.Task":
		task := reflect.ValueOf(data).Interface().(models.Task)

		return mvc.Response{
			Code: iris.StatusOK,
			Object: Response{
				Code: iris.StatusOK,
				Message: message,
				Data: &task,
			},
		}
	case "models.Log":
		log := reflect.ValueOf(data).Interface().(models.Log)

		return mvc.Response{
			Code: iris.StatusOK,
			Object: Response{
				Code: iris.StatusOK,
				Message: message,
				Data: &log,
			},
		}
	case "models.Team":
		team := reflect.ValueOf(data).Interface().(models.Team)

		return mvc.Response{
			Code: iris.StatusOK,
			Object: Response{
				Code: iris.StatusOK,
				Message: message,
				Data: &team,
			},
		}
	}

	return mvc.Response{
		Code: iris.StatusOK,
		Object: Response{
			Code: iris.StatusOK,
			Message: message,
			Data: data,
		},
	}
}

// 认证失败响应
func UnAuthenticated(message string) mvc.Response {
	return mvc.Response{
		Code: iris.StatusUnauthorized,
		Object: Response{
			Code: iris.StatusUnauthorized,
			Message: message,
			Data: make(map[string]interface{}),
		},
	}
}

func NotFound(message string) mvc.Response {
	return mvc.Response{
		Code: iris.StatusNotFound,
		Object: Response{
			Code: iris.StatusNotFound,
			Message: message,
			Data: make(map[string]interface{}),
		},
	}
}

func ValidationError(message string) mvc.Response {
	return mvc.Response{
		Code: iris.StatusUnprocessableEntity,
		Object: Response{
			Code: iris.StatusUnprocessableEntity,
			Message: message,
			Data: make(map[string]interface{}),
		},
	}
}

func InternalServerError(message string, err error) mvc.Response {
	return mvc.Response{
		Code: iris.StatusInternalServerError,
		Object: Response{
			Code: iris.StatusInternalServerError,
			Message: message,
			Data: err,
		},
	}
}

func Send(code int, message string, data interface{}) mvc.Response {
	return mvc.Response{
		Code: code,
		Object: Response{
			Code: code,
			Message: message,
			Data: data,
		},
	}
}