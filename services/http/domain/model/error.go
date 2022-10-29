package model

import (
	"fmt"
)

const statusError = "error"

// ErrorResponse structured error description
type ErrorResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error is error interface implementation
func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// SwapCode returns a copy of supplied ErrorResponse with code swapped to given in argument `code`
func (e ErrorResponse) SwapCode(code int) ErrorResponse {
	e.Code = code
	return e
}

// NewErrorResponse is a constructor for ErrorResponse
func NewErrorResponse(code int, err error) ErrorResponse {
	var msg string
	if err != nil {
		msg = err.Error()
	}
	return ErrorResponse{
		Code:    code,
		Message: msg,
		Status:  statusError,
	}
}

// NewErrorResponseWithMessage is a constructor for ErrorResponse
func NewErrorResponseWithMessage(code int, message string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
		Status:  statusError,
	}
}
