package main

import (
	"godemo/internal/golib/httpdemo/proxy"
)

func main() {
	proxy.CommitDeviceSshReq("127.0.0.1:8090")
}
