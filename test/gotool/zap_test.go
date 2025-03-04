package gotool

import (
	"godemo/internal/gotool/loggingdemo/zapdemo"
	"testing"
)

func TestZapLogging(t *testing.T) {
	err := zapdemo.UseZapLogging()
	if err != nil {
		t.Fatal(err)
	}
	// zapdemo.UseZapExample()
	zapdemo.UseZapProduction()
	// zapdemo.UseZapProductionSuger()
}
