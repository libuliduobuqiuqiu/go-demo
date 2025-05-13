package errors

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	depth = 32
)

type StackError struct {
	msg   string
	stack []uintptr
}

func (s *StackError) Error() string {
	return s.msg
}

func (s *StackError) StackTrace() string {
	var sb strings.Builder
	frames := runtime.CallersFrames(s.stack)

	for {
		frame, more := frames.Next()
		sb.WriteString(fmt.Sprintf("%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line))

		if !more {
			break
		}
	}

	return sb.String()
}

func NewStackError(msg string) error {
	pcs := make([]uintptr, depth)

	n := runtime.Callers(2, pcs)
	return &StackError{
		msg:   msg,
		stack: pcs[:n],
	}
}

func UseStackError() {
	s := NewStackError("test stack error")

	if v, ok := s.(*StackError); ok {
		fmt.Println("Error:", v.Error())
		fmt.Println("StackErrors:")
		fmt.Println(v.StackTrace())
	}
}
