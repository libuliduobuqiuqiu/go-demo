package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"godemo/pkg/traceroute"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("用法: %s <目标主机>\n", os.Args[0])
		os.Exit(1)
	}

	dest := os.Args[1]
	fmt.Printf("traceroute to %s (%s), %d hops max\n", dest, runtime.GOOS, traceroute.DefaultConfig().MaxHops)
	fmt.Printf("当前系统: %s/%s\n", runtime.GOOS, runtime.GOARCH)

	start := time.Now()
	results, err := traceroute.Traceroute(dest, traceroute.DefaultConfig())
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n追踪结果:\n")
	for _, result := range results {
		if result.Success {
			fmt.Printf("%d  %s  %v\n", result.Hop, result.IP, result.RTT)
		} else {
			fmt.Printf("%d  * * *\n", result.Hop)
		}
	}

	fmt.Printf("\n总耗时: %v\n", time.Since(start))
}
