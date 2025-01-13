package golib

import (
	"godemo/internal/golib/stringsdemo"
	"testing"
)

func TestUseStringsIndex(t *testing.T) {
	stringsdemo.CheckStringIndex()
}

func TestUseUrlIndex(t *testing.T) {
	stringsdemo.CheckUrlIndex()
}

func TestConvert(t *testing.T) {
	stringsdemo.ConvertStringToInt64()
}
