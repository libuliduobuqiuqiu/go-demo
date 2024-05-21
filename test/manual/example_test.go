package manual

import (
	"fmt"
	"godemo/internal/gomanual/genericsdemo"
	"testing"
)

func SayHello() {
	fmt.Println("hello")
}

func SayBye() {
	fmt.Println("bye")
}

func ExampleSayHello() {
	SayHello()
	// Output:
	// hello
}

func ExampleSayBye() {
	SayBye()
	// Output:
	// bye
}

type myInt int

func TestSumInt(t *testing.T) {
	a, b := 1, 101
	expected := 103

	actual := genericsdemo.SumInt[int](a, b)
	if expected != actual {
		t.Errorf("SumInt(%d, %d) expected %d, actual is %d\n", a, b, expected, actual)
	}
}

func TestEqual(t *testing.T) {
	t.Cleanup(func() {
		t.Helper()
		t.Logf("finished.")
	})
	a, b := myInt(101), myInt(101)

	if !genericsdemo.Equal[myInt](a, b) {
		t.Errorf("equal(%d, %d) error", a, b)
	}
}
