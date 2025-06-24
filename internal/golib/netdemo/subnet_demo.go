package netdemo

import (
	"fmt"
	"net"
)

func getBroadCast(subnet *net.IPNet) net.IP {
	broadcast := make(net.IP, len(subnet.IP))
	copy(broadcast, subnet.IP)
	for i := range broadcast {
		broadcast[i] |= ^subnet.Mask[i]
	}

	return broadcast
}

// 检查subnet1是否包含subnet2
func isSubnetContained(subnet1, subnet2 *net.IPNet) bool {

	// 判断IP版本
	if len(subnet1.IP) != len(subnet2.IP) {
		return false
	}

	mask1, _ := subnet1.Mask.Size()
	mask2, _ := subnet2.Mask.Size()

	// 只有subnet1子网比subnet2短（范围更大）才能包含
	if mask1 > mask2 {
		return false
	}

	// 判断subnet1是否包含subnet2起始IP
	if !(subnet1.Contains(subnet2.IP)) {
		return false
	}

	broadcast := getBroadCast(subnet2)
	return subnet1.Contains(broadcast)

}

// 检查两个子网是否重复
func isOverlapping(subnet1, subnet2 *net.IPNet) bool {

	if subnet1.Contains(subnet2.IP) {
		return true
	}

	if subnet2.Contains(subnet1.IP) {
		return true
	}

	broadcast1 := getBroadCast(subnet1)
	if subnet2.Contains(broadcast1) {
		return true
	}

	broadcast2 := getBroadCast(subnet2)
	return subnet1.Contains(broadcast2)
}

// CheckSubnetRelationship 检查两个子网的关系
func CheckSubnetRelationship(cidr1, cidr2 string) (string, error) {
	// 解析CIDR
	_, subnet1, err := net.ParseCIDR(cidr1)
	if err != nil {
		return "", fmt.Errorf("解析CIDR1失败: %v", err)
	}

	_, subnet2, err := net.ParseCIDR(cidr2)
	if err != nil {
		return "", fmt.Errorf("解析CIDR2失败: %v", err)
	}

	// 检查IP版本是否相同
	if len(subnet1.IP) != len(subnet2.IP) {
		return "IP版本不同(IPv4与IPv6)", nil
	}

	// 判断关系
	switch {
	case subnet1.String() == subnet2.String():
		return "两个子网完全相同", nil
	case isSubnetContained(subnet1, subnet2):
		return fmt.Sprintf("%s 完全包含 %s", cidr1, cidr2), nil
	case isSubnetContained(subnet2, subnet1):
		return fmt.Sprintf("%s 完全包含 %s", cidr2, cidr1), nil
	case isOverlapping(subnet1, subnet2):
		return "两个子网部分重叠但不互相包含", nil
	default:
		return "两个子网无重叠", nil
	}
}

func CheckSubnetContained() error {

	// 测试用例包含IPv4和IPv6
	testCases := []struct {
		cidr1 string
		cidr2 string
	}{
		// IPv4测试
		{"192.168.1.0/24", "192.168.1.128/25"},
		{"10.0.0.0/8", "10.1.2.0/24"},
		{"172.16.0.0/16", "172.17.0.0/16"},
		{"192.168.1.0/24", "192.168.1.0/24"},
		{"192.168.1.0/25", "192.168.1.128/25"},

		// IPv6测试
		{"2001:db8::/32", "2001:db8:abcd::/48"},
		{"2001:db8:abcd::/48", "2001:db8:abcd:1234::/64"},
		{"2001:db8::/32", "2001:db9::/32"},
		{"2001:db8:abcd::/48", "2001:db8:abcd::/48"},
		{"2001:db8:abcd::/48", "2001:db8:abcd:1234::/64"},

		// IPv4与IPv6混合测试
		{"192.168.1.0/24", "2001:db8::/32"},
	}

	for _, tc := range testCases {
		result, err := CheckSubnetRelationship(tc.cidr1, tc.cidr2)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			continue
		}
		fmt.Printf("%-18s %-22s => %s\n", tc.cidr1, tc.cidr2, result)
	}

	return nil
}
