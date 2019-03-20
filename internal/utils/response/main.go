package response

import "github.com/kataras/iris"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 发送成功响应
func Success(message string, data interface{}) *Response {
	return &Response{
		Code: iris.StatusOK,
		Message: message,
		Data: data,
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