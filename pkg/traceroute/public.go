package traceroute

// calculateChecksum 计算ICMP校验和
// description: 计算ICMP校验和
// param: data []byte 数据
// return: uint16 校验和
// author:
// date: 2025-04-28
// 逻辑：
// 1. 将数据分成两个字节一组
// 2. 定义一个变量sum，初始值为0
// 3. 遍历数据，将每两个字节转换为uint16，并累加到sum中
// 4. 将sum右移16位，并和sum进行按位或运算
// 5. 将sum右移16位，并和sum进行按位或运算
// 6. 返回sum的补码
func calculateChecksum(data []byte) uint16 {
	var sum uint32
	for i := 0; i < len(data)-1; i += 2 {
		sum += uint32(data[i])<<8 | uint32(data[i+1])
	}
	if len(data)%2 != 0 {
		sum += uint32(data[len(data)-1]) << 8
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += sum >> 16
	return ^uint16(sum)
}
