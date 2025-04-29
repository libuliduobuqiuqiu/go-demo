package pkg

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const (
	SEPARATOR    = ":"
	DEFAULTERROR = iota + 3000
	INTERNALERROR
	BADREQUESTERROR
)

type StackTraceInterface interface {
	StackTrace() errors.StackTrace
}

type CustomError struct {
	code     int
	messages []string
	err      error
	stacks   errors.StackTrace
}

func (c *CustomError) Error() string {
	if len(c.messages) > 0 {
		msg := strings.Join(c.messages, SEPARATOR)
		return fmt.Sprintf("%s%s%s", msg, SEPARATOR, c.err.Error())
	} else {
		return c.err.Error()
	}
}

func (c *CustomError) Code() int {
	return c.code
}

func (c *CustomError) StackTrace() errors.StackTrace {
	return c.stacks
}

func NewCustomErr(code int, msg string, err error) *CustomError {
	tmp := &CustomError{
		code: code,
		err:  err,
	}

	if v, ok := err.(StackTraceInterface); ok {
		tmp.stacks = v.StackTrace()
	}

	tmp.messages = append(tmp.messages, err.Error(), msg)
	return tmp
}

func NewErrWrf(msg string, err error) error {
	return NewCustomErr(
		INTERNALERROR, msg, err,
	)
}
