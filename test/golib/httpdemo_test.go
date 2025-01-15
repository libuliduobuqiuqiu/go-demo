package golib

import (
	"context"
	"godemo/internal/golib/httpdemo"
	"godemo/internal/golib/httpdemo/proxy"
	"testing"
)

func TestHttpClientGet(t *testing.T) {
	url := "http://127.0.0.1:8989/person"
	data, err := httpdemo.GetRequest(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}

func TestHttpServer(t *testing.T) {
	httpdemo.HandleHttpRequest()
}

func TestReverseService(t *testing.T) {
	proxy.StartReverseProxy()
}

func TestCommitHttpReq(t *testing.T) {
	tmpUrl := "http://127.0.0.1:8090/netac?proxy_pass=https://10.21.21.64:443/mgmt/tm/ltm/monitor/http/~Common~mo_http5055"
	// tmpUrl := "http://127.0.0.1:8090/netac?proxy_pass=http://127.0.0.1:8989/person"
	if err := proxy.CommitDeviceHttpReq(tmpUrl); err != nil {
		t.Fatal(err)
	}
}

func TestCommitSshReq(t *testing.T) {
	addr := "127.0.0.1:8090"
	proxy.CommitDeviceSshReq(addr)
}
