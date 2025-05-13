package errors

import (
	"fmt"
	"runtime/debug"

	"github.com/pkg/errors"
)

func RaiseError() error {
	return fmt.Errorf("test raise error.")
}

func WrapError() error {
	err := RaiseError()
	return errors.Wrap(err, "Wrap Error: ")
}

type MyError struct {
	msg string
	err error
}

func (m MyError) Error() string {
	return fmt.Sprintf("error occurred: %s\nStack: \n%s", m.msg, debug.Stack())
}

func (m MyError) UnWrap() error {
	return m.err
}

func NewMyError() error {
	fmt.Println(string(debug.Stack()))
	return MyError{
		msg: "test my Error",
		err: errors.New(fmt.Sprintf("test my Wrap Error")),
	}
}

func HandleError() error {
	return fmt.Errorf("%w", NewMyError())
}
