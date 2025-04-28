package traceroute

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// handlers 存储不同操作系统的处理逻辑
	handlers = make(map[string]TracerouteInterface)
	// mu 用于保护handlers的并发访问
	mu sync.RWMutex
)

// Register 注册一个操作系统的处理逻辑
func Register(os string, handler TracerouteInterface) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := handlers[os]; exists {
		return fmt.Errorf("操作系统 %s 的处理逻辑已经注册", os)
	}

	handlers[os] = handler
	return nil
}

// GetHandler 获取当前操作系统的处理逻辑
func GetHandler() (TracerouteInterface, error) {
	mu.RLock()
	defer mu.RUnlock()

	handler, exists := handlers[runtime.GOOS]
	if !exists {
		return nil, fmt.Errorf("不支持的操作系统: %s", runtime.GOOS)
	}

	return handler, nil
}
