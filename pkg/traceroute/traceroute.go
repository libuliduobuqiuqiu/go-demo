package traceroute

import (
	"fmt"
	"net"
	"time"
)

// Result 表示一个跳点的结果
type Result struct {
	Hop     int
	IP      net.IP
	RTT     time.Duration
	Success bool
}

// TracerouteInterface 定义traceroute的接口
type TracerouteInterface interface {
	// SetTTL 设置连接的TTL值
	SetTTL(conn net.Conn, ttl int) error
	// SendICMP 发送ICMP回显请求包
	SendICMP(conn net.Conn) error
	// ParseReplyIP 从ICMP响应中解析IP地址
	ParseReplyIP(reply []byte) net.IP
}

// TracerouteConfig 定义traceroute的配置
type TracerouteConfig struct {
	MaxHops    int
	TimeoutMs  int
	PacketSize int
}

// DefaultConfig 返回默认配置
func DefaultConfig() TracerouteConfig {
	return TracerouteConfig{
		MaxHops:    30,
		TimeoutMs:  1000,
		PacketSize: 52,
	}
}

// Traceroute 执行traceroute操作
func Traceroute(dest string, config TracerouteConfig) ([]Result, error) {

	impl, err := GetHandler()
	if err != nil {
		return nil, fmt.Errorf("获取处理器失败: %v", err)
	}

	destAddr, err := net.ResolveIPAddr("ip4", dest)
	if err != nil {
		return nil, fmt.Errorf("无法解析目标地址: %v", err)
	}
	fmt.Println("destAddr", destAddr)

	results := make([]Result, 0, config.MaxHops)
	conn, err := net.DialIP("ip4:icmp", nil, destAddr)
	if err != nil {
		return nil, fmt.Errorf("无法创建ICMP连接: %v", err)
	}
	defer conn.Close()

	// 设置TTL从1开始
	for ttl := 1; ttl <= config.MaxHops; ttl++ {
		// 设置TTL
		if err := impl.SetTTL(conn, ttl); err != nil {
			return results, fmt.Errorf("设置TTL失败: %v", err)
		}

		// 发送ICMP包
		start := time.Now()
		if err := impl.SendICMP(conn); err != nil {
			return results, fmt.Errorf("发送ICMP包失败: %v", err)
		}
		fmt.Println("发送ICMP包成功")

		// 接收响应
		reply := make([]byte, config.PacketSize)
		conn.SetReadDeadline(time.Now().Add(time.Duration(config.TimeoutMs) * time.Millisecond))
		_, err := conn.Read(reply)
		if err != nil {
			fmt.Println("接收响应失败", err)
		}
		rtt := time.Since(start)

		result := Result{
			Hop:     ttl,
			RTT:     rtt,
			Success: err == nil,
		}

		if err == nil {
			// 解析响应IP
			if ip := impl.ParseReplyIP(reply); ip != nil {
				result.IP = ip
			}
		}

		if result.Success {
			fmt.Printf("%d  %s  %v\n", result.Hop, result.IP, result.RTT)
		} else {
			fmt.Printf("%d  * * *\n", result.Hop)
		}

		results = append(results, result)

		// 如果到达目标，结束traceroute
		if result.Success && result.IP.Equal(destAddr.IP) {
			break
		}
	}

	return results, nil
}
