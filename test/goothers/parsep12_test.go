package goothers

import (
	"godemo/internal/goothers"
	"testing"
)

func TestParseP12(t *testing.T) {
	password := "111111"
	err := goothers.ParseSM2P12("sm2enc.p12", password)
	if err != nil {
		t.Fatal(err)
	}

	err = goothers.ParseSM2P12Bydata("sm2sign.p12", password)
	if err != nil {
		t.Fatal(err)
	}

	// err = goothers.ParseRegular("certificate.p12", "123456")
	// if err != nil {
	// 	t.Fatal(err)
	// }
}
