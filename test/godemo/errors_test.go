package godemo

import (
	"godemo/internal/godemo/errors"
	"godemo/pkg"
	"testing"
)

func TestHandleError(t *testing.T) {

	err := errors.HandleError()
	if err != nil {
		t.Log(err)
	}

	err = errors.NewMyError()
	t.Error(err)

}

func TestStackError(t *testing.T) {

	errors.UseStackError()

}

func TestCustomError(t *testing.T) {
	err := pkg.NewErr("test custom err")
	pkg.PrintErr(err)
}
