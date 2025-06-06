package response

import (
	"net/http"
)

type (
	Meta struct {
		Page  int   `json:"page"`
		Limit int   `json:"limit"`
		Total int64 `json:"total"`
		Start int
	}
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta    *Meta       `json:"meta,omitempty"`
	}
)

// Success Sending a successful response
func Success(message string, data interface{}, meta *Meta) Response {
	if data == nil {
		return Response{
			Code:    http.StatusOK,
			Message: message,
			Data:    struct{}{},
		}
	}

	return Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

// UnAuthenticated Authentication Failure
func UnAuthenticated(message string) Response {
	return Response{
		Code:    http.StatusUnauthorized,
		Message: message,
		Data:    struct{}{},
	}
}

// NotFound Sending a not found response
func NotFound(message string) Response {
	return Response{
		Code:    http.StatusNotFound,
		Message: message,
		Data:    struct{}{},
	}
}

// ValidationError Sending a validation error response
func ValidationError(message string, err error) Response {
	return Response{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
		Data:    struct{}{},
	}
}

// InternalServerError Sending a internal server error response
func InternalServerError(message string, err error) Response {
	return Response{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    struct{}{},
	}
}

func (r *Response) setMeta(meta Meta) {
	r.Meta = &meta
}

func (r *Response) setData(data interface{}) {
	r.Data = data
}

func WithMeta(meta Meta) func(*Response) {
	return func(response *Response) {
		response.setMeta(meta)
	}
}

func WithData(data interface{}) func(*Response) {
	return func(response *Response) {
		response.setData(data)
	}
}

// Send Sending a basic response
func Send(options ...func(response *Response)) *Response {
	response := Response{}
	for _, option := range options {
		option(&response)
	}
	return &response
}
