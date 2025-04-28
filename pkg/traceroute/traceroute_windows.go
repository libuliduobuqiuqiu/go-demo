//go:build windows

package traceroute

import (
	"log"
	"net"
	"syscall"
)

func init() {
	// 注册Windows处理逻辑
	if err := Register("windows", NewWindowsTraceroute()); err != nil {
		log.Fatal(err)
	}
}

// WindowsTraceroute Windows系统的traceroute实现
type WindowsTraceroute struct{}

// NewWindowsTraceroute 创建Windows系统的traceroute实现
func NewWindowsTraceroute() *WindowsTraceroute {
	return &WindowsTraceroute{}
}

// SetTTL 设置连接的TTL值
func (w *WindowsTraceroute) SetTTL(conn net.Conn, ttl int) error {
	// 获取底层文件描述符
	file, err := conn.(*net.IPConn).File()
	if err != nil {
		return err
	}
	defer file.Close()

	// 设置TTL
	return syscall.SetsockoptInt(syscall.Handle(file.Fd()), syscall.IPPROTO_IP, syscall.IP_TTL, ttl)
}

// SendICMP 发送ICMP回显请求包
func (w *WindowsTraceroute) SendICMP(conn net.Conn) error {
	// 构造ICMP包
	icmp := []byte{
		8, 0, 0, 0, // Type=8 (Echo Request), Code=0, Checksum=0
		0, 1, // Identifier
		0, 1, // Sequence Number
	}

	// 计算校验和
	checksum := calculateChecksum(icmp)
	icmp[2] = byte(checksum >> 8)
	icmp[3] = byte(checksum & 0xff)

	// 发送ICMP包
	_, err := conn.Write(icmp)
	return err
}

// ParseReplyIP 从ICMP响应中解析IP地址
func (w *WindowsTraceroute) ParseReplyIP(reply []byte) net.IP {
	if len(reply) < 20 {
		return nil
	}

	// 检查ICMP类型
	icmpType := reply[0]
	if icmpType != 0 && icmpType != 11 { // 0=Echo Reply, 11=Time Exceeded
		return nil
	}

	// 从IP头中提取源IP
	return net.IP(reply[12:16])
}
