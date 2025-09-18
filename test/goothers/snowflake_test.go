package goothers

import (
	"godemo/internal/goothers"
	"testing"
)

func TestReplaceSnowFlake(t *testing.T) {
	err := goothers.ReplaceUUIDWithSnowflakeID("H3C-SW.txt")
	if err != nil {
		t.Fatal(err)
	}
}
