package golib_test

import (
	"godemo/internal/golib/netdemo"
	"net"
	"testing"
)

func TestScanALLIp(t *testing.T) {

	// subnet := "192.168.10.0/24"
	subnet := "2001:0db8:85a3::/64"
	startIP := "2001:0db8:85a3:0000:0000:0000:0001:0000"
	endIP := "2001:0db8:85a3:0000:0000:0000:0002:0000"
	netdemo.GetAllIP(subnet, startIP, endIP)
}

func TestCountIPs(t *testing.T) {
	// subnet := "2001:0db8:85a3::/64"
	subnet := "192.168.1.0/24"
	startIP := "192.168.1.1"
	endIP := "192.168.1.10"
	netdemo.GetAllIP(subnet, startIP, endIP)
}

func TestIncrement(t *testing.T) {
	a := net.ParseIP("192.168.255.255")
	netdemo.Increment(a)
}

func TestNetParseIP(t *testing.T) {
	ip := "240e:6b1:10:1::46"
	tmpIP := net.ParseIP(ip)
	t.Log(tmpIP)
	t.Log(netdemo.IPToBinary(tmpIP))
	t.Log(netdemo.IPToString(tmpIP))
}

func TestParseUrl(t *testing.T) {
	url := "http://127.0.0.1:8090/netac"
	netdemo.ParseUrlString(url)

}
