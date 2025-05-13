package pkg

import (
	"fmt"
	"runtime"
)

const (
	DEPTH        = 32
	SEPARATOR    = ":"
	DEFAULTERROR = iota + 3000
	INTERNALERROR
	BADREQUESTERROR
)

type Frame struct {
	File string
	Line int
	Name string
}

type StackTraceInterface interface {
	StackTrace() []Frame
}

type CustomError struct {
	code   int
	msg    string
	cause  error
	stacks []uintptr
}

func (c *CustomError) Error() string {
	if c.cause != nil {
		return fmt.Sprintf("Code=%d, Message=%s, Cauese=%v", c.code, c.msg, c.cause)
	}
	return fmt.Sprintf("Code=%d, Message=%s", c.code, c.msg)
}

func (c *CustomError) Code() int {
	return c.code
}

func (c *CustomError) StackTrace() (stackFrame []Frame) {
	frames := runtime.CallersFrames(c.stacks)

	for {
		frame, more := frames.Next()

		tmp := Frame{
			File: frame.File,
			Line: frame.Line,
			Name: frame.Function,
		}
		stackFrame = append(stackFrame, tmp)

		if !more {
			break
		}
	}

	return
}

func (c *CustomError) Unwrap() error {
	return c.cause
}

func newCustomErr(code int, msg string, cause error) *CustomError {
	pcs := make([]uintptr, DEPTH)
	n := runtime.Callers(2, pcs)

	return &CustomError{
		code:   code,
		cause:  cause,
		msg:    msg,
		stacks: pcs[:n],
	}
}

func NewErr(msg string) error {
	return newCustomErr(INTERNALERROR, msg, nil)
}
func WrapErr(code int, msg string, err error) error {
	return newCustomErr(code, msg, err)
}
