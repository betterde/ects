package response

import (
	"net/http"
)

type (
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

// Success Sending a successful response
func Success(message string, data interface{}) Response {
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

// Send Sending a basic response
func Send(code int, message string, data interface{}) Response {
	if data == nil {
		return Response{
			Code:    code,
			Message: message,
			Data:    struct{}{},
		}
	}

	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
