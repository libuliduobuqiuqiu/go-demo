package godemo

import (
	"fmt"
	"godemo/internal/godemo"
	"testing"
	"time"
)

func TestPaserUnixTime(t *testing.T) {

	godemo.FromUnixTime()

}

func TestPrintTime(t *testing.T) {

	now := time.Now()
	fmt.Println(now.Format("20060102150405.000"))
	fmt.Printf("%s%03d\n", now.Format("20060102150405"), now.Nanosecond()/1e6)

}
