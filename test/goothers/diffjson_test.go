package goothers

import (
	"godemo/internal/goothers"
	"testing"
)

func TestDiffJson(t *testing.T) {

	// goothers.DiffJson("/mnt/d/Company/provider_json/local_critx.json", "/mnt/d/Company/provider_json/remote_critx.json")

	goothers.DiffJson(
		"/mnt/d/Company/provider_json/local_hillstone_ngfw.json",
		"/mnt/d/Company/provider_json/remote_hillstone_ngfw.json",
	)

	DiffNginxJson()

}

func DiffNginxJson() {
	goothers.DiffJson(
		"/mnt/d/Company/provider_json/local_nginx.json",
		"/mnt/d/Company/provider_json/remote_nginx.json",
	)

}
