package netdemo

import (
	"crypto/rand"
	"net"
)

// 生成随机 IPv6 地址
func randomIPv6() net.IP {
	ip := make([]byte, 16) // IPv6 是 16 字节
	_, err := rand.Read(ip)
	if err != nil {
		panic(err)
	}
	return net.IP(ip)
}

// 根据随机地址生成子网，例如 /64
func RandomIPv6Subnet(prefixLen int) *net.IPNet {
	ip := randomIPv6()
	mask := net.CIDRMask(prefixLen, 128)
	return &net.IPNet{
		IP:   ip.Mask(mask),
		Mask: mask,
	}
}
