package core

import (
	"errors"
	"fmt"
)

type ErrorCode string

const (
	BadRequestError    ErrorCode = "BadRequestError"
	ValidationError    ErrorCode = "ValidationError"
	NotFoundError      ErrorCode = "NotFoundError"
	AlreadyExistsError ErrorCode = "AlreadyExistsError"
	SystemError        ErrorCode = "SystemError"
)

type appError struct {
	code ErrorCode
	err  error
}

func NewError(code ErrorCode, message string) error {
	return &appError{
		code: code,
		err:  fmt.Errorf(message),
	}
}

func AsAppError(err error) *appError {
	var dst *appError
	if ok := errors.As(err, &dst); !ok {
		return &appError{
			code: SystemError,
			err:  fmt.Errorf("致命的なシステムエラーが発生しました。"),
		}
	}
	return dst
}

func (e *appError) Code() ErrorCode {
	return e.code
}

func (e *appError) Error() string {
	return e.err.Error()
}
