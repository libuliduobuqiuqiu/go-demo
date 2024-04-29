package osdemo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var lines = []string{
	"zhangsan",
	"lis",
	"wangwu",
	"zhaosan",
	"ifconfig",
	"top",
	"netstat",
}

func UseFileUtil() error {

	f, err := os.Create("new_file.temp")
	if err != nil {
		return err
	}

	defer f.Close()

	// 全部写入
	f.Write([]byte("hello,world\n"))

	// 逐行写入
	for _, line := range lines {
		f.WriteString(line + "\n")
	}

	// 缓存写入
	bf := bufio.NewWriter(f)
	for _, line := range lines {
		bf.WriteString(line + "\n")
	}

	// 将缓存写入文件
	if err = bf.Flush(); err != nil {
		return err
	}

	return nil
}

// 超过缓存空间大小，直接插入
// 优势对比os包，在处理小文件时，减少系统调用次数，提高文件读写效率
// 同时在读取大文件时，也能够避免一次性读取大文件到内存，导致内存占用过高，影响其他服务
func UseBufioMaxInsert() {
	var tmpStrs []string

	for i := 0; i < 1000; i++ {
		tmpStrs = append(tmpStrs, lines...)
	}

	f, err := os.Create("max_file.temp")
	if err != nil {
		fmt.Println(err)
		return
	}

	bf := bufio.NewWriter(f)
	str := strings.Join(tmpStrs, "-")
	fmt.Println(len(str))
	count, err := bf.WriteString(str)
	fmt.Println("Insert: ", count)

	time.Sleep(20 * time.Second)

	if err := bf.Flush(); err != nil {
		fmt.Println(err)
		return
	}
}
