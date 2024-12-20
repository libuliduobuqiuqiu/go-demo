package golib

import (
	"godemo/internal/golib/jsondemo"
	"testing"
)

func TestJsonUnmarshal(t *testing.T) {

	err := jsondemo.UnmarshalObj()
	if err != nil {
		t.Fatal(err)
	}

}

