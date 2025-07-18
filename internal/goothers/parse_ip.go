package goothers

import (
	"fmt"
	"strconv"
	"strings"
)

// PortRange 表示一个端口范围
type PortRange struct {
	Start int
	End   int
}

// PortProtocol 表示一个端口或端口范围与协议的组合
type PortProtocol struct {
	Port    PortRange
	Proto   string
	IsRange bool
}

// HasDuplicates 检查两个端口协议列表是否有重复项
func HasDuplicates(list1, list2 []string) (bool, error) {
	// 解析第一个列表
	ppList1, err := parsePortProtocolList(list1)
	if err != nil {
		return false, err
	}

	// 解析第二个列表
	ppList2, err := parsePortProtocolList(list2)
	if err != nil {
		return false, err
	}

	// 检查是否有重叠
	for _, pp1 := range ppList1 {
		for _, pp2 := range ppList2 {
			if pp1.Proto == pp2.Proto && hasPortOverlap(pp1.Port, pp2.Port) {
				return true, nil
			}
		}
	}

	return false, nil
}

// parsePortProtocolList 解析端口协议列表
func parsePortProtocolList(list []string) ([]PortProtocol, error) {
	var result []PortProtocol

	for _, item := range list {
		pp, err := parsePortProtocol(item)
		if err != nil {
			return nil, err
		}
		result = append(result, pp)
	}

	return result, nil
}

// parsePortProtocol 解析单个端口协议字符串
func parsePortProtocol(s string) (PortProtocol, error) {
	// 标准化字符串
	s = strings.TrimSpace(s)
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		return PortProtocol{}, fmt.Errorf("invalid format: %s, expected port/protocol", s)
	}

	// 解析协议
	proto := strings.TrimSpace(strings.ToLower(parts[1]))
	if proto != "tcp" && proto != "udp" {
		return PortProtocol{}, fmt.Errorf("invalid protocol: %s, only tcp/udp supported", proto)
	}

	// 解析端口部分
	portStr := strings.TrimSpace(parts[0])
	if portStr == "" {
		return PortProtocol{}, fmt.Errorf("empty port in: %s", s)
	}

	// 检查是否是端口范围
	if strings.Contains(portStr, "-") {
		rangeParts := strings.Split(portStr, "-")
		if len(rangeParts) != 2 {
			return PortProtocol{}, fmt.Errorf("invalid port range format: %s", portStr)
		}

		start, err := strconv.Atoi(strings.TrimSpace(rangeParts[0]))
		if err != nil {
			return PortProtocol{}, fmt.Errorf("invalid start port: %s", rangeParts[0])
		}

		end, err := strconv.Atoi(strings.TrimSpace(rangeParts[1]))
		if err != nil {
			return PortProtocol{}, fmt.Errorf("invalid end port: %s", rangeParts[1])
		}

		if start > end {
			return PortProtocol{}, fmt.Errorf("start port %d is greater than end port %d", start, end)
		}

		if !isValidPort(start) || !isValidPort(end) {
			return PortProtocol{}, fmt.Errorf("port out of range (1-65535): %s", portStr)
		}

		return PortProtocol{
			Port:    PortRange{Start: start, End: end},
			Proto:   proto,
			IsRange: true,
		}, nil
	}

	// 单个端口
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return PortProtocol{}, fmt.Errorf("invalid port: %s", portStr)
	}

	if !isValidPort(port) {
		return PortProtocol{}, fmt.Errorf("port out of range (1-65535): %d", port)
	}

	return PortProtocol{
		Port:    PortRange{Start: port, End: port},
		Proto:   proto,
		IsRange: false,
	}, nil
}

// hasPortOverlap 检查两个端口范围是否有重叠
func hasPortOverlap(r1, r2 PortRange) bool {
	return r1.Start <= r2.End && r2.Start <= r1.End
}

// isValidPort 检查端口是否有效
func isValidPort(port int) bool {
	return port >= 1 && port <= 65535
}
