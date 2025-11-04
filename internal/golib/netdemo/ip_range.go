package netdemo

import (
	"fmt"
	"math/big"
	"net"
	"strings"
)

// IP → big.Int
func ipToBigInt(ip net.IP, byteLen int) *big.Int {
	ip = ip.To16()
	if byteLen == 4 {
		ip = ip.To4()
	}

	return new(big.Int).SetBytes(ip)
}

// big.Int → IP
func bigIntToIP(ipInt *big.Int, byteLen int) net.IP {
	bytes := ipInt.Bytes()
	if len(bytes) < byteLen {
		pad := make([]byte, byteLen-len(bytes))
		bytes = append(pad, bytes...)
	}
	return net.IP(bytes)
}

func CountIPRange(startStr, endStr string) {
	var byteLen int
	if strings.Contains(startStr, ":") {
		byteLen = 16
	} else {
		byteLen = 4
	}

	startIP := net.ParseIP(startStr)
	endIP := net.ParseIP(endStr)
	if startIP == nil || endIP == nil {
		fmt.Println("无效 IP")
		return
	}

	start := ipToBigInt(startIP, byteLen)
	end := ipToBigInt(endIP, byteLen)

	diff := big.NewInt(0).Sub(end, start)
	fmt.Println(diff.String())
	diff.Add(diff, big.NewInt(1))

	fmt.Printf("范围 %s - %s 共 %s 个 IPv6 地址\n", startStr, endStr, diff.String())

	for i := big.NewInt(0); i.Cmp(big.NewInt(256)) <= 0 && start.Cmp(end) <= 0; i.Add(i, big.NewInt(1)) {
		fmt.Println(bigIntToIP(start, byteLen).String())
		start.Add(start, big.NewInt(1))
	}
}
