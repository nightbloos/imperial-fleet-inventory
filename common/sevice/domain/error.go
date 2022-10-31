package domain

import (
	"errors"
	"fmt"
)

const (
	// ErrNotFound used to mark error as not found
	ErrNotFound string = "not_found"
	// ErrInternal used to mark error as internal
	ErrInternal string = "internal"
)

// NewNotFoundError creates new Not Found error
func NewNotFoundError(msg string) Error {
	return WrapWithNotFoundError(nil, msg)
}

// WrapWithNotFoundError wraps existing error with Not Found error
func WrapWithNotFoundError(err error, msg string) Error {
	return Error{errorType: ErrNotFound, err: err, msg: msg}
}

// NewInternalError creates new Internal error
func NewInternalError(msg string) Error {
	return WrapWithInternalError(nil, msg)
}

// WrapWithInternalError wraps existing error with Internal error
func WrapWithInternalError(err error, msg string) Error {
	return Error{errorType: ErrInternal, err: err, msg: msg}
}

// Error as common errors for our services
type Error struct {
	msg       string
	err       error
	errorType string
}

// GetErrorType returns error type
func (e Error) GetErrorType() string {
	return e.errorType
}

func (e Error) Error() string {
	if e.err == nil {
		return e.msg
	}
	if e.msg == "" {
		return e.err.Error()
	}
	return fmt.Sprintf("%s: %s", e.msg, e.err.Error())
}

// Cause returns errors cause
func (e Error) Cause() error {
	return e.err
}

// Unwrap unwraps error
func (e Error) Unwrap() error {
	return e.err
}

// IsNotFoundError checks if NotFound type error provided
func IsNotFoundError(e error) bool {
	var err Error
	if errors.As(e, &err) && err.GetErrorType() == ErrNotFound {
		return true
	}
	return false
}

// IsInternalError checks if Internal type error provided
func IsInternalError(e error) bool {
	var err Error
	if errors.As(e, &err) && err.GetErrorType() == ErrInternal {
		return true
	}
	return false
}
